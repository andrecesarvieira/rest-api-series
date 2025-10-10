# 📚 Manual Completo - Projeto 1: Java + Spring Boot (API de Biblioteca)

> **Documentação Completa** de tudo que foi desenvolvido, aprendido e configurado no primeiro projeto da série REST API Multi-linguagem.

---

## 📑 Índice

1. [Visão Geral](#visão-geral)
2. [Ambiente e Instalação](#ambiente-e-instalação)
3. [Criação do Projeto](#criação-do-projeto)
4. [Estrutura do Projeto](#estrutura-do-projeto)
5. [Configurações](#configurações)
6. [Código Completo](#código-completo)
7. [Conceitos Aprendidos](#conceitos-aprendidos)
8. [Testes e Validação](#testes-e-validação)
9. [Git e GitHub](#git-e-github)
10. [Troubleshooting](#troubleshooting)
11. [Referências](#referências)

---

## 🎯 Visão Geral

### Objetivo do Projeto
Criar uma API REST completa para gerenciar um acervo de livros, implementando operações CRUD (Create, Read, Update, Delete) usando as melhores práticas de desenvolvimento Java com Spring Boot.

### Tecnologias Utilizadas
| Tecnologia | Versão | Finalidade |
|------------|--------|------------|
| Java | 17+ | Linguagem de programação |
| Spring Boot | 3.3.5 | Framework web |
| Spring Data JPA | 3.5.4 | Persistência de dados |
| Hibernate | 6.6.x | ORM |
| H2 Database | 2.3.x | Banco em memória |
| Bean Validation | 3.0 | Validação de dados |
| Maven | 3.6+ | Build e dependências |

### Funcionalidades Implementadas
✅ Listar todos os livros  
✅ Buscar livro por ID  
✅ Criar novo livro  
✅ Atualizar livro existente  
✅ Deletar livro  
✅ Buscar livros por título  
✅ Validação automática de dados  
✅ Tratamento global de exceções  
✅ Console H2 para visualizar banco  

---

## 💻 Ambiente e Instalação

### Pré-requisitos

#### 1. Java Development Kit (JDK)
```bash
# Verificar instalação
java -version
javac -version

# Instalar OpenJDK 17 (Ubuntu/Debian)
sudo apt update
sudo apt install openjdk-17-jdk

# Verificar instalação
java -version
# Deve mostrar: openjdk version "17.0.x"
```

#### 2. Apache Maven
```bash
# Verificar instalação
mvn -version

# Instalar Maven (Ubuntu/Debian)
sudo apt install maven

# Verificar
mvn -version
# Deve mostrar: Apache Maven 3.6.x ou superior
```

#### 3. IDE Recomendada
- **IntelliJ IDEA Community Edition** (preferida)
- Download: https://www.jetbrains.com/idea/download/
- Alternativa: VS Code com Extension Pack for Java

### Variáveis de Ambiente (se necessário)
```bash
# Adicionar ao ~/.bashrc ou ~/.zshrc
export JAVA_HOME=/usr/lib/jvm/java-17-openjdk-amd64
export PATH=$JAVA_HOME/bin:$PATH
```

---

## 🏗️ Criação do Projeto

### Método 1: Spring Initializr (Web)

1. **Acessar:** https://start.spring.io/

2. **Configurações do Projeto:**
```
Project: Maven
Language: Java
Spring Boot: 3.3.5
```

3. **Metadados:**
```
Group: com.biblioteca
Artifact: biblioteca-api
Name: biblioteca-api
Description: API REST para gerenciar biblioteca
Package name: com.biblioteca.api
Packaging: Jar
Java: 17
```

4. **Dependências Adicionadas:**
- Spring Web
- Spring Data JPA
- H2 Database
- Validation

5. **Generate** → Download → Extrair → Abrir no IntelliJ

### Método 2: Maven Command Line
```bash
mvn archetype:generate \
  -DgroupId=com.biblioteca \
  -DartifactId=biblioteca-api \
  -DarchetypeArtifactId=maven-archetype-quickstart \
  -DinteractiveMode=false
```

### Método 3: IntelliJ IDEA
```
File → New → Project → Spring Initializr
Seguir mesmo wizard do Spring Initializr web
```

---

## 📂 Estrutura do Projeto

### Árvore de Diretórios Completa
```
01-java-springboot-biblioteca/
├── src/
│   ├── main/
│   │   ├── java/com/biblioteca/api/
│   │   │   ├── controller/
│   │   │   │   └── LivroController.java
│   │   │   ├── dto/
│   │   │   │   ├── LivroRequestDTO.java
│   │   │   │   └── LivroResponseDTO.java
│   │   │   ├── exception/
│   │   │   │   ├── ErrorResponse.java
│   │   │   │   └── GlobalExceptionHandler.java
│   │   │   ├── model/
│   │   │   │   └── Livro.java
│   │   │   ├── repository/
│   │   │   │   └── LivroRepository.java
│   │   │   ├── service/
│   │   │   │   └── LivroService.java
│   │   │   └── BibliotecaApiApplication.java
│   │   └── resources/
│   │       ├── application.properties
│   │       └── static/
│   │           └── index.html (opcional)
│   └── test/
│       └── java/com/biblioteca/api/
│           └── BibliotecaApiApplicationTests.java
├── target/ (gerado pelo Maven - ignorado no Git)
├── .gitignore
├── pom.xml
├── mvnw
├── mvnw.cmd
└── README.md
```

### Arquitetura em Camadas

```
┌─────────────────────────────────────┐
│    Camada de Apresentação           │
│    (Controller)                     │  ← REST Endpoints
│    • LivroController                │
└──────────────┬──────────────────────┘
               │ DTOs
               ↓
┌─────────────────────────────────────┐
│    Camada de Negócio                │
│    (Service)                        │  ← Lógica de Negócio
│    • LivroService                   │
└──────────────┬──────────────────────┘
               │ Entities
               ↓
┌─────────────────────────────────────┐
│    Camada de Persistência           │
│    (Repository)                     │  ← Acesso a Dados
│    • LivroRepository                │
└──────────────┬──────────────────────┘
               │ SQL
               ↓
┌─────────────────────────────────────┐
│    Banco de Dados                   │
│    • H2 (in-memory)                 │
└─────────────────────────────────────┘
```

### Responsabilidades de Cada Camada

| Camada | Responsabilidade | Não Deve Fazer |
|--------|------------------|----------------|
| **Controller** | Receber requisições HTTP, validar entrada, retornar respostas | Lógica de negócio, acesso direto ao banco |
| **Service** | Regras de negócio, validações complexas, orquestração | Detalhes HTTP, queries SQL diretas |
| **Repository** | Acesso ao banco de dados, queries | Lógica de negócio, validações |
| **Model** | Representar dados, estrutura | Lógica, comportamentos complexos |

---

## ⚙️ Configurações

### pom.xml (Dependências Maven)
```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 
         https://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>
    
    <parent>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-parent</artifactId>
        <version>3.3.5</version>
        <relativePath/>
    </parent>
    
    <groupId>com.biblioteca</groupId>
    <artifactId>biblioteca-api</artifactId>
    <version>0.0.1-SNAPSHOT</version>
    <name>biblioteca-api</name>
    <description>API REST para gerenciar biblioteca</description>
    
    <properties>
        <java.version>17</java.version>
    </properties>
    
    <dependencies>
        <!-- Spring Boot Web -->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        
        <!-- Spring Data JPA -->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-data-jpa</artifactId>
        </dependency>
        
        <!-- H2 Database -->
        <dependency>
            <groupId>com.h2database</groupId>
            <artifactId>h2</artifactId>
            <scope>runtime</scope>
        </dependency>
        
        <!-- Validation -->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-validation</artifactId>
        </dependency>
        
        <!-- Test -->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
    </dependencies>
    
    <build>
        <plugins>
            <plugin>
                <groupId>org.springframework.boot</groupId>
                <artifactId>spring-boot-maven-plugin</artifactId>
            </plugin>
        </plugins>
    </build>
</project>
```

### application.properties
```properties
# ===== SERVIDOR =====
server.port=8081

# ===== H2 DATABASE =====
spring.datasource.url=jdbc:h2:mem:bibliotecadb
spring.datasource.driverClassName=org.h2.Driver
spring.datasource.username=sa
spring.datasource.password=

# ===== JPA / HIBERNATE =====
spring.jpa.database-platform=org.hibernate.dialect.H2Dialect
spring.jpa.hibernate.ddl-auto=create-drop
spring.jpa.show-sql=true
spring.jpa.properties.hibernate.format_sql=true

# ===== H2 CONSOLE =====
spring.h2.console.enabled=true
spring.h2.console.path=/h2-console

# ===== LOGGING =====
logging.level.org.springframework.web=INFO
logging.level.org.hibernate=INFO
```

**Explicação das Configurações:**

- `server.port=8081` → API roda na porta 8081
- `jdbc:h2:mem:` → Banco em memória (dados somem ao parar)
- `ddl-auto=create-drop` → Cria tabelas ao iniciar, apaga ao parar
- `show-sql=true` → Mostra SQLs executados no console
- `h2.console.enabled=true` → Ativa interface web do H2

---

## 💾 Código Completo

### 1. Entidade (Model)

**Livro.java**
```java
package com.biblioteca.api.model;

import jakarta.persistence.*;
import java.time.LocalDateTime;

@Entity
@Table(name = "livros")
public class Livro {
    
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    
    @Column(nullable = false, length = 200)
    private String titulo;
    
    @Column(nullable = false, length = 150)
    private String autor;
    
    @Column(unique = true, length = 20)
    private String isbn;
    
    @Column(name = "ano_publicacao")
    private Integer anoPublicacao;
    
    @Column(nullable = false)
    private Boolean disponivel = true;
    
    @Column(name = "data_cadastro", updatable = false)
    private LocalDateTime dataCadastro;
    
    @Column(name = "data_atualizacao")
    private LocalDateTime dataAtualizacao;
    
    @PrePersist
    protected void onCreate() {
        dataCadastro = LocalDateTime.now();
        dataAtualizacao = LocalDateTime.now();
    }
    
    @PreUpdate
    protected void onUpdate() {
        dataAtualizacao = LocalDateTime.now();
    }
    
    // Construtores
    public Livro() {}
    
    public Livro(String titulo, String autor, String isbn, Integer anoPublicacao) {
        this.titulo = titulo;
        this.autor = autor;
        this.isbn = isbn;
        this.anoPublicacao = anoPublicacao;
    }
    
    // Getters e Setters
    public Long getId() { return id; }
    public void setId(Long id) { this.id = id; }
    
    public String getTitulo() { return titulo; }
    public void setTitulo(String titulo) { this.titulo = titulo; }
    
    public String getAutor() { return autor; }
    public void setAutor(String autor) { this.autor = autor; }
    
    public String getIsbn() { return isbn; }
    public void setIsbn(String isbn) { this.isbn = isbn; }
    
    public Integer getAnoPublicacao() { return anoPublicacao; }
    public void setAnoPublicacao(Integer anoPublicacao) { 
        this.anoPublicacao = anoPublicacao; 
    }
    
    public Boolean getDisponivel() { return disponivel; }
    public void setDisponivel(Boolean disponivel) { 
        this.disponivel = disponivel; 
    }
    
    public LocalDateTime getDataCadastro() { return dataCadastro; }
    public void setDataCadastro(LocalDateTime dataCadastro) { 
        this.dataCadastro = dataCadastro; 
    }
    
    public LocalDateTime getDataAtualizacao() { return dataAtualizacao; }
    public void setDataAtualizacao(LocalDateTime dataAtualizacao) { 
        this.dataAtualizacao = dataAtualizacao; 
    }
}
```

**Anotações JPA Explicadas:**

| Anotação | Função |
|----------|--------|
| `@Entity` | Marca classe como entidade JPA (tabela) |
| `@Table(name = "livros")` | Define nome da tabela |
| `@Id` | Define chave primária |
| `@GeneratedValue` | Auto-incremento |
| `@Column` | Configurações da coluna (nullable, length, unique, etc) |
| `@PrePersist` | Executado antes de INSERT |
| `@PreUpdate` | Executado antes de UPDATE |

---

### 2. DTOs (Data Transfer Objects)

**LivroRequestDTO.java**
```java
package com.biblioteca.api.dto;

import jakarta.validation.constraints.*;

public class LivroRequestDTO {
    
    @NotBlank(message = "Título é obrigatório")
    @Size(min = 1, max = 200, message = "Título deve ter entre 1 e 200 caracteres")
    private String titulo;
    
    @NotBlank(message = "Autor é obrigatório")
    @Size(min = 1, max = 150, message = "Autor deve ter entre 1 e 150 caracteres")
    private String autor;
    
    @Pattern(regexp = "^(?=(?:\\D*\\d){10}(?:(?:\\D*\\d){3})?$)[\\d-]+$", 
             message = "ISBN inválido")
    private String isbn;
    
    @Min(value = 1000, message = "Ano deve ser maior que 1000")
    @Max(value = 2100, message = "Ano deve ser menor que 2100")
    private Integer anoPublicacao;
    
    private Boolean disponivel = true;
    
    // Construtores
    public LivroRequestDTO() {}
    
    public LivroRequestDTO(String titulo, String autor, String isbn, 
                          Integer anoPublicacao) {
        this.titulo = titulo;
        this.autor = autor;
        this.isbn = isbn;
        this.anoPublicacao = anoPublicacao;
    }
    
    // Getters e Setters
    public String getTitulo() { return titulo; }
    public void setTitulo(String titulo) { this.titulo = titulo; }
    
    public String getAutor() { return autor; }
    public void setAutor(String autor) { this.autor = autor; }
    
    public String getIsbn() { return isbn; }
    public void setIsbn(String isbn) { this.isbn = isbn; }
    
    public Integer getAnoPublicacao() { return anoPublicacao; }
    public void setAnoPublicacao(Integer anoPublicacao) { 
        this.anoPublicacao = anoPublicacao; 
    }
    
    public Boolean getDisponivel() { return disponivel; }
    public void setDisponivel(Boolean disponivel) { 
        this.disponivel = disponivel; 
    }
}
```

**LivroResponseDTO.java**
```java
package com.biblioteca.api.dto;

import java.time.LocalDateTime;

public class LivroResponseDTO {
    
    private Long id;
    private String titulo;
    private String autor;
    private String isbn;
    private Integer anoPublicacao;
    private Boolean disponivel;
    private LocalDateTime dataCadastro;
    private LocalDateTime dataAtualizacao;
    
    public LivroResponseDTO() {}
    
    public LivroResponseDTO(Long id, String titulo, String autor, String isbn,
                           Integer anoPublicacao, Boolean disponivel,
                           LocalDateTime dataCadastro, LocalDateTime dataAtualizacao) {
        this.id = id;
        this.titulo = titulo;
        this.autor = autor;
        this.isbn = isbn;
        this.anoPublicacao = anoPublicacao;
        this.disponivel = disponivel;
        this.dataCadastro = dataCadastro;
        this.dataAtualizacao = dataAtualizacao;
    }
    
    // Getters e Setters (todos os campos)
    public Long getId() { return id; }
    public void setId(Long id) { this.id = id; }
    
    public String getTitulo() { return titulo; }
    public void setTitulo(String titulo) { this.titulo = titulo; }
    
    public String getAutor() { return autor; }
    public void setAutor(String autor) { this.autor = autor; }
    
    public String getIsbn() { return isbn; }
    public void setIsbn(String isbn) { this.isbn = isbn; }
    
    public Integer getAnoPublicacao() { return anoPublicacao; }
    public void setAnoPublicacao(Integer anoPublicacao) { 
        this.anoPublicacao = anoPublicacao; 
    }
    
    public Boolean getDisponivel() { return disponivel; }
    public void setDisponivel(Boolean disponivel) { 
        this.disponivel = disponivel; 
    }
    
    public LocalDateTime getDataCadastro() { return dataCadastro; }
    public void setDataCadastro(LocalDateTime dataCadastro) { 
        this.dataCadastro = dataCadastro; 
    }
    
    public LocalDateTime getDataAtualizacao() { return dataAtualizacao; }
    public void setDataAtualizacao(LocalDateTime dataAtualizacao) { 
        this.dataAtualizacao = dataAtualizacao; 
    }
}
```

**Validações Bean Validation:**

| Anotação | Função |
|----------|--------|
| `@NotBlank` | Não pode ser null, vazio ou só espaços |
| `@NotNull` | Não pode ser null |
| `@Size(min, max)` | Tamanho da string |
| `@Min(value)` | Valor mínimo para números |
| `@Max(value)` | Valor máximo para números |
| `@Pattern(regexp)` | Validação por regex |
| `@Email` | Formato de email válido |

---

### 3. Repository

**LivroRepository.java**
```java
package com.biblioteca.api.repository;

import com.biblioteca.api.model.Livro;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface LivroRepository extends JpaRepository<Livro, Long> {
    
    // Spring Data JPA cria automaticamente as implementações!
    
    // Métodos herdados (já existem):
    // - List<Livro> findAll()
    // - Optional<Livro> findById(Long id)
    // - Livro save(Livro livro)
    // - void deleteById(Long id)
    // - boolean existsById(Long id)
    // - long count()
    
    // Métodos customizados (Spring cria automaticamente baseado no nome):
    Optional<Livro> findByIsbn(String isbn);
    
    List<Livro> findByTituloContainingIgnoreCase(String titulo);
    
    List<Livro> findByAutorContainingIgnoreCase(String autor);
}
```

**Query Methods - Convenção de Nomes:**

| Nome do Método | SQL Gerado |
|----------------|------------|
| `findByIsbn(String isbn)` | `WHERE isbn = ?` |
| `findByTituloContaining(String titulo)` | `WHERE titulo LIKE %?%` |
| `findByTituloContainingIgnoreCase(String titulo)` | `WHERE LOWER(titulo) LIKE LOWER(%?%)` |
| `findByAnoPublicacaoGreaterThan(Integer ano)` | `WHERE ano_publicacao > ?` |
| `findByDisponivelTrue()` | `WHERE disponivel = true` |

---

### 4. Service

**LivroService.java**
```java
package com.biblioteca.api.service;

import com.biblioteca.api.dto.LivroRequestDTO;
import com.biblioteca.api.dto.LivroResponseDTO;
import com.biblioteca.api.model.Livro;
import com.biblioteca.api.repository.LivroRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.stream.Collectors;

@Service
public class LivroService {
    
    @Autowired
    private LivroRepository livroRepository;
    
    public List<LivroResponseDTO> listarTodos() {
        List<Livro> livros = livroRepository.findAll();
        return livros.stream()
                .map(this::convertToResponseDTO)
                .collect(Collectors.toList());
    }
    
    public LivroResponseDTO buscarPorId(Long id) {
        Livro livro = livroRepository.findById(id)
                .orElseThrow(() -> new RuntimeException(
                    "Livro não encontrado com ID: " + id));
        return convertToResponseDTO(livro);
    }
    
    public LivroResponseDTO criar(LivroRequestDTO requestDTO) {
        livroRepository.findByIsbn(requestDTO.getIsbn())
                .ifPresent(livro -> {
                    throw new RuntimeException(
                        "ISBN já cadastrado: " + requestDTO.getIsbn());
                });
        
        Livro livro = convertToEntity(requestDTO);
        Livro livroSalvo = livroRepository.save(livro);
        return convertToResponseDTO(livroSalvo);
    }
    
    public LivroResponseDTO atualizar(Long id, LivroRequestDTO requestDTO) {
        Livro livroExistente = livroRepository.findById(id)
                .orElseThrow(() -> new RuntimeException(
                    "Livro não encontrado com ID: " + id));
        
        if (requestDTO.getIsbn() != null && 
            !requestDTO.getIsbn().equals(livroExistente.getIsbn())) {
            livroRepository.findByIsbn(requestDTO.getIsbn())
                    .ifPresent(livro -> {
                        throw new RuntimeException(
                            "ISBN já cadastrado: " + requestDTO.getIsbn());
                    });
        }
        
        livroExistente.setTitulo(requestDTO.getTitulo());
        livroExistente.setAutor(requestDTO.getAutor());
        livroExistente.setIsbn(requestDTO.getIsbn());
        livroExistente.setAnoPublicacao(requestDTO.getAnoPublicacao());
        livroExistente.setDisponivel(requestDTO.getDisponivel());
        
        Livro livroAtualizado = livroRepository.save(livroExistente);
        return convertToResponseDTO(livroAtualizado);
    }
    
    public void deletar(Long id) {
        if (!livroRepository.existsById(id)) {
            throw new RuntimeException("Livro não encontrado com ID: " + id);
        }
        livroRepository.deleteById(id);
    }
    
    public List<LivroResponseDTO> buscarPorTitulo(String titulo) {
        List<Livro> livros = livroRepository
            .findByTituloContainingIgnoreCase(titulo);
        return livros.stream()
                .map(this::convertToResponseDTO)
                .collect(Collectors.toList());
    }
    
    // Métodos de conversão
    private LivroResponseDTO convertToResponseDTO(Livro livro) {
        return new LivroResponseDTO(
                livro.getId(),
                livro.getTitulo(),
                livro.getAutor(),
                livro.getIsbn(),
                livro.getAnoPublicacao(),
                livro.getDisponivel(),
                livro.getDataCadastro(),
                livro.getDataAtualizacao()
        );
    }
    
    private Livro convertToEntity(LivroRequestDTO requestDTO) {
        Livro livro = new Livro();
        livro.setTitulo(requestDTO.getTitulo());
        livro.setAutor(requestDTO.getAutor());
        livro.setIsbn(requestDTO.getIsbn());
        livro.setAnoPublicacao(requestDTO.getAnoPublicacao());
        livro.setDisponivel(requestDTO.getDisponivel());
        return livro;
    }
}
```

---

### 5. Exception Handling

**ErrorResponse.java**
```java
package com.biblioteca.api.exception;

import java.time.LocalDateTime;

public class ErrorResponse {
    
    private String error;
    private String message;
    private LocalDateTime timestamp;
    private String path;
    
    public ErrorResponse(String error, String message, String path) {
        this.error = error;
        this.message = message;
        this.timestamp = LocalDateTime.now();
        this.path = path;
    }
    
    // Getters e Setters
    public String getError() { return error; }
    public void setError(String error) { this.error = error; }
    
    public String getMessage() { return message; }
    public void setMessage(String message) { this.message = message; }
    
    public LocalDateTime getTimestamp() { return timestamp; }
    public void setTimestamp(LocalDateTime timestamp) { 
        this.timestamp = timestamp; 
    }
    
    public String getPath() { return path; }
    public void setPath(String path) { this.path = path; }
}
```

**GlobalExceptionHandler.java**
```java
package com.biblioteca.api.exception;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.FieldError;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;
import org.springframework.web.context.request.WebRequest;

import java.util.HashMap;
import java.util.Map;

@RestControllerAdvice
public class GlobalExceptionHandler {
    
    @ExceptionHandler(RuntimeException.class)
    public ResponseEntity<ErrorResponse> handleRuntimeException(
            RuntimeException ex, WebRequest request) {
        
        ErrorResponse error = new ErrorResponse(
                "Bad Request",
                ex.getMessage(),
                request.getDescription(false).replace("uri=", "")
        );
        
        return new ResponseEntity<>(error, HttpStatus.BAD_REQUEST);
    }
    
    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<Map<String, String>> handleValidationExceptions(
            MethodArgumentNotValidException ex) {
        
        Map<String, String> errors = new HashMap<>();
        
        ex.getBindingResult().getAllErrors().forEach((error) -> {
            String fieldName = ((FieldError) error).getField();
            String errorMessage = error.getDefaultMessage();
            errors.put(fieldName, errorMessage);
        });
        
        return new ResponseEntity<>(errors, HttpStatus.BAD_REQUEST);
    }
    
    @ExceptionHandler(Exception.class)
    public ResponseEntity<ErrorResponse> handleGlobalException(
            Exception ex, WebRequest request) {
        
        ErrorResponse error = new ErrorResponse(
                "Internal Server Error",
                ex.getMessage(),
                request.getDescription(false).replace("uri=", "")
        );
        
        return new ResponseEntity<>(error, HttpStatus.INTERNAL_SERVER_ERROR);
    }
}
```

---

### 6. Controller

**LivroController.java**
```java
package com.biblioteca.api.controller;

import com.biblioteca.api.dto.LivroRequestDTO;
import com.biblioteca.api.dto.LivroResponseDTO;
import com.biblioteca.api.service.LivroService;
import jakarta.validation.Valid;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/livros")
public class LivroController {
    
    @Autowired
    private LivroService livroService;
    
    @GetMapping
    public ResponseEntity<List<LivroResponseDTO>> listarTodos() {
        List<LivroResponseDTO> livros = livroService.listarTodos();
        return ResponseEntity.ok(livros);
    }
    
    @GetMapping("/{id}")
    public ResponseEntity<LivroResponseDTO> buscarPorId(@PathVariable Long id) {
        LivroResponseDTO livro = livroService.buscarPorId(id);
        return ResponseEntity.ok(livro);
    }
    
    @PostMapping
    public ResponseEntity<LivroResponseDTO> criar(
            @Valid @RequestBody LivroRequestDTO requestDTO) {
        LivroResponseDTO livro = livroService.criar(requestDTO);
        return ResponseEntity.status(HttpStatus.CREATED).body(livro);
    }
    
    @PutMapping("/{id}")
    public ResponseEntity<LivroResponseDTO> atualizar(
            @PathVariable Long id,
            @Valid @RequestBody LivroRequestDTO requestDTO) {
        
        LivroResponseDTO livro = livroService.atualizar(id, requestDTO);
        return ResponseEntity.ok(livro);
    }
    
    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deletar(@PathVariable Long id) {
        livroService.deletar(id);
        return ResponseEntity.noContent().build();
    }
    
    @GetMapping("/search")
    public ResponseEntity<List<LivroResponseDTO>> buscarPorTitulo(
            @RequestParam(required = false) String titulo) {
        
        List<LivroResponseDTO> livros = livroService.buscarPorTitulo(titulo);
        return ResponseEntity.ok(livros);
    }
}
```

**Anotações Spring MVC:**

| Anotação | Função |
|----------|--------|
| `@RestController` | Controller que retorna JSON/XML |
| `@RequestMapping("/api/livros")` | Base URL do controller |
| `@GetMapping` | HTTP GET |
| `@PostMapping` | HTTP POST |
| `@PutMapping` | HTTP PUT |
| `@DeleteMapping` | HTTP DELETE |
| `@PathVariable` | Captura variável da URL |
| `@RequestBody` | Captura corpo da requisição |
| `@RequestParam` | Captura query parameter |
| `@Valid` | Ativa validação Bean Validation |

---

### 7. Aplicação Principal

**BibliotecaApiApplication.java**
```java
package com.biblioteca.api;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class BibliotecaApiApplication {
    
    public static void main(String[] args) {
        SpringApplication.run(BibliotecaApiApplication.class, args);
    }
}
```

---

## 🎓 Conceitos Aprendidos

### 1. Inversão de Controle (IoC) e Injeção de Dependência (DI)

**Problema sem DI:**
```java
public class LivroService {
    private LivroRepository repository = new LivroRepositoryImpl();
    // Acoplamento forte! Service conhece implementação concreta
}
```

**Solução com DI (Spring):**
```java
@Service
public class LivroService {
    @Autowired  // Spring injeta automaticamente
    private LivroRepository repository;
}
```

**Benefícios:**
- ✅ Desacoplamento
- ✅ Facilita testes (mocks)
- ✅ Flexibilidade (trocar implementações)

---

### 2. Optional<T> - Evitando NullPointerException

**Antes (Java 7):**
```java
Livro livro = repository.findById(id);
if (livro != null) {
    // usar livro
} else {
    throw new Exception();
}
```

**Agora (Java 8+):**
```java
Livro livro = repository.findById(id)
    .orElseThrow(() -> new RuntimeException("Não encontrado"));
```

**Outros métodos úteis:**
```java
optional.isPresent()           // true se tem valor
optional.ifPresent(livro -> {...})  // executa se tiver valor
optional.orElse(defaultValue)  // retorna default se vazio
optional.orElseThrow(...)      // lança exceção se vazio
```

---

### 3. Streams API - Processamento de Coleções

**Antes (Java 7):**
```java
List<LivroResponseDTO> dtos = new ArrayList<>();
for (Livro livro : livros) {
    dtos.add(convertToDTO(livro));
}
return dtos;
```

**Agora (Java 8+):**
```java
return livros.stream()
    .map(this::convertToDTO)
    .collect(Collectors.toList());
```

**Operações úteis:**
```java
.filter(livro -> livro.getDisponivel())  // Filtrar
.map(this::convertToDTO)                 // Transformar
.sorted(Comparator.comparing(Livro::getTitulo))  // Ordenar
.limit(10)                               // Limitar quantidade
.collect(Collectors.toList())            // Coletar resultado
```

---

### 4. Method Reference (::)

**Forma longa (lambda):**
```java
.map(livro -> this.convertToDTO(livro))
```

**Forma curta (method reference):**
```java
.map(this::convertToDTO)
```

**Tipos de Method Reference:**
```java
String::toUpperCase          // Método de instância
Math::max                    // Método estático
ArrayList::new               // Construtor
this::convertToDTO           // Método do objeto atual
```

---

### 5. JPA Entity Lifecycle Callbacks

```java
@PrePersist   // Antes de INSERT
@PostPersist  // Depois de INSERT
@PreUpdate    // Antes de UPDATE
@PostUpdate   // Depois de UPDATE
@PreRemove    // Antes de DELETE
@PostRemove   // Depois de DELETE
@PostLoad     // Depois de SELECT
```

**Exemplo prático:**
```java
@PrePersist
protected void onCreate() {
    dataCadastro = LocalDateTime.now();
    if (status == null) status = "ATIVO";
}
```

---

### 6. ResponseEntity<T> - Controle Total da Resposta HTTP

```java
// 200 OK
return ResponseEntity.ok(data);

// 201 CREATED
return ResponseEntity.status(HttpStatus.CREATED).body(data);

// 204 NO CONTENT
return ResponseEntity.noContent().build();

// 404 NOT FOUND
return ResponseEntity.notFound().build();

// 400 BAD REQUEST
return ResponseEntity.badRequest().body(error);

// Custom headers
return ResponseEntity.ok()
    .header("X-Custom-Header", "value")
    .body(data);
```

---

### 7. Padrão DTO (Data Transfer Object)

**Por que usar DTOs?**

1. **Segurança**: Não expor campos sensíveis da entidade
2. **Controle**: Diferentes representações para entrada/saída
3. **Desacoplamento**: Cliente não conhece estrutura interna
4. **Validação**: Validar entrada separadamente

**Fluxo:**
```
Cliente → RequestDTO → Service → Entity → Repository → Database
Database → Entity → Service → ResponseDTO → Cliente
```

---

### 8. Spring Data JPA - Query Methods

**Convenções de nomes que geram queries automaticamente:**

```java
// WHERE campo = ?
findByCampo(valor)

// WHERE campo LIKE %?%
findByCampoContaining(valor)

// WHERE campo > ?
findByCampoGreaterThan(valor)

// WHERE campo < ?
findByCampoLessThan(valor)

// WHERE campo1 = ? AND campo2 = ?
findByCampo1AndCampo2(valor1, valor2)

// WHERE campo1 = ? OR campo2 = ?
findByCampo1OrCampo2(valor1, valor2)

// ORDER BY campo ASC
findByOrderByCampoAsc()

// WHERE disponivel = true
findByDisponivelTrue()
```

---

## 🧪 Testes e Validação

### Endpoints da API

| Método | URL | Body | Resposta |
|--------|-----|------|----------|
| GET | `/api/livros` | - | 200 OK + Array |
| GET | `/api/livros/1` | - | 200 OK + Objeto / 404 |
| POST | `/api/livros` | JSON | 201 Created + Objeto / 400 |
| PUT | `/api/livros/1` | JSON | 200 OK + Objeto / 404 / 400 |
| DELETE | `/api/livros/1` | - | 204 No Content / 404 |
| GET | `/api/livros/search?titulo=X` | - | 200 OK + Array |

### Exemplos de Requisições (curl)

**1. Criar Livro:**
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

**2. Listar Todos:**
```bash
curl http://localhost:8081/api/livros
```

**3. Buscar por ID:**
```bash
curl http://localhost:8081/api/livros/1
```

**4. Atualizar:**
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

**5. Deletar:**
```bash
curl -X DELETE http://localhost:8081/api/livros/1
```

**6. Buscar por Título:**
```bash
curl "http://localhost:8081/api/livros/search?titulo=clean"
```

### Validação de Erros

**Teste 1: Criar sem título**
```bash
curl -X POST http://localhost:8081/api/livros \
  -H "Content-Type: application/json" \
  -d '{
    "autor": "Autor",
    "isbn": "123",
    "anoPublicacao": 2020
  }'

# Resposta esperada: 400 Bad Request
{
  "titulo": "Título é obrigatório"
}
```

**Teste 2: ISBN duplicado**
```bash
# Criar livro
curl -X POST http://localhost:8081/api/livros \
  -H "Content-Type: application/json" \
  -d '{"titulo":"Livro 1","autor":"Autor","isbn":"123"}'

# Tentar criar com mesmo ISBN
curl -X POST http://localhost:8081/api/livros \
  -H "Content-Type: application/json" \
  -d '{"titulo":"Livro 2","autor":"Autor","isbn":"123"}'

# Resposta: 400 Bad Request
{
  "error": "Bad Request",
  "message": "ISBN já cadastrado: 123",
  ...
}
```

**Teste 3: ID inexistente**
```bash
curl http://localhost:8081/api/livros/999

# Resposta: 400 Bad Request
{
  "error": "Bad Request",
  "message": "Livro não encontrado com ID: 999",
  ...
}
```

### H2 Console

**Acessar:** http://localhost:8081/h2-console

**Login:**
- JDBC URL: `jdbc:h2:mem:bibliotecadb`
- User: `sa`
- Password: *(vazio)*

**Queries úteis:**
```sql
-- Ver todos os livros
SELECT * FROM livros;

-- Contar livros
SELECT COUNT(*) FROM livros;

-- Livros disponíveis
SELECT * FROM livros WHERE disponivel = true;

-- Livros por autor
SELECT * FROM livros WHERE autor LIKE '%Martin%';

-- Livros ordenados por data
SELECT * FROM livros ORDER BY data_cadastro DESC;

-- Ver estrutura da tabela
SHOW COLUMNS FROM livros;
```

---

## 📦 Git e GitHub

### Configuração Inicial

```bash
# Configurar nome e email (se ainda não fez)
git config --global user.name "Seu Nome"
git config --global user.email "seu-email@example.com"

# Ver configurações
git config --list
```

### Comandos Principais

```bash
# Inicializar repositório
git init

# Ver status
git status

# Adicionar arquivos
git add .                    # Todos os arquivos
git add arquivo.java         # Arquivo específico
git add src/                 # Pasta específica

# Commit
git commit -m "mensagem"

# Ver histórico
git log
git log --oneline            # Formato resumido
git log --graph --oneline    # Com gráfico

# Ver diferenças
git diff                     # Mudanças não staged
git diff --staged            # Mudanças staged

# Conectar ao remoto
git remote add origin URL
git remote -v                # Ver remotes

# Push
git push -u origin main      # Primeira vez
git push                     # Próximas vezes

# Pull (atualizar local)
git pull
```

### .gitignore Configurado

```gitignore
# Maven
target/
pom.xml.tag
pom.xml.releaseBackup
pom.xml.versionsBackup
pom.xml.next
release.properties

# IDE
.idea/
*.iml
.vscode/
.project
.classpath
.settings/

# OS
.DS_Store
Thumbs.db

# Logs
*.log

# Application
application-local.properties
```

### SSH Setup (Feito)

```bash
# Gerar chave
ssh-keygen -t ed25519 -C "seu-email@example.com"

# Adicionar ao agent
eval "$(ssh-agent -s)"
ssh-add ~/.ssh/id_ed25519

# Copiar chave pública
cat ~/.ssh/id_ed25519.pub

# Testar conexão
ssh -T git@github.com
```

### Comandos Úteis Git

```bash
# Desfazer último commit (mantém mudanças)
git reset --soft HEAD~1

# Desfazer mudanças em arquivo
git checkout -- arquivo.java

# Ver branches
git branch

# Criar e mudar para nova branch
git checkout -b feature/nova-funcionalidade

# Merge de branch
git checkout main
git merge feature/nova-funcionalidade

# Remover arquivo do Git (mas manter localmente)
git rm --cached arquivo

# Atualizar .gitignore e remover arquivos já commitados
git rm -r --cached .
git add .
git commit -m "fix: atualiza .gitignore"
```

---

## 🛠️ Troubleshooting

### Problemas Comuns e Soluções

#### 1. Porta 8080/8081 já em uso

**Erro:**
```
Port 8081 is already in use
```

**Solução 1: Mudar porta**
```properties
# application.properties
server.port=8082
```

**Solução 2: Matar processo**
```bash
# Ver processo usando a porta
lsof -i :8081

# Matar processo
kill -9 PID
```

---

#### 2. Erro "Failed to configure a DataSource"

**Erro:**
```
Failed to configure a DataSource: 'url' attribute is not specified
```

**Solução:**
```properties
# Verificar se está configurado em application.properties
spring.datasource.url=jdbc:h2:mem:bibliotecadb
spring.datasource.driverClassName=org.h2.Driver
```

---

#### 3. Hibernate não cria tabelas

**Erro:** Tabela não existe ao fazer requisição

**Solução:**
```properties
# Verificar configuração
spring.jpa.hibernate.ddl-auto=create-drop

# Para produção use:
# spring.jpa.hibernate.ddl-auto=validate
```

---

#### 4. Erro 404 ao acessar endpoints

**Problema:** URL errada ou controller não sendo escaneado

**Verificar:**
```java
// Controller deve estar no mesmo package ou subpackage do Application
com.biblioteca.api            // Application aqui
com.biblioteca.api.controller // Controller aqui (OK)

// Se não estiver, adicionar scan explícito:
@SpringBootApplication
@ComponentScan(basePackages = "com.biblioteca")
public class BibliotecaApiApplication {...}
```

---

#### 5. JSON não serializa/deserializa corretamente

**Problema:** Faltam getters/setters

**Solução:**
```java
// Toda classe DTO/Entity precisa de:
// - Construtor vazio
// - Getters e Setters de todos os campos

public class LivroDTO {
    private String titulo;
    
    public LivroDTO() {}  // OBRIGATÓRIO
    
    public String getTitulo() { return titulo; }
    public void setTitulo(String titulo) { this.titulo = titulo; }
}
```

---

#### 6. Maven não baixa dependências

**Solução:**
```bash
# Limpar cache do Maven
mvn clean

# Forçar download
mvn clean install -U

# Deletar pasta .m2 e tentar novamente
rm -rf ~/.m2/repository
mvn clean install
```

---

#### 7. IntelliJ não reconhece classes

**Solução:**
```
File → Invalidate Caches → Invalidate and Restart
```

---

#### 8. Erro "java: package jakarta.persistence does not exist"

**Problema:** Maven não importou dependências

**Solução:**
```
Botão direito no projeto → Maven → Reload Project
```

---

## 📚 Comandos Maven Úteis

```bash
# Compilar projeto
mvn compile

# Executar testes
mvn test

# Limpar target/
mvn clean

# Criar JAR
mvn package

# Instalar no repositório local
mvn install

# Executar aplicação
mvn spring-boot:run

# Ver dependências
mvn dependency:tree

# Atualizar dependências
mvn clean install -U

# Pular testes
mvn clean install -DskipTests
```

---

## 🎯 Boas Práticas Aplicadas

### 1. Separação de Responsabilidades
✅ Controller → apenas HTTP  
✅ Service → lógica de negócio  
✅ Repository → acesso a dados  

### 2. DTO Pattern
✅ Nunca expor entidades diretamente  
✅ RequestDTO para entrada  
✅ ResponseDTO para saída  

### 3. Validação em Múltiplas Camadas
✅ Bean Validation no DTO (`@NotBlank`, etc)  
✅ Validações de negócio no Service  
✅ Constraints no banco (unique, not null)  

### 4. Tratamento de Exceções
✅ GlobalExceptionHandler centralizado  
✅ Mensagens de erro claras  
✅ HTTP status codes apropriados  

### 5. Nomenclatura
✅ Classes em PascalCase: `LivroService`  
✅ Métodos em camelCase: `buscarPorId()`  
✅ URLs RESTful: `/api/livros`, não `/api/getLivros`  
✅ Substantivos no plural: `/livros`, não `/livro`  

### 6. Versionamento
✅ Commits descritivos  
✅ Usar convenções: `feat:`, `fix:`, `docs:`  
✅ README atualizado  

---

## 📖 Referências e Documentação

### Documentação Oficial
- [Spring Boot Reference](https://docs.spring.io/spring-boot/docs/current/reference/html/)
- [Spring Data JPA](https://spring.io/projects/spring-data-jpa)
- [Bean Validation](https://beanvalidation.org/)
- [H2 Database](https://www.h2database.com/)

### Tutoriais Úteis
- [Baeldung - Spring Boot](https://www.baeldung.com/spring-boot)
- [Spring Guides](https://spring.io/guides)
- [REST API Best Practices](https://restfulapi.net/)

### Livros Recomendados
- "Spring Boot in Action" - Craig Walls
- "Spring in Action" - Craig Walls
- "RESTful Web Services" - Leonard Richardson

---

## ✅ Checklist de Conclusão

### Funcionalidades
- [x] CRUD completo de livros
- [x] Validação de dados
- [x] Tratamento de erros
- [x] Busca por título
- [x] H2 Console funcionando

### Qualidade de Código
- [x] Arquitetura em camadas
- [x] DTOs implementados
- [x] Repository pattern
- [x] Service layer
- [x] Exception handling

### Documentação
- [x] README.md completo
- [x] Comentários no código
- [x] Endpoints documentados

### Versionamento
- [x] Git inicializado
- [x] .gitignore configurado
- [x] Commits organizados
- [x] Push para GitHub
- [x] SSH configurado

### Testes
- [x] Testado todos os endpoints
- [x] Validações funcionando
- [x] H2 Console acessível

---

## 🎓 Resumo dos Aprendizados

### Conceitos Técnicos
✅ Spring Boot e suas convenções  
✅ Spring Data JPA e query methods  
✅ Bean Validation  
✅ REST API design  
✅ Layered Architecture  
✅ Dependency Injection  
✅ Optional e Streams API  
✅ Exception handling  

### Ferramentas
✅ Maven (build e dependências)  
✅ IntelliJ IDEA  
✅ H2 Database  
✅ Git e GitHub (SSH)  
✅ curl / Postman  

### Boas Práticas
✅ Separação de responsabilidades  
✅ DTOs para entrada/saída  
✅ Validação em múltiplas camadas  
✅ Tratamento centralizado de erros  
✅ Nomenclatura RESTful  
✅ Versionamento com Git  

---

## 🚀 Próximos Passos

Para aprofundar este projeto, considere:

1. **Testes Unitários**
   - JUnit 5
   - Mockito para mocks
   - TestContainers

2. **Documentação API**
   - Swagger/OpenAPI
   - SpringDoc

3. **Segurança**
   - Spring Security
   - JWT Authentication

4. **Banco de Dados Real**
   - PostgreSQL
   - MySQL

5. **Docker**
   - Dockerfile
   - docker-compose.yml

6. **CI/CD**
   - GitHub Actions
   - Testes automatizados

7. **Observabilidade**
   - Logs estruturados (Logback)
   - Métricas (Actuator)
   - Monitoring (Prometheus + Grafana)

---

## 📝 Notas Finais

Este projeto serviu como base sólida para entender:
- Como construir APIs REST profissionais
- Arquitetura de aplicações Spring Boot
- Boas práticas de desenvolvimento Java
- Versionamento com Git/GitHub

O conhecimento adquirido aqui será aplicado nos próximos projetos da série, permitindo comparar diferentes abordagens e tecnologias.

---

**Manual criado em:** Outubro 2025  
**Versão:** 1.0  
**Projeto:** REST API Series - Projeto 1  
**Autor:** André

---

*"O código é lido muito mais vezes do que é escrito. Escreva pensando em quem vai ler."*
