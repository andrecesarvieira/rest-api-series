package com.biblioteca.api.repository;

import com.biblioteca.api.model.Livro;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface LivroRepository extends JpaRepository<Livro, Long> {

    // Spring cria a implementação automaticamente!
    // Métodos básicos já existem: findAll(), findById(), save(), delete()

    // Você pode adicionar queries customizadas:
    Optional<Livro> findByIsbn(String isbn);

    List<Livro> findByTituloContainingIgnoreCase(String titulo);

    List<Livro> findByAutorContainingIgnoreCase(String autor);
}