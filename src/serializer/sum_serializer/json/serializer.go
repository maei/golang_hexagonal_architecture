package json

import (
	"encoding/json"
	"github.com/maei/golang_hexagonal_architecture/src/service/sum_service"
	"log"
)

type Serial struct{}

func (s *Serial) Decode(input []byte) (*sum_service.SumRequest, error) {
	sumRequest := &sum_service.SumRequest{}
	err := json.Unmarshal(input, sumRequest)
	if err != nil {
		log.Println("error while decode byte")
		return nil, err
	}
	return sumRequest, nil
}

func (s *Serial) Encode(input *sum_service.SumResponse) ([]byte, error) {
	rawResponse, err := json.Marshal(input)
	if err != nil {
		log.Println("error while encode sum_service-response struct")
		return nil, err
	}
	return rawResponse, nil
}
