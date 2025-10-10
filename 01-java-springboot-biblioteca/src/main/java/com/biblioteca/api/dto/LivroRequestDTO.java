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

    @Min(value = 1000, message = "Ano de publicação deve ser maior que 1000")
    @Max(value = 2100, message = "Ano de publicação deve ser menor que 2100")
    private Integer anoPublicacao;

    private Boolean disponivel = true;

    // Construtores
    public LivroRequestDTO() {
    }

    public LivroRequestDTO(String titulo, String autor, String isbn, Integer anoPublicacao) {
        this.titulo = titulo;
        this.autor = autor;
        this.isbn = isbn;
        this.anoPublicacao = anoPublicacao;
    }

    // Getters e Setters
    public String getTitulo() {
        return titulo;
    }

    public void setTitulo(String titulo) {
        this.titulo = titulo;
    }

    public String getAutor() {
        return autor;
    }

    public void setAutor(String autor) {
        this.autor = autor;
    }

    public String getIsbn() {
        return isbn;
    }

    public void setIsbn(String isbn) {
        this.isbn = isbn;
    }

    public Integer getAnoPublicacao() {
        return anoPublicacao;
    }

    public void setAnoPublicacao(Integer anoPublicacao) {
        this.anoPublicacao = anoPublicacao;
    }

    public Boolean getDisponivel() {
        return disponivel;
    }

    public void setDisponivel(Boolean disponivel) {
        this.disponivel = disponivel;
    }
}