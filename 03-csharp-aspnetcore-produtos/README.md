# API de Produtos - C# ASP.NET Core

Uma API REST completa para gerenciamento de produtos desenvolvida em C# com ASP.NET Core 9.0.

## 🏗️ Arquitetura

Este projeto segue os princípios de **Clean Architecture** com separação clara de responsabilidades:

```
03-csharp-aspnetcore-produtos/
├── Controllers/          # Endpoints da API REST
├── Models/              # Entidades do domínio
├── DTOs/                # Data Transfer Objects
├── Services/            # Lógica de negócio
├── Repositories/        # Acesso a dados
├── Data/                # Contexto EF e Seeding
├── Exceptions/          # Tratamento de erros
├── Validators/          # Validações FluentValidation
├── Program.cs           # Configuração da aplicação
└── ProdutosApi.csproj   # Dependências do projeto
```

## 🚀 Tecnologias

- **Framework**: ASP.NET Core 9.0
- **Banco de Dados**: Entity Framework In-Memory
- **Validação**: FluentValidation
- **Mapeamento**: AutoMapper
- **Documentação**: Swagger/OpenAPI
- **Logging**: ILogger integrado

## 📦 Dependências

```xml
<PackageReference Include="Microsoft.AspNetCore.OpenApi" Version="9.0.0" />
<PackageReference Include="Microsoft.EntityFrameworkCore.InMemory" Version="9.0.0" />
<PackageReference Include="Swashbuckle.AspNetCore" Version="6.8.1" />
<PackageReference Include="FluentValidation.AspNetCore" Version="11.3.0" />
<PackageReference Include="AutoMapper.Extensions.Microsoft.DependencyInjection" Version="12.0.1" />
```

## 🏃‍♂️ Como Executar

### Pré-requisitos
- .NET 9.0 SDK
- Visual Studio 2022 ou VS Code (opcional)

### Executando o Projeto
```bash
# Navegar para o diretório do projeto
cd 03-csharp-aspnetcore-produtos

# Restaurar dependências
dotnet restore

# Executar a aplicação
dotnet run
```

A API estará disponível em:
- **HTTP**: http://localhost:5000
- **Swagger UI**: http://localhost:5000/swagger

## 🎯 Modelo de Dados

### Produto
```csharp
public class Produto
{
    public int Id { get; set; }
    public string Nome { get; set; }           // 2-100 caracteres
    public string? Descricao { get; set; }     // Máx. 500 caracteres
    public decimal Preco { get; set; }         // > 0
    public string Categoria { get; set; }      // 2-50 caracteres
    public DateTime DataCriacao { get; set; }
    public bool Ativo { get; set; }
}
```

### DTOs
- **ProdutoRequestDTO**: Para criação/atualização
- **ProdutoResponseDTO**: Para respostas da API

## 🌐 Endpoints da API

### Produtos
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/api/produtos` | Lista todos os produtos |
| GET | `/api/produtos/{id}` | Obtém produto por ID |
| POST | `/api/produtos` | Cria novo produto |
| PUT | `/api/produtos/{id}` | Atualiza produto |
| DELETE | `/api/produtos/{id}` | Remove produto |

### Filtros e Consultas
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/api/produtos/categoria/{categoria}` | Filtra por categoria |
| GET | `/api/produtos/preco?minPreco={min}&maxPreco={max}` | Filtra por faixa de preço |
| GET | `/api/produtos/ativos` | Lista apenas produtos ativos |
| GET | `/api/produtos/paginado?page={p}&size={s}` | Paginação |
| GET | `/api/produtos/estatisticas` | Estatísticas consolidadas |

### Paginação Avançada
```http
GET /api/produtos/paginado?page=1&size=10&categoria=Eletrônicos&nome=Samsung&ativo=true
```

**Parâmetros:**
- `page`: Número da página (padrão: 1)
- `size`: Itens por página (padrão: 10, máx: 100)
- `categoria`: Filtro por categoria (opcional)
- `nome`: Busca no nome do produto (opcional)
- `ativo`: Filtro por status (opcional)

## 📝 Exemplos de Uso

### Criar Produto
```bash
curl -X POST "http://localhost:5000/api/produtos" \\
  -H "Content-Type: application/json" \\
  -d '{
    "nome": "Notebook Gamer",
    "descricao": "Notebook para jogos com placa dedicada",
    "preco": 3499.90,
    "categoria": "Informática",
    "ativo": true
  }'
```

