package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	listenAddr string
	engine     *gin.Engine
}

func NewServer(listenAddr string, engine *gin.Engine) *Server {
	return &Server{
		listenAddr: listenAddr,
		engine:     engine,
	}
}

func (s *Server) Run() error {
	return s.engine.Run(s.listenAddr)
}
