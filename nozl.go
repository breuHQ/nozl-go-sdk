package nozlgosdk

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Nozl struct {
	TenantAPIKey string
	Url          string
}

func (n *Nozl) SendMessage(msg *Message) ([]byte, error) {
	payload := msg.ToJson()

	if payload != nil {
		client := &http.Client{}
		req, err := http.NewRequest("POST", n.Url, bytes.NewReader(payload))

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", n.TenantAPIKey))

		resp, err := client.Do(req)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		return body, nil
		
	} else {
		fmt.Println("Error: payload is empty")
		return nil, errors.New("Error: payload is empty")
	}
}