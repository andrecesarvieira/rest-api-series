package service

import (
	"fmt"
	"sort"
	"strings"

	"github.com/google/uuid"
	"inventario-api/internal/database"
	"inventario-api/internal/dtos"
	"inventario-api/internal/models"
	"inventario-api/internal/repository"
)

// ProductService implementa a lógica de negócio para produtos
type ProductService struct {
	repo repository.ProductRepository
}

// NewProductService cria uma nova instância do service
func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

// CreateProduct cria um novo produto com validações de negócio
func (s *ProductService) CreateProduct(req *dtos.CreateProductRequest) (*dtos.ProductResponse, error) {
	// Validações de negócio
	if err := s.validateCreateRequest(req); err != nil {
		return nil, err
	}

	// Cria o modelo
	product := &models.Product{
		Nome:       strings.TrimSpace(req.Nome),
		Descricao:  strings.TrimSpace(req.Descricao),
		Preco:      req.Preco,
		Quantidade: req.Quantidade,
		Categoria:  req.Categoria,
		Ativo:      true, // Padrão é ativo
	}

	// Se ativo foi especificado na requisição, usa o valor
	if req.Ativo != nil {
		product.Ativo = *req.Ativo
	}

	// Salva no repositório
	if err := s.repo.Create(product); err != nil {
		return nil, fmt.Errorf("erro ao criar produto: %w", err)
	}

	// Retorna o produto criado
	return s.toProductResponse(product), nil
}

// GetProductByID busca um produto por ID
func (s *ProductService) GetProductByID(id uuid.UUID) (*dtos.ProductResponse, error) {
	product, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("produto não encontrado: %w", err)
	}

	return s.toProductResponse(product), nil
}

// GetAllProducts retorna todos os produtos
func (s *ProductService) GetAllProducts() (*dtos.ProductListResponse, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar produtos: %w", err)
	}

	responses := make([]dtos.ProductResponse, len(products))
	for i, product := range products {
		responses[i] = *s.toProductResponse(product)
	}

	return &dtos.ProductListResponse{
		Produtos: responses,
		Paginacao: dtos.PaginationMeta{
			PaginaAtual:    1,
			ItensPorPagina: len(responses),
			TotalItens:     len(responses),
			TotalPaginas:   1,
			TemProxima:     false,
			TemAnterior:    false,
		},
	}, nil
}

// GetProductsFiltered retorna produtos filtrados e paginados
func (s *ProductService) GetProductsFiltered(
	categoria *models.ProductCategory,
	precoMin, precoMax *float64,
	apenasAtivos, apenasEstoque *bool,
	nome *string,
	page, size int,
) (*dtos.ProductListResponse, error) {

	// Validações dos parâmetros
	if page <= 0 {
		page = 1
	}
	if size <= 0 || size > 100 {
		size = 10
	}

	options := database.FilterOptions{
		Categoria:     categoria,
		PrecoMinimo:   precoMin,
		PrecoMaximo:   precoMax,
		ApenasAtivos:  apenasAtivos,
		ApenasEstoque: apenasEstoque,
		Nome:          nome,
		Page:          page,
		Size:          size,
	}

	products, total, err := s.repo.GetFiltered(options)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar produtos filtrados: %w", err)
	}

	responses := make([]dtos.ProductResponse, len(products))
	for i, product := range products {
		responses[i] = *s.toProductResponse(product)
	}

	totalPages := (total + size - 1) / size
	hasNext := page < totalPages
	hasPrevious := page > 1

	return &dtos.ProductListResponse{
		Produtos: responses,
		Paginacao: dtos.PaginationMeta{
			PaginaAtual:    page,
			ItensPorPagina: size,
			TotalItens:     total,
			TotalPaginas:   totalPages,
			TemProxima:     hasNext,
			TemAnterior:    hasPrevious,
		},
		Filtros: dtos.FilterInfo{
			Categoria:     categoria,
			PrecoMinimo:   precoMin,
			PrecoMaximo:   precoMax,
			ApenasAtivos:  apenasAtivos,
			ApenasEstoque: apenasEstoque,
			Nome:          nome,
		},
	}, nil
}

