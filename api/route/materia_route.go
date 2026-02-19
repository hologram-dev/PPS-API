package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func MateriaRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.MateriaController{
		MateriaRepository: &usecase.MateriaUseCase{},
	}
	MateriaRouter := group.Group("/materia")
	MateriaRouter.POST("/", ec.Create)
	MateriaRouter.GET("/", ec.Fetch)
	MateriaRouter.GET("/:id", ec.FetchById)
	MateriaRouter.PUT("/", ec.Update)
	MateriaRouter.DELETE("/:id", ec.Delete)
}
