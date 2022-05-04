package asyncjob

import (
	"context"
	"time"
)

type Job interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	State() JobState
}

type jobHandler func(ctx context.Context) error

type JobState int

const (
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeout
	StateCompleted
	StateRetryFailed
)

const (
	defaultMaxTimeout = 10 * time.Second
)

var (
	defaultRetryTimes = []time.Duration{time.Second, 5 * time.Second, 10 * time.Second}
)

func (js JobState) String() string {
	return []string{"Init", "Running", "Failed", "Timeout", "Completed", "RetryFailed"}[js]
}

type jobConfig struct {
	maxTimeout time.Duration
	retries    []time.Duration
}

type job struct {
	config     jobConfig
	handler    jobHandler
	state      JobState
	retryIndex int
	stopChan   chan bool
}

func NewJob(handler jobHandler) *job {
	return &job{
		config: jobConfig{
			maxTimeout: defaultMaxTimeout,
			retries:    defaultRetryTimes,
		},
		handler:    handler,
		state:      StateInit,
		retryIndex: -1,
		stopChan:   make(chan bool),
	}
}

func (j *job) Execute(ctx context.Context) error {
	j.state = StateRunning

	err := j.handler(ctx)
	if err != nil {
		j.state = StateFailed
		return err
	}

	j.state = StateCompleted

	return nil
}

func (j *job) Retry(ctx context.Context) error {
	j.retryIndex += 1
	time.Sleep(j.config.retries[j.retryIndex])

	err := j.Execute(ctx)
	if err == nil {
		j.state = StateCompleted
		return nil
	}

	if j.retryIndex == len(j.config.retries)-1 {
		j.state = StateRetryFailed
		return err
	}

	j.state = StateFailed
	return err
}

func (j *job) State() JobState {
	return j.state
}
