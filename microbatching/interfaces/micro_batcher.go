package interfaces

type MicroBatcher[T any, R any] interface {
	Submit(job Job[T, R]) error
	Shutdown()
	ReadResult() <-chan JobResult[R]
}
