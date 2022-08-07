package cron

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/jasonlvhit/gocron"
	"log"
	"notepad-slack/configuration"
	"notepad-slack/handler"
	"strings"
)

type PostAddUserToSlackChannel struct {
	Users   string `json:"users"`
	Channel string `json:"channel"`
}

type SlackError struct {
	User  string `json:"user,omitempty"`
	Ok    bool   `json:"ok,omitempty"`
	Error string `json:"error,omitempty"`
}

type PostAddUserToSlackChannelResponse struct {
	Ok      bool                   `json:"ok"`
	Error   string                 `json:"error,omitempty"`
	Channel map[string]interface{} `json:"channel,omitempty"`
	Errors  []SlackError           `json:"errors,omitempty"`
}

// SlackAddUserToChannel checks if user is in channel and includes if not present
//
// userIds:
//
//	provide list of userIds. e.g. UXXXXXXX
//
// Steps include:
//   - set array of userIds
//   - set channel
//
// As an example, this program calls SlackAddUserToChannel:
//
//	 func SampleCall(message string) {
//			SlackAddUserToChannel([]string{"UXXXXXX"}, "test-channel")(map[string]string{ "app.slack-url", "https://slack.com/"})
//	 }
func SlackAddUserToChannel(userIds []string, channel string) func(map[string]interface{}) {
	return func(config map[string]interface{}) {
		if userIds == nil || len(userIds) == configuration.ZERO {
			userIds = []string{config["app.slack-user-default"].(string)}
		}
		if marshal, postError := json.Marshal(&PostAddUserToSlackChannel{
			Users:   strings.Join(userIds[:], ","),
			Channel: channel,
		}); postError == nil {
			client := resty.New()
			if resp, err := client.R().
				SetHeader("Authorization", "Bearer "+configuration.GetSlackToken()).
				SetHeader("Content-Type", "application/json").
				SetBody(marshal).
				Post(configuration.GetSlackUrl() + "api/conversations.invite"); err == nil {

				result := &PostAddUserToSlackChannelResponse{}
				if err := json.Unmarshal(resp.Body(), result); err == nil {
					if result.Ok {
						log.Println("created user successfully")
					} else {
						log.Println("could not create user successfully: " + result.Error)
					}
				}
			} else {
				log.Println("error: " + err.Error())
			}
		} else {
			log.Println("error: " + postError.Error())
		}
	}
}

// SlackMessageCreatorCron cron message creator via Slack
//
// message:
//
//	text to send and interval to push
//
// interval:
//
//	how frequent to send, e.g. 5
//
// Steps include:
//   - set message and trigger interval
//   - pass app configuration data on calling returned func
//
// As an example, this program calls SlackMessageCreatorCron:
//
//	 func SampleCall(message string) {
//			SlackMessageCreatorCron(message, 5)(map[string]string{ "app.slack-url", "https://slack.com/"})
//	 }
func SlackMessageCreatorCron(message string, interval uint64) func(map[string]interface{}) {

	return func(config map[string]interface{}) {
		payload := fmt.Sprintf("{\"channel\":\"#test-channel\",\"text\":\"%s\"}", message)
		_ = gocron.Every(interval).Second().Do(func() {
			log.Println("triggering CreateSlackMessageHandler")
			handler.CreateSlackMessageHandler(payload)
			log.Println("completed triggering CreateSlackMessageHandler")
		})
		<-gocron.Start()
	}
}