// UpdateProduct atualiza um produto existente
func (s *ProductService) UpdateProduct(id uuid.UUID, req *dtos.UpdateProductRequest) (*dtos.ProductResponse, error) {
	// Busca o produto existente
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("produto não encontrado: %w", err)
	}

	// Aplica as atualizações
	updated := *existing
	
	if req.Nome != nil {
		if err := s.validateNome(*req.Nome); err != nil {
			return nil, err
		}
		updated.Nome = strings.TrimSpace(*req.Nome)
	}
	
	if req.Descricao != nil {
		if err := s.validateDescricao(*req.Descricao); err != nil {
			return nil, err
		}
		updated.Descricao = strings.TrimSpace(*req.Descricao)
	}
	
	if req.Preco != nil {
		if err := s.validatePreco(*req.Preco); err != nil {
			return nil, err
		}
		updated.Preco = *req.Preco
	}
	
	if req.Quantidade != nil {
		if err := s.validateQuantidade(*req.Quantidade); err != nil {
			return nil, err
		}
		updated.Quantidade = *req.Quantidade
	}
	
	if req.Categoria != nil {
		updated.Categoria = *req.Categoria
	}
	
	if req.Ativo != nil {
		updated.Ativo = *req.Ativo
	}

	// Salva as alterações
	if err := s.repo.Update(id, &updated); err != nil {
		return nil, fmt.Errorf("erro ao atualizar produto: %w", err)
	}

	return s.toProductResponse(&updated), nil
}

// DeleteProduct remove um produto
func (s *ProductService) DeleteProduct(id uuid.UUID) error {
	// Verifica se o produto existe
	_, err := s.repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("produto não encontrado: %w", err)
	}

	// Remove o produto
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("erro ao deletar produto: %w", err)
	}

	return nil
}

// GetProductsByCategory retorna produtos de uma categoria específica
func (s *ProductService) GetProductsByCategory(category models.ProductCategory) (*dtos.ProductListResponse, error) {
	products, err := s.repo.GetByCategory(category)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar produtos por categoria: %w", err)
	}

	responses := make([]dtos.ProductResponse, len(products))
	for i, product := range products {
		responses[i] = *s.toProductResponse(product)
	}

	return &dtos.ProductListResponse{
		Produtos: responses,
		Paginacao: dtos.PaginationMeta{
			PaginaAtual:    1,
			ItensPorPagina: len(responses),
			TotalItens:     len(responses),
			TotalPaginas:   1,
			TemProxima:     false,
			TemAnterior:    false,
		},
		Filtros: dtos.FilterInfo{
			Categoria: &category,
		},
	}, nil
}

// GetActiveProducts retorna apenas produtos ativos
func (s *ProductService) GetActiveProducts() (*dtos.ProductListResponse, error) {
	products, err := s.repo.GetActiveProducts()
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar produtos ativos: %w", err)
	}

	responses := make([]dtos.ProductResponse, len(products))
	for i, product := range products {
		responses[i] = *s.toProductResponse(product)
	}

	ativo := true
	return &dtos.ProductListResponse{
		Produtos: responses,
		Paginacao: dtos.PaginationMeta{
			PaginaAtual:    1,
			ItensPorPagina: len(responses),
			TotalItens:     len(responses),
			TotalPaginas:   1,
			TemProxima:     false,
			TemAnterior:    false,
		},
		Filtros: dtos.FilterInfo{
			ApenasAtivos: &ativo,
		},
	}, nil
}

// GetInStockProducts retorna apenas produtos em estoque
func (s *ProductService) GetInStockProducts() (*dtos.ProductListResponse, error) {
	products, err := s.repo.GetInStockProducts()
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar produtos em estoque: %w", err)
	}

	responses := make([]dtos.ProductResponse, len(products))
	for i, product := range products {
		responses[i] = *s.toProductResponse(product)
	}

	emEstoque := true
	return &dtos.ProductListResponse{
		Produtos: responses,
		Paginacao: dtos.PaginationMeta{
			PaginaAtual:    1,
			ItensPorPagina: len(responses),
			TotalItens:     len(responses),
			TotalPaginas:   1,
			TemProxima:     false,
			TemAnterior:    false,
		},
		Filtros: dtos.FilterInfo{
			ApenasEstoque: &emEstoque,
		},
	}, nil
}

// UpdateStock atualiza apenas a quantidade de um produto
func (s *ProductService) UpdateStock(id uuid.UUID, novaQuantidade int) (*dtos.ProductResponse, error) {
	if err := s.validateQuantidade(novaQuantidade); err != nil {
		return nil, err
	}

	// Busca o produto existente
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("produto não encontrado: %w", err)
	}

	// Atualiza apenas a quantidade
	updated := *existing
	updated.Quantidade = novaQuantidade

	// Salva as alterações
	if err := s.repo.Update(id, &updated); err != nil {
		return nil, fmt.Errorf("erro ao atualizar estoque: %w", err)
	}

	return s.toProductResponse(&updated), nil
}

