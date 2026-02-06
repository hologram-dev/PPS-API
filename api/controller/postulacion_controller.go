package controller

import (
	"net/http"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PostulacionController struct {
	PostulacionRepository domain.PostulacionRepository
}

func (pc *PostulacionController) Create(c *gin.Context) {
	var postulacion domain.Postulacion

	err := c.ShouldBind(&postulacion)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	postulacion.ID = uuid.New()

	if postulacion.NumeroPostulacion == 0 {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "NumeroPostulacion is required"})
		return
	}

	err = pc.PostulacionRepository.Create(c, postulacion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Postulacion created successfully",
	})
}

func (pc *PostulacionController) Fetch(c *gin.Context) {
	postulaciones, err := pc.PostulacionRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, postulaciones)
}

func (pc *PostulacionController) FetchById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid UUID format"})
		return
	}

	postulacion, err := pc.PostulacionRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, postulacion)
}

func (pc *PostulacionController) Update(c *gin.Context) {
	updatedPostulacion := &domain.Postulacion{}

	err := c.ShouldBind(updatedPostulacion)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedPostulacion.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID Postulacion is required to update"})
		return
	}

	err = pc.PostulacionRepository.Update(c, *updatedPostulacion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Postulacion updated successfully"})
}

func (pc *PostulacionController) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid UUID format"})
		return
	}

	err = pc.PostulacionRepository.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Postulacion deleted successfully"})
}
