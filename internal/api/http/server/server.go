package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	internalProm "github.com/avraam311/analysis-utility/internal/infra/prometheus"
)

func NewRouter() *gin.Engine {
	e := gin.New()
	e.Use(gin.Recovery(), gin.Logger())

	e.GET("/metrics", gin.WrapH(promhttp.HandlerFor(internalProm.Registry, promhttp.HandlerOpts{})))

	return e
}

func NewServer(addr string, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
