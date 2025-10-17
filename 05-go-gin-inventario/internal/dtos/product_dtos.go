package dtos

import (
	"time"

	"github.com/google/uuid"
	"inventario-api/internal/models"
)

// CreateProductRequest representa a requisição para criar um novo produto
type CreateProductRequest struct {
	Nome       string                  `json:"nome" binding:"required,min=2,max=100" example:"Smartphone Samsung Galaxy"`
	Descricao  string                  `json:"descricao" binding:"max=500" example:"Smartphone com tela de 6.1 polegadas e câmera de 64MP"`
	Preco      float64                 `json:"preco" binding:"required,min=0" example:"1299.99"`
	Quantidade int                     `json:"quantidade" binding:"min=0" example:"50"`
	Categoria  models.ProductCategory  `json:"categoria" binding:"required,oneof=eletronicos roupas casa livros esportes beleza brinquedos automotivo alimentos outros" example:"eletronicos"`
	Ativo      *bool                   `json:"ativo,omitempty" example:"true"`
}

// UpdateProductRequest representa a requisição para atualizar um produto
type UpdateProductRequest struct {
	Nome       *string                 `json:"nome,omitempty" binding:"omitempty,min=2,max=100" example:"Smartphone Samsung Galaxy S24"`
	Descricao  *string                 `json:"descricao,omitempty" binding:"omitempty,max=500" example:"Smartphone com tela de 6.1 polegadas, câmera de 64MP e 5G"`
	Preco      *float64                `json:"preco,omitempty" binding:"omitempty,min=0" example:"1399.99"`
	Quantidade *int                    `json:"quantidade,omitempty" binding:"omitempty,min=0" example:"45"`
	Categoria  *models.ProductCategory `json:"categoria,omitempty" binding:"omitempty,oneof=eletronicos roupas casa livros esportes beleza brinquedos automotivo alimentos outros" example:"eletronicos"`
	Ativo      *bool                   `json:"ativo,omitempty" example:"true"`
}

