package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/themellowpea/godo/internal/db/albums"
)

type AlbumHandler struct {
	db *pgxpool.Pool
}

func NewAlbumHandler(db *pgxpool.Pool) *AlbumHandler {
	return &AlbumHandler{db: db}
}

func (h *AlbumHandler) GetAllAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums.GetAlbums())
}
