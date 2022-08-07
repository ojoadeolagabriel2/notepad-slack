package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"notepad-slack/consts"
)

// LivelinessHandler handler
func LivelinessHandler() func(context *gin.Context) {
	return func(context *gin.Context) {
		if result, err := json.Marshal(map[string]string{"code": string(consts.SUCCESS)}); err == nil {
			context.String(200, string(result))
		}
	}
}
