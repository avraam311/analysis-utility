package server

import (
	"net/http"

	"github.com/avraam311/analysis-utility/internal/api/http/handlers/analysis"
	"github.com/avraam311/analysis-utility/internal/infra/config"
	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *config.Config, handlerAn *analysis.Handler) *gin.Engine {
	e := gin.Default()

	anGroup := e.Group("/analysis")
	{
		anGroup.GET("/get", handlerAn.GetAnalysis)
	}

	return e
}

func NewServer(addr string, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
