package mrrun

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"time"
)

const (
	// keyProcessID - название ключа ID процесса, добавляемого в контекст.
	keyProcessID = "process_id"
)

type (
	// processExecuter - обёртка для запуска и остановки процесса с поддержкой синхронизации.
	// Содержит функции Execute (запуск) и Interrupt (остановка), а также
	// канал Synchronizer для уведомления о готовности процесса.
	//
	// Используется внутри AppRunner для управления жизненным циклом процессов.
	processExecuter struct {
		// Execute - функция запуска процесса (блокируется до завершения).
		Execute func() error

		// Interrupt - функция принудительной остановки процесса.
		Interrupt func(error)

		// Synchronizer - канал синхронизации для уведомления о готовности процесса.
		Synchronizer ProcessSync // OPTIONAL
	}
)

// ErrSystemWaitingTimeForPreviousProcessHasExpired - период ожидания запуска предыдущего процесса истёк.
var ErrSystemWaitingTimeForPreviousProcessHasExpired = errors.New("the waiting time for the previous process has expired")

func (r *AppRunner) contextWithProcessID(ctx context.Context, process Process) context.Context {
	ctx = r.traceManager.WithGeneratedProcessID(ctx, keyProcessID)
	r.logger.Info(ctx, "Start new process", "process_name", process.Caption())

	return ctx
}

// makeExecuter - создаёт processExecuter для процесса без синхронизации.
// Запуск этого процесса не зависит от других процессов.
func (r *AppRunner) makeExecuter(ctx context.Context, process Process) processExecuter {
	ctx = r.contextWithProcessID(ctx, process)

	return processExecuter{
		Execute: func() error {
			return process.Start(ctx, func() {})
		},
		Interrupt: func(_ error) {
			if err := process.Shutdown(ctx); err != nil {
				r.logger.Error(ctx, "AppRunner.makeExecuter", "error", err)
			}
		},
	}
}

// makeNextExecuter - создаёт processExecuter для процесса с поддержкой синхронизации.
// Если prev задан (непустой ProcessSync), запуск процесса ожидает сигнала
// готовности от предыдущего процесса; при пустом prev процесс стартует сразу.
//
// Логика работы:
//  1. Ожидает сигнала от prev.ready в пределах prev.readyTimeout;
//  2. При истечении таймаута возвращает ошибку ErrSystemWaitingTimeForPreviousProcessHasExpired;
//  3. При отмене контекста завершает процесс без ошибки (игнорирует ошибку контекста);
//  4. После готовности вызывает process.Start() с сигналом ready;
//  5. При Interrupt ожидает готовности процесса, затем вызывает Shutdown;
func (r *AppRunner) makeNextExecuter(ctx context.Context, process Process, prev ProcessSync) processExecuter {
	ctx = r.contextWithProcessID(ctx, process)

	isStartCalled := atomic.Bool{}
	chCurrentReady := make(chan struct{})

	return processExecuter{
		Execute: func() error {
			if prev.ready != nil {
				select {
				case <-time.NewTimer(prev.readyTimeout).C:
					close(chCurrentReady)

					return fmt.Errorf(
						"%w [process=%s, previous=%s]",
						ErrSystemWaitingTimeForPreviousProcessHasExpired, process.Caption(), prev.Caption,
					)
				case <-prev.ready:
				}
			}

			if err := ctx.Err(); err != nil {
				close(chCurrentReady)

				// ignore errors from other processes (context canceled)
				return nil //nolint:nilerr
			}

			isStartCalled.Store(true)

			return process.Start(
				ctx,
				func() {
					close(chCurrentReady)
				},
			)
		},
		Interrupt: func(_ error) {
			// ожидается готовность функции Execute
			select {
			case <-time.NewTimer(process.ReadyTimeout() + prev.readyTimeout).C:
				r.logger.Error(
					ctx,
					"the waiting time to interrupt the process has expired",
					"process", process.Caption(),
				)
			case <-chCurrentReady:
			}

			if !isStartCalled.Load() {
				return
			}

			if err := process.Shutdown(ctx); err != nil {
				r.logger.Error(ctx, "AppRunner.makeNextExecuter", "error", err)
			}
		},
		Synchronizer: ProcessSync{
			Caption:      process.Caption(),
			readyTimeout: process.ReadyTimeout(),
			ready:        chCurrentReady,
		},
	}
}
