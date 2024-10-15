package interfaces

type Job[T any, R any] interface {
	GetID() int
	Execute() JobResult[R]
}
