package structs

type JsonResponse struct {
	ApiStatus  int64       `json:"api_status"`
	ApiMessage string      `json:"api_message"`
	Data       interface{} `json:"data"`
}

type ReadInterface interface {
	Id() int64
}
