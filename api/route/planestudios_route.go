package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func PlanEstudiosRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	pec := &controller.PlanEstudiosController{
		PlanEstudiosRepository: &usecase.PlanEstudiosUseCase{},
	}
	planEstudiosRouter := group.Group("/planestudios")
	planEstudiosRouter.POST("/", pec.Create)
	planEstudiosRouter.GET("/", pec.Fetch)
	planEstudiosRouter.GET("/:id", pec.FetchById)
	planEstudiosRouter.PUT("/", pec.Update)
	planEstudiosRouter.DELETE("/:id", pec.Delete)
}
