# nozl-go-sdk
Golang SDK for Nozl

This reposistory contains the official code of Nozl's Golang SDK.

## Componets
Nozl Go SDK consists of 2 components present in `ngs` package
1. `ngs.Message`
2. `ngs.Client`


## Installation
Go SDK of Nozl can be downloaded directly from the Github repository.

```bash
go get -u github.com/breuHQ/nozl-go-sdk@v0.0.3 # replace v0.0.3 with relevant version 
```
## Quick Start
Below is a sample code for sending SMS via Twilio. This assumes that Twilio service has
already been registered in Nozl.

```golang
package main

import (
	"fmt"

	"github.com/breuHQ/nozl-go-sdk"
)

func main() {
	pathParams := map[string]interface{}{
		"AccountSid": "<TWILIO-ACCCOUNT-SID>",
	}
	queryParams := map[string]interface{}{}
	requestBody := map[string]interface{}{
		"user_id": "sdfdsfdg", // user_id is a custom filter for tracking messages
		"To":      "<RECEIVER-PHONE-NUMBER>",
		"Body":    "hello world from nozl golang SDK",
		"From":    "<SENDER-PHONE-NUMBER>",
	}

	msg := &ngs.Message{
		ServiceID:   "a13b3453-3e20-44f1-9790-f866cf029d71", // get this from Nozl dashboard
		OperationID: "CreateMessage", // OpenAPI schema operation ID
		PathParams:  pathParams,
		QueryParams: queryParams,
		RequestBody: requestBody,
	}

	nozl := &ngs.Client{
		TenantAPIKey: "0e3be246-0f20-4809-9afe-f6daf40cc9ec", // get this from Nozl dashboard
		Url:          "http://dx.nozl.dev/sdk/message",
	}

	res, err := nozl.SendMessage(msg)

	if err != nil {
		fmt.Errorf("Test failed! %s", err.Error())
	} else {
		fmt.Println(string(res))
	}
}
```
