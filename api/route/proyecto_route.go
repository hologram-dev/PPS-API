package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func ProyectoRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	pc := &controller.ProyectoController{
		ProyectoRepository: &usecase.ProyectoUseCase{},
	}
	ProyectoRouter := group.Group("/proyecto")
	ProyectoRouter.POST("/", pc.Create)
	ProyectoRouter.GET("/", pc.Fetch)
	ProyectoRouter.GET("/:id", pc.FetchById)
	ProyectoRouter.PUT("/", pc.Update)
	ProyectoRouter.DELETE("/:id", pc.Delete)
}
