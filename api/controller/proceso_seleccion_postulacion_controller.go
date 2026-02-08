package controller

import (
	"net/http"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProcesoSeleccionPostulacionController struct {
	ProcesoSeleccionPostulacionRepository domain.ProcesoSeleccionPostulacionRepository
}

func (spe *ProcesoSeleccionPostulacionController) Create(c *gin.Context) {
	var ProcesoSeleccionPostulacion domain.ProcesoSeleccionPostulacion

	err := c.ShouldBind(&ProcesoSeleccionPostulacion)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	ProcesoSeleccionPostulacion.ID = uuid.New()

	err = spe.ProcesoSeleccionPostulacionRepository.Create(c, ProcesoSeleccionPostulacion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "ProcesoSeleccionPostulacion created successfully",
	})
}

func (spe *ProcesoSeleccionPostulacionController) Fetch(c *gin.Context) {
	ProcesoSeleccionPostulacions, err := spe.ProcesoSeleccionPostulacionRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ProcesoSeleccionPostulacions)
}

func (spe *ProcesoSeleccionPostulacionController) FetchById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ProcesoSeleccionPostulacions, err := spe.ProcesoSeleccionPostulacionRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, ProcesoSeleccionPostulacions)
}

func (spe *ProcesoSeleccionPostulacionController) Update(c *gin.Context) {
	updatedProcesoSeleccionPostulacion := &domain.ProcesoSeleccionPostulacion{}

	err := c.ShouldBind(updatedProcesoSeleccionPostulacion)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedProcesoSeleccionPostulacion.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID ProcesoSeleccionPostulacion is requiered to update"})
		return
	}

	err = spe.ProcesoSeleccionPostulacionRepository.Update(c, *updatedProcesoSeleccionPostulacion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "ProcesoSeleccionPostulacion updated succesfully"})
}

func (spe *ProcesoSeleccionPostulacionController) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = spe.ProcesoSeleccionPostulacionRepository.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "ProcesoSeleccionPostulacion delete succesfully"})
}
