package runner

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/choraio/server/idx/config"
	"github.com/choraio/server/idx/process"
)

// Runner runs continuous background processes.
type Runner struct {
	ctx       context.Context
	cfg       config.Config
	waitGroup sync.WaitGroup
}

// NewRunner creates a new runner.
func NewRunner(ctx context.Context, cfg config.Config) Runner {
	return Runner{
		ctx: ctx,
		cfg: cfg,
	}
}

// Close shuts down the runner.
func (r *Runner) Close() {
	// wait for processes to finish
	r.waitGroup.Wait()

	fmt.Println("runner", "shutdown complete")
}

// RunProcess runs a process using the provided process function.
func (r *Runner) RunProcess(f process.Function, p process.Params) {
	// add process to wait group
	r.waitGroup.Add(1)

	go func() {
		// decrement wait group counter on exit
		defer r.waitGroup.Done()

		defer fmt.Println("runner", "stopping process", p.Name)

		// initialize runner retry count
		runnerRetryCount := uint64(0)

		fmt.Println("runner", "starting process", p.Name)

		for {
			// log retry count
			if runnerRetryCount > 0 {
				fmt.Println("runner", "retry count", p.Name, runnerRetryCount)
			}

			// exit on exceeding max retries
			if runnerRetryCount > r.cfg.IdxRunnerMaxRetries {
				fmt.Println("runner", "maximum retries", p.Name, r.cfg.IdxRunnerMaxRetries)
				return // exit process
			}

			// set process start time
			processStart := time.Now()

			// execute process function
			err := f(r.ctx, p)

			// set process duration
			processDuration := time.Since(processStart)

			fmt.Println("runner", "process duration", p.Name, processDuration)

			if err != nil {
				fmt.Println("runner", "process error", p.Name, err)

				// update retry count
				runnerRetryCount++
			}

			// wait for context done or backoff duration
			select {
			case <-r.ctx.Done():
				return // exit process
			case <-time.After(r.cfg.IdxRunnerBackoff):
				fmt.Println("runner", "backing off", p.Name, r.cfg.IdxRunnerBackoff)
				continue // continue process
			}
		}
	}()
}
