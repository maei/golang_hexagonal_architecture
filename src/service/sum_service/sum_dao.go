package sum_service

import "context"

type SumRepositoryInterface interface {
	Store(req *SumResult) error
	Find(code string) (*SumResult, error)
	Disconnect(context.Context)
}
