package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestLivelinessHandler(t *testing.T) {
	r := SetUpRouter()
	r.GET("/liveliness", LivelinessHandler())
	mockResponse := "{\"code\":\"90000\"}"

	req, _ := http.NewRequest("GET", "/liveliness", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData := fmt.Sprint(w.Body)
	assert.Equal(t, mockResponse, responseData)
	assert.Equal(t, http.StatusOK, w.Code)
}
