package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func EmpresaRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.EmpresaController{
		EmpresaRepository: &usecase.EmpresaUseCase{},
	}
	EmpresaRouter := group.Group("/empresa")
	EmpresaRouter.POST("/", ec.Create)
	EmpresaRouter.GET("/", ec.Fetch)
	EmpresaRouter.GET("/:id", ec.FetchById)
	EmpresaRouter.PUT("/", ec.Update)
	EmpresaRouter.DELETE("/:id", ec.Delete)
}
