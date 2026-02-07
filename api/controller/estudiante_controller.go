package controller

import (
	"net/http"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EstudianteController struct {
	EstudianteRepository domain.EstudianteRepository
}

func (ec *EstudianteController) Create(c *gin.Context) {
	var estudiante domain.Estudiante

	err := c.ShouldBind(&estudiante)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	estudiante.ID = uuid.New()

	if estudiante.ApellidoEstudiante == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ApellidoEstudiante is required"})
		return
	}

	if estudiante.CorreoEstudiante == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "CorreoEstudiante is required"})
		return
	}

	if estudiante.CorreoInstitucional == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "CorreoInstitucional is required"})
		return
	}

	if estudiante.CuilEstudiante == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "CuilEstudiante is required"})
		return
	}

	if estudiante.DniEstudiante == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "DniEstudiante is required"})
		return
	}

	if estudiante.FechaNacimientoEstudiante.IsZero() {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "FechaNacimientoEstudiante is required"})
		return
	}

	if estudiante.NombreEstudiante == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "NombreEstudiante is required"})
		return
	}

	if estudiante.TipoDNI == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "TipoDNI is required"})
		return
	}

	err = ec.EstudianteRepository.Create(c, estudiante)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Estudiante creado exitosamente"})
}

func (ec *EstudianteController) Fetch(c *gin.Context) {
	estudiantes, err := ec.EstudianteRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, estudiantes)
}

func (ec *EstudianteController) FetchById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid UUID format"})
		return
	}

	estudiante, err := ec.EstudianteRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, estudiante)
}

func (ec *EstudianteController) Update(c *gin.Context) {
	updatedEstudiante := &domain.Estudiante{}

	err := c.ShouldBind(updatedEstudiante)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedEstudiante.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID Estudiante is required to update"})
		return
	}

	err = ec.EstudianteRepository.Update(c, *updatedEstudiante)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Estudiante updated successfully"})
}

func (ec *EstudianteController) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid UUID format"})
		return
	}

	err = ec.EstudianteRepository.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Estudiante deleted successfully"})
}
