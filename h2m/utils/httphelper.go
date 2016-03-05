package utils

import (
	"bytes"
	"net/http"
)

func SendRequest(url, payload string) (*http.Response, error) {
	body := &bytes.Buffer{}
	body.Write([]byte(payload))
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	return client.Do(req)
}
