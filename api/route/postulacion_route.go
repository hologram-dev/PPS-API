package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func PostulacionRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	pc := &controller.PostulacionController{
		PostulacionRepository: &usecase.PostulacionUseCase{},
	}
	postulacionRouter := group.Group("/postulacion")
	postulacionRouter.POST("/", pc.Create)
	postulacionRouter.GET("/", pc.Fetch)
	postulacionRouter.GET("/:id", pc.FetchById)
	postulacionRouter.PUT("/", pc.Update)
	postulacionRouter.DELETE("/:id", pc.Delete)
}
