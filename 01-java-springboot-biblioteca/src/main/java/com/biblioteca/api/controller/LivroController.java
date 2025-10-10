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

@RestController  // ← Indica que é um Controller REST
@RequestMapping("/api/livros")  // ← Base URL: /api/livros
public class LivroController {

    @Autowired
    private LivroService livroService;

    // GET /api/livros - Listar todos
    @GetMapping
    public ResponseEntity<List<LivroResponseDTO>> listarTodos() {
        List<LivroResponseDTO> livros = livroService.listarTodos();
        return ResponseEntity.ok(livros);  // ← HTTP 200 OK
    }

    // GET /api/livros/{id} - Buscar por ID
    @GetMapping("/{id}")
    public ResponseEntity<LivroResponseDTO> buscarPorId(@PathVariable Long id) {
        LivroResponseDTO livro = livroService.buscarPorId(id);
        return ResponseEntity.ok(livro);  // ← HTTP 200 OK
    }

    // POST /api/livros - Criar novo livro
    @PostMapping
    public ResponseEntity<LivroResponseDTO> criar(@Valid @RequestBody LivroRequestDTO requestDTO) {
        LivroResponseDTO livro = livroService.criar(requestDTO);
        return ResponseEntity.status(HttpStatus.CREATED).body(livro);  // ← HTTP 201 CREATED
    }

    // PUT /api/livros/{id} - Atualizar livro
    @PutMapping("/{id}")
    public ResponseEntity<LivroResponseDTO> atualizar(
            @PathVariable Long id,
            @Valid @RequestBody LivroRequestDTO requestDTO) {

        LivroResponseDTO livro = livroService.atualizar(id, requestDTO);
        return ResponseEntity.ok(livro);  // ← HTTP 200 OK
    }

    // DELETE /api/livros/{id} - Deletar livro
    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deletar(@PathVariable Long id) {
        livroService.deletar(id);
        return ResponseEntity.noContent().build();  // ← HTTP 204 NO CONTENT
    }

    // GET /api/livros/search?titulo=xxx - Buscar por título
    @GetMapping("/search")
    public ResponseEntity<List<LivroResponseDTO>> buscarPorTitulo(
            @RequestParam(required = false) String titulo) {

        List<LivroResponseDTO> livros = livroService.buscarPorTitulo(titulo);
        return ResponseEntity.ok(livros);
    }
}