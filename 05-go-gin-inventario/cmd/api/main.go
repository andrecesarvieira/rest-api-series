package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"inventario-api/internal/database"
	"inventario-api/internal/handlers"
	"inventario-api/internal/middleware"
	"inventario-api/internal/repository"
	"inventario-api/internal/service"
)

func main() {
	// Inicializa banco de dados em memória
	db := database.NewInMemoryDatabase()
	
	// Inicializa repository
	repo := repository.NewInMemoryProductRepository(db)
	
	// Inicializa service
	productService := service.NewProductService(repo)
	
	// Inicializa handler
	productHandler := handlers.NewProductHandler(productService)
	
	// Configura Gin
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	
	// Middlewares globais
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())
	router.Use(middleware.CORS())
	router.Use(middleware.Security())
	router.Use(middleware.RequestID())
	router.Use(middleware.RateLimiter())
	
	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "OK",
			"service":   "Inventário API",
			"version":   "1.0.0",
			"timestamp": "2023-01-15T10:30:00Z",
		})
	})
	
	// API routes
	api := router.Group("/api")
	{
		produtos := api.Group("/produtos")
		{
			// CRUD básico
			produtos.POST("", productHandler.CreateProduct)
			produtos.GET("", productHandler.GetAllProducts)
			produtos.GET("/:id", productHandler.GetProduct)
			produtos.PUT("/:id", productHandler.UpdateProduct)
			produtos.DELETE("/:id", productHandler.DeleteProduct)
			
			// Endpoints especializados
			produtos.GET("/filtros", productHandler.GetProductsFiltered)
			produtos.GET("/categoria/:categoria", productHandler.GetProductsByCategory)
			produtos.GET("/ativos", productHandler.GetActiveProducts)
			produtos.GET("/estoque", productHandler.GetInStockProducts)
			produtos.PATCH("/:id/estoque", productHandler.UpdateStock)
			produtos.GET("/estatisticas", productHandler.GetStatistics)
		}
	}
	
	// Endpoint para documentação da API
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":     "API de Inventário - Go + Gin",
			"version":     "1.0.0",
			"description": "API REST para gerenciamento de inventário/estoque",
			"endpoints": gin.H{
				"health":              "GET /health",
				"produtos":            "GET /api/produtos",
				"criar_produto":       "POST /api/produtos",
				"buscar_produto":      "GET /api/produtos/{id}",
				"atualizar_produto":   "PUT /api/produtos/{id}",
				"deletar_produto":     "DELETE /api/produtos/{id}",
				"filtrar_produtos":    "GET /api/produtos/filtros",
				"produtos_categoria":  "GET /api/produtos/categoria/{categoria}",
				"produtos_ativos":     "GET /api/produtos/ativos",
				"produtos_estoque":    "GET /api/produtos/estoque",
				"atualizar_estoque":   "PATCH /api/produtos/{id}/estoque",
				"estatisticas":        "GET /api/produtos/estatisticas",
			},
			"categories": []string{
				"eletronicos", "roupas", "casa", "livros", 
				"esportes", "beleza", "brinquedos", 
				"automotivo", "alimentos", "outros",
			},
		})
	})
	
	// Inicia o servidor
	port := ":8000"
	log.Printf("🚀 Servidor iniciado na porta %s", port)
	log.Printf("📖 Documentação da API: http://localhost%s", port)
	log.Printf("🏥 Health Check: http://localhost%s/health", port)
	log.Printf("📦 Endpoint de Produtos: http://localhost%s/api/produtos", port)
	
	if err := router.Run(port); err != nil {
		log.Fatal("Falha ao iniciar servidor:", err)
	}
}