package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/themellowpea/godo/internal/handlers"
)

type AlbumRoutes struct {
	db *pgxpool.Pool
}

func NewAlbumRoutes(db *pgxpool.Pool) *AlbumRoutes {
	return &AlbumRoutes{db: db}
}

func (a *AlbumRoutes) RegisterRoutes(router *gin.Engine) {
	albumHandler := handlers.NewAlbumHandler(a.db)

	albumRoutes := router.Group("/albums")
	{
		albumRoutes.GET("/", albumHandler.GetAllAlbums)
	}
}
