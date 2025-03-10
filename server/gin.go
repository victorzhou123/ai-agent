package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/victorzhou123/ai-agent/agent"
	"github.com/victorzhou123/ai-agent/config"
)

const BasePath = "/api"

func StartWebServer(cfg *config.Config) error {
	engine := gin.New()
	engine.Use(gin.Recovery())

	engine.UseRawPath = true

	if err := setRouters(engine, cfg); err != nil {
		return err
	}

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:           engine,
		ReadTimeout:       time.Duration(cfg.Server.ReadTimeout) * time.Millisecond,
		ReadHeaderTimeout: time.Duration(cfg.Server.ReadHeaderTimeout) * time.Millisecond,
	}

	return server.ListenAndServe()
}

func setRouters(engine *gin.Engine, cfg *config.Config) error {

	rg := engine.Group(BasePath)

	agentService := agent.NewAgentService(cfg.Agent)
	handlerService := NewHandler(agentService)

	// health check
	rg.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// abstract router
	rg.POST("/v1/abstract", handlerService.AbstractHandler())

	// polish router
	rg.POST("/v1/polish", handlerService.PolishHandler())

	return nil
}
