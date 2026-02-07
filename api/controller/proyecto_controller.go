package controller

import (
	"net/http"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProyectoController struct {
	ProyectoRepository domain.ProyectoRepository
}

func (pc *ProyectoController) Create(c *gin.Context) {
	var proyecto domain.Proyecto

	err := c.ShouldBind(&proyecto)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if proyecto.NombreProyecto == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "NombreProyecto is required"})
		return
	}

	proyecto.ID = uuid.New()

	err = pc.ProyectoRepository.Create(c, proyecto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Proyecto created successfully",
	})
}

func (pc *ProyectoController) Fetch(c *gin.Context) {
	proyectos, err := pc.ProyectoRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, proyectos)
}

func (pc *ProyectoController) FetchById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	proyectos, err := pc.ProyectoRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, proyectos)
}

func (pc *ProyectoController) Update(c *gin.Context) {
	updatedProyecto := &domain.Proyecto{}

	err := c.ShouldBind(updatedProyecto)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedProyecto.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID Proyecto is requiered to update"})
		return
	}

	err = pc.ProyectoRepository.Update(c, *updatedProyecto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Proyecto updated succesfully"})
}

func (pc *ProyectoController) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = pc.ProyectoRepository.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Proyecto delete succesfully"})
}
