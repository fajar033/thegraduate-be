package model

type ResponseModel struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type ResponseModelFailed struct {
	Message any    `json:"message"`
	Status  string `json:"status"`
}