### Atualizar Produto
```bash
curl -X PUT "http://localhost:5000/api/produtos/1" \\
  -H "Content-Type: application/json" \\
  -d '{
    "nome": "Smartphone Premium",
    "descricao": "Versão atualizada com mais memória",
    "preco": 1599.90,
    "categoria": "Eletrônicos",
    "ativo": true
  }'
```

### Buscar com Filtros
```bash
# Produtos da categoria Eletrônicos
curl "http://localhost:5000/api/produtos/categoria/Eletrônicos"

# Produtos entre R$ 1000 e R$ 2000
curl "http://localhost:5000/api/produtos/preco?minPreco=1000&maxPreco=2000"

# Paginação com filtros
curl "http://localhost:5000/api/produtos/paginado?page=1&size=5&categoria=Informática"
```

## 🧪 Testes Automatizados

Execute o script de testes incluído:

```powershell
.\\testar-produtos-api.ps1
```

**Cenários testados:**
- ✅ CRUD completo (Create, Read, Update, Delete)
- ✅ Filtros por categoria, preço e status
- ✅ Paginação com filtros combinados
- ✅ Validações de entrada
- ✅ Tratamento de erros (404, 400)
- ✅ Estatísticas e agregações

## 🛡️ Validações

### Regras de Negócio
- Nome do produto deve ser único
- Preço deve ser positivo
- Categoria obrigatória com mínimo 2 caracteres

### Validações de Entrada (FluentValidation)
```csharp
RuleFor(x => x.Nome)
    .NotEmpty().WithMessage("Nome é obrigatório")
    .Length(2, 100).WithMessage("Nome deve ter entre 2 e 100 caracteres");

RuleFor(x => x.Preco)
    .GreaterThan(0).WithMessage("Preço deve ser maior que zero");
```

## 🚨 Tratamento de Erros

O middleware global captura e trata exceções:

- **400 Bad Request**: Dados de entrada inválidos
- **404 Not Found**: Produto não encontrado
- **422 Unprocessable Entity**: Erro de validação de negócio
- **500 Internal Server Error**: Erros não tratados

Exemplo de resposta de erro:
```json
{
  "statusCode": 400,
  "message": "Nome é obrigatório",
  "details": "Requisição inválida",
  "timestamp": "2025-10-17T16:38:54.870Z"
}
```

## 📊 Estatísticas

O endpoint `/api/produtos/estatisticas` retorna:
```json
{
  "totalProdutos": 5,
  "produtosAtivos": 5,
  "produtosInativos": 0,
  "precoMedio": 1649.9,
  "precoMinimo": 349.9,
  "precoMaximo": 2499.9,
  "categorias": [
    {
      "categoria": "Eletrônicos",
      "quantidade": 2,
      "precoMedio": 1749.9
    }
  ],
  "produtoMaisCaro": { /* dados do produto */ },
  "produtoMaisBarato": { /* dados do produto */ }
}
```

## 🔧 Configuração

### Program.cs
```csharp
// Entity Framework In-Memory
builder.Services.AddDbContext<ProdutoContext>(options =>
    options.UseInMemoryDatabase("ProdutosDb"));

// Injeção de Dependências
builder.Services.AddScoped<IProdutoRepository, ProdutoRepository>();
builder.Services.AddScoped<IProdutoService, ProdutoService>();

// AutoMapper e FluentValidation
builder.Services.AddAutoMapper(typeof(Program));
builder.Services.AddFluentValidationAutoValidation();
```

## 📚 Dados de Exemplo

A aplicação inicia com 5 produtos pré-cadastrados:
- Smartphone Samsung Galaxy (Eletrônicos)
- Notebook Dell Inspiron (Informática)
- Fone de Ouvido Sony (Áudio)
- Câmera Canon EOS (Fotografia)
- Smart TV LG 55" (Eletrônicos)

## 🎨 Swagger/OpenAPI

Acesse a documentação interativa em `/swagger` quando a aplicação estiver rodando.

## 📈 Status do Projeto

✅ **Projeto Completo e Funcional**
- Todas as funcionalidades implementadas
- Testes automatizados passando
- Documentação completa
- Pronto para produção

---

**Projeto 3 de 5** da série "REST APIs em Diferentes Linguagens"
- ✅ Projeto 1: Java + Spring Boot
- ✅ Projeto 2: Node.js + Express + TypeScript  
- ✅ **Projeto 3: C# + ASP.NET Core** ← Atual
- 🚧 Projeto 4: Python + FastAPI
- 🚧 Projeto 5: Go + Gin