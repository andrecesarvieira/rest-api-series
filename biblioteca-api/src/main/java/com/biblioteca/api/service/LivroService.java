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

    // Listar todos os livros
    public List<LivroResponseDTO> listarTodos() {
        List<Livro> livros = livroRepository.findAll();

        // Converte List<Livro> para List<LivroResponseDTO>
        return livros.stream()
                .map(this::convertToResponseDTO)
                .collect(Collectors.toList());
    }

    // Buscar livro por ID
    public LivroResponseDTO buscarPorId(Long id) {
        Livro livro = livroRepository.findById(id)
                .orElseThrow(() -> new RuntimeException("Livro não encontrado com ID: " + id));

        return convertToResponseDTO(livro);
    }

    // Criar novo livro
    public LivroResponseDTO criar(LivroRequestDTO requestDTO) {
        // Verificar se ISBN já existe
        livroRepository.findByIsbn(requestDTO.getIsbn())
                .ifPresent(livro -> {
                    throw new RuntimeException("ISBN já cadastrado: " + requestDTO.getIsbn());
                });

        // Converter DTO para Entidade
        Livro livro = convertToEntity(requestDTO);

        // Salvar no banco
        Livro livroSalvo = livroRepository.save(livro);

        // Retornar DTO de resposta
        return convertToResponseDTO(livroSalvo);
    }

    // Atualizar livro existente
    public LivroResponseDTO atualizar(Long id, LivroRequestDTO requestDTO) {
        // Verificar se livro existe
        Livro livroExistente = livroRepository.findById(id)
                .orElseThrow(() -> new RuntimeException("Livro não encontrado com ID: " + id));

        // Verificar se ISBN mudou e se já existe
        if (requestDTO.getIsbn() != null && !requestDTO.getIsbn().equals(livroExistente.getIsbn())) {
            livroRepository.findByIsbn(requestDTO.getIsbn())
                    .ifPresent(livro -> {
                        throw new RuntimeException("ISBN já cadastrado: " + requestDTO.getIsbn());
                    });
        }

        // Atualizar campos
        livroExistente.setTitulo(requestDTO.getTitulo());
        livroExistente.setAutor(requestDTO.getAutor());
        livroExistente.setIsbn(requestDTO.getIsbn());
        livroExistente.setAnoPublicacao(requestDTO.getAnoPublicacao());
        livroExistente.setDisponivel(requestDTO.getDisponivel());

        // Salvar alterações
        Livro livroAtualizado = livroRepository.save(livroExistente);

        return convertToResponseDTO(livroAtualizado);
    }

    // Deletar livro
    public void deletar(Long id) {
        // Verificar se existe
        if (!livroRepository.existsById(id)) {
            throw new RuntimeException("Livro não encontrado com ID: " + id);
        }

        livroRepository.deleteById(id);
    }

    // Buscar por título
    public List<LivroResponseDTO> buscarPorTitulo(String titulo) {
        List<Livro> livros = livroRepository.findByTituloContainingIgnoreCase(titulo);

        return livros.stream()
                .map(this::convertToResponseDTO)
                .collect(Collectors.toList());
    }

    // ========== MÉTODOS DE CONVERSÃO ==========

    // Converter Entidade para ResponseDTO
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

    // Converter RequestDTO para Entidade
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