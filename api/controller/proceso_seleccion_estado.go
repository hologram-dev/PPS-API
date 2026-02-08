package controller

import (
	"net/http"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProcesoSeleccionEstadoController struct {
	ProcesoSeleccionEstadoRepository domain.ProcesoSeleccionEstadoRepository
}

func (sec *ProcesoSeleccionEstadoController) Create(c *gin.Context) {
	var ProcesoSeleccionEstado domain.ProcesoSeleccionEstado

	err := c.ShouldBind(&ProcesoSeleccionEstado)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	ProcesoSeleccionEstado.ID = uuid.New()

	err = sec.ProcesoSeleccionEstadoRepository.Create(c, ProcesoSeleccionEstado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "ProcesoSeleccionEstado created successfully",
	})
}

func (sec *ProcesoSeleccionEstadoController) Fetch(c *gin.Context) {
	ProcesoSeleccionEstados, err := sec.ProcesoSeleccionEstadoRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ProcesoSeleccionEstados)
}

func (sec *ProcesoSeleccionEstadoController) FetchById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ProcesoSeleccionEstados, err := sec.ProcesoSeleccionEstadoRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, ProcesoSeleccionEstados)
}

func (sec *ProcesoSeleccionEstadoController) Update(c *gin.Context) {
	updatedProcesoSeleccionEstado := &domain.ProcesoSeleccionEstado{}

	err := c.ShouldBind(updatedProcesoSeleccionEstado)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedProcesoSeleccionEstado.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID ProcesoSeleccionEstado is requiered to update"})
		return
	}

	err = sec.ProcesoSeleccionEstadoRepository.Update(c, *updatedProcesoSeleccionEstado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "ProcesoSeleccionEstado updated succesfully"})
}

func (sec *ProcesoSeleccionEstadoController) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = sec.ProcesoSeleccionEstadoRepository.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "ProcesoSeleccionEstado delete succesfully"})
}
