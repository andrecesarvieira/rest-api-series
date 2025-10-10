# 🚀 REST API Series - Jornada Multi-linguagem

[![Status](https://img.shields.io/badge/Status-Em%20Desenvolvimento-yellow)]()
[![License](https://img.shields.io/badge/License-MIT-blue.svg)]()
[![Java](https://img.shields.io/badge/Java-17-orange)]()
[![Node.js](https://img.shields.io/badge/Node.js-18+-green)]()
[![.NET](https://img.shields.io/badge/.NET-8-purple)]()
[![Python](https://img.shields.io/badge/Python-3.10+-blue)]()
[![Go](https://img.shields.io/badge/Go-1.21+-cyan)]()

Uma série completa de **5 APIs REST** implementadas em diferentes linguagens e frameworks modernos. Cada projeto implementa os mesmos conceitos REST para facilitar comparações e demonstrar versatilidade técnica.

## 🎯 Objetivo

Dominar o desenvolvimento de APIs REST através da prática com diferentes tecnologias, compreendendo:

- ✅ Princípios REST e boas práticas
- ✅ CRUD completo (Create, Read, Update, Delete)
- ✅ Validação de dados e tratamento de erros
- ✅ Padrões arquiteturais (MVC, Layered Architecture)
- ✅ Documentação de APIs
- ✅ Especificidades de cada ecossistema

---

## 📚 Projetos da Série

| # | Projeto | Linguagem | Framework | Domínio | Status |
|---|---------|-----------|-----------|---------|--------|
| 1 | [**Biblioteca**](./01-java-springboot-biblioteca/) | Java | Spring Boot | Gerenciamento de livros | ✅ **Completo** |
| 2 | [**Tarefas**](./02-nodejs-express-tarefas/) | Node.js | Express + TypeScript | Sistema To-Do | 🚧 Em progresso |
| 3 | [**Pedidos**](./03-csharp-aspnet-pedidos/) | C# | ASP.NET Core | E-commerce | ⏳ Planejado |
| 4 | [**Produtos**](./04-python-fastapi-produtos/) | Python | FastAPI | Catálogo | ⏳ Planejado |
| 5 | [**Usuários**](./05-go-gin-usuarios/) | Go | Gin | Gestão de usuários | ⏳ Planejado |

---

## 🛠️ Stack Tecnológico

### 1️⃣ Java + Spring Boot (Biblioteca)
```
Framework: Spring Boot 3.3.5
ORM: JPA/Hibernate
Validação: Bean Validation
Banco de Dados: H2 (em memória)
Build Tool: Maven
```

### 2️⃣ Node.js + Express (Tarefas)
```
Runtime: Node.js 18+
Framework: Express.js
Linguagem: TypeScript
Validação: Zod
Storage: Em memória (array)
```

### 3️⃣ C# + ASP.NET Core (Pedidos)
```
Framework: .NET 8
ORM: Entity Framework Core
Validação: FluentValidation
Banco de Dados: In-Memory
```

### 4️⃣ Python + FastAPI (Produtos)
```
Framework: FastAPI
Validação: Pydantic
Server: Uvicorn
Storage: Em memória
```

### 5️⃣ Go + Gin (Usuários)
```
Framework: Gin
Linguagem: Go 1.21+
Storage: Map em memória
```

---

## 📋 Padrões REST Implementados

Todos os projetos seguem os mesmos padrões REST:

### HTTP Methods
- **GET** - Recuperar recursos
- **POST** - Criar novos recursos
- **PUT** - Atualizar recursos completos
- **PATCH** - Atualizar recursos parcialmente
- **DELETE** - Remover recursos

### Status Codes
- **200 OK** - Sucesso em GET, PUT, PATCH
- **201 Created** - Recurso criado com POST
- **204 No Content** - Sucesso em DELETE (sem corpo)
- **400 Bad Request** - Dados inválidos
- **404 Not Found** - Recurso não encontrado
- **500 Internal Server Error** - Erro do servidor

### Estrutura de URLs
```
GET    /api/{recursos}           # Listar todos
GET    /api/{recursos}/{id}      # Buscar por ID
POST   /api/{recursos}           # Criar novo
PUT    /api/{recursos}/{id}      # Atualizar completo
PATCH  /api/{recursos}/{id}      # Atualizar parcial
DELETE /api/{recursos}/{id}      # Deletar
GET    /api/{recursos}/search    # Buscar com filtros
```

---

## 🚀 Como Executar

Cada projeto tem instruções específicas no seu próprio README. Exemplo geral:

```bash
# 1. Clonar o repositório
git clone https://github.com/SEU-USUARIO/rest-api-series.git
cd rest-api-series

# 2. Entrar no projeto desejado
cd 01-java-springboot-biblioteca

# 3. Seguir instruções específicas do README do projeto
```

---

## 📖 Recursos Adicionais

### Documentação
- 📁 [Collections do Postman](./docs/postman/) - Coleções prontas para importar
- 📊 [Comparação entre Linguagens](./docs/comparacoes.md) - Análise detalhada
- 💡 [Aprendizados e Insights](./docs/aprendizados.md) - O que aprendi em cada projeto

### Ferramentas Recomendadas
- [Postman](https://www.postman.com/) - Testar APIs
- [Insomnia](https://insomnia.rest/) - Alternativa ao Postman
- [Docker](https://www.docker.com/) - Containerização
- [DBeaver](https://dbeaver.io/) - Cliente de banco de dados

---

## 🎓 Conceitos Aprendidos

### Arquitetura
- ✅ Layered Architecture (Controller → Service → Repository)
- ✅ Separation of Concerns
- ✅ Dependency Injection
- ✅ DTO Pattern (Data Transfer Objects)

### Boas Práticas
- ✅ Validação de entrada de dados
- ✅ Tratamento centralizado de exceções
- ✅ Status codes HTTP apropriados
- ✅ Nomenclatura RESTful de URLs
- ✅ Versionamento de APIs

### Segurança
- ✅ Validação de dados de entrada
- ✅ Tratamento de erros sem expor detalhes internos
- ✅ Sanitização de inputs

---

## 📊 Comparativo de Tecnologias

| Critério | Java | Node.js | C# | Python | Go |
|----------|------|---------|----|---------|----|
| **Curva de Aprendizado** | Média | Baixa | Média | Baixa | Média |
| **Produtividade** | Alta | Muito Alta | Alta | Muito Alta | Alta |
| **Performance** | Muito Alta | Alta | Muito Alta | Média | Excelente |
| **Ecossistema** | Maduro | Vasto | Maduro | Rico | Crescente |
| **Mercado** | Muito Forte | Forte | Forte | Crescente | Crescente |

---

## 🎯 Próximos Passos

- [ ] **Projeto 2**: Node.js + Express (Tarefas)
- [ ] **Projeto 3**: C# + ASP.NET Core (Pedidos)
- [ ] **Projeto 4**: Python + FastAPI (Produtos)
- [ ] **Projeto 5**: Go + Gin (Usuários)

### Melhorias Futuras
- [ ] Adicionar autenticação JWT em todos os projetos
- [ ] Implementar testes unitários e de integração
- [ ] Migrar para bancos de dados reais (PostgreSQL)
- [ ] Dockerizar todas as aplicações
- [ ] Criar docker-compose para rodar tudo junto
- [ ] Implementar CI/CD com GitHub Actions
- [ ] Adicionar rate limiting
- [ ] Implementar cache com Redis
- [ ] Criar documentação Swagger/OpenAPI
- [ ] Deploy em cloud (AWS, Azure, Heroku)

---

## 👨‍💻 Sobre o Autor

**André** - Desenvolvedor com experiência em Cobol e C#, expandindo conhecimentos para o ecossistema moderno de APIs e microsserviços.

### Conecte-se comigo
- 💼 [LinkedIn](#)
- 🐱 [GitHub](#)
- 📧 [Email](#)

---

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

## ⭐ Apoie o Projeto

Se este projeto te ajudou nos estudos ou serviu como referência, considere:
- ⭐ Dar uma estrela no repositório
- 🔀 Fazer um fork para suas próprias experiências
- 📢 Compartilhar com outros desenvolvedores

---

## 🙏 Agradecimentos

Inspirado pela necessidade de entender profundamente como diferentes linguagens e frameworks abordam os mesmos problemas de desenvolvimento de APIs REST.

---

**Última atualização**: Outubro 2025

**Status do Projeto**: 🟡 Em Desenvolvimento Ativo (1/5 projetos concluídos)Teste SSH
