/*
 * @Author: yujiajie
 * @Date: 2024-03-18 16:25:50
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-18 17:05:24
 * @FilePath: /gateway/core/executors/periodicalexecutor.go
 * @Description:
 */
package executors

import (
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

const idleRound = 10

type (
	TaskContainer interface {
		AddTask(task any) bool
		Execute(tasks any)
		RemoveAll() any
	}

	PeriodicalExecutor struct {
		commander   chan any
		interval    time.Duration
		container   TaskContainer
		wg          sync.WaitGroup
		wgBarrier   sync.Mutex
		confirmChan chan struct{}
		inflight    int32
		guarded     bool
		newTicker   func(duration time.Duration) time.Ticker
		mu          sync.Mutex
	}
)

func NewPeriodicalExecutor(interval time.Duration, container TaskContainer) *PeriodicalExecutor {
	executor := &PeriodicalExecutor{
		commander:   make(chan any, 1),
		interval:    interval,
		container:   container,
		confirmChan: make(chan struct{}),
		newTicker: func(d time.Duration) time.Ticker {
			return *time.NewTicker(d)
		},
	}

	return executor
}

func (pe *PeriodicalExecutor) Add(task any) {
	if vals, ok := pe.addAndCheck(task); ok {
		pe.commander <- vals
		<-pe.confirmChan
	}
}

func (pe *PeriodicalExecutor) Flush() bool {
	pe.enterExecution()
	return pe.executeTasks(func() any {
		pe.mu.Lock()
		defer pe.mu.Unlock()
		return pe.container.RemoveAll()
	}())
}

func (pe *PeriodicalExecutor) Sync(fn func()) {
	pe.mu.Lock()
	defer pe.mu.Unlock()
	fn()
}

func (pe *PeriodicalExecutor) Wait() {
	pe.Flush()
	pe.wgBarrier.Lock()
	pe.wgBarrier.Unlock()
	pe.wg.Wait()
}

func (pe *PeriodicalExecutor) addAndCheck(task any) (any, bool) {
	pe.mu.Lock()
	defer func() {
		if !pe.guarded {
			pe.guarded = true
		}
		pe.mu.Unlock()
	}()

	if pe.container.AddTask(task) {
		atomic.AddInt32(&pe.inflight, 1)
		return pe.container.RemoveAll(), true
	}

	return nil, false
}

func (pe *PeriodicalExecutor) backgroundFlush() {
	go func() {
		defer pe.Flush()

		ticker := pe.newTicker(pe.interval)
		defer ticker.Stop()

		var commander bool
		last := time.Now()
		for {
			select {
			case vals := <-pe.commander:
				commander = true
				atomic.AddInt32(&pe.inflight, -1)
				pe.enterExecution()
				pe.confirmChan <- struct{}{}
				pe.executeTasks(vals)
				last = time.Now()
			case <-ticker.C:
				if commander {
					commander = false
				} else if pe.Flush() {
					last = time.Now()
				} else if pe.shallQuit(last) {
					return
				}
			}
		}
	}()
}

func (pe *PeriodicalExecutor) doneExecution() {
	pe.wg.Done()
}

func (pe *PeriodicalExecutor) enterExecution() {
	pe.wgBarrier.Lock()
	defer pe.wgBarrier.Unlock()
	pe.wg.Add(1)
}

func (pe *PeriodicalExecutor) executeTasks(tasks any) bool {
	defer pe.doneExecution()

	ok := pe.hasTasks(tasks)
	if ok {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		pe.container.Execute(tasks)
	}

	return ok
}

func (pe *PeriodicalExecutor) hasTasks(tasks any) bool {
	if tasks == nil {
		return false
	}

	val := reflect.ValueOf(tasks)
	switch val.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return val.Len() > 0
	default:
		return true
	}
}

func (pe *PeriodicalExecutor) shallQuit(last time.Time) (stop bool) {
	if time.Since(last) <= pe.interval*idleRound {
		return
	}

	pe.mu.Lock()
	if atomic.LoadInt32(&pe.inflight) == 0 {
		pe.guarded = false
		stop = true
	}
	pe.mu.Unlock()
	return
}
