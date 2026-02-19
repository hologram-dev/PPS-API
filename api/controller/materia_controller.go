package controller

import (
	"net/http"
	"strconv"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
)

type MateriaController struct {
	MateriaRepository domain.MateriaRepository
}

func (me *MateriaController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var Materia domain.Materia

	err := c.ShouldBind(&Materia)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if Materia.Name == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Name is required"})
		return
	}

	err = me.MateriaRepository.Create(c, Materia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Materia created successfully",
	})
}

func (me *MateriaController) Fetch(c *gin.Context) {
	Materias, err := me.MateriaRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Materias)
}

func (me *MateriaController) FetchById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	Materias, err := me.MateriaRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Materias)
}

func (me *MateriaController) Update(c *gin.Context) {
	updatedMateria := &domain.Materia{}

	err := c.ShouldBind(updatedMateria)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedMateria.ID == 0 {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID Materia is requiered to update"})
		return
	}

	err = me.MateriaRepository.Update(c, *updatedMateria)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Materia updated succesfully"})
}

func (me *MateriaController) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = me.MateriaRepository.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Materia delete succesfully"})
}
