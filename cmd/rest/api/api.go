package api

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/rlarkin212/bjj-cs/configs"
)

const defaultPort = "5000"

func Generate() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.ForceConsoleColor()

	return router
}

func Start(router *gin.Engine, config *configs.Config) {
	port := config.Rest.Port
	if port == "" {
		port = defaultPort
	}

	errs := make(chan error, 2)

	go func() {
		fmt.Printf("listening in port: %s", port)
		errs <- router.Run(":" + port)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)

		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)
}
