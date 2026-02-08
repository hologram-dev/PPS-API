package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func ProcesoSeleccionPostulacionRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.ProcesoSeleccionPostulacionController{
		ProcesoSeleccionPostulacionRepository: &usecase.ProcesoSeleccionPostulacionUseCase{},
	}
	ProcesoSeleccionPostulacionRouter := group.Group("/proceso_seleccion_postulacion")
	ProcesoSeleccionPostulacionRouter.POST("/", ec.Create)
	ProcesoSeleccionPostulacionRouter.GET("/", ec.Fetch)
	ProcesoSeleccionPostulacionRouter.GET("/:id", ec.FetchById)
	ProcesoSeleccionPostulacionRouter.PUT("/", ec.Update)
	ProcesoSeleccionPostulacionRouter.DELETE("/:id", ec.Delete)
}
