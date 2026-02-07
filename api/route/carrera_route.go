package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	//	"gorm-template/api/controller"

	"github.com/gin-gonic/gin"
)

func CarreraRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.CarreraController{
		CarreraRepository: &usecase.CarreraUseCase{},
	}
	CarreraRouter := group.Group("/Carrera")
	CarreraRouter.POST("/", ec.Create)
	CarreraRouter.GET("/", ec.Fetch)
	CarreraRouter.GET("/:id", ec.FetchById)
	CarreraRouter.PUT("/", ec.Update)
	CarreraRouter.DELETE("/:id", ec.Delete)
}
