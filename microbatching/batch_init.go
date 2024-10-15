package microbatching

import (
	inf "github.com/bimalkeeth/upguard/microbatching/interfaces"
	"sync"
	"time"
)

// BatchConfig represents the configuration for batching operations.
type BatchConfig struct {
	BatchSize            int
	BatchTimeOutDuration time.Duration
}

// MicroBatched represents a micro-batching system which processes jobs of type T
// and produces results of type R. It controls the batching through a jobChannel
// and returns results through a resultChannel.
type microBatched[T any, R any] struct {
	jobChannel     chan inf.Job[T, R]
	resultChannel  chan inf.JobResult[R]
	batchConfig    BatchConfig
	batchProcessor inf.BatchProcessor[T, R]
	shutDownChan   chan struct{}
	waitGroup      sync.WaitGroup
	batchTimer     *time.Timer
}

// NewMicroBatched initializes and starts a micro-batching
// system to process jobs with specified batch configuration.
func NewMicroBatched[T any, R any](
	batchConfig BatchConfig,
	batchProcessor inf.BatchProcessor[T, R]) inf.MicroBatcher[T, R] {
	microBatched := &microBatched[T, R]{
		jobChannel:     make(chan inf.Job[T, R]),
		resultChannel:  make(chan inf.JobResult[R]),
		batchConfig:    batchConfig,
		batchProcessor: batchProcessor,
		shutDownChan:   make(chan struct{}),
		waitGroup:      sync.WaitGroup{},
	}

	microBatched.assignTimer()

	microBatched.waitGroup.Add(1)

	microBatched.start()

	return microBatched
}
