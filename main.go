package main

import (
	"notepad-slack/application"
	"notepad-slack/cron"
)

const SlackInterval = 2

func main() {
	app := application.Create()
	app.StartConfiguration()

	go cron.SlackMessageCreatorCron("Testing 1,2,3..!", SlackInterval)(app.Configuration)
	app.StartAPI()
}
