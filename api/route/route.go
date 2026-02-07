package route

import (
	"gorm-template/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, gin *gin.Engine) {

	publicRouter := gin.Group("/api")

	EmpresaRouter(env, timeout, publicRouter)
	UniversidadRouter(env, timeout, publicRouter)
	ContratoRouter(env, timeout, publicRouter)
	PostulacionRouter(env, timeout, publicRouter)
	ProyectoPuestoRouter(env, timeout, publicRouter)
	PuestoRouter(env, timeout, publicRouter)
}
