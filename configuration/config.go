package configuration

import (
	"os"
	"strconv"
	"time"
	"unicode/utf8"
)

var GlobalConfig *Configuration

const (
	ZERO = 0
)

type Configuration struct {
	SourceType string
	Data       map[string]interface{}
}

func Start() {

}

// Initialize starts configuration read
func Initialize() *Configuration {
	configuration := &Configuration{
		SourceType: "env",
		Data:       map[string]interface{}{},
	}
	GlobalConfig = configuration

	configuration.Data["app.id"] = "notepad-slack"
	configuration.Data["app.port"] = 12345
	configuration.Data["app.default_timer_sec"] = 1
	configuration.Data["app.slack-lookup-interval"] = toInteger(os.Getenv("ENV_SLACK_LOOKUP_INTERVAL"), 3)
	configuration.Data["app.slack-url"] = toString(os.Getenv("ENV_SLACK_URL"), "https://slack.com/")
	configuration.Data["app.slack-token"] = toString(os.Getenv("ENV_SLACK_TOKEN"), "xoxb-2453596686656-2453645651088-Up5TjqtU3YsX9tUJDsL2eW62")
	configuration.Data["app.slack-user-default"] = toString(os.Getenv("ENV_SLACK_USER_DEFAULT"), "U03SMH9NW4T")
	configuration.Data["app.startup-time"] = time.Now()
	configuration.Data["database.host"] = toString(os.Getenv("ENV_DB_HOSTNAME"), "localhost")
	configuration.Data["database.name"] = toString(os.Getenv("ENV_DB_NAME"), "postgres")
	configuration.Data["database.port"] = toInteger(os.Getenv("ENV_DB_PORT"), 5432)
	configuration.Data["database.username"] = toString(os.Getenv("ENV_DB_USERNAME"), "postgres")
	configuration.Data["database.password"] = toString(os.Getenv("ENV_DB_PASSWORD"), "postgres")
	return configuration
}

// GetSlackLookupInterval get slack lookup interval
func GetSlackLookupInterval() int {
	return GlobalConfig.Data["app.slack-lookup-interval"].(int)
}

func GetAppDefaultTimer() int {
	return GlobalConfig.Data["app.default_timer_sec"].(int)
}

// GetSlackUrl get slack url
func GetSlackUrl() string {
	return GlobalConfig.Data["app.slack-url"].(string)
}

// GetSlackToken oauth token
func GetSlackToken() string {
	return GlobalConfig.Data["app.slack-token"].(string)
}

// toString safe string translation
func toString(data string, defaultValue string) string {
	if utf8.RuneCountInString(data) == ZERO {
		return defaultValue
	}
	return data
}

// toInteger safe int translation
func toInteger(data string, defaultValue int) int {
	if utf8.RuneCountInString(data) == ZERO {
		return defaultValue
	}
	if i, err := strconv.Atoi(data); err == nil {
		return i
	} else {
		return defaultValue
	}
}

// toBoolean safe bool translation
func toBoolean(data string, defaultValue bool) bool {
	if utf8.RuneCountInString(data) == ZERO {
		return defaultValue
	} else {
		if result, err := strconv.ParseBool(data); err == nil {
			return result
		}
		return defaultValue
	}
}
