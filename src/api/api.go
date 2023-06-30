package api

import (
	"net/http"

	"github.com/dominic-wassef/golander/src/database"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	db     *database.Database
}

func NewServer(db *database.Database) *Server {
	router := gin.Default()
	server := &Server{router: router, db: db}
	server.routes()
	return server
}

func (s *Server) Start(port string) {
	s.router.Run(port)
}

func (s *Server) routes() {
	s.router.GET("/ping", s.pingHandler)
}

func (s *Server) pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
