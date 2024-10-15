package interfaces

import "context"

// BatchProcessor is an interface that generalizes processing
// of worker jobs with inputs of type T and outputs of type R.
type BatchProcessor[T any, R any] interface {
	// ProcessBatch processes a slice of jobs of type T
	// And returns a slice of job results of type R.
	ProcessBatch(ctx context.Context, jobs []Job[T, R]) []JobResult[R]
}
