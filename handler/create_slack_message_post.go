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
	Error   bool   `json:"error,omitempty"`
	Channel string `json:"channel"`
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
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(configuration.GetSlackUrl() + "api/chat.postMessage"); err == nil {

		result := &CreateSlackHttpResponse{}
		if err := json.Unmarshal(resp.Body(), result); err == nil {
			if marshal, toHttpResponseErr := json.MarshalIndent(result, "", " "); toHttpResponseErr == nil {
				log.Println(string(marshal))
			}
			if !result.Ok {
				log.Printf("could not send message %s", payload)
			}
		} else {
			log.Println(err)
		}
	} else {
		log.Println("Skipping... " + err.Error())
	}
}
