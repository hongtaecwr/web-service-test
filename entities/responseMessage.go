package entities

type ResponseMessage struct {
	Status  int         `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}
