package handler

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"log"
	"notepad-slack/configuration"
)

// CreateSlackHttpResponse http response
//
//	see [handler.CreateSlackMessageHandler] for example usage
type CreateSlackHttpResponse struct {
	Ok      bool   `json:"ok"`
	Error   string `json:"error,omitempty"`
	Channel string `json:"channel,omitempty"`
	Message map[string]interface{}
}

// CreateSlackMessageHandler post to Slack for message creation
//
//	payload:
//		is the message to be sent
//	returns:
//		nothing
func CreateSlackMessageHandler(payload string) {
	client := resty.New()
	if resp, err := client.R().
		SetHeader("Authorization", "Bearer "+configuration.GetSlackToken()).
		SetHeader("Content-Type", "application/json;charset=utf-8").
		SetBody(payload).
		Post(configuration.GetSlackUrl() + "api/chat.postMessage"); err == nil {

		resultObj := &CreateSlackHttpResponse{}
		log.Println(string(resp.Body()))

		if err := json.Unmarshal(resp.Body(), resultObj); err == nil {
			if marshal, toHttpResponseErr := json.MarshalIndent(resultObj, "", " "); toHttpResponseErr == nil {
				log.Println(string(marshal))
			}
			if !resultObj.Ok {
				log.Printf("message not sent: %s", payload)
			}
		} else {
			log.Println(err)
		}
	} else {
		log.Println("Skipping... " + err.Error())
	}
}
