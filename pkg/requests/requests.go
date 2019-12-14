package requests

import (
	"encoding/json"
	"net/http"
)

// Request ...
type Request struct {
	URL     string
	Headers map[string]string
}

var httpClient = &http.Client{}

// Get make a http get call
func (req Request) Get(result interface{}) error {
	return req.Call("GET", result)
}

// Post make a http get call
func (req Request) Post(result interface{}) error {
	return req.Call("POST", result)
}

// Call make a http call
func (req Request) Call(callType string, result interface{}) error {
	httpReq, err := http.NewRequest(callType, req.URL, nil)
	if req.Headers == nil {
		req.Headers = map[string]string{}
	}
	for headerKey, headerValue := range req.Headers {
		httpReq.Header.Add(headerKey, headerValue)
	}
	if err != nil {
		return err
	}
	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	if result == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(result)
}
