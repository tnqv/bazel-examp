package router

import (
	"bazel_exam/pkg/shared/health"

	"github.com/gin-gonic/gin"
)

func GetEngine() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/", health.Health("Main app"))
	return r
}
