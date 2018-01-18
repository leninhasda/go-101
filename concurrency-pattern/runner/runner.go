package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner runs given list of task within given timeline
// can be shut down
type Runner struct {
	// interrupt channel reports a singnal from the os
	interrupt chan os.Signal

	// complete channel reports that processing is done
	complete chan error

	// timeout reports that time has run out
	timeout <-chan time.Time

	tasks []func(int)
}

// ErrTimeout is returnen when a value is received on timeout
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt is returnen when a value is received on interrupt
var ErrInterrupt = errors.New("received interrupt")

// New returns a new runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add attaches tasks to the runner
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start ...
func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
