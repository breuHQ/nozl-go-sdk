package main

import (
	"fmt"
	"testing"
)

func TestSendingTwilioMsg(t *testing.T) {
	pathParams := map[string]interface{}{
		"AccountSid": "AC9f560ea30baaaf8013e4e44284eb6769",
	}
	queryParams := map[string]interface{}{}
	requestBody := map[string]interface{}{
		"user_id": "sdfdsfdg",
		"To":      "+923244253153",
		"Body":    "hello world from nozl golang SDK",
		"From":    "+19034598701",
	}

	msg := &Message{
		ServiceID:   "a13b3453-3e20-44f1-9790-f866cf029d71",
		OperationID: "CreateMessage",
		PathParams:  pathParams,
		QueryParams: queryParams,
		RequestBody: requestBody,
	}

	nozl := Nozl{
		TenantAPIKey: "0e3be246-0f20-4809-9afe-f6daf40cc9ec",
		Url:          "http://localhost:1323/sdk/message",
	}

	res, err := nozl.SendMessage(msg)

	if err != nil {
		t.Errorf("Test failed! %s", err.Error())
	} else {
		fmt.Println(string(res))
	}
}
