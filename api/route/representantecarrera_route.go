package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func RepresentanteCarreraRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	rcc := &controller.RepresentanteCarreraController{
		RepresentanteCarreraRepository: &usecase.RepresentanteCarreraUseCase{},
	}
	representanteCarreraRouter := group.Group("/representante-carrera")
	representanteCarreraRouter.POST("/", rcc.Create)
	representanteCarreraRouter.GET("/", rcc.Fetch)
	representanteCarreraRouter.GET("/:id", rcc.FetchById)
	representanteCarreraRouter.PUT("/", rcc.Update)
	representanteCarreraRouter.DELETE("/:id", rcc.Delete)
}
