package controller

import (
	"net/http"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProyectoPuestoController struct {
	ProyectoPuestoRepository domain.ProyectoPuestoRepository
}

func (ppc *ProyectoPuestoController) Create(c *gin.Context) {
	var proyectoPuesto domain.ProyectoPuesto

	err := c.ShouldBind(&proyectoPuesto)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	proyectoPuesto.ID = uuid.New()

	if proyectoPuesto.CantidadVacantes == 0 {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "CantidadVacantes is required"})
		return
	}

	err = ppc.ProyectoPuestoRepository.Create(c, proyectoPuesto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "ProyectoPuesto created successfully",
	})
}

func (ppc *ProyectoPuestoController) Fetch(c *gin.Context) {
	proyectoPuestos, err := ppc.ProyectoPuestoRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, proyectoPuestos)
}

func (ppc *ProyectoPuestoController) FetchById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid UUID format"})
		return
	}

	proyectoPuesto, err := ppc.ProyectoPuestoRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, proyectoPuesto)
}

func (ppc *ProyectoPuestoController) Update(c *gin.Context) {
	updatedProyectoPuesto := &domain.ProyectoPuesto{}

	err := c.ShouldBind(updatedProyectoPuesto)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedProyectoPuesto.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID ProyectoPuesto is required to update"})
		return
	}

	err = ppc.ProyectoPuestoRepository.Update(c, *updatedProyectoPuesto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "ProyectoPuesto updated successfully"})
}

func (ppc *ProyectoPuestoController) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid UUID format"})
		return
	}

	err = ppc.ProyectoPuestoRepository.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "ProyectoPuesto deleted successfully"})
}
