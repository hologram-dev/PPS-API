package controller

import (
	"net/http"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RepresentanteCarreraController struct {
	RepresentanteCarreraRepository domain.RepresentanteCarreraRepository
}

func (rcc *RepresentanteCarreraController) Create(c *gin.Context) {
	var representanteCarrera domain.RepresentanteCarrera
	err := c.ShouldBindJSON(&representanteCarrera)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	representanteCarrera.ID = uuid.New()

	if representanteCarrera.ApellidoRepresentanteCarrera == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ApellidoRepresentanteCarrera is required"})
		return
	}

	if representanteCarrera.CorreoRepresentanteCarrera == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "CorreoRepresentanteCarrera is required"})
		return
	}

	if representanteCarrera.NombreRepresentanteCarrera == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "NombreRepresentanteCarrera is required"})
		return
	}

	err = rcc.RepresentanteCarreraRepository.Create(c, representanteCarrera)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "RepresentanteCarrera created successfully"})
}

func (rcc *RepresentanteCarreraController) Fetch(c *gin.Context) {
	representanteCarrera, err := rcc.RepresentanteCarreraRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, representanteCarrera)
}

func (rcc *RepresentanteCarreraController) FetchById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid UUID format"})
		return
	}

	representanteCarrera, err := rcc.RepresentanteCarreraRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, representanteCarrera)
}

func (rcc *RepresentanteCarreraController) Update(c *gin.Context) {
	updatedRepresentanteCarrera := &domain.RepresentanteCarrera{}
	err := c.ShouldBind(updatedRepresentanteCarrera)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedRepresentanteCarrera.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID RepresentanteCarrera is required to update"})
		return
	}

	err = rcc.RepresentanteCarreraRepository.Update(c, *updatedRepresentanteCarrera)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "RepresentanteCarrera updated successfully"})
}

func (rcc *RepresentanteCarreraController) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid UUID format"})
		return
	}

	err = rcc.RepresentanteCarreraRepository.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "RepresentanteCarrera deleted successfully"})
}
