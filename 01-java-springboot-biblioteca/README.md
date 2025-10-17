# 📚 Biblioteca API - Java + Spring Boot

[![Spring Boot](https://img.shields.io/badge/Spring%20Boot-3.3.5-brightgreen)](https://spring.io/projects/spring-boot)
[![Java](https://img.shields.io/badge/Java-17-orange)](https://www.java.com)
[![Maven](https://img.shields.io/badge/Maven-3.6+-blue)](https://maven.apache.org)
[![License](https://img.shields.io/badge/License-MIT-yellow)](../../LICENSE)

API REST para gerenciamento de biblioteca desenvolvida com **Java** e **Spring Boot**, implementando operações CRUD completas para livros.

> **Parte 1 de 5** da série [REST API Multi-linguagem](../)

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

API REST que permite gerenciar um acervo de livros com as seguintes funcionalidades:

✅ Listar todos os livros  
✅ Buscar livro por ID  
✅ Criar novo livro  
✅ Atualizar informações de livro  
✅ Remover livro  
✅ Buscar livros por título  
✅ Validação de dados de entrada  
✅ Tratamento centralizado de exceções  
✅ Banco de dados em memória (H2)  

---

## 🛠️ Tecnologias

| Tecnologia | Versão | Finalidade |
|------------|--------|------------|
| Java | 17+ | Linguagem de programação |
| Spring Boot | 3.3.5 | Framework web |
| Spring Data JPA | 3.3.5 | Persistência de dados |
| Hibernate | 6.x | ORM (Object-Relational Mapping) |
| H2 Database | 2.x | Banco de dados em memória |
| Bean Validation | 3.0 | Validação de dados |
| Maven | 3.6+ | Gerenciamento de dependências |

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
│          Service Layer                  │  ← Lógica de negócio
│  (Regras e validações)                  │
└─────────────────┬───────────────────────┘
                  │
                  ↓
┌─────────────────────────────────────────┐
│        Repository Layer                 │  ← Acesso a dados
│  (Spring Data JPA)                      │
└─────────────────┬───────────────────────┘
                  │
                  ↓
┌─────────────────────────────────────────┐
│          H2 Database                    │  ← Persistência
└─────────────────────────────────────────┘
```

### Fluxo de Dados

```
Cliente → Controller → DTO → Service → Entity → Repository → Database
```

---

## 🔗 Endpoints da API

**Base URL**: `http://localhost:8081/api`

### Livros

| Método | Endpoint | Descrição | Request Body | Response |
|--------|----------|-----------|--------------|----------|
| `GET` | `/livros` | Lista todos os livros | - | 200 OK |
| `GET` | `/livros/{id}` | Busca livro por ID | - | 200 OK / 404 Not Found |
| `POST` | `/livros` | Cria novo livro | JSON | 201 Created / 400 Bad Request |
| `PUT` | `/livros/{id}` | Atualiza livro | JSON | 200 OK / 404 Not Found |
| `DELETE` | `/livros/{id}` | Remove livro | - | 204 No Content / 404 Not Found |
| `GET` | `/livros/search?titulo=X` | Busca por título | - | 200 OK |

### Modelo de Dados

**Request (POST/PUT):**
```json
{
  "titulo": "Clean Code",
  "autor": "Robert C. Martin",
  "isbn": "978-0132350884",
  "anoPublicacao": 2008,
  "disponivel": true
}
```

**Response:**
```json
{
  "id": 1,
  "titulo": "Clean Code",
  "autor": "Robert C. Martin",
  "isbn": "978-0132350884",
  "anoPublicacao": 2008,
  "disponivel": true,
  "dataCadastro": "2025-10-10T18:40:00",
  "dataAtualizacao": "2025-10-10T18:40:00"
}
```

**Validações:**
- `titulo`: obrigatório, 1-200 caracteres
- `autor`: obrigatório, 1-150 caracteres
- `isbn`: formato válido, único
- `anoPublicacao`: entre 1000 e 2100

---

## 🚀 Como Executar

### Pré-requisitos

Certifique-se de ter instalado:
- ☕ **Java 17** ou superior ([Download OpenJDK](https://adoptium.net/))
- 📦 **Maven 3.6+** ([Download Maven](https://maven.apache.org/download.cgi))
- 🛠️ **IntelliJ IDEA** ou IDE de sua preferência

### Instalação

```bash
# 1. Clone o repositório
git clone https://github.com/SEU-USUARIO/rest-api-series.git

# 2. Navegue até o projeto
cd rest-api-series/01-java-springboot-biblioteca

# 3. Execute com Maven
mvn spring-boot:run

# Ou compile e execute o JAR
mvn clean package
java -jar target/biblioteca-api-0.0.1-SNAPSHOT.jar
```

### Verificar se está funcionando

A API estará disponível em: **http://localhost:8081**

**Teste rápido:**
```bash
curl http://localhost:8081/api/livros
# Resposta esperada: []
```

---

## 📝 Exemplos de Uso

### 1. Criar um livro

```bash
curl -X POST http://localhost:8081/api/livros \
  -H "Content-Type: application/json" \
  -d '{
    "titulo": "Clean Code",
    "autor": "Robert C. Martin",
    "isbn": "978-0132350884",
    "anoPublicacao": 2008,
    "disponivel": true
  }'
```

**Resposta (201 Created):**
```json
{
  "id": 1,
  "titulo": "Clean Code",
  "autor": "Robert C. Martin",
  "isbn": "978-0132350884",
  "anoPublicacao": 2008,
  "disponivel": true,
  "dataCadastro": "2025-10-10T18:40:00.123456",
  "dataAtualizacao": "2025-10-10T18:40:00.123456"
}
```

---

### 2. Listar todos os livros

```bash
curl http://localhost:8081/api/livros
```

---

### 3. Buscar livro por ID

```bash
curl http://localhost:8081/api/livros/1
```

---

### 4. Atualizar livro

```bash
curl -X PUT http://localhost:8081/api/livros/1 \
  -H "Content-Type: application/json" \
  -d '{
    "titulo": "Clean Code",
    "autor": "Robert C. Martin",
    "isbn": "978-0132350884",
    "anoPublicacao": 2008,
    "disponivel": false
  }'
```

---

### 5. Buscar por título

```bash
curl "http://localhost:8081/api/livros/search?titulo=clean"
```

---

### 6. Deletar livro

```bash
curl -X DELETE http://localhost:8081/api/livros/1
```

---

## 🔍 H2 Database Console

O projeto inclui o console web do H2 para visualizar e consultar o banco de dados.

**URL**: http://localhost:8081/h2-console

**Credenciais:**
- **JDBC URL**: `jdbc:h2:mem:bibliotecadb`
- **Username**: `sa`
- **Password**: *(deixe em branco)*

**Exemplo de Query:**
```sql
SELECT * FROM livros ORDER BY data_cadastro DESC;
```

---

## 📂 Estrutura do Projeto

```
src/
├── main/
│   ├── java/com/biblioteca/api/
│   │   ├── controller/
│   │   │   └── LivroController.java          ← Endpoints REST
│   │   ├── dto/
│   │   │   ├── LivroRequestDTO.java          ← Request payload
│   │   │   └── LivroResponseDTO.java         ← Response payload
│   │   ├── exception/
│   │   │   ├── ErrorResponse.java            ← Formato de erro
│   │   │   └── GlobalExceptionHandler.java   ← Tratamento global
│   │   ├── model/
│   │   │   └── Livro.java                    ← Entidade JPA
│   │   ├── repository/
│   │   │   └── LivroRepository.java          ← Spring Data JPA
│   │   ├── service/
│   │   │   └── LivroService.java             ← Lógica de negócio
│   │   └── BibliotecaApiApplication.java     ← Ponto de entrada
│   └── resources/
│       ├── application.properties             ← Configurações
│       └── static/
│           └── index.html                     ← Página inicial
└── test/                                      ← Testes (futuro)
```

---

## 🎓 Conceitos Aplicados

### Design Patterns
- ✅ **DTO Pattern** - Separação entre entidade e representação externa
- ✅ **Repository Pattern** - Abstração de acesso a dados
- ✅ **Dependency Injection** - Inversão de controle
- ✅ **Layered Architecture** - Separação de responsabilidades

### Spring Boot Features
- ✅ **@RestController** - Controllers REST
- ✅ **@Service** - Camada de serviço
- ✅ **@Repository** - Camada de persistência
- ✅ **@Autowired** - Injeção de dependências
- ✅ **@Valid** - Validação automática
- ✅ **@RestControllerAdvice** - Tratamento global de exceções

### JPA/Hibernate
- ✅ **@Entity** - Mapeamento objeto-relacional
- ✅ **@Id, @GeneratedValue** - Chave primária auto-incremento
- ✅ **@Column** - Configuração de colunas
- ✅ **@PrePersist, @PreUpdate** - Callbacks de ciclo de vida
- ✅ **Spring Data JPA** - Queries automáticas

### Validações
- ✅ **@NotBlank** - Campo não vazio
- ✅ **@Size** - Tamanho min/max
- ✅ **@Min, @Max** - Valores numéricos
- ✅ **@Pattern** - Validação por regex

### HTTP
- ✅ Status codes apropriados (200, 201, 204, 400, 404, 500)
- ✅ Content-Type: application/json
- ✅ RESTful URL design
- ✅ Métodos HTTP semânticos

---

## 🧪 Testes

### Testar com Postman

Importe a collection disponível em: [`../../docs/postman/biblioteca-api.json`](../../docs/postman/)

### Testar com Curl

Execute os exemplos da seção [Exemplos de Uso](#exemplos-de-uso)

---

## 🚧 Melhorias Futuras

### Curto Prazo
- [ ] Adicionar paginação e ordenação
- [ ] Implementar filtros avançados
- [ ] Adicionar campo de categoria
- [ ] Implementar soft delete

### Médio Prazo
- [ ] Adicionar autenticação JWT
- [ ] Implementar testes unitários (JUnit + Mockito)
- [ ] Adicionar testes de integração
- [ ] Documentação Swagger/OpenAPI
- [ ] Migrar para PostgreSQL

### Longo Prazo
- [ ] Dockerizar aplicação
- [ ] CI/CD com GitHub Actions
- [ ] Deploy em cloud (AWS/Heroku)
- [ ] Implementar cache com Redis
- [ ] Rate limiting
- [ ] Logs estruturados (ELK Stack)
- [ ] Métricas e monitoring (Prometheus + Grafana)

---

## 📚 Recursos de Aprendizado

- 📖 [Spring Boot Reference](https://docs.spring.io/spring-boot/docs/current/reference/html/)
- 📖 [Spring Data JPA](https://spring.io/projects/spring-data-jpa)
- 📖 [RESTful API Design Best Practices](https://restfulapi.net/)
- 📖 [Bean Validation Documentation](https://beanvalidation.org/)

---

## 🔗 Links Relacionados

- [⬆️ Voltar para a série completa](../)
- [➡️ Próximo projeto: Node.js + Express](../02-nodejs-express-tarefas/)
- [📁 Collections do Postman](../../docs/postman/)

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