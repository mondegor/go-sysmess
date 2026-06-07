package helper_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/mondegor/go-sysmess/mrprocess"
	"github.com/mondegor/go-sysmess/mrprocess/helper"
	"github.com/mondegor/go-sysmess/mrprocess/helper/mock"
)

//go:generate mockgen -source=item_batch_player.go -destination=./mock/helper.go -package=mock

func TestItemBatchPlayer_ValidatesBatchSize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		batchSize int
		wantErr   error
	}{
		{name: "zero", batchSize: 0, wantErr: helper.ErrInternalBatchSizeIsZeroOrNegative},
		{name: "negative", batchSize: -1, wantErr: helper.ErrInternalBatchSizeIsZeroOrNegative},
		{name: "greater than total limit", batchSize: 200000, wantErr: helper.ErrInternalBatchSizeIsGreaterThanTotalLimit},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			// обработчик и эмиттер не должны вызываться при невалидном batchSize
			h := mock.NewMockhandler(ctrl)
			e := mock.NewMockeventEmitter(ctrl)
			player := helper.NewItemBatchPlayer(h, e)

			err := player.Execute(context.Background(), tt.batchSize)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestItemBatchPlayer_NegativeCountIsRejected(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	h := mock.NewMockhandler(ctrl)
	h.EXPECT().Execute(gomock.Any(), 10).Return(-1, nil)

	e := mock.NewMockeventEmitter(ctrl) // событие не эмитится при ошибке
	player := helper.NewItemBatchPlayer(h, e)

	err := player.Execute(context.Background(), 10)
	require.ErrorIs(t, err, helper.ErrInternalNegativeProcessedCount)
}

func TestItemBatchPlayer_HandlerError(t *testing.T) {
	t.Parallel()

	errHandler := errors.New("handler failed")

	tests := []struct {
		name    string
		respErr error
		wantErr error
	}{
		{name: "generic error is propagated", respErr: errHandler, wantErr: errHandler},
		{name: "deadline becomes temporary problem", respErr: context.DeadlineExceeded, wantErr: mrprocess.ErrSystemTemporaryProblemHasOccurred},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			h := mock.NewMockhandler(ctrl)
			h.EXPECT().Execute(gomock.Any(), 10).Return(0, tt.respErr)

			e := mock.NewMockeventEmitter(ctrl) // событие не эмитится при ошибке
			player := helper.NewItemBatchPlayer(h, e)

			err := player.Execute(context.Background(), 10)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestItemBatchPlayer_StopsOnPartialBatch(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	h := mock.NewMockhandler(ctrl)
	// полный пакет, полный пакет, затем неполный (последняя пачка) - обработка завершается
	gomock.InOrder(
		h.EXPECT().Execute(gomock.Any(), 10).Return(10, nil),
		h.EXPECT().Execute(gomock.Any(), 10).Return(10, nil),
		h.EXPECT().Execute(gomock.Any(), 10).Return(3, nil),
	)

	e := mock.NewMockeventEmitter(ctrl)
	e.EXPECT().Emit(gomock.Any(), "Execute", "total", 23, "duration_sec", gomock.Any(), "batch_size", 10)
	player := helper.NewItemBatchPlayer(h, e)

	err := player.Execute(context.Background(), 10)
	require.NoError(t, err)
}

func TestItemBatchPlayer_StopsOnEmptyResult(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	h := mock.NewMockhandler(ctrl)
	h.EXPECT().Execute(gomock.Any(), 10).Return(0, nil)

	e := mock.NewMockeventEmitter(ctrl)
	e.EXPECT().Emit(gomock.Any(), "Execute", "total", 0, "duration_sec", gomock.Any(), "batch_size", 10)
	player := helper.NewItemBatchPlayer(h, e)

	err := player.Execute(context.Background(), 10)
	require.NoError(t, err)
}

func TestItemBatchPlayer_StopsOnTotalLimit(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	h := mock.NewMockhandler(ctrl)
	// totalLimit=25, batchSize=10: после 30 обработанных элементов лимит превышен
	gomock.InOrder(
		h.EXPECT().Execute(gomock.Any(), 10).Return(10, nil),
		h.EXPECT().Execute(gomock.Any(), 10).Return(10, nil),
		h.EXPECT().Execute(gomock.Any(), 10).Return(10, nil),
	)

	e := mock.NewMockeventEmitter(ctrl)
	e.EXPECT().Emit(gomock.Any(), "Execute", "total", 30, "duration_sec", gomock.Any(), "batch_size", 10)
	player := helper.NewItemBatchPlayerWithTotalLimit(h, e, 25)

	err := player.Execute(context.Background(), 10)
	require.NoError(t, err)
}
