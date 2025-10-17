package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"inventario-api/internal/dtos"
	"inventario-api/internal/models"
	"inventario-api/internal/service"
)

// ProductHandler gerencia os endpoints de produtos
type ProductHandler struct {
	service *service.ProductService
}

// NewProductHandler cria uma nova instância do handler
func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

// CreateProduct godoc
// @Summary Criar novo produto
// @Description Cria um novo produto no inventário
// @Tags produtos
// @Accept json
// @Produce json
// @Param produto body dtos.CreateProductRequest true "Dados do produto"
// @Success 201 {object} dtos.ProductResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 422 {object} dtos.ValidationErrorResponse
// @Router /api/produtos [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req dtos.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.handleValidationError(c, err)
		return
	}

	product, err := h.service.CreateProduct(&req)
	if err != nil {
		h.handleError(c, http.StatusBadRequest, "CREATION_ERROR", err.Error())
		return
	}

	c.JSON(http.StatusCreated, product)
}

// GetProduct godoc
// @Summary Buscar produto por ID
// @Description Retorna um produto específico pelo seu ID
// @Tags produtos
// @Accept json
// @Produce json
// @Param id path string true "ID do produto"
// @Success 200 {object} dtos.ProductResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Router /api/produtos/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, err := h.parseUUID(c.Param("id"))
	if err != nil {
		h.handleError(c, http.StatusBadRequest, "INVALID_ID", "ID do produto inválido")
		return
	}

	product, err := h.service.GetProductByID(id)
	if err != nil {
		h.handleError(c, http.StatusNotFound, "PRODUCT_NOT_FOUND", "Produto não encontrado")
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetAllProducts godoc
// @Summary Listar todos os produtos
// @Description Retorna uma lista de todos os produtos
// @Tags produtos
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ProductListResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /api/produtos [get]
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		h.handleError(c, http.StatusInternalServerError, "FETCH_ERROR", "Erro ao buscar produtos")
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetProductsFiltered godoc
// @Summary Buscar produtos com filtros e paginação
// @Description Retorna produtos filtrados por categoria, preço, status, nome com paginação
// @Tags produtos
// @Accept json
// @Produce json
// @Param categoria query string false "Categoria do produto" Enums(eletronicos,roupas,casa,livros,esportes,beleza,brinquedos,automotivo,alimentos,outros)
// @Param preco_minimo query number false "Preço mínimo"
// @Param preco_maximo query number false "Preço máximo"
// @Param apenas_ativos query boolean false "Apenas produtos ativos"
// @Param apenas_estoque query boolean false "Apenas produtos em estoque"
// @Param nome query string false "Busca por nome ou descrição"
// @Param page query int false "Número da página" default(1)
// @Param size query int false "Itens por página" default(10)
// @Success 200 {object} dtos.ProductListResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /api/produtos/filtros [get]
func (h *ProductHandler) GetProductsFiltered(c *gin.Context) {
	// Parse dos parâmetros de query
	var categoria *models.ProductCategory
	if catStr := c.Query("categoria"); catStr != "" {
		cat := models.ProductCategory(catStr)
		categoria = &cat
	}

	var precoMin, precoMax *float64
	if minStr := c.Query("preco_minimo"); minStr != "" {
		if min, err := strconv.ParseFloat(minStr, 64); err == nil {
			precoMin = &min
		}
	}
	if maxStr := c.Query("preco_maximo"); maxStr != "" {
		if max, err := strconv.ParseFloat(maxStr, 64); err == nil {
			precoMax = &max
		}
	}

	var apenasAtivos, apenasEstoque *bool
	if ativosStr := c.Query("apenas_ativos"); ativosStr != "" {
		if ativos, err := strconv.ParseBool(ativosStr); err == nil {
			apenasAtivos = &ativos
		}
	}
	if estoqueStr := c.Query("apenas_estoque"); estoqueStr != "" {
		if estoque, err := strconv.ParseBool(estoqueStr); err == nil {
			apenasEstoque = &estoque
		}
	}

	var nome *string
	if nomeStr := c.Query("nome"); nomeStr != "" {
		nome = &nomeStr
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	products, err := h.service.GetProductsFiltered(categoria, precoMin, precoMax, apenasAtivos, apenasEstoque, nome, page, size)
	if err != nil {
		h.handleError(c, http.StatusInternalServerError, "FETCH_ERROR", "Erro ao buscar produtos")
		return
	}

	c.JSON(http.StatusOK, products)
}

// UpdateProduct godoc
// @Summary Atualizar produto
// @Description Atualiza um produto existente
// @Tags produtos
// @Accept json
// @Produce json
// @Param id path string true "ID do produto"
// @Param produto body dtos.UpdateProductRequest true "Dados para atualização"
// @Success 200 {object} dtos.ProductResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 422 {object} dtos.ValidationErrorResponse
// @Router /api/produtos/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := h.parseUUID(c.Param("id"))
	if err != nil {
		h.handleError(c, http.StatusBadRequest, "INVALID_ID", "ID do produto inválido")
		return
	}

	var req dtos.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.handleValidationError(c, err)
		return
	}

	product, err := h.service.UpdateProduct(id, &req)
	if err != nil {
		if err.Error() == "produto não encontrado" {
			h.handleError(c, http.StatusNotFound, "PRODUCT_NOT_FOUND", "Produto não encontrado")
		} else {
			h.handleError(c, http.StatusBadRequest, "UPDATE_ERROR", err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary Deletar produto
// @Description Remove um produto do inventário
// @Tags produtos
// @Accept json
// @Produce json
// @Param id path string true "ID do produto"
// @Success 204 "Produto deletado com sucesso"
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Router /api/produtos/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := h.parseUUID(c.Param("id"))
	if err != nil {
		h.handleError(c, http.StatusBadRequest, "INVALID_ID", "ID do produto inválido")
		return
	}

	if err := h.service.DeleteProduct(id); err != nil {
		if err.Error() == "produto não encontrado" {
			h.handleError(c, http.StatusNotFound, "PRODUCT_NOT_FOUND", "Produto não encontrado")
		} else {
			h.handleError(c, http.StatusInternalServerError, "DELETE_ERROR", "Erro ao deletar produto")
		}
		return
	}

	c.Status(http.StatusNoContent)
}

// GetProductsByCategory godoc
// @Summary Buscar produtos por categoria
// @Description Retorna produtos de uma categoria específica
// @Tags produtos
// @Accept json
// @Produce json
// @Param categoria path string true "Categoria" Enums(eletronicos,roupas,casa,livros,esportes,beleza,brinquedos,automotivo,alimentos,outros)
// @Success 200 {object} dtos.ProductListResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /api/produtos/categoria/{categoria} [get]
func (h *ProductHandler) GetProductsByCategory(c *gin.Context) {
	categoria := models.ProductCategory(c.Param("categoria"))
	
	// Valida se a categoria é válida
	if !h.isValidCategory(categoria) {
		h.handleError(c, http.StatusBadRequest, "INVALID_CATEGORY", "Categoria inválida")
		return
	}

	products, err := h.service.GetProductsByCategory(categoria)
	if err != nil {
		h.handleError(c, http.StatusInternalServerError, "FETCH_ERROR", "Erro ao buscar produtos por categoria")
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetActiveProducts godoc
// @Summary Buscar produtos ativos
// @Description Retorna apenas produtos marcados como ativos
// @Tags produtos
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ProductListResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /api/produtos/ativos [get]
func (h *ProductHandler) GetActiveProducts(c *gin.Context) {
	products, err := h.service.GetActiveProducts()
	if err != nil {
		h.handleError(c, http.StatusInternalServerError, "FETCH_ERROR", "Erro ao buscar produtos ativos")
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetInStockProducts godoc
// @Summary Buscar produtos em estoque
// @Description Retorna apenas produtos que estão em estoque (quantidade > 0 e ativo)
// @Tags produtos
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ProductListResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /api/produtos/estoque [get]
func (h *ProductHandler) GetInStockProducts(c *gin.Context) {
	products, err := h.service.GetInStockProducts()
	if err != nil {
		h.handleError(c, http.StatusInternalServerError, "FETCH_ERROR", "Erro ao buscar produtos em estoque")
		return
	}

	c.JSON(http.StatusOK, products)
}

// UpdateStock godoc
// @Summary Atualizar estoque do produto
// @Description Atualiza apenas a quantidade em estoque de um produto
// @Tags produtos
// @Accept json
// @Produce json
// @Param id path string true "ID do produto"
// @Param estoque body dtos.StockUpdateRequest true "Nova quantidade"
// @Success 200 {object} dtos.ProductResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 422 {object} dtos.ValidationErrorResponse
// @Router /api/produtos/{id}/estoque [patch]
func (h *ProductHandler) UpdateStock(c *gin.Context) {
	id, err := h.parseUUID(c.Param("id"))
	if err != nil {
		h.handleError(c, http.StatusBadRequest, "INVALID_ID", "ID do produto inválido")
		return
	}

	var req dtos.StockUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.handleValidationError(c, err)
		return
	}

	product, err := h.service.UpdateStock(id, req.Quantidade)
	if err != nil {
		if err.Error() == "produto não encontrado" {
			h.handleError(c, http.StatusNotFound, "PRODUCT_NOT_FOUND", "Produto não encontrado")
		} else {
			h.handleError(c, http.StatusBadRequest, "UPDATE_ERROR", err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetStatistics godoc
// @Summary Obter estatísticas do inventário
// @Description Retorna estatísticas completas do inventário incluindo valores, categorias e rankings
// @Tags estatísticas
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ProductStatistics
// @Failure 500 {object} dtos.ErrorResponse
// @Router /api/produtos/estatisticas [get]
func (h *ProductHandler) GetStatistics(c *gin.Context) {
	stats, err := h.service.GetStatistics()
	if err != nil {
		h.handleError(c, http.StatusInternalServerError, "STATS_ERROR", "Erro ao buscar estatísticas")
		return
	}

	c.JSON(http.StatusOK, stats)
}

// Métodos auxiliares privados

func (h *ProductHandler) parseUUID(idStr string) (uuid.UUID, error) {
	return uuid.Parse(idStr)
}

func (h *ProductHandler) isValidCategory(categoria models.ProductCategory) bool {
	validCategories := []models.ProductCategory{
		models.CategoryEletronicos,
		models.CategoryRoupas,
		models.CategoryCasa,
		models.CategoryLivros,
		models.CategoryEsportes,
		models.CategoryBeleza,
		models.CategoryBrinquedos,
		models.CategoryAutomotivo,
		models.CategoryAlimentos,
		models.CategoryOutros,
	}

	for _, valid := range validCategories {
		if categoria == valid {
			return true
		}
	}
	return false
}

func (h *ProductHandler) handleError(c *gin.Context, statusCode int, codigo string, mensagem string) {
	c.JSON(statusCode, dtos.ErrorResponse{
		Erro:      mensagem,
		Codigo:    codigo,
		Timestamp: time.Now(),
	})
}

func (h *ProductHandler) handleValidationError(c *gin.Context, err error) {
	c.JSON(http.StatusUnprocessableEntity, dtos.ValidationErrorResponse{
		Erro:      "Dados inválidos",
		Codigo:    "VALIDATION_ERROR",
		Timestamp: time.Now(),
		Campos:    []dtos.ValidationError{
			{
				Campo: "request_body",
				Erro:  err.Error(),
				Tag:   "binding",
			},
		},
	})
}