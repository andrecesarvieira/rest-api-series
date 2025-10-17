# 📋 Tarefas API - Node.js + Express

[![Node.js](https://img.shields.io/badge/Node.js-18+-brightgreen)](https://nodejs.org)
[![Express](https://img.shields.io/badge/Express-4.18+-blue)](https://expressjs.com)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.2+-blue)](https://www.typescriptlang.org)
[![Zod](https://img.shields.io/badge/Zod-3.22+-purple)](https://zod.dev)
[![License](https://img.shields.io/badge/License-MIT-yellow)](../../LICENSE)

API REST para gerenciamento de tarefas desenvolvida com **Node.js**, **Express** e **TypeScript**, implementando operações CRUD completas com validação robusta e funcionalidades avançadas.

> **Parte 2 de 5** da série [REST API Multi-linguagem](../)

---

## 📋 Índice

- [Sobre o Projeto](#sobre-o-projeto)
- [Tecnologias](#tecnologias)
- [Arquitetura](#arquitetura)
- [Endpoints da API](#endpoints-da-api)
- [Como Executar](#como-executar)
- [Exemplos de Uso](#exemplos-de-uso)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Conceitos Aplicados](#conceitos-aplicados)
- [Melhorias Futuras](#melhorias-futuras)

---

## 📖 Sobre o Projeto

API REST que permite gerenciar tarefas com as seguintes funcionalidades:

✅ Listar todas as tarefas com filtros e paginação  
✅ Buscar tarefa por ID  
✅ Criar nova tarefa  
✅ Atualizar informações de tarefa  
✅ Remover tarefa  
✅ Marcar tarefa como concluída/reativar  
✅ Filtrar por status, prioridade e categoria  
✅ Estatísticas e dashboard de métricas  
✅ Validação robusta com Zod  
✅ Tratamento centralizado de exceções  
✅ Armazenamento em memória com dados de exemplo  
✅ Segurança com Helmet e CORS  

---

## 🛠️ Tecnologias

| Tecnologia | Versão | Finalidade |
|------------|--------|------------|
| Node.js | 18+ | Runtime JavaScript |
| Express | 4.18+ | Framework web |
| TypeScript | 5.2+ | Tipagem estática |
| Zod | 3.22+ | Validação de schemas |
| UUID | 9.0+ | Geração de IDs únicos |
| Helmet | 7.1+ | Segurança HTTP headers |
| CORS | 2.8+ | Cross-Origin Resource Sharing |
| Morgan | 1.10+ | Logging de requisições |
| Nodemon | 3.0+ | Auto-reload em desenvolvimento |
| ts-node | 10.9+ | Execução TypeScript |

---

## 🏗️ Arquitetura

O projeto segue o padrão **Layered Architecture** com separação clara de responsabilidades:

```
┌─────────────────────────────────────────┐
│          Controller Layer               │  ← Endpoints REST
│  (Recebe requisições HTTP)              │
└─────────────────┬───────────────────────┘
                  │
                  ↓
┌─────────────────────────────────────────┐
│         Middleware Layer                │  ← Validação, Erros, Logs
│  (Zod, Error Handler, Morgan)           │
└─────────────────┬───────────────────────┘
                  │
                  ↓
┌─────────────────────────────────────────┐
│          Service Layer                  │  ← Lógica de negócio
│  (Regras e validações)                  │
└─────────────────┬───────────────────────┘
                  │
                  ↓
┌─────────────────────────────────────────┐
│         Storage Layer                   │  ← Armazenamento de dados
│  (In-Memory DataStore)                  │
└─────────────────────────────────────────┘
```

### Fluxo de Dados

```
Cliente → Routes → Middleware → Controller → Service → DataStore
```

---

## 🔗 Endpoints da API

**Base URL**: `http://localhost:3000/api`

### Health Check

| Método | Endpoint | Descrição | Response |
|--------|----------|-----------|----------|
| `GET` | `/health` | Status da API | 200 OK |

### Tarefas

| Método | Endpoint | Descrição | Request Body | Response |
|--------|----------|-----------|--------------|----------|
| `GET` | `/tarefas` | Lista todas as tarefas | - | 200 OK |
| `GET` | `/tarefas/{id}` | Busca tarefa por ID | - | 200 OK / 404 Not Found |
| `POST` | `/tarefas` | Cria nova tarefa | JSON | 201 Created / 400 Bad Request |
| `PUT` | `/tarefas/{id}` | Atualiza tarefa | JSON | 200 OK / 404 Not Found |
| `DELETE` | `/tarefas/{id}` | Remove tarefa | - | 200 OK / 404 Not Found |
| `PATCH` | `/tarefas/{id}/concluir` | Marca como concluída | - | 200 OK / 404 Not Found |
| `PATCH` | `/tarefas/{id}/reativar` | Reativa tarefa | - | 200 OK / 404 Not Found |
| `GET` | `/tarefas/estatisticas` | Estatísticas e métricas | - | 200 OK |

### Query Parameters (Filtros e Paginação)

| Parâmetro | Tipo | Descrição | Exemplo |
|-----------|------|-----------|---------|
| `concluida` | boolean | Filtrar por status | `?concluida=false` |
| `prioridade` | string | Filtrar por prioridade | `?prioridade=alta` |
| `categoria` | string | Filtrar por categoria | `?categoria=Trabalho` |
| `page` | number | Número da página | `?page=1` |
| `limit` | number | Itens por página (max 100) | `?limit=10` |

### Modelo de Dados

**Request (POST):**
```json
{
  "titulo": "Implementar nova funcionalidade",
  "descricao": "Desenvolver sistema de notificações",
  "prioridade": "alta",
  "categoria": "Desenvolvimento",
  "dataVencimento": "2025-10-25T17:00:00.000Z"
}
```

**Request (PUT):**
```json
{
  "titulo": "Título atualizado",
  "descricao": "Nova descrição",
  "concluida": true,
  "prioridade": "media",
  "categoria": "Nova Categoria",
  "dataVencimento": "2025-11-01T12:00:00.000Z"
}
```

**Response:**
```json
{
  "message": "Tarefa criada com sucesso",
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "titulo": "Implementar nova funcionalidade",
    "descricao": "Desenvolver sistema de notificações",
    "concluida": false,
    "prioridade": "alta",
    "categoria": "Desenvolvimento",
    "dataCriacao": "2025-10-17T16:00:00.000Z",
    "dataAtualizacao": "2025-10-17T16:00:00.000Z",
    "dataVencimento": "2025-10-25T17:00:00.000Z"
  },
  "timestamp": "2025-10-17T16:00:00.000Z"
}
```

**Validações:**
- `titulo`: obrigatório, 1-200 caracteres
- `descricao`: opcional, máximo 1000 caracteres
- `prioridade`: 'baixa' | 'media' | 'alta' (padrão: 'media')
- `categoria`: opcional, máximo 50 caracteres
- `dataVencimento`: opcional, formato ISO 8601, não pode ser no passado

---

## 🚀 Como Executar

### Pré-requisitos

Certifique-se de ter instalado:
- 🟢 **Node.js 18+** ([Download Node.js](https://nodejs.org/))
- 📦 **npm** (incluído com Node.js)
- 🛠️ **VS Code** ou IDE de sua preferência

### Instalação

```bash
# 1. Clone o repositório
git clone https://github.com/SEU-USUARIO/rest-api-series.git

# 2. Navegue até o projeto
cd rest-api-series/02-nodejs-express-tarefas

# 3. Instale as dependências
npm install

# 4. Execute em modo desenvolvimento
npm run dev

# Ou compile e execute em produção
npm run build
npm start
```

### Verificar se está funcionando

A API estará disponível em: **http://localhost:3000**

**Teste rápido:**
```bash
curl http://localhost:3000/health
# Resposta esperada: {"status":"OK","message":"API de Tarefas está funcionando!"}
```

**Teste das tarefas:**
```bash
curl http://localhost:3000/api/tarefas
# Resposta: Lista com 3 tarefas de exemplo
```

---

## 📝 Exemplos de Uso

### 1. Health Check

```bash
curl http://localhost:3000/health
```

**Resposta (200 OK):**
```json
{
  "status": "OK",
  "message": "API de Tarefas está funcionando!",
  "timestamp": "2025-10-17T16:00:00.000Z",
  "environment": "development"
}
```

---

### 2. Criar uma tarefa

```bash
curl -X POST http://localhost:3000/api/tarefas \
  -H "Content-Type: application/json" \
  -d '{
    "titulo": "Estudar Node.js",
    "descricao": "Aprofundar conhecimentos em Express e TypeScript",
    "prioridade": "alta",
    "categoria": "Educação",
    "dataVencimento": "2025-10-30T18:00:00.000Z"
  }'
```

**Resposta (201 Created):**
```json
{
  "message": "Tarefa criada com sucesso",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "titulo": "Estudar Node.js",
    "descricao": "Aprofundar conhecimentos em Express e TypeScript",
    "concluida": false,
    "prioridade": "alta",
    "categoria": "Educação",
    "dataCriacao": "2025-10-17T16:00:00.000Z",
    "dataAtualizacao": "2025-10-17T16:00:00.000Z",
    "dataVencimento": "2025-10-30T18:00:00.000Z"
  },
  "timestamp": "2025-10-17T16:00:00.000Z"
}
```

---

### 3. Listar todas as tarefas

```bash
curl http://localhost:3000/api/tarefas
```

---

### 4. Filtrar tarefas não concluídas

```bash
curl "http://localhost:3000/api/tarefas?concluida=false"
```

---

### 5. Filtrar por prioridade alta com paginação

```bash
curl "http://localhost:3000/api/tarefas?prioridade=alta&page=1&limit=5"
```

---

### 6. Buscar tarefa por ID

```bash
curl http://localhost:3000/api/tarefas/550e8400-e29b-41d4-a716-446655440000
```

---

### 7. Atualizar tarefa

```bash
curl -X PUT http://localhost:3000/api/tarefas/550e8400-e29b-41d4-a716-446655440000 \
  -H "Content-Type: application/json" \
  -d '{
    "titulo": "Estudar Node.js - ATUALIZADO",
    "descricao": "Foco em APIs REST e validação",
    "prioridade": "media",
    "concluida": false
  }'
```

---

### 8. Marcar tarefa como concluída

```bash
curl -X PATCH http://localhost:3000/api/tarefas/550e8400-e29b-41d4-a716-446655440000/concluir
```

---

### 9. Reativar tarefa

```bash
curl -X PATCH http://localhost:3000/api/tarefas/550e8400-e29b-41d4-a716-446655440000/reativar
```

---

### 10. Obter estatísticas

```bash
curl http://localhost:3000/api/tarefas/estatisticas
```

**Resposta:**
```json
{
  "message": "Estatísticas recuperadas com sucesso",
  "data": {
    "total": 5,
    "concluidas": 2,
    "pendentes": 3,
    "porPrioridade": {
      "baixa": 0,
      "media": 2,
      "alta": 3
    },
    "porCategoria": {
      "Desenvolvimento": 2,
      "Educação": 1,
      "DevOps": 1,
      "Teste": 1
    },
    "proximosVencimentos": [
      {
        "id": "abc123",
        "titulo": "Tarefa com vencimento próximo",
        "dataVencimento": "2025-10-20T23:59:59.000Z"
      }
    ]
  }
}
```

---

### 11. Deletar tarefa

```bash
curl -X DELETE http://localhost:3000/api/tarefas/550e8400-e29b-41d4-a716-446655440000
```

---

## 📂 Estrutura do Projeto

```
src/
├── app.ts                           ← Servidor Express principal
├── models/
│   └── Tarefa.ts                    ← Interfaces TypeScript
├── schemas/
│   └── tarefaSchema.ts              ← Validações Zod
├── controllers/
│   └── tarefaController.ts          ← Controllers REST
├── services/
│   └── tarefaService.ts             ← Lógica de negócio
├── storage/
│   └── dataStore.ts                 ← Armazenamento em memória
├── middlewares/
│   ├── errorHandler.ts              ← Tratamento de erros
│   └── validateRequest.ts           ← Validação de requisições
└── routes/
    └── tarefaRoutes.ts              ← Definição das rotas

Arquivos de configuração:
├── package.json                     ← Dependências e scripts
├── tsconfig.json                    ← Configuração TypeScript
├── .gitignore                       ← Arquivos ignorados
└── README.md                        ← Documentação
```

---

## 🎓 Conceitos Aplicados

### Design Patterns
- ✅ **Repository Pattern** - Abstração de acesso a dados
- ✅ **Service Layer** - Lógica de negócio centralizada
- ✅ **Middleware Pattern** - Interceptadores de requisição
- ✅ **Singleton Pattern** - DataStore e Service instâncias únicas
- ✅ **DTO Pattern** - Separação entre dados internos e externos

### TypeScript Features
- ✅ **Interfaces** - Contratos de dados
- ✅ **Types** - Tipagem personalizada
- ✅ **Enums** - Valores constantes
- ✅ **Generics** - Tipos flexíveis
- ✅ **Strict Mode** - Validação rigorosa

### Express Features
- ✅ **Router** - Organização de rotas
- ✅ **Middleware Stack** - Pipeline de processamento
- ✅ **Error Handling** - Tratamento centralizado
- ✅ **JSON Parser** - Processamento automático
- ✅ **CORS** - Cross-origin support

### Zod Validations
- ✅ **Schema Definition** - Validação declarativa
- ✅ **Transform** - Conversão de dados
- ✅ **Refinement** - Validações customizadas
- ✅ **Error Messages** - Mensagens personalizadas
- ✅ **Type Inference** - Tipos automáticos

### Security & Quality
- ✅ **Helmet** - Headers de segurança
- ✅ **Input Validation** - Validação rigorosa
- ✅ **Error Sanitization** - Informações controladas
- ✅ **Type Safety** - Prevenção de erros
- ✅ **Request Logging** - Auditoria

### HTTP Best Practices
- ✅ Status codes apropriados (200, 201, 400, 404, 500)
- ✅ Content-Type: application/json
- ✅ RESTful URL design
- ✅ Métodos HTTP semânticos
- ✅ Consistent error format

---

## 🧪 Testes

### Testar com cURL

Execute os exemplos da seção [Exemplos de Uso](#exemplos-de-uso)

### Testar com PowerShell

```powershell
# Health Check
Invoke-RestMethod -Uri "http://localhost:3000/health" -Method GET

# Listar tarefas
Invoke-RestMethod -Uri "http://localhost:3000/api/tarefas" -Method GET

# Criar tarefa
$body = @{
    titulo = "Nova tarefa"
    prioridade = "alta"
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:3000/api/tarefas" -Method POST -Body $body -ContentType "application/json"
```

### Dados de Exemplo

O sistema já inicia com 3 tarefas pré-cadastradas:

1. **"Implementar API REST"** (Alta prioridade, Desenvolvimento, Pendente)
2. **"Escrever documentação"** (Média prioridade, Documentação, Pendente)
3. **"Configurar CI/CD"** (Alta prioridade, DevOps, Concluída)

---

## 🔍 Funcionalidades Especiais

### 1. Filtros Avançados

```bash
# Combinar múltiplos filtros
curl "http://localhost:3000/api/tarefas?concluida=false&prioridade=alta&categoria=Desenvolvimento"
```

### 2. Paginação

```bash
# Paginar resultados
curl "http://localhost:3000/api/tarefas?page=2&limit=5"
```

### 3. Dashboard de Estatísticas

Inclui métricas completas:
- Total de tarefas
- Distribuição por status
- Distribuição por prioridade
- Distribuição por categoria
- Próximos vencimentos (7 dias)

### 4. Validação Robusta

- Títulos únicos (não permite duplicação)
- Datas de vencimento não podem ser no passado
- Validação de UUID em parâmetros
- Sanitização de entrada

### 5. Tratamento de Erros

Formato consistente de erro:
```json
{
  "error": "Tipo do erro",
  "message": "Mensagem descritiva",
  "timestamp": "2025-10-17T16:00:00.000Z",
  "path": "/api/tarefas/id-invalido",
  "method": "GET"
}
```

---

## 🚧 Melhorias Futuras

### Curto Prazo
- [ ] Adicionar testes unitários (Jest)
- [ ] Implementar ordenação personalizada
- [ ] Adicionar campo de tags
- [ ] Implementar busca textual

### Médio Prazo
- [ ] Adicionar autenticação JWT
- [ ] Migrar para banco PostgreSQL/MongoDB
- [ ] Implementar upload de anexos
- [ ] Documentação Swagger/OpenAPI
- [ ] Rate limiting

### Longo Prazo
- [ ] Dockerizar aplicação
- [ ] CI/CD com GitHub Actions
- [ ] Deploy em cloud (AWS/Vercel)
- [ ] Implementar WebSockets para real-time
- [ ] Cache com Redis
- [ ] Métricas e monitoring
- [ ] Notificações por email/push

---

## 📚 Recursos de Aprendizado

- 📖 [Express.js Guide](https://expressjs.com/en/guide/)
- 📖 [TypeScript Handbook](https://www.typescriptlang.org/docs/)
- 📖 [Zod Documentation](https://zod.dev/)
- 📖 [Node.js Best Practices](https://github.com/goldbergyoni/nodebestpractices)
- 📖 [RESTful API Design](https://restfulapi.net/)

---

## 🔗 Links Relacionados

- [⬅️ Projeto anterior: Java + Spring Boot](../01-java-springboot-biblioteca/)
- [⬆️ Voltar para a série completa](../)
- [➡️ Próximo projeto: C# + ASP.NET Core](../03-csharp-aspnet-produtos/)

---

## 👨‍💻 Autor

**André Cesar Vieira** - [GitHub](https://github.com/andrecesarvieira)

---

## 📝 Licença

Este projeto está sob a licença MIT. Consulte o arquivo [LICENSE](../../LICENSE) para mais detalhes.

---

**⭐ Se este projeto foi útil, considere dar uma estrela no repositório!**

---

*Desenvolvido como parte da série REST API Multi-linguagem - Outubro 2025*