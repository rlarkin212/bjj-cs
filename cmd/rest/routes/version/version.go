package version

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlarkin212/bjj-cs/configs"
)

type versionHandler struct {
	config *configs.Config
}

func NewVersionHandler(config *configs.Config) *versionHandler {
	return &versionHandler{
		config: config,
	}
}

func (v *versionHandler) Version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		v.config.Version: "",
	})
}
