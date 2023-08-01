package ngs

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	TenantAPIKey string
	Url          string
}

func (c *Client) SendMessage(msg *Message) ([]byte, error) {
	payload := msg.ToJson()

	if payload != nil {
		client := &http.Client{}
		req, err := http.NewRequest("POST", c.Url, bytes.NewReader(payload))

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.TenantAPIKey))

		resp, err := client.Do(req)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		return body, nil
		
	} else {
		fmt.Println("error: payload is empty")
		return nil, errors.New("error: payload is empty")
	}
}