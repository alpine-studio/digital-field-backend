package server

import (
	"hackaton/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *server) Test(c *gin.Context) {
	c.JSON(http.StatusOK, "test")
}

// Vineyard

func (s *server) PostVineyard(c *gin.Context) {
	var vineyard models.Vineyard
	c.ShouldBind(&vineyard)

	s.store.VineyardRepo.Post(&vineyard)
	c.JSON(http.StatusOK, s.store.VineyardRepo.Vineyards)
}

func (s *server) GetVineyard(c *gin.Context) {
	c.JSON(http.StatusOK, s.store.VineyardRepo.Vineyards)
}
