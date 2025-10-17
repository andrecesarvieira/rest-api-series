package repository

import (
	"github.com/google/uuid"
	"inventario-api/internal/database"
	"inventario-api/internal/models"
)

// ProductRepository define a interface para operações de produto
type ProductRepository interface {
	// CRUD básico
	Create(product *models.Product) error
	GetByID(id uuid.UUID) (*models.Product, error)
	GetAll() ([]*models.Product, error)
	Update(id uuid.UUID, product *models.Product) error
	Delete(id uuid.UUID) error

	// Consultas especializadas
	GetByCategory(category models.ProductCategory) ([]*models.Product, error)
	GetActiveProducts() ([]*models.Product, error)
	GetInStockProducts() ([]*models.Product, error)
	GetFiltered(options database.FilterOptions) ([]*models.Product, int, error)
	
	// Estatísticas
	GetStatistics() (map[string]interface{}, error)
}

// InMemoryProductRepository implementa ProductRepository usando banco em memória
type InMemoryProductRepository struct {
	db *database.InMemoryDatabase
}

// NewInMemoryProductRepository cria uma nova instância do repository
func NewInMemoryProductRepository(db *database.InMemoryDatabase) *InMemoryProductRepository {
	return &InMemoryProductRepository{
		db: db,
	}
}

// Create adiciona um novo produto
func (r *InMemoryProductRepository) Create(product *models.Product) error {
	return r.db.Create(product)
}

// GetByID busca um produto por ID
func (r *InMemoryProductRepository) GetByID(id uuid.UUID) (*models.Product, error) {
	return r.db.GetByID(id)
}

// GetAll retorna todos os produtos
func (r *InMemoryProductRepository) GetAll() ([]*models.Product, error) {
	return r.db.GetAll()
}

// Update atualiza um produto existente
func (r *InMemoryProductRepository) Update(id uuid.UUID, product *models.Product) error {
	return r.db.Update(id, product)
}

// Delete remove um produto
func (r *InMemoryProductRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(id)
}

// GetByCategory retorna produtos de uma categoria específica
func (r *InMemoryProductRepository) GetByCategory(category models.ProductCategory) ([]*models.Product, error) {
	return r.db.GetByCategory(category)
}

// GetActiveProducts retorna apenas produtos ativos
func (r *InMemoryProductRepository) GetActiveProducts() ([]*models.Product, error) {
	return r.db.GetActiveProducts()
}

// GetInStockProducts retorna apenas produtos em estoque
func (r *InMemoryProductRepository) GetInStockProducts() ([]*models.Product, error) {
	return r.db.GetInStockProducts()
}

// GetFiltered retorna produtos filtrados e paginados
func (r *InMemoryProductRepository) GetFiltered(options database.FilterOptions) ([]*models.Product, int, error) {
	return r.db.GetFiltered(options)
}

// GetStatistics retorna estatísticas dos produtos
func (r *InMemoryProductRepository) GetStatistics() (map[string]interface{}, error) {
	return r.db.GetStatistics()
}