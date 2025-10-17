package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// ProductCategory representa as categorias disponíveis para produtos
type ProductCategory string

const (
	CategoryEletronicos  ProductCategory = "eletronicos"
	CategoryRoupas      ProductCategory = "roupas"
	CategoryCasa        ProductCategory = "casa"
	CategoryLivros      ProductCategory = "livros"
	CategoryEsportes    ProductCategory = "esportes"
	CategoryBeleza      ProductCategory = "beleza"
	CategoryBrinquedos  ProductCategory = "brinquedos"
	CategoryAutomotivo  ProductCategory = "automotivo"
	CategoryAlimentos   ProductCategory = "alimentos"
	CategoryOutros      ProductCategory = "outros"
)

// Product representa um produto no inventário
type Product struct {
	ID             uuid.UUID       `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Nome           string          `json:"nome" gorm:"not null;size:100" validate:"required,min=2,max=100"`
	Descricao      string          `json:"descricao" gorm:"size:500" validate:"max=500"`
	Preco          float64         `json:"preco" gorm:"not null;check:preco >= 0" validate:"required,min=0"`
	Quantidade     int             `json:"quantidade" gorm:"not null;default:0;check:quantidade >= 0" validate:"min=0"`
	Categoria      ProductCategory `json:"categoria" gorm:"not null;size:50" validate:"required,oneof=eletronicos roupas casa livros esportes beleza brinquedos automotivo alimentos outros"`
	Ativo          bool            `json:"ativo" gorm:"not null;default:true"`
	DataCriacao    time.Time       `json:"data_criacao" gorm:"autoCreateTime"`
	DataAtualizacao time.Time      `json:"data_atualizacao" gorm:"autoUpdateTime"`
}

// TableName especifica o nome da tabela para GORM
func (Product) TableName() string {
	return "produtos"
}

// IsValid verifica se o produto tem dados válidos
func (p *Product) IsValid() bool {
	return p.Nome != "" && 
		   p.Preco >= 0 && 
		   p.Quantidade >= 0 && 
		   p.Categoria != ""
}

// IsInStock verifica se o produto está em estoque
func (p *Product) IsInStock() bool {
	return p.Quantidade > 0 && p.Ativo
}

// GetDisplayPrice retorna o preço formatado para exibição
func (p *Product) GetDisplayPrice() string {
	return fmt.Sprintf("R$ %.2f", p.Preco)
}

// UpdateStock atualiza a quantidade em estoque
func (p *Product) UpdateStock(novaQuantidade int) error {
	if novaQuantidade < 0 {
		return fmt.Errorf("quantidade não pode ser negativa")
	}
	p.Quantidade = novaQuantidade
	p.DataAtualizacao = time.Now()
	return nil
}

// Deactivate desativa o produto
func (p *Product) Deactivate() {
	p.Ativo = false
	p.DataAtualizacao = time.Now()
}

// Activate ativa o produto
func (p *Product) Activate() {
	p.Ativo = true
	p.DataAtualizacao = time.Now()
}