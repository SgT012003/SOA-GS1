package controller

import (
	"net/http"
	"strconv"

	"upskilling-api/model"
	"upskilling-api/service"

	"github.com/gin-gonic/gin"
)

var matriculaService = service.NewMatriculaService()

// MatricularRequest é o DTO para a requisição de matrícula.
type MatricularRequest struct {
	UsuarioID int64 `json:"usuario_id" binding:"required,gt=0"`
	TrilhaID  int64 `json:"trilha_id" binding:"required,gt=0"`
}

// MatricularUsuario godoc
// @Summary Matricular usuário em uma trilha
// @Description Realiza a inscrição de um usuário em uma trilha de aprendizagem.
// @Tags Matriculas
// @Accept json
// @Produce json
// @Param matricula body MatricularRequest true "Dados da Matrícula"
// @Success 201 {object} model.Matricula
// @Failure 400 {object} model.ErrorResponse
// @Failure 422 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /matriculas [post]
func MatricularUsuario(c *gin.Context) {
	var req MatricularRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Dados de entrada inválidos.", Details: err.Error()})
		return
	}

	res, err := matriculaService.Matricular(req.UsuarioID, req.TrilhaID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetMatriculasByUsuario godoc
// @Summary Lista matrículas de um usuário
// @Description Retorna todas as matrículas de um usuário específico.
// @Tags Matriculas
// @Produce json
// @Param id path int true "ID do Usuário"
// @Success 200 {array} model.Matricula
// @Failure 400 {object} model.ErrorResponse
// @Failure 422 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /usuarios/{id}/matriculas [get]
func GetMatriculasByUsuario(c *gin.Context) {
	usuarioID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "ID de Usuário inválido.", Details: "O ID deve ser um número inteiro."})
		return
	}

	res, err := matriculaService.GetMatriculasByUsuario(usuarioID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
