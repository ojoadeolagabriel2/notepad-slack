package main

import (
	"notepad-slack/application"
	"notepad-slack/configuration"
	"notepad-slack/cron"
)

func main() {
	app := application.Create()
	app.StartConfiguration()

	go cron.SlackMessageCreatorCron("Testing 1,2,3..!", uint64(configuration.GetSlackLookupInterval()))(app.Configuration)
	app.StartAPI()
}
