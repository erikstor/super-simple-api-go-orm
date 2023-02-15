package types

import (
	"container/list"
	"encoding/json"
)

type ResponseMessageHttp struct {
	Status int `json:"status"`
	//Data    interface{} `json:"data"`
	Message string     `json:"message"`
	Error   *list.List `json:"error"`
}

func GetMessageError(status int, message string, error *list.List) []byte {
	resp := ResponseMessageHttp{Status: status, Message: message, Error: error}
	output, _ := json.Marshal(resp)

	return output
}
