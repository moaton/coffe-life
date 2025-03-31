package v1

type Response struct {
	Meta        interface{} `json:"meta" swaggertype:"object"`
	ErrorCode   string      `json:"errorCode"`
	Description string      `json:"description"`
	Payload     interface{} `json:"payload" swaggertype:"object"`
}

type Success struct {
	Message string `json:"message"`
}
