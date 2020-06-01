package sum_service

type SumRequest struct {
	A int64 `json:"a"`
	B int64 `json:"b"`
}

type SumResponse struct {
	Result int64  `json:"result"`
	Code   string `json:"code"`
}

type SumResult struct {
	Result int64  `bson:"result"`
	A      int64  `bson:"a"`
	B      int64  `bson:"b"`
	Code   string `bson:"code"`
}
