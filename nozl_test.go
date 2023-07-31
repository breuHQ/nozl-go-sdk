package main

import (
	"fmt"
	"testing"
)

func TestSendingSGMsg(t *testing.T) {

	pathParams := map[string]interface{}{}
	queryParams := map[string]interface{}{}
	requestBody := map[string]interface{}{
		"user_id": "sdfdsfdg",
		"personalizations": []map[string]interface{}{
			{
				"to": []map[string]string{
					{"email": "demon6656@gmail.com"},
				},
			},
		},
		"from":    map[string]string{"email": "taimoor@breu.io"},
		"subject": "Sending with SendGrid is Fun recursion",
		"content": []map[string]string{
			{
				"type":  "text/plain",
				"value": "test email from golang nozl SDK",
			},
		},
	}

	msg := &Message{
		ServiceID:   "de9cfe4c-2179-49fb-82c3-de7ff400c167",
		OperationID: "POST_mail-send",
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

func TestSendingTwilioMsg(t *testing.T) {

	pathParams := map[string]interface{}{
		"AccountSid": "AC9f560ea30baaaf8013e4e44284eb6769",
	}
	queryParams := map[string]interface{}{}
	requestBody := map[string]interface{}{
		"user_id": "sdfdsfdg",
		"To": "+923244253153",
        "Body": "hello world from nozl golang SDK",
        "From": "+19034598701",
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