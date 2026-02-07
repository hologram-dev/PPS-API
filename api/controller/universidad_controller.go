package controller

import (
	"net/http"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UniversidadController struct {
	UniversidadRepository domain.UniversidadRepository
}

func (uc *UniversidadController) Create(c *gin.Context) {
	var universidad domain.Universidad

	err := c.ShouldBind(&universidad)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	universidad.ID = uuid.New()

	if universidad.NombreUniversidad == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "NombreUniversidad is required"})
		return
	}

	if universidad.CuitUniversidad == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "CuitUniversidad is required"})
		return
	}

	if universidad.DireccionUniversidad == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "DireccionUniversidad is required"})
		return
	}

	if universidad.CorreoUniversidad == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "CorreoUniversidad is required"})
		return
	}

	err = uc.UniversidadRepository.Create(c, universidad)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Universidad created successfully",
	})
}

func (uc *UniversidadController) Fetch(c *gin.Context) {
	universidades, err := uc.UniversidadRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, universidades)
}

func (uc *UniversidadController) FetchById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid UUID format"})
		return
	}

	universidad, err := uc.UniversidadRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, universidad)
}

func (uc *UniversidadController) Update(c *gin.Context) {
	updatedUniversidad := &domain.Universidad{}

	err := c.ShouldBind(updatedUniversidad)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedUniversidad.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID Universidad is required to update"})
		return
	}

	err = uc.UniversidadRepository.Update(c, *updatedUniversidad)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Universidad updated successfully"})
}

func (uc *UniversidadController) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid UUID format"})
		return
	}

	err = uc.UniversidadRepository.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Universidad deleted successfully"})
}