// GetStatistics retorna estatísticas dos produtos
func (s *ProductService) GetStatistics() (*dtos.ProductStatistics, error) {
	stats, err := s.repo.GetStatistics()
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar estatísticas: %w", err)
	}

	// Busca produtos para rankings
	allProducts, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar produtos para estatísticas: %w", err)
	}

	// Ordena para top 5
	sort.Slice(allProducts, func(i, j int) bool {
		return allProducts[i].Preco > allProducts[j].Preco
	})
	top5Caros := s.getTop5(allProducts)

	sort.Slice(allProducts, func(i, j int) bool {
		return allProducts[i].Preco < allProducts[j].Preco
	})
	top5Baratos := s.getTop5(allProducts)

	sort.Slice(allProducts, func(i, j int) bool {
		return allProducts[i].Quantidade > allProducts[j].Quantidade
	})
	top5Estoque := s.getTop5(allProducts)

	// Converte estatísticas por categoria
	categoryStatsRaw := stats["por_categoria"].(map[models.ProductCategory]*database.CategoryStats)
	categoryStats := make([]dtos.CategoryStatistics, 0, len(categoryStatsRaw))
	for _, catStat := range categoryStatsRaw {
		categoryStats = append(categoryStats, dtos.CategoryStatistics{
			Categoria:       catStat.Categoria,
			TotalProdutos:   catStat.TotalProdutos,
			ProdutosAtivos:  catStat.ProdutosAtivos,
			ValorTotal:      catStat.ValorTotal,
			PrecoMedio:      catStat.PrecoMedio,
			QuantidadeTotal: catStat.QuantidadeTotal,
		})
	}

	return &dtos.ProductStatistics{
		TotalProdutos:        stats["total_produtos"].(int),
		ProdutosAtivos:       stats["produtos_ativos"].(int),
		ProdutosInativos:     stats["produtos_inativos"].(int),
		ProdutosEmEstoque:    stats["produtos_em_estoque"].(int),
		ProdutosSemEstoque:   stats["produtos_sem_estoque"].(int),
		ValorTotalInventario: stats["valor_total_inventario"].(float64),
		PrecoMedio:           stats["preco_medio"].(float64),
		PrecoMinimo:          stats["preco_minimo"].(float64),
		PrecoMaximo:          stats["preco_maximo"].(float64),
		QuantidadeTotal:      stats["quantidade_total"].(int),
		PorCategoria:         categoryStats,
		Top5MaisCaros:        top5Caros,
		Top5MaisBaratos:      top5Baratos,
		Top5MaisEstoque:      top5Estoque,
	}, nil
}

// Métodos auxiliares privados

func (s *ProductService) toProductResponse(product *models.Product) *dtos.ProductResponse {
	return &dtos.ProductResponse{
		ID:              product.ID,
		Nome:            product.Nome,
		Descricao:       product.Descricao,
		Preco:           product.Preco,
		PrecoFormatado:  product.GetDisplayPrice(),
		Quantidade:      product.Quantidade,
		Categoria:       product.Categoria,
		Ativo:           product.Ativo,
		EmEstoque:       product.IsInStock(),
		DataCriacao:     product.DataCriacao,
		DataAtualizacao: product.DataAtualizacao,
	}
}

func (s *ProductService) getTop5(products []*models.Product) []dtos.ProductResponse {
	max := 5
	if len(products) < max {
		max = len(products)
	}

	result := make([]dtos.ProductResponse, max)
	for i := 0; i < max; i++ {
		result[i] = *s.toProductResponse(products[i])
	}

	return result
}

// Validações de negócio

func (s *ProductService) validateCreateRequest(req *dtos.CreateProductRequest) error {
	if err := s.validateNome(req.Nome); err != nil {
		return err
	}
	if err := s.validateDescricao(req.Descricao); err != nil {
		return err
	}
	if err := s.validatePreco(req.Preco); err != nil {
		return err
	}
	if err := s.validateQuantidade(req.Quantidade); err != nil {
		return err
	}
	return nil
}

func (s *ProductService) validateNome(nome string) error {
	nome = strings.TrimSpace(nome)
	if len(nome) < 2 {
		return fmt.Errorf("nome deve ter pelo menos 2 caracteres")
	}
	if len(nome) > 100 {
		return fmt.Errorf("nome deve ter no máximo 100 caracteres")
	}
	return nil
}

func (s *ProductService) validateDescricao(descricao string) error {
	if len(descricao) > 500 {
		return fmt.Errorf("descrição deve ter no máximo 500 caracteres")
	}
	return nil
}

func (s *ProductService) validatePreco(preco float64) error {
	if preco < 0 {
		return fmt.Errorf("preço deve ser maior ou igual a zero")
	}
	return nil
}

func (s *ProductService) validateQuantidade(quantidade int) error {
	if quantidade < 0 {
		return fmt.Errorf("quantidade deve ser maior ou igual a zero")
	}
	return nil
}