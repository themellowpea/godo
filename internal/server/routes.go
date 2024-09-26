package server

import (
	"github.com/gin-gonic/gin"
	"github.com/themellowpea/godo/internal/handlers"
)

type Router struct{}

func NewRouter() *gin.Engine {
	router := gin.Default()
	albumHandler := handlers.NewAlbumHandler()

	router.GET("/albums", albumHandler.GetAllAlbums)

	return router
}
