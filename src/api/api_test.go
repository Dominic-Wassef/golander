package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dominic-wassef/golander/src/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupMockServer() *Server {
	// Here we are mocking the database, you can replace this with an actual mock if you wish
	mockDB := &database.Database{}

	// Setup routes
	server := NewServer(mockDB)

	return server
}

func TestPingHandler(t *testing.T) {
	server := setupMockServer()
	recorder := httptest.NewRecorder()

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/ping", nil)

	// Create a router and attach the pingHandler
	router := gin.Default()
	router.GET("/ping", server.pingHandler)

	// Serve the request
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", recorder.Body.String())
}
