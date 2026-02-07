package controller

import (
	"net/http"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ContratoController struct {
	ContratoRepository domain.ContratoRepository
}

func (cc *ContratoController) Create(c *gin.Context) {
	var contrato domain.Contrato

	err := c.ShouldBind(&contrato)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	contrato.ID = uuid.New()

	if contrato.NumeroContrato == 0 {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "NumeroContrato is required"})
		return
	}

	if contrato.FechaInicioContrato.IsZero() {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "FechaInicioContrato is required"})
		return
	}

	if contrato.FechaFinContrato.IsZero() {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "FechaFinContrato is required"})
		return
	}

	err = cc.ContratoRepository.Create(c, contrato)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Contrato created successfully",
	})
}

func (cc *ContratoController) Fetch(c *gin.Context) {
	contratos, err := cc.ContratoRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, contratos)
}

func (cc *ContratoController) FetchById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid UUID format"})
		return
	}

	contrato, err := cc.ContratoRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, contrato)
}

func (cc *ContratoController) Update(c *gin.Context) {
	updatedContrato := &domain.Contrato{}

	err := c.ShouldBind(updatedContrato)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedContrato.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID Contrato is required to update"})
		return
	}

	err = cc.ContratoRepository.Update(c, *updatedContrato)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Contrato updated successfully"})
}

func (cc *ContratoController) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid UUID format"})
		return
	}

	err = cc.ContratoRepository.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Contrato deleted successfully"})
}
