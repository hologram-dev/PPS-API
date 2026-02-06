package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func UniversidadRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	uc := &controller.UniversidadController{
		UniversidadRepository: &usecase.UniversidadUseCase{},
	}
	universidadRouter := group.Group("/universidad")
	universidadRouter.POST("/", uc.Create)
	universidadRouter.GET("/", uc.Fetch)
	universidadRouter.GET("/:id", uc.FetchById)
	universidadRouter.PUT("/", uc.Update)
	universidadRouter.DELETE("/:id", uc.Delete)
}
