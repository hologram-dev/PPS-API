package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func ProyectoPuestoRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ppc := &controller.ProyectoPuestoController{
		ProyectoPuestoRepository: &usecase.ProyectoPuestoUseCase{},
	}
	proyectoPuestoRouter := group.Group("/proyecto-puesto")
	proyectoPuestoRouter.POST("/", ppc.Create)
	proyectoPuestoRouter.GET("/", ppc.Fetch)
	proyectoPuestoRouter.GET("/:id", ppc.FetchById)
	proyectoPuestoRouter.PUT("/", ppc.Update)
	proyectoPuestoRouter.DELETE("/:id", ppc.Delete)
}
