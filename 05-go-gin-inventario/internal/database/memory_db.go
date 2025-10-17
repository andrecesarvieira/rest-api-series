package database

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"inventario-api/internal/models"
)

// InMemoryDatabase implementa um banco de dados em memória thread-safe
type InMemoryDatabase struct {
	products map[uuid.UUID]*models.Product
	mutex    sync.RWMutex
	lastID   int
}

// NewInMemoryDatabase cria uma nova instância do banco em memória
func NewInMemoryDatabase() *InMemoryDatabase {
	db := &InMemoryDatabase{
		products: make(map[uuid.UUID]*models.Product),
		lastID:   0,
	}
	
	// Inicializa com dados de exemplo
	db.seedData()
	
	return db
}

// Create adiciona um novo produto ao banco
func (db *InMemoryDatabase) Create(product *models.Product) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	// Gera um novo UUID se não existir
	if product.ID == uuid.Nil {
		product.ID = uuid.New()
	}

	// Verifica se o produto já existe
	if _, exists := db.products[product.ID]; exists {
		return fmt.Errorf("produto com ID %s já existe", product.ID)
	}

	// Define timestamps
	now := time.Now()
	product.DataCriacao = now
	product.DataAtualizacao = now

	// Copia o produto para evitar modificações externas
	productCopy := *product
	db.products[product.ID] = &productCopy

	return nil
}

// GetByID busca um produto por ID
func (db *InMemoryDatabase) GetByID(id uuid.UUID) (*models.Product, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	product, exists := db.products[id]
	if !exists {
		return nil, fmt.Errorf("produto com ID %s não encontrado", id)
	}

	// Retorna uma cópia para evitar modificações externas
	productCopy := *product
	return &productCopy, nil
}

// GetAll retorna todos os produtos
func (db *InMemoryDatabase) GetAll() ([]*models.Product, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	products := make([]*models.Product, 0, len(db.products))
	for _, product := range db.products {
		productCopy := *product
		products = append(products, &productCopy)
	}

	// Ordena por data de criação (mais recentes primeiro)
	sort.Slice(products, func(i, j int) bool {
		return products[i].DataCriacao.After(products[j].DataCriacao)
	})

	return products, nil
}

// FilterOptions define opções de filtro para busca
type FilterOptions struct {
	Categoria     *models.ProductCategory
	PrecoMinimo   *float64
	PrecoMaximo   *float64
	ApenasAtivos  *bool
	ApenasEstoque *bool
	Nome          *string
	Page          int
	Size          int
}

// GetFiltered retorna produtos filtrados e paginados
func (db *InMemoryDatabase) GetFiltered(options FilterOptions) ([]*models.Product, int, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	var filtered []*models.Product

	for _, product := range db.products {
		if db.matchesFilter(product, options) {
			productCopy := *product
			filtered = append(filtered, &productCopy)
		}
	}

	// Ordena por data de criação (mais recentes primeiro)
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].DataCriacao.After(filtered[j].DataCriacao)
	})

	total := len(filtered)

	// Aplica paginação
	if options.Page <= 0 {
		options.Page = 1
	}
	if options.Size <= 0 {
		options.Size = 10
	}

	start := (options.Page - 1) * options.Size
	end := start + options.Size

	if start >= total {
		return []*models.Product{}, total, nil
	}

	if end > total {
		end = total
	}

	return filtered[start:end], total, nil
}

