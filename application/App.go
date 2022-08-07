package application

import (
	"database/sql"
	"notepad-slack/configuration"
	"notepad-slack/handler"
)
import "github.com/gin-gonic/gin"

type App struct {
	ApplicationId string
	Configuration map[string]interface{}
	WebEngine     *gin.Engine
	Database      *sql.DB
}

// Create initiates base application context
//
//	app:
//		applies to [application.App] struct
//
// Usage:
//
//	func sample() {
//		app := &App{}
//	}
func Create() *App {
	return &App{
		ApplicationId: "notepad-slack",
	}
}

// StartConfiguration initiates application configuration
//
//	app:
//		applies to [application.App] struct
//
// Usage:
//
//	func sample() {
//		app := &App{}
//		app.StartConfiguration()
//	}
func (app *App) StartConfiguration() {
	app.Configuration = configuration.Initialize().Data
}

// StartAPI initiates application http server
//
//	app:
//		applies to [application.App] struct
//
// Usage:
//
//	func sample() {
//		app := &App{}
//		app.StartAPI()
//	}
func (app *App) StartAPI() {
	app.WebEngine = handler.InitializeAPI()
}
