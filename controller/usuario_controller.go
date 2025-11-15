package controller

import (
	"net/http"
	"strconv"

	"upskilling-api/model"
	"upskilling-api/service"

	"github.com/gin-gonic/gin"
)

var usuarioService = service.NewUsuarioService()

// CreateUsuario godoc
// @Summary Cria um novo usuário
// @Description Cria um novo usuário na plataforma de upskilling/reskilling.
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param usuario body model.CreateUsuarioRequest true "Dados do Usuário"
// @Success 201 {object} model.UsuarioResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 409 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /usuarios [post]
func CreateUsuario(c *gin.Context) {
	var req model.CreateUsuarioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Dados de entrada inválidos.", Details: err.Error()})
		return
	}

	res, err := usuarioService.Create(&req)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetAllUsuarios godoc
// @Summary Lista todos os usuários
// @Description Retorna uma lista de todos os usuários cadastrados.
// @Tags Usuarios
// @Produce json
// @Success 200 {array} model.UsuarioResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /usuarios [get]
func GetAllUsuarios(c *gin.Context) {
	res, err := usuarioService.FindAll()
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetUsuarioByID godoc
// @Summary Busca um usuário por ID
// @Description Retorna os detalhes de um usuário específico.
// @Tags Usuarios
// @Produce json
// @Param id path int true "ID do Usuário"
// @Success 200 {object} model.UsuarioResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /usuarios/{id} [get]
func GetUsuarioByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "ID inválido.", Details: "O ID deve ser um número inteiro."})
		return
	}

	res, err := usuarioService.FindByID(id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateUsuario godoc
// @Summary Atualiza um usuário
// @Description Atualiza os dados de um usuário existente.
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param id path int true "ID do Usuário"
// @Param usuario body model.UpdateUsuarioRequest true "Dados do Usuário para atualização"
// @Success 200 {object} model.UsuarioResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /usuarios/{id} [put]
func UpdateUsuario(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "ID inválido.", Details: "O ID deve ser um número inteiro."})
		return
	}

	var req model.UpdateUsuarioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Dados de entrada inválidos.", Details: err.Error()})
		return
	}

	res, err := usuarioService.Update(id, &req)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteUsuario godoc
// @Summary Deleta um usuário
// @Description Remove um usuário da plataforma.
// @Tags Usuarios
// @Produce json
// @Param id path int true "ID do Usuário"
// @Success 204 "No Content"
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /usuarios/{id} [delete]
func DeleteUsuario(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "ID inválido.", Details: "O ID deve ser um número inteiro."})
		return
	}

	err = usuarioService.Delete(id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
