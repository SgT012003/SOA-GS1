package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"upskilling-api/controller"
	"upskilling-api/db"

	_ "upskilling-api/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Plataforma de Upskilling/Reskilling API
// @version 1.0
// @description API RESTful para uma plataforma de Upskilling/Reskilling voltada ao futuro do trabalho (2030+).
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Carrega variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Não foi possível carregar o arquivo .env. Usando variáveis de ambiente do sistema.")
	}

	// Inicializa a conexão com o banco de dados
	db.InitDB()
	defer db.CloseDB()

	// Configura o Gin
	router := gin.Default()

	// Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Endpoint raiz
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "/swagger/index.html")
	})

	// Middleware de tratamento de erros customizado
	router.Use(controller.ErrorHandlerMiddleware())

	// Rotas da API
	v1 := router.Group("/api/v1")
	{
		// Rotas de Usuários (CRUD)
		usuarios := v1.Group("/usuarios")
		{
			usuarios.POST("/", controller.CreateUsuario)
			usuarios.GET("/", controller.GetAllUsuarios)
			usuarios.GET("/:id", controller.GetUsuarioByID)
			usuarios.PUT("/:id", controller.UpdateUsuario)
			usuarios.DELETE("/:id", controller.DeleteUsuario)
		}

		// Rotas de Trilhas (CRUD)
		trilhas := v1.Group("/trilhas")
		{
			trilhas.POST("/", controller.CreateTrilha)
			trilhas.GET("/", controller.GetAllTrilhas)
			trilhas.GET("/:id", controller.GetTrilhaByID)
			trilhas.PUT("/:id", controller.UpdateTrilha)
			trilhas.DELETE("/:id", controller.DeleteTrilha)
		}

		// Rotas de Inscrição (Extra)
		v1.POST("/matriculas", controller.MatricularUsuario)
		v1.GET("/usuarios/:id/matriculas", controller.GetMatriculasByUsuario)
	}

	// Rota para documentação Swagger (se gerada localmente)
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciado na porta: %s", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
