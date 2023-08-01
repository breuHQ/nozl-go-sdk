package nozlgosdk

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	ServiceID   string                 `json:"service_id"`
	OperationID string                 `json:"operation_id"`
	PathParams  map[string]interface{} `json:"path_params"`
	QueryParams map[string]interface{} `json:"query_params"`
	RequestBody map[string]interface{} `json:"req_body"`
}

func (m *Message) ToJson() []byte {
	payload, err := json.Marshal(&m)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return payload
}