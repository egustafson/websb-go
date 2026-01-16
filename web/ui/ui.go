package ui

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/egustafson/websb-go/pkg/config"
)

func Init(ctx context.Context, srvCfg *config.ServerConfig, router *gin.RouterGroup) {

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	// TODO: more initialization
}
