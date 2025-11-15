package controller

import (
	"net/http"
	"strconv"

	"upskilling-api/model"
	"upskilling-api/service"

	"github.com/gin-gonic/gin"
)

var trilhaService = service.NewTrilhaService()

// CreateTrilha godoc
// @Summary Cria uma nova trilha de aprendizagem
// @Description Cria uma nova trilha de upskilling/reskilling.
// @Tags Trilhas
// @Accept json
// @Produce json
// @Param trilha body model.CreateTrilhaRequest true "Dados da Trilha"
// @Success 201 {object} model.TrilhaResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /trilhas [post]
func CreateTrilha(c *gin.Context) {
	var req model.CreateTrilhaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Dados de entrada inválidos.", Details: err.Error()})
		return
	}

	res, err := trilhaService.Create(&req)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetAllTrilhas godoc
// @Summary Lista todas as trilhas de aprendizagem
// @Description Retorna uma lista de todas as trilhas cadastradas.
// @Tags Trilhas
// @Produce json
// @Success 200 {array} model.TrilhaResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /trilhas [get]
func GetAllTrilhas(c *gin.Context) {
	res, err := trilhaService.FindAll()
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetTrilhaByID godoc
// @Summary Busca uma trilha por ID
// @Description Retorna os detalhes de uma trilha específica.
// @Tags Trilhas
// @Produce json
// @Param id path int true "ID da Trilha"
// @Success 200 {object} model.TrilhaResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /trilhas/{id} [get]
func GetTrilhaByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "ID inválido.", Details: "O ID deve ser um número inteiro."})
		return
	}

	res, err := trilhaService.FindByID(id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateTrilha godoc
// @Summary Atualiza uma trilha
// @Description Atualiza os dados de uma trilha existente.
// @Tags Trilhas
// @Accept json
// @Produce json
// @Param id path int true "ID da Trilha"
// @Param trilha body model.UpdateTrilhaRequest true "Dados da Trilha para atualização"
// @Success 200 {object} model.TrilhaResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /trilhas/{id} [put]
func UpdateTrilha(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "ID inválido.", Details: "O ID deve ser um número inteiro."})
		return
	}

	var req model.UpdateTrilhaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Dados de entrada inválidos.", Details: err.Error()})
		return
	}

	res, err := trilhaService.Update(id, &req)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteTrilha godoc
// @Summary Deleta uma trilha
// @Description Remove uma trilha de aprendizagem.
// @Tags Trilhas
// @Produce json
// @Param id path int true "ID da Trilha"
// @Success 204 "No Content"
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /trilhas/{id} [delete]
func DeleteTrilha(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "ID inválido.", Details: "O ID deve ser um número inteiro."})
		return
	}

	err = trilhaService.Delete(id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
