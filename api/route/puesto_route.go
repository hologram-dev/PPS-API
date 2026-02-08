package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func PuestoRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.PuestoController{
		PuestoRepository: &usecase.PuestoUseCase{},
	}
	PuestoRouter := group.Group("/puesto")
	PuestoRouter.POST("/", ec.Create)
	PuestoRouter.GET("/", ec.Fetch)
	PuestoRouter.GET("/:id", ec.FetchById)
	PuestoRouter.PUT("/", ec.Update)
	PuestoRouter.DELETE("/:id", ec.Delete)
}
