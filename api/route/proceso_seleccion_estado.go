package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func ProcesoSeleccionEstadoRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.ProcesoSeleccionEstadoController{
		ProcesoSeleccionEstadoRepository: &usecase.ProcesoSeleccionEstadoUseCase{},
	}
	ProcesoSeleccionEstadoRouter := group.Group("/proceso-seleccion-estado")
	ProcesoSeleccionEstadoRouter.POST("/", ec.Create)
	ProcesoSeleccionEstadoRouter.GET("/", ec.Fetch)
	ProcesoSeleccionEstadoRouter.GET("/:id", ec.FetchById)
	ProcesoSeleccionEstadoRouter.PUT("/", ec.Update)
	ProcesoSeleccionEstadoRouter.DELETE("/:id", ec.Delete)
}
