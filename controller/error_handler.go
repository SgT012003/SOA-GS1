package controller

import (
	"log"
	"net/http"

	"upskilling-api/model"

	"github.com/gin-gonic/gin"
)

// ErrorHandlerMiddleware é um middleware para tratamento centralizado de erros.
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Verifica se houve algum erro durante o processamento da requisição
		if len(c.Errors) > 0 {
			err := c.Errors[0].Err

			// Tenta converter o erro para a interface CustomError
			if customErr, ok := err.(model.CustomError); ok {
				c.JSON(customErr.StatusCode(), model.ErrorResponse{
					Message: customErr.Message(),
					Details: customErr.Error(),
				})
				return
			}

			// Trata erros de validação do Gin (binding)
			if c.Writer.Status() == http.StatusBadRequest {
				// O Gin já deve ter retornado o erro 400, apenas garantimos o formato
				c.JSON(http.StatusBadRequest, model.ErrorResponse{
					Message: "Dados de entrada inválidos.",
					Details: err.Error(),
				})
				return
			}

			// Erro interno genérico (500)
			log.Printf("Erro interno não tratado: %v", err)
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: "Ocorreu um erro interno no servidor.",
				Details: err.Error(),
			})
		}
	}
}

// handleError é uma função auxiliar para lidar com erros de serviço.
func handleError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	// Adiciona o erro ao contexto para ser tratado pelo middleware
	c.Error(err)
	c.Abort()
}
