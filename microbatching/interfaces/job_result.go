package interfaces

// JobResult represents the result of a job execution.
type JobResult[R any] interface {
	// GetJobID returns the unique identifier of the executed job.
	GetJobID() int
	// GetJobName returns the name of the job.
	GetJobName() string
	// IsSuccess returns true if the job execution was successful; otherwise, it returns false.
	IsSuccess() bool
	// GetError returns an error if the job execution failed; otherwise, it returns nil.
	GetError() error
	// GetResult returns the result of the job execution.
	GetResult() R
}
