package route

import (
	"gorm-template/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, gin *gin.Engine) {

	publicRouter := gin.Group("/api")


	EmpresaRouter(env, timeout, publicRouter)
}
