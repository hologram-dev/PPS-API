package controller

import (
	"net/http"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PuestoController struct {
	PuestoRepository domain.PuestoRepository
}

func (te *PuestoController) Create(c *gin.Context) {
	var Puesto domain.Puesto

	err := c.ShouldBind(&Puesto)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if Puesto.NombrePuesto == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Nombre Puesto is required"})
		return
	}

	Puesto.ID = uuid.New()

	err = te.PuestoRepository.Create(c, Puesto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Puesto created successfully",
	})
}

func (te *PuestoController) Fetch(c *gin.Context) {
	Puestos, err := te.PuestoRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Puestos)
}

func (te *PuestoController) FetchById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	Puestos, err := te.PuestoRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Puestos)
}

func (te *PuestoController) Update(c *gin.Context) {
	updatedPuesto := &domain.Puesto{}

	err := c.ShouldBind(updatedPuesto)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedPuesto.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID Puesto is requiered to update"})
		return
	}

	err = te.PuestoRepository.Update(c, *updatedPuesto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Puesto updated succesfully"})
}

func (te *PuestoController) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = te.PuestoRepository.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Puesto delete succesfully"})
}
