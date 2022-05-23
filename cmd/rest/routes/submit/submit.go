package submit

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlarkin212/bjj-cs/configs"
	"github.com/rlarkin212/bjj-cs/internal/entities/instructionals"
	"github.com/rlarkin212/bjj-cs/internal/service/submit"
)

type submitHandler struct {
	servive submit.SubmitService
}

func NewSubmitHandler(config *configs.Config) *submitHandler {
	return &submitHandler{
		servive: *submit.NewSubmitService(config),
	}
}

func (h *submitHandler) Submit(c *gin.Context) {
	var input *instructionals.NewInstructional
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	instructional, err := h.servive.Submit(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"instructional": instructional,
	})
}
