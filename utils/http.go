package utils

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

type HttpError struct {
	Code    string
	Message string
}

func (error *HttpError) Error() string {
	return "boom"
}

// HttpPost http post call
func HttpPost(url string, headers map[string]string, body interface{}) (string, *HttpError) {
	if marshalledBody, err := json.Marshal(&body); err == nil {
		client := resty.New()
		for key, item := range headers {
			client.SetHeader(key, item)
		}
		if resp, err := client.R().
			SetBody(marshalledBody).
			Post(url); err == nil {
			return string(resp.Body()), nil
		} else {
			return "", &HttpError{Code: "500", Message: err.Error()}
		}
	}
	return "", &HttpError{Code: "500", Message: "error marshalling http-post request body"}
}

// HttpGet http get call
func HttpGet(url string, headers map[string]string) (string, *HttpError) {
	client := resty.New()
	for key, item := range headers {
		client.SetHeader(key, item)
	}
	if resp, err := client.R().
		Get(url); err == nil {
		return string(resp.Body()), nil
	} else {
		return "", &HttpError{Code: "500", Message: "error processing http-get request"}
	}
}