// matchesFilter verifica se um produto atende aos critérios de filtro
func (db *InMemoryDatabase) matchesFilter(product *models.Product, options FilterOptions) bool {
	// Filtro por categoria
	if options.Categoria != nil && product.Categoria != *options.Categoria {
		return false
	}

	// Filtro por preço mínimo
	if options.PrecoMinimo != nil && product.Preco < *options.PrecoMinimo {
		return false
	}

	// Filtro por preço máximo
	if options.PrecoMaximo != nil && product.Preco > *options.PrecoMaximo {
		return false
	}

	// Filtro por status ativo
	if options.ApenasAtivos != nil && *options.ApenasAtivos && !product.Ativo {
		return false
	}

	// Filtro por produtos em estoque
	if options.ApenasEstoque != nil && *options.ApenasEstoque && !product.IsInStock() {
		return false
	}

	// Filtro por nome (busca parcial, case-insensitive)
	if options.Nome != nil && *options.Nome != "" {
		nome := strings.ToLower(*options.Nome)
		produtoNome := strings.ToLower(product.Nome)
		produtoDesc := strings.ToLower(product.Descricao)
		
		if !strings.Contains(produtoNome, nome) && !strings.Contains(produtoDesc, nome) {
			return false
		}
	}

	return true
}

// Update atualiza um produto existente
func (db *InMemoryDatabase) Update(id uuid.UUID, product *models.Product) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	existing, exists := db.products[id]
	if !exists {
		return fmt.Errorf("produto com ID %s não encontrado", id)
	}

	// Preserva campos que não devem ser alterados
	product.ID = id
	product.DataCriacao = existing.DataCriacao
	product.DataAtualizacao = time.Now()

	// Atualiza o produto
	productCopy := *product
	db.products[id] = &productCopy

	return nil
}

// Delete remove um produto do banco
func (db *InMemoryDatabase) Delete(id uuid.UUID) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if _, exists := db.products[id]; !exists {
		return fmt.Errorf("produto com ID %s não encontrado", id)
	}

	delete(db.products, id)
	return nil
}

// GetByCategory retorna produtos de uma categoria específica
func (db *InMemoryDatabase) GetByCategory(category models.ProductCategory) ([]*models.Product, error) {
	options := FilterOptions{
		Categoria: &category,
	}
	products, _, err := db.GetFiltered(options)
	return products, err
}

// GetActiveProducts retorna apenas produtos ativos
func (db *InMemoryDatabase) GetActiveProducts() ([]*models.Product, error) {
	active := true
	options := FilterOptions{
		ApenasAtivos: &active,
	}
	products, _, err := db.GetFiltered(options)
	return products, err
}

// GetInStockProducts retorna apenas produtos em estoque
func (db *InMemoryDatabase) GetInStockProducts() ([]*models.Product, error) {
	inStock := true
	options := FilterOptions{
		ApenasEstoque: &inStock,
	}
	products, _, err := db.GetFiltered(options)
	return products, err
}

// GetStatistics retorna estatísticas dos produtos
func (db *InMemoryDatabase) GetStatistics() (map[string]interface{}, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	stats := make(map[string]interface{})
	categoryStats := make(map[models.ProductCategory]*CategoryStats)
	
	var totalProdutos, produtosAtivos, produtosInativos, produtosEmEstoque, produtosSemEstoque int
	var valorTotal, precoMinimo, precoMaximo float64
	var quantidadeTotal int
	
	precoMinimo = -1 // Inicializa com -1 para detectar primeiro produto
	
	for _, product := range db.products {
		totalProdutos++
		
		if product.Ativo {
			produtosAtivos++
		} else {
			produtosInativos++
		}
		
		if product.IsInStock() {
			produtosEmEstoque++
		} else {
			produtosSemEstoque++
		}
		
		valorTotal += product.Preco * float64(product.Quantidade)
		quantidadeTotal += product.Quantidade
		
		// Atualiza preços mínimo e máximo
		if precoMinimo == -1 || product.Preco < precoMinimo {
			precoMinimo = product.Preco
		}
		if product.Preco > precoMaximo {
			precoMaximo = product.Preco
		}
		
		// Estatísticas por categoria
		if categoryStats[product.Categoria] == nil {
			categoryStats[product.Categoria] = &CategoryStats{
				Categoria: product.Categoria,
			}
		}
		cat := categoryStats[product.Categoria]
		cat.TotalProdutos++
		if product.Ativo {
			cat.ProdutosAtivos++
		}
		cat.ValorTotal += product.Preco * float64(product.Quantidade)
		cat.QuantidadeTotal += product.Quantidade
	}
	
	// Calcula preço médio
	var precoMedio float64
	if totalProdutos > 0 {
		precoMedio = valorTotal / float64(quantidadeTotal)
	}
	
	// Calcula preço médio por categoria
	for _, cat := range categoryStats {
		if cat.QuantidadeTotal > 0 {
			cat.PrecoMedio = cat.ValorTotal / float64(cat.QuantidadeTotal)
		}
	}
	
	stats["total_produtos"] = totalProdutos
	stats["produtos_ativos"] = produtosAtivos
	stats["produtos_inativos"] = produtosInativos
	stats["produtos_em_estoque"] = produtosEmEstoque
	stats["produtos_sem_estoque"] = produtosSemEstoque
	stats["valor_total_inventario"] = valorTotal
	stats["preco_medio"] = precoMedio
	stats["preco_minimo"] = precoMinimo
	stats["preco_maximo"] = precoMaximo
	stats["quantidade_total"] = quantidadeTotal
	stats["por_categoria"] = categoryStats
	
	return stats, nil
}

