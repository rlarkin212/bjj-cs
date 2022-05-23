package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rlarkin212/bjj-cs/cmd/rest/routes/health"
	"github.com/rlarkin212/bjj-cs/cmd/rest/routes/search"
	"github.com/rlarkin212/bjj-cs/cmd/rest/routes/submit"
	"github.com/rlarkin212/bjj-cs/cmd/rest/routes/version"
	"github.com/rlarkin212/bjj-cs/configs"
)

func RegisterRoutes(router *gin.Engine, config *configs.Config) {
	v1 := router.Group("v1")
	register(v1, config)
}

func register(rg *gin.RouterGroup, config *configs.Config) {
	search := search.NewSearchHandler(config)
	submit := submit.NewSubmitHandler(config)

	health := health.NewHealthHandler()
	version := version.NewVersionHandler(config)

	//*search endpoints
	rg.GET("/instructionals", search.Instructionals)
	rg.GET("/instructional/:id", search.Instructional)
	rg.GET("/count", search.Count)

	//*submit endpoints
	rg.POST("/submit", submit.Submit)

	//*info endpoints
	rg.GET("/health", health.Health)
	rg.GET("/version", version.Version)
}
