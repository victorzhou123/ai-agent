package agent

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerService struct {
	agent AgentService
}

func NewHandler(agent AgentService) handlerService {
	return handlerService{agent: agent}
}

func (s *handlerService) AbstractHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req *AbstractRequest
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "request failed"})
			return
		}

		resp, err := s.agent.Abstract(req.Content)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("agent failed, err: %s", err)})
			return
		}

		ctx.JSON(http.StatusOK, resp)
	}
}

func (s *handlerService) PolishHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req *PolishRequest
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "request failed"})
			return
		}

		resp, err := s.agent.Polish(req.Content)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("agent failed, err: %s", err)})
			return
		}

		ctx.JSON(http.StatusOK, resp)
	}
}
