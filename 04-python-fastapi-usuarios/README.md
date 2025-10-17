# API de Usuários - Python FastAPI

Uma API REST completa para gerenciamento de usuários desenvolvida em Python com FastAPI.

## 🏗️ Arquitetura

Este projeto segue os princípios de **Clean Architecture** com separação clara de responsabilidades:

```
04-python-fastapi-usuarios/
├── app/
│   ├── models/              # Modelos de dados
│   ├── schemas/             # Schemas Pydantic (validação)
│   ├── services/            # Lógica de negócio
│   ├── routers/             # Endpoints da API
│   ├── database/            # Banco de dados em memória
│   └── core/                # Middleware e configurações
├── main.py                  # Ponto de entrada da aplicação
├── requirements.txt         # Dependências Python
└── testar-usuarios-api.ps1  # Script de testes automatizados
```

## 🚀 Tecnologias

- **Framework**: FastAPI (Python)
- **Validação**: Pydantic com validadores customizados
- **Banco de Dados**: Sistema em memória (InMemoryDatabase)
- **Documentação**: Swagger/OpenAPI automático
- **Server**: Uvicorn ASGI

## 📦 Dependências

```txt
fastapi
uvicorn[standard]
email-validator
```

## 🏃‍♂️ Como Executar

### Pré-requisitos
- Python 3.11+ (testado com Python 3.14)
- pip (gerenciador de pacotes Python)

### Instalação e Execução
```bash
# Navegar para o diretório do projeto
cd 04-python-fastapi-usuarios

# Instalar dependências
pip install -r requirements.txt

# Executar a aplicação
python main.py

# Ou usando uvicorn diretamente
uvicorn main:app --host 0.0.0.0 --port 8000 --reload
```

A API estará disponível em:
- **HTTP**: http://localhost:8000
- **Swagger UI**: http://localhost:8000/docs
- **ReDoc**: http://localhost:8000/redoc

## 🎯 Modelo de Dados

### Usuário
```python
class User:
    id: int                    # ID único (auto-incremento)
    nome: str                  # Nome completo (2-100 caracteres)
    email: str                 # Email único e válido
    idade: int                 # Idade (0-120 anos)
    ativo: bool                # Status ativo (padrão: True)
    data_criacao: datetime     # Data/hora de criação (automático)
    perfil: UserProfile        # Perfil do usuário
```

### Perfis de Usuário
```python
class UserProfile(Enum):
    USER = "user"              # Usuário padrão
    ADMIN = "admin"            # Administrador
    MODERATOR = "moderator"    # Moderador
```

## 🌐 Endpoints da API

### Usuários Básicos
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/api/usuarios/` | Lista todos os usuários |
| GET | `/api/usuarios/{id}` | Obtém usuário por ID |
| POST | `/api/usuarios/` | Cria novo usuário |
| PUT | `/api/usuarios/{id}` | Atualiza usuário |
| DELETE | `/api/usuarios/{id}` | Remove usuário |

### Filtros e Consultas Especializadas
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/api/usuarios/perfil/{profile}` | Filtra por perfil |
| GET | `/api/usuarios/status/ativos` | Lista apenas usuários ativos |
| GET | `/api/usuarios/idade/faixa?min_idade={min}&max_idade={max}` | Filtra por faixa etária |
| GET | `/api/usuarios/paginado/lista` | Paginação com filtros avançados |
| GET | `/api/usuarios/estatisticas/dados` | Estatísticas consolidadas |

### Paginação Avançada
```http
GET /api/usuarios/paginado/lista?page=1&size=10&nome=João&email=silva&perfil=user&ativo=true
```

**Parâmetros de Paginação:**
- `page`: Número da página (padrão: 1, mínimo: 1)
- `size`: Itens por página (padrão: 10, máximo: 100)
- `nome`: Busca parcial no nome (case-insensitive)
- `email`: Busca parcial no email (case-insensitive)
- `perfil`: Filtro por perfil (user, admin, moderator)
- `ativo`: Filtro por status ativo (true/false)

## 📝 Exemplos de Uso

### Criar Usuário
```bash
curl -X POST "http://localhost:8000/api/usuarios/" \\
  -H "Content-Type: application/json" \\
  -d '{
    "nome": "Maria Silva",
    "email": "maria.silva@exemplo.com",
    "idade": 28,
    "perfil": "user",
    "ativo": true
  }'
```

### Atualizar Usuário
```bash
curl -X PUT "http://localhost:8000/api/usuarios/1" \\
  -H "Content-Type: application/json" \\
  -d '{
    "nome": "Maria Silva Santos",
    "idade": 29,
    "perfil": "moderator"
  }'
```

### Buscar com Filtros
```bash
# Usuários do perfil admin
curl "http://localhost:8000/api/usuarios/perfil/admin"

# Usuários entre 25 e 35 anos
curl "http://localhost:8000/api/usuarios/idade/faixa?min_idade=25&max_idade=35"

# Paginação com nome "João"
curl "http://localhost:8000/api/usuarios/paginado/lista?page=1&size=5&nome=João"

# Usuários ativos do perfil user
curl "http://localhost:8000/api/usuarios/paginado/lista?perfil=user&ativo=true"
```

