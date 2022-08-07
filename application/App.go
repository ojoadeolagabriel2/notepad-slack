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

// Create starts up application context
func Create() *App {
	app := &App{
		ApplicationId: "notepad-slack",
	}
	return app
}

func (app *App) StartConfiguration() {
	app.Configuration = configuration.Initialize().Data
}

func (app *App) StartAPI() {
	app.WebEngine = handler.InitializeAPI()
}

func (app *App) StartPostgres() {

}
