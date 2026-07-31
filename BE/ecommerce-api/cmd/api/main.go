package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/horadoqa/ecommerce-api/config"
	"github.com/horadoqa/ecommerce-api/internal/handler"
	"github.com/horadoqa/ecommerce-api/internal/repository"
	"github.com/horadoqa/ecommerce-api/internal/routes"
	"github.com/horadoqa/ecommerce-api/internal/service"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Println("Arquivo .env não encontrado")
	}

	// Conexão com banco
	db, err := config.ConnectDatabase()

	if err != nil {
		log.Fatal(err)
	}

	// Repository
	clienteRepository := repository.ClienteRepository{
		DB: db,
	}

	// Service
	clienteService := service.ClienteService{
		Repository: &clienteRepository,
	}

	// Handler
	clienteHandler := handler.ClienteHandler{
		Service: &clienteService,
	}

	// Gin
	router := gin.Default()

	// Healthcheck
	// router.GET("/healthcheck", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"status": "ok",
	// 		"api":    "ecommerce-api",
	// 	})
	// })

	// Healthcheck
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "WORKING",
		})
	})

	// Rota raiz
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "online",
			"api":    "ecommerce-api",
		})
	})

	// Rotas da aplicação
	routes.ClienteRoutes(
		router,
		&clienteHandler,
	)

	// Inicia API
	router.Run(":8080")
}
