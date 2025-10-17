# ⚠️ IMPORTANTE: INSTALAÇÃO DO GO NECESSÁRIA

Este projeto requer **Go 1.21+** para ser executado. Como o Go não está instalado no sistema, siga as instruções abaixo:

## 🚀 Como Instalar e Executar

### 1. Instalar Go
```bash
# Windows (usando winget)
winget install GoLang.Go

# Ou baixe em: https://golang.org/dl/
```

### 2. Verificar Instalação
```bash
go version
```

### 3. Instalar Dependências
```bash
cd 05-go-gin-inventario
go mod tidy
```

### 4. Executar a Aplicação
```bash
go run cmd/api/main.go
```

### 5. Executar Testes
```powershell
./testar-inventario-api.ps1
```

## 📂 Estrutura do Projeto

```
05-go-gin-inventario/
├── cmd/api/main.go              # Ponto de entrada
├── internal/
│   ├── models/product.go        # Modelo de dados
│   ├── dtos/product_dtos.go     # DTOs e validações
│   ├── database/memory_db.go    # Banco em memória
│   ├── repository/              # Repository pattern
│   ├── service/                 # Lógica de negócio
│   ├── handlers/                # Handlers HTTP
│   └── middleware/              # Middlewares
├── go.mod                       # Dependências
└── testar-inventario-api.ps1    # Testes automatizados
```

## 🔧 Funcionalidades Implementadas

✅ **CRUD Completo**: Create, Read, Update, Delete
✅ **Filtros Avançados**: Por categoria, preço, status, nome
✅ **Paginação**: Com metadados completos
✅ **Validações**: Dados de entrada e regras de negócio
✅ **Estatísticas**: Rankings e análises do inventário
✅ **Middlewares**: CORS, Logging, Rate Limiting, Security
✅ **Thread-Safe**: Operações concorrentes seguras
✅ **Repository Pattern**: Arquitetura limpa e testável

## 🌐 Endpoints Disponíveis

- `GET /health` - Health check
- `GET /api/produtos` - Listar produtos
- `POST /api/produtos` - Criar produto
- `GET /api/produtos/{id}` - Buscar por ID
- `PUT /api/produtos/{id}` - Atualizar produto
- `DELETE /api/produtos/{id}` - Deletar produto
- `GET /api/produtos/filtros` - Filtros avançados
- `GET /api/produtos/categoria/{categoria}` - Por categoria
- `GET /api/produtos/ativos` - Apenas ativos
- `GET /api/produtos/estoque` - Em estoque
- `PATCH /api/produtos/{id}/estoque` - Atualizar estoque
- `GET /api/produtos/estatisticas` - Estatísticas

## 📊 Tecnologias

- **Go 1.21+**: Linguagem de programação
- **Gin Framework**: Web framework rápido e minimalista
- **UUID**: Identificadores únicos
- **Validator**: Validação de dados
- **In-Memory Database**: Armazenamento thread-safe

---

**Projeto 5 de 5** da série "REST APIs em Diferentes Linguagens"