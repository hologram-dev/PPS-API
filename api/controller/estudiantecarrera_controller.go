package controller

import (
	"gorm-template/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EstudianteCarreraController struct {
	EstudianteCarreraRepository domain.EstudianteCarreraRepository
}

func (ecc *EstudianteCarreraController) Create(c *gin.Context) {
	var estudianteCarrera domain.EstudianteCarrera

	err := c.ShouldBind(&estudianteCarrera)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	estudianteCarrera.ID = uuid.New()

	if estudianteCarrera.NroLegajo == 0 {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "NroLegajo is required"})
		return
	}

	err = ecc.EstudianteCarreraRepository.Create(c, estudianteCarrera)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "EstudianteCarrera created successfully",
	})
}

func (ecc *EstudianteCarreraController) Fetch(c *gin.Context) {
	estudianteCarreras, err := ecc.EstudianteCarreraRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, estudianteCarreras)
}

func (ecc *EstudianteCarreraController) FetchById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	estudianteCarrera, err := ecc.EstudianteCarreraRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, estudianteCarrera)
}

func (ecc *EstudianteCarreraController) Update(c *gin.Context) {
	updatedEstudianteCarrera := &domain.EstudianteCarrera{}

	err := c.ShouldBind(updatedEstudianteCarrera)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedEstudianteCarrera.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID EstudianteCarrera is requiered to update"})
		return
	}

	err = ecc.EstudianteCarreraRepository.Update(c, *updatedEstudianteCarrera)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "EstudianteCarrera updated succesfully"})
}

func (ecc *EstudianteCarreraController) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = ecc.EstudianteCarreraRepository.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "EstudianteCarrera deleted successfully"})
}
