package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/themellowpea/godo/internal/db/albums"
)

type AlbumHandler struct{}

func NewAlbumHandler() *AlbumHandler {
	return &AlbumHandler{}
}

func (h *AlbumHandler) GetAllAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums.GetAlbums())
}
