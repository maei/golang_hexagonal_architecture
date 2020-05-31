package service

import (
	"github.com/teris-io/shortid"
	"log"
)

type SumServiceInterface interface {
	Compute(req *SumRequest) error
	FindResult(code string) (*SumResponse, error)
}

type sumService struct {
	sumRepo SumRepositoryInterface
}

func NewSumService(repo SumRepositoryInterface) SumServiceInterface {
	return &sumService{
		sumRepo: repo,
	}
}

func (s *sumService) Compute(req *SumRequest) error {
	a := req.A
	b := req.B
	sum := a + b
	code := shortid.MustGenerate()
	sendRes := &SumResult{
		Result: sum,
		A:      a,
		B:      b,
		Code:   code,
	}
	saveErr := s.sumRepo.Store(sendRes)
	if saveErr != nil {
		log.Println("error while writing to db")
		return saveErr
	}
	return nil
}

func (s *sumService) FindResult(code string) (*SumResponse, error) {
	res, err := s.sumRepo.Find(code)
	if err != nil {
		return nil, err
	}
	resultRes := &SumResponse{
		Result: res.Result,
		Code:   res.Code,
	}
	return resultRes, nil
}
