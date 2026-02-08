package controller

import (
	"net/http"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PostulacionEstadoController struct {
	PostulacionEstadoRepository domain.PostulacionEstadoRepository
}

func (pee *PostulacionEstadoController) Create(c *gin.Context) {
	var PostulacionEstado domain.PostulacionEstado

	err := c.ShouldBind(&PostulacionEstado)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	PostulacionEstado.ID = uuid.New()

	err = pee.PostulacionEstadoRepository.Create(c, PostulacionEstado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "PostulacionEstado created successfully",
	})
}

func (pee *PostulacionEstadoController) Fetch(c *gin.Context) {
	PostulacionEstados, err := pee.PostulacionEstadoRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, PostulacionEstados)
}

func (pee *PostulacionEstadoController) FetchById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	PostulacionEstados, err := pee.PostulacionEstadoRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, PostulacionEstados)
}

func (pee *PostulacionEstadoController) Update(c *gin.Context) {
	updatedEntity1 := &domain.PostulacionEstado{}

	err := c.ShouldBind(updatedEntity1)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedEntity1.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID PostulacionEstado is requiered to update"})
		return
	}

	err = pee.PostulacionEstadoRepository.Update(c, *updatedEntity1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "PostulacionEstado updated succesfully"})
}

func (pee *PostulacionEstadoController) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = pee.PostulacionEstadoRepository.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "PostulacionEstado delete succesfully"})
}
