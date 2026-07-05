package helper_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"github.com/mondegor/go-core/mrprocess"
	"github.com/mondegor/go-core/mrprocess/helper"
	"github.com/mondegor/go-core/mrprocess/helper/mock"
)

//go:generate mockgen -source=item_batch_player.go -destination=./mock/helper.go -package=mock

type ItemBatchPlayerSuite struct {
	suite.Suite

	ctrl *gomock.Controller
	h    *mock.Mockhandler
	e    *mock.MockeventEmitter
}

func TestItemBatchPlayerSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ItemBatchPlayerSuite))
}

func (ts *ItemBatchPlayerSuite) SetupTest() {
	ts.initMocks()
}

func (ts *ItemBatchPlayerSuite) SetupSubTest() {
	ts.initMocks()
}

func (ts *ItemBatchPlayerSuite) initMocks() {
	ts.ctrl = gomock.NewController(ts.T())
	ts.h = mock.NewMockhandler(ts.ctrl)
	ts.e = mock.NewMockeventEmitter(ts.ctrl)
}

func (ts *ItemBatchPlayerSuite) TestValidatesBatchSize() {
	type testCase struct {
		name      string
		batchSize int
		wantErr   error
	}

	tests := []testCase{
		{name: "zero", batchSize: 0, wantErr: helper.ErrInternalBatchSizeIsZeroOrNegative},
		{name: "negative", batchSize: -1, wantErr: helper.ErrInternalBatchSizeIsZeroOrNegative},
		{name: "greater than total limit", batchSize: 200000, wantErr: helper.ErrInternalBatchSizeIsGreaterThanTotalLimit},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			// обработчик и эмиттер не должны вызываться при невалидном batchSize
			player := helper.NewItemBatchPlayer(ts.h, ts.e)

			err := player.Execute(context.Background(), tt.batchSize)
			ts.Require().ErrorIs(err, tt.wantErr)
		})
	}
}

func (ts *ItemBatchPlayerSuite) TestNegativeCountIsRejected() {
	ts.h.EXPECT().Execute(gomock.Any(), 10).Return(-1, nil)

	player := helper.NewItemBatchPlayer(ts.h, ts.e) // событие не эмитится при ошибке

	err := player.Execute(context.Background(), 10)
	ts.Require().ErrorIs(err, helper.ErrInternalNegativeProcessedCount)
}

func (ts *ItemBatchPlayerSuite) TestHandlerError() {
	errHandler := errors.New("handler failed")

	type testCase struct {
		name    string
		respErr error
		wantErr error
	}

	tests := []testCase{
		{name: "generic error is propagated", respErr: errHandler, wantErr: errHandler},
		{name: "deadline becomes temporary problem", respErr: context.DeadlineExceeded, wantErr: mrprocess.ErrSystemTemporaryProblemHasOccurred},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			ts.h.EXPECT().Execute(gomock.Any(), 10).Return(0, tt.respErr)

			player := helper.NewItemBatchPlayer(ts.h, ts.e) // событие не эмитится при ошибке

			err := player.Execute(context.Background(), 10)
			ts.Require().ErrorIs(err, tt.wantErr)
		})
	}
}

func (ts *ItemBatchPlayerSuite) TestStopsOnPartialBatch() {
	// полный пакет, полный пакет, затем неполный (последняя пачка) - обработка завершается
	gomock.InOrder(
		ts.h.EXPECT().Execute(gomock.Any(), 10).Return(10, nil),
		ts.h.EXPECT().Execute(gomock.Any(), 10).Return(10, nil),
		ts.h.EXPECT().Execute(gomock.Any(), 10).Return(3, nil),
	)

	ts.e.EXPECT().Emit(gomock.Any(), "Execute", "total", 23, "duration_sec", gomock.Any(), "batch_size", 10)

	player := helper.NewItemBatchPlayer(ts.h, ts.e)

	err := player.Execute(context.Background(), 10)
	ts.Require().NoError(err)
}

func (ts *ItemBatchPlayerSuite) TestStopsOnEmptyResult() {
	ts.h.EXPECT().Execute(gomock.Any(), 10).Return(0, nil)
	ts.e.EXPECT().Emit(gomock.Any(), "Execute", "total", 0, "duration_sec", gomock.Any(), "batch_size", 10)

	player := helper.NewItemBatchPlayer(ts.h, ts.e)

	err := player.Execute(context.Background(), 10)
	ts.Require().NoError(err)
}

func (ts *ItemBatchPlayerSuite) TestStopsOnTotalLimit() {
	// totalLimit=25, batchSize=10: после 30 обработанных элементов лимит превышен
	gomock.InOrder(
		ts.h.EXPECT().Execute(gomock.Any(), 10).Return(10, nil),
		ts.h.EXPECT().Execute(gomock.Any(), 10).Return(10, nil),
		ts.h.EXPECT().Execute(gomock.Any(), 10).Return(10, nil),
	)

	ts.e.EXPECT().Emit(gomock.Any(), "Execute", "total", 30, "duration_sec", gomock.Any(), "batch_size", 10)

	player := helper.NewItemBatchPlayerWithTotalLimit(ts.h, ts.e, 25)

	err := player.Execute(context.Background(), 10)
	ts.Require().NoError(err)
}
