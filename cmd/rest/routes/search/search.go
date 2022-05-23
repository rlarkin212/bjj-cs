package search

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlarkin212/bjj-cs/configs"
	"github.com/rlarkin212/bjj-cs/internal/service/search"
)

type searchHandler struct {
	service search.SearchService
}

func NewSearchHandler(config *configs.Config) *searchHandler {
	return &searchHandler{
		service: *search.NewSearchService(config),
	}
}

func (h *searchHandler) Instructionals(c *gin.Context) {
	instructionsls, err := h.service.Instructionals()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"instructionals": instructionsls,
	})
}

func (h *searchHandler) Instructional(c *gin.Context) {
	id := c.Param("id")

	instructional, err := h.service.Instructional(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"instructional": instructional,
	})
}

func (h *searchHandler) Count(c *gin.Context) {
	count, err := h.service.Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"count": count,
	})
}