## 🧪 Testes Automatizados

Execute o script de testes incluído:

```powershell
.\\testar-usuarios-api.ps1
```

**Cenários testados (14 testes):**
- ✅ CRUD completo (Create, Read, Update, Delete)
- ✅ Filtros por perfil, idade, status ativo
- ✅ Paginação com filtros múltiplos
- ✅ Busca parcial por nome e email
- ✅ Validações de entrada (email único, idade válida)
- ✅ Tratamento de erros (404, 409, 422)
- ✅ Estatísticas dinâmicas

## 🛡️ Validações

### Schemas Pydantic
```python
class UserCreate(BaseModel):
    nome: str = Field(..., min_length=2, max_length=100)
    email: EmailStr = Field(...)
    idade: int = Field(..., ge=0, le=120)
    perfil: UserProfile = Field(default=UserProfile.USER)
    ativo: bool = Field(default=True)
```

### Regras de Negócio
- **Email único**: Não permite emails duplicados
- **Nome obrigatório**: Mínimo 2, máximo 100 caracteres
- **Idade válida**: Entre 0 e 120 anos
- **Email válido**: Validação automática via EmailStr

### Validadores Customizados
```python
@field_validator('nome')
def validate_nome(cls, v):
    if not v.strip():
        raise ValueError('Nome não pode estar vazio')
    return v.strip()

@field_validator('email')
def validate_email(cls, v):
    return v.lower()  # Normaliza para minúsculas
```

## 🚨 Tratamento de Erros

A API retorna códigos HTTP apropriados:

- **200 OK**: Sucesso na operação
- **201 Created**: Usuário criado com sucesso
- **204 No Content**: Usuário removido com sucesso
- **400 Bad Request**: Dados de entrada inválidos
- **404 Not Found**: Usuário não encontrado
- **409 Conflict**: Email já existe (duplicado)
- **422 Unprocessable Entity**: Erro de validação

Exemplo de resposta de erro:
```json
{
  "detail": "Email 'joao@email.com' já está em uso"
}
```

## 📊 Estatísticas

O endpoint `/api/usuarios/estatisticas/dados` retorna:
```json
{
  "total_usuarios": 6,
  "usuarios_ativos": 5,
  "usuarios_inativos": 1,
  "idade_media": 31.5,
  "idade_minima": 25,
  "idade_maxima": 40,
  "perfis": [
    {
      "perfil": "admin",
      "quantidade": 2,
      "idade_media": 34.0
    },
    {
      "perfil": "user", 
      "quantidade": 3,
      "idade_media": 32.0
    }
  ],
  "usuario_mais_velho": {
    "id": 7,
    "nome": "Admin Teste",
    "email": "admin@teste.com",
    "idade": 40
  },
  "usuario_mais_novo": {
    "id": 3,
    "nome": "Pedro Oliveira", 
    "email": "pedro.oliveira@email.com",
    "idade": 25
  }
}
```

## 🔧 Arquitetura Interna

### Banco de Dados em Memória
```python
class InMemoryDatabase:
    def __init__(self):
        self.users: Dict[int, User] = {}
        self.next_id = 1
```

### Dependency Injection
```python
def get_user_service(db: InMemoryDatabase = Depends(get_database)):
    return UserService(db)
```

### Middleware Configurado
- **CORS**: Permite requisições de qualquer origem
- **Logging**: Log detalhado das requisições
- **Exception Handler**: Tratamento global de erros

## 📚 Dados de Exemplo

A aplicação inicia com 5 usuários pré-cadastrados:
- João Silva (Admin, 28 anos)
- Maria Santos (User, 32 anos)
- Pedro Oliveira (Moderator, 25 anos)
- Ana Costa (User, 29 anos, inativa)
- Carlos Lima (User, 35 anos)

## 🎨 FastAPI Features

### Documentação Automática
- **Swagger UI**: Interface interativa para testar endpoints
- **ReDoc**: Documentação elegante da API
- **OpenAPI Schema**: Geração automática da especificação

### Validação Automática
- **Request/Response Models**: Validação via Pydantic
- **Type Hints**: Documentação automática de tipos
- **Error Messages**: Mensagens descritivas de erro

### Performance
- **ASGI**: Suporte assíncrono nativo
- **Uvicorn**: Servidor de alta performance
- **Fast JSON**: Serialização otimizada

## 📈 Status do Projeto

✅ **Projeto Completo e Funcional**
- Todas as funcionalidades implementadas
- 14 cenários de teste automatizados passando
- Documentação completa e interativa
- Pronto para produção

---

**Projeto 4 de 5** da série "REST APIs em Diferentes Linguagens"
- ✅ Projeto 1: Java + Spring Boot
- ✅ Projeto 2: Node.js + Express + TypeScript  
- ✅ Projeto 3: C# + ASP.NET Core
- ✅ **Projeto 4: Python + FastAPI** ← Atual
- 🚧 Projeto 5: Go + Gin