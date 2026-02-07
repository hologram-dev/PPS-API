package controller

import (
	"net/http"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CarreraController struct {
	CarreraRepository domain.CarreraRepository
}

func (cc *CarreraController) Create(c *gin.Context) {
	var Carrera domain.Carrera

	err := c.ShouldBind(&Carrera)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if Carrera.Nombre == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Name is required"})
		return
	}

	Carrera.ID = uuid.New()

	err = cc.CarreraRepository.Create(c, Carrera)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Carrera created successfully",
	})
}

func (cc *CarreraController) Fetch(c *gin.Context) {
	Carreras, err := cc.CarreraRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Carreras)
}

func (cc *CarreraController) FetchById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	Carreras, err := cc.CarreraRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Carreras)
}

func (cc *CarreraController) Update(c *gin.Context) {
	updatedCarrera := &domain.Carrera{}

	err := c.ShouldBind(updatedCarrera)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedCarrera.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID Carrera is requiered to update"})
		return
	}

	err = cc.CarreraRepository.Update(c, *updatedCarrera)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Carrera updated succesfully"})
}

func (cc *CarreraController) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = cc.CarreraRepository.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Carrera delete succesfully"})
}
