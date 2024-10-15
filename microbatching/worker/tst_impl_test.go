package worker

import inf "github.com/bimalkeeth/upguard/microbatching/interfaces"

type StrJob[T any, R any] struct {
	Id   int
	Data string
}

func (s StrJob[T, R]) GetID() int {
	return s.Id
}

func (s StrJob[T, R]) Execute() inf.JobResult[string] {
	return StrJobResult[string]{
		JobId:   1,
		JobName: "StrJob",
		Success: true,
		Error:   nil,
		Result:  s.Data,
	}
}

type StrJobResult[R any] struct {
	JobId   int
	JobName string
	Success bool
	Error   error
	Result  R
}

func (rs StrJobResult[R]) GetJobID() int {
	return rs.JobId
}

func (rs StrJobResult[R]) GetJobName() string {
	return rs.JobName
}

func (rs StrJobResult[R]) IsSuccess() bool {
	return rs.Success
}

func (rs StrJobResult[R]) GetError() error {
	return rs.Error
}

func (rs StrJobResult[R]) GetResult() R {
	return rs.Result
}
