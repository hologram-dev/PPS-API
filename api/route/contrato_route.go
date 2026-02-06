package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func ContratoRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	cc := &controller.ContratoController{
		ContratoRepository: &usecase.ContratoUseCase{},
	}
	contratoRouter := group.Group("/contrato")
	contratoRouter.POST("/", cc.Create)
	contratoRouter.GET("/", cc.Fetch)
	contratoRouter.GET("/:id", cc.FetchById)
	contratoRouter.PUT("/", cc.Update)
	contratoRouter.DELETE("/:id", cc.Delete)
}
