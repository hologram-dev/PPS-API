package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func RepresentanteEmpresaRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.RepresentanteEmpresaController{
		RepresentanteEmpresaRepository: &usecase.RepresentanteEmpresaUseCase{},
	}
	RepresentanteEmpresaRouter := group.Group("/representante-empresa")
	RepresentanteEmpresaRouter.POST("/", ec.Create)
	RepresentanteEmpresaRouter.GET("/", ec.Fetch)
	RepresentanteEmpresaRouter.GET("/:id", ec.FetchById)
	RepresentanteEmpresaRouter.PUT("/", ec.Update)
	RepresentanteEmpresaRouter.DELETE("/:id", ec.Delete)
}
