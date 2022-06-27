package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetLinksHandler(t *testing.T) {
	r := SetUpRouter()
	r.GET("/links", GetLinks)
	req, _ := http.NewRequest("GET", "/links", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var links []Link
	json.Unmarshal(w.Body.Bytes(), &links)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, links)
}
