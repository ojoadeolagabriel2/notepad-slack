package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"notepad-slack/consts"
)

// ReadinessHandler handler
func ReadinessHandler() func(context *gin.Context) {
	return func(context *gin.Context) {
		if result, err := json.Marshal(map[string]interface{}{
			"code": string(consts.SUCCESS),
			"systems": map[string]string{
				"db":        "ok",
				"messaging": "ok",
			},
		}); err == nil {
			context.String(200, string(result))
		}
	}
}
