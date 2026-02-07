package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func PostulacionEstadoRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.PostulacionEstadoController{
		PostulacionEstadoRepository: &usecase.PostulacionEstadoUseCase{},
	}
	PostulacionEstadoRouter := group.Group("/postulacion_estado")
	PostulacionEstadoRouter.POST("/", ec.Create)
	PostulacionEstadoRouter.GET("/", ec.Fetch)
	PostulacionEstadoRouter.GET("/:id", ec.FetchById)
	PostulacionEstadoRouter.PUT("/", ec.Update)
	PostulacionEstadoRouter.DELETE("/:id", ec.Delete)
}
