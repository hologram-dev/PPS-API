package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func EstudianteCarreraRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ecc := &controller.EstudianteCarreraController{
		EstudianteCarreraRepository: &usecase.EstudianteCarreraUseCase{},
	}
	EstudianteCarreraRouter := group.Group("/estudiantecarrera")
	EstudianteCarreraRouter.POST("/", ecc.Create)
	EstudianteCarreraRouter.GET("/", ecc.Fetch)
	EstudianteCarreraRouter.GET("/:id", ecc.FetchById)
	EstudianteCarreraRouter.PUT("/", ecc.Update)
	EstudianteCarreraRouter.DELETE("/:id", ecc.Delete)
}
