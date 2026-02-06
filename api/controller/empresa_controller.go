package controller

import (
	"net/http"
	"strconv"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EmpresaController struct {
	EmpresaRepository domain.EmpresaRepository
}

func (te *EmpresaController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var Empresa domain.Empresa

	err := c.ShouldBind(&Empresa)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	Empresa.ID = uuid.New()

	if Empresa.NombreEmpresa == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "NombreEmpresa is required"})
		return
	}

	if Empresa.CodPostalEmpresa == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "CodPostalEmpresa is required"})
		return
	}	

	if Empresa.CuitEmpresa == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "CuitEmpresa is required"})
		return
	}

	if Empresa.DireccionEmpresa == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "DireccionEmpresa is required"})
		return
	}

	if Empresa.NumeroTelefonoEmpresa == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "NumeroTelefonoEmpresa is required"})
		return
	}

	err = te.EmpresaRepository.Create(c, Empresa)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Empresa created successfully",
	})
}

func (te *EmpresaController) Fetch(c *gin.Context) {
	Empresas, err := te.EmpresaRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Empresas)
}

func (te *EmpresaController) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	Empresas, err := te.EmpresaRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Empresas)
}

func (te *EmpresaController) Update(c *gin.Context) {
	updatedEmpresa := &domain.Empresa{}

	err := c.ShouldBind(updatedEmpresa)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedEmpresa.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID Empresa is requiered to update"})
		return
	}

	err = te.EmpresaRepository.Update(c, *updatedEmpresa)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Empresa updated succesfully"})
}

func (te *EmpresaController) Delete(c *gin.Context) {
	EmpresaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = te.EmpresaRepository.Delete(c, EmpresaID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Empresa delete succesfully"})
}
