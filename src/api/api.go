package api

import (
	"net/http"
	"strconv"

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
	s.router.GET("/repositories", s.getRepositoriesHandler)
	s.router.GET("/repository/:id", s.getRepositoryByIDHandler)
	s.router.DELETE("/repository/:id", s.deleteRepositoryHandler)
}

func (s *Server) pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// Handler for getting all repositories
func (s *Server) getRepositoriesHandler(c *gin.Context) {
	repositories, err := s.db.GetAllRepos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, repositories)
}

// Handler for getting a repository by ID
func (s *Server) getRepositoryByIDHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
		return
	}
	repo, err := s.db.GetRepoByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, repo)
}

// Handler for deleting a repository
func (s *Server) deleteRepositoryHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
		return
	}
	err = s.db.DeleteRepoByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
