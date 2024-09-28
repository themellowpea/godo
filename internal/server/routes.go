package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Route interface {
	RegisterRoutes(router *gin.Engine)
}

func SetupRoutes(router *gin.Engine, db *pgxpool.Pool) {
	routes := []Route{
		NewAlbumRoutes(db),
	}

	for _, r := range routes {
		r.RegisterRoutes(router)
	}
}