// ProductResponse representa a resposta de um produto
type ProductResponse struct {
	ID              uuid.UUID               `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Nome            string                  `json:"nome" example:"Smartphone Samsung Galaxy"`
	Descricao       string                  `json:"descricao" example:"Smartphone com tela de 6.1 polegadas e câmera de 64MP"`
	Preco           float64                 `json:"preco" example:"1299.99"`
	PrecoFormatado  string                  `json:"preco_formatado" example:"R$ 1.299,99"`
	Quantidade      int                     `json:"quantidade" example:"50"`
	Categoria       models.ProductCategory  `json:"categoria" example:"eletronicos"`
	Ativo           bool                    `json:"ativo" example:"true"`
	EmEstoque       bool                    `json:"em_estoque" example:"true"`
	DataCriacao     time.Time               `json:"data_criacao" example:"2023-01-15T10:30:00Z"`
	DataAtualizacao time.Time               `json:"data_atualizacao" example:"2023-01-15T10:30:00Z"`
}

// ProductListResponse representa a resposta paginada de produtos
type ProductListResponse struct {
	Produtos   []ProductResponse `json:"produtos"`
	Paginacao  PaginationMeta    `json:"paginacao"`
	Filtros    FilterInfo        `json:"filtros"`
}

// PaginationMeta contém metadados da paginação
type PaginationMeta struct {
	PaginaAtual     int `json:"pagina_atual" example:"1"`
	ItensPorPagina  int `json:"itens_por_pagina" example:"10"`
	TotalItens      int `json:"total_itens" example:"150"`
	TotalPaginas    int `json:"total_paginas" example:"15"`
	TemProxima      bool `json:"tem_proxima" example:"true"`
	TemAnterior     bool `json:"tem_anterior" example:"false"`
}

// FilterInfo contém informações sobre os filtros aplicados
type FilterInfo struct {
	Categoria     *models.ProductCategory `json:"categoria,omitempty" example:"eletronicos"`
	PrecoMinimo   *float64                `json:"preco_minimo,omitempty" example:"100.00"`
	PrecoMaximo   *float64                `json:"preco_maximum,omitempty" example:"2000.00"`
	ApenasAtivos  *bool                   `json:"apenas_ativos,omitempty" example:"true"`
	ApenasEstoque *bool                   `json:"apenas_estoque,omitempty" example:"true"`
	Nome          *string                 `json:"nome,omitempty" example:"samsung"`
}

// StockUpdateRequest representa a requisição para atualizar estoque
type StockUpdateRequest struct {
	Quantidade int `json:"quantidade" binding:"required,min=0" example:"100"`
}

// ProductStatistics representa as estatísticas dos produtos
type ProductStatistics struct {
	TotalProdutos         int                            `json:"total_produtos" example:"150"`
	ProdutosAtivos        int                            `json:"produtos_ativos" example:"140"`
	ProdutosInativos      int                            `json:"produtos_inativos" example:"10"`
	ProdutosEmEstoque     int                            `json:"produtos_em_estoque" example:"130"`
	ProdutosSemEstoque    int                            `json:"produtos_sem_estoque" example:"20"`
	ValorTotalInventario  float64                        `json:"valor_total_inventario" example:"125000.50"`
	PrecoMedio            float64                        `json:"preco_medio" example:"850.25"`
	PrecoMinimo           float64                        `json:"preco_minimo" example:"15.99"`
	PrecoMaximo           float64                        `json:"preco_maximo" example:"5999.99"`
	QuantidadeTotal       int                            `json:"quantidade_total" example:"2500"`
	PorCategoria          []CategoryStatistics           `json:"por_categoria"`
	Top5MaisCaros         []ProductResponse              `json:"top5_mais_caros"`
	Top5MaisBaratos       []ProductResponse              `json:"top5_mais_baratos"`
	Top5MaisEstoque       []ProductResponse              `json:"top5_mais_estoque"`
}

// CategoryStatistics representa estatísticas por categoria
type CategoryStatistics struct {
	Categoria            models.ProductCategory `json:"categoria" example:"eletronicos"`
	TotalProdutos        int                    `json:"total_produtos" example:"25"`
	ProdutosAtivos       int                    `json:"produtos_ativos" example:"23"`
	ValorTotal           float64                `json:"valor_total" example:"45000.00"`
	PrecoMedio           float64                `json:"preco_medio" example:"1800.00"`
	QuantidadeTotal      int                    `json:"quantidade_total" example:"350"`
}

// ErrorResponse representa uma resposta de erro
type ErrorResponse struct {
	Erro      string    `json:"erro" example:"Produto não encontrado"`
	Codigo    string    `json:"codigo" example:"PRODUCT_NOT_FOUND"`
	Timestamp time.Time `json:"timestamp" example:"2023-01-15T10:30:00Z"`
	Detalhes  string    `json:"detalhes,omitempty" example:"Produto com ID 123 não foi encontrado no sistema"`
}

// ValidationErrorResponse representa uma resposta de erro de validação
type ValidationErrorResponse struct {
	Erro      string              `json:"erro" example:"Dados inválidos"`
	Codigo    string              `json:"codigo" example:"VALIDATION_ERROR"`
	Timestamp time.Time           `json:"timestamp" example:"2023-01-15T10:30:00Z"`
	Campos    []ValidationError   `json:"campos"`
}

// ValidationError representa um erro específico de campo
type ValidationError struct {
	Campo   string `json:"campo" example:"preco"`
	Valor   interface{} `json:"valor" example:"-10.5"`
	Erro    string `json:"erro" example:"Preço deve ser maior ou igual a zero"`
	Tag     string `json:"tag" example:"min"`
}

// SuccessResponse representa uma resposta de sucesso genérica
type SuccessResponse struct {
	Sucesso   bool      `json:"sucesso" example:"true"`
	Mensagem  string    `json:"mensagem" example:"Produto criado com sucesso"`
	Timestamp time.Time `json:"timestamp" example:"2023-01-15T10:30:00Z"`
	Dados     interface{} `json:"dados,omitempty"`
}