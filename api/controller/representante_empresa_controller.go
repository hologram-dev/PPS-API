package controller

import (
	"net/http"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RepresentanteEmpresaController struct {
	RepresentanteEmpresaRepository domain.RepresentanteEmpresaRepository
}

func (ree *RepresentanteEmpresaController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var RepresentanteEmpresa domain.RepresentanteEmpresa

	err := c.ShouldBind(&RepresentanteEmpresa)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if RepresentanteEmpresa.Nombre == "" || RepresentanteEmpresa.Apellido == "" || RepresentanteEmpresa.Correo == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Nombre, Apellido o Correo es requerido"})
		return
	}

	RepresentanteEmpresa.ID = uuid.New()

	err = ree.RepresentanteEmpresaRepository.Create(c, RepresentanteEmpresa)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "RepresentanteEmpresa created successfully",
	})
}

func (ree *RepresentanteEmpresaController) Fetch(c *gin.Context) {
	RepresentanteEmpresas, err := ree.RepresentanteEmpresaRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, RepresentanteEmpresas)
}

func (ree *RepresentanteEmpresaController) FetchById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	RepresentanteEmpresas, err := ree.RepresentanteEmpresaRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, RepresentanteEmpresas)
}

func (ree *RepresentanteEmpresaController) Update(c *gin.Context) {
	updatedRepresentanteEmpresa := &domain.RepresentanteEmpresa{}

	err := c.ShouldBind(updatedRepresentanteEmpresa)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedRepresentanteEmpresa.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID RepresentanteEmpresa is requiered to update"})
		return
	}

	err = ree.RepresentanteEmpresaRepository.Update(c, *updatedRepresentanteEmpresa)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "RepresentanteEmpresa updated succesfully"})
}

func (ree *RepresentanteEmpresaController) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = ree.RepresentanteEmpresaRepository.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "RepresentanteEmpresa delete succesfully"})
}
