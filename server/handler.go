package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/victorzhou123/ai-agent/agent"
)

type handlerService struct {
	agent agent.AgentService
}

func NewHandler(agent agent.AgentService) handlerService {
	return handlerService{agent: agent}
}

func (s *handlerService) AbstractHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req AbstractRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("request failed, err: %s", err)})
			return
		}

		abs, err := s.agent.Abstract(req.Content)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("agent failed, err: %s", err)})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"content": abs})
	}
}

func (s *handlerService) PolishHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req PolishRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("request failed, err: %s", err)})
			return
		}

		polished, err := s.agent.Polish(req.Content)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("agent failed, err: %s", err)})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"content": polished})
	}
}
