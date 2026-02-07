package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func EstudianteRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.EstudianteController{
		EstudianteRepository: &usecase.EstudianteUseCase{},
	}

	estudianteRouter := group.Group("/estudiante")
	estudianteRouter.POST("/", ec.Create)
	estudianteRouter.GET("/", ec.Fetch)
	estudianteRouter.GET("/:id", ec.FetchById)
	estudianteRouter.PUT("/", ec.Update)
	estudianteRouter.DELETE("/:id", ec.Delete)
}
