package controller

import (
	"net/http"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PlanEstudiosController struct {
	PlanEstudiosRepository domain.PlanEstudiosRepository
}

func (pec *PlanEstudiosController) Create(c *gin.Context) {
	var planEstudios domain.PlanEstudios

	err := c.ShouldBind(&planEstudios)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	planEstudios.ID = uuid.New()

	if planEstudios.FechaInicioPlanEstudios.IsZero() {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "FechaInicioPlanEstudios is required"})
		return
	}

	if planEstudios.NombrePlanEstudios == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "NombrePlanEstudios is required"})
		return
	}

	err = pec.PlanEstudiosRepository.Create(c, planEstudios)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "PlanEstudios created successfully",
	})
}

func (pec *PlanEstudiosController) Fetch(c *gin.Context) {
	planEstudios, err := pec.PlanEstudiosRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, planEstudios)
}

func (pec *PlanEstudiosController) FetchById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid UUID format"})
		return
	}

	planEstudio, err := pec.PlanEstudiosRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, planEstudio)
}

func (pec *PlanEstudiosController) Update(c *gin.Context) {
	updatedPlanEstudios := &domain.PlanEstudios{}

	err := c.ShouldBind(updatedPlanEstudios)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedPlanEstudios.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID PlanEstudios is required to update"})
		return
	}

	err = pec.PlanEstudiosRepository.Update(c, *updatedPlanEstudios)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "PlanEstudios updated successfully"})
}

func (pec *PlanEstudiosController) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid UUID format"})
		return
	}

	err = pec.PlanEstudiosRepository.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "PlanEstudios deleted successfully"})
}