// CategoryStats representa estatísticas de uma categoria
type CategoryStats struct {
	Categoria       models.ProductCategory `json:"categoria"`
	TotalProdutos   int                    `json:"total_produtos"`
	ProdutosAtivos  int                    `json:"produtos_ativos"`
	ValorTotal      float64                `json:"valor_total"`
	PrecoMedio      float64                `json:"preco_medio"`
	QuantidadeTotal int                    `json:"quantidade_total"`
}

// seedData inicializa o banco com dados de exemplo
func (db *InMemoryDatabase) seedData() {
	produtos := []*models.Product{
		{
			ID:          uuid.New(),
			Nome:        "Smartphone Samsung Galaxy S24",
			Descricao:   "Smartphone com tela de 6.1 polegadas, câmera de 50MP e 5G",
			Preco:       2299.99,
			Quantidade:  25,
			Categoria:   models.CategoryEletronicos,
			Ativo:       true,
		},
		{
			ID:          uuid.New(),
			Nome:        "Notebook Dell Inspiron",
			Descricao:   "Notebook com processador Intel i7, 16GB RAM e SSD 512GB",
			Preco:       3499.99,
			Quantidade:  10,
			Categoria:   models.CategoryEletronicos,
			Ativo:       true,
		},
		{
			ID:          uuid.New(),
			Nome:        "Camiseta Nike Dri-FIT",
			Descricao:   "Camiseta esportiva com tecnologia que remove o suor",
			Preco:       89.99,
			Quantidade:  50,
			Categoria:   models.CategoryRoupas,
			Ativo:       true,
		},
		{
			ID:          uuid.New(),
			Nome:        "Livro Clean Code",
			Descricao:   "Manual de programação limpa por Robert C. Martin",
			Preco:       65.90,
			Quantidade:  30,
			Categoria:   models.CategoryLivros,
			Ativo:       true,
		},
		{
			ID:          uuid.New(),
			Nome:        "Bicicleta Mountain Bike",
			Descricao:   "Bicicleta 21 marchas para trilhas e aventuras",
			Preco:       1299.99,
			Quantidade:  8,
			Categoria:   models.CategoryEsportes,
			Ativo:       true,
		},
		{
			ID:          uuid.New(),
			Nome:        "Perfume Masculino Hugo Boss",
			Descricao:   "Fragrância sofisticada de 100ml",
			Preco:       189.99,
			Quantidade:  0, // Sem estoque
			Categoria:   models.CategoryBeleza,
			Ativo:       false, // Inativo
		},
		{
			ID:          uuid.New(),
			Nome:        "Sofá 3 Lugares",
			Descricao:   "Sofá confortável para sala de estar",
			Preco:       899.99,
			Quantidade:  5,
			Categoria:   models.CategoryCasa,
			Ativo:       true,
		},
	}

	for _, produto := range produtos {
		now := time.Now()
		produto.DataCriacao = now
		produto.DataAtualizacao = now
		db.products[produto.ID] = produto
	}
}