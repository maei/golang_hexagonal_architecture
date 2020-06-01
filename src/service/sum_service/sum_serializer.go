package sum_service

type SumSerializerInterface interface {
	Decode(input []byte) (*SumRequest, error)
	Encode(request *SumResponse) ([]byte, error)
}
