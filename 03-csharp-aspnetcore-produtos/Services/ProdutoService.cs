using AutoMapper;
using ProdutosApi.DTOs;
using ProdutosApi.Repositories;
using ProdutosApi.Models;
using ProdutosApi.Exceptions;

namespace ProdutosApi.Services
{
    public class ProdutoService : IProdutoService
    {
        private readonly IProdutoRepository _repository;
        private readonly IMapper _mapper;

        public ProdutoService(IProdutoRepository repository, IMapper mapper)
        {
            _repository = repository;
            _mapper = mapper;
        }

        public async Task<IEnumerable<ProdutoResponseDTO>> GetAllAsync()
        {
            var produtos = await _repository.GetAllAsync();
            return _mapper.Map<IEnumerable<ProdutoResponseDTO>>(produtos);
        }

        public async Task<ProdutoResponseDTO?> GetByIdAsync(int id)
        {
            if (id <= 0)
                throw new BadRequestException("ID deve ser maior que zero");

            var produto = await _repository.GetByIdAsync(id);
            
            if (produto == null)
                return null;

            return _mapper.Map<ProdutoResponseDTO>(produto);
        }

        public async Task<ProdutoResponseDTO> CreateAsync(ProdutoRequestDTO produtoRequest)
        {
            var produto = _mapper.Map<Produto>(produtoRequest);
            
            // Validações de negócio
            await ValidateBusinessRules(produto);

            var createdProduto = await _repository.CreateAsync(produto);
            return _mapper.Map<ProdutoResponseDTO>(createdProduto);
        }

        public async Task<ProdutoResponseDTO?> UpdateAsync(int id, ProdutoRequestDTO produtoRequest)
        {
            if (id <= 0)
                throw new BadRequestException("ID deve ser maior que zero");

            var existingProduto = await _repository.GetByIdAsync(id);
            if (existingProduto == null)
                return null;

            var produto = _mapper.Map<Produto>(produtoRequest);
            
            // Validações de negócio
            await ValidateBusinessRules(produto);

            var updatedProduto = await _repository.UpdateAsync(id, produto);
            return _mapper.Map<ProdutoResponseDTO>(updatedProduto);
        }

        public async Task<bool> DeleteAsync(int id)
        {
            if (id <= 0)
                throw new BadRequestException("ID deve ser maior que zero");

            return await _repository.DeleteAsync(id);
        }

        public async Task<IEnumerable<ProdutoResponseDTO>> GetByCategoryAsync(string categoria)
        {
            if (string.IsNullOrWhiteSpace(categoria))
                throw new BadRequestException("Categoria não pode ser vazia");

            var produtos = await _repository.GetByCategoryAsync(categoria);
            return _mapper.Map<IEnumerable<ProdutoResponseDTO>>(produtos);
        }

        public async Task<IEnumerable<ProdutoResponseDTO>> GetByPriceRangeAsync(decimal minPreco, decimal maxPreco)
        {
            if (minPreco < 0)
                throw new BadRequestException("Preço mínimo não pode ser negativo");
            
            if (maxPreco < minPreco)
                throw new BadRequestException("Preço máximo deve ser maior que o preço mínimo");

            var produtos = await _repository.GetByPriceRangeAsync(minPreco, maxPreco);
            return _mapper.Map<IEnumerable<ProdutoResponseDTO>>(produtos);
        }

        public async Task<IEnumerable<ProdutoResponseDTO>> GetActiveAsync()
        {
            var produtos = await _repository.GetActiveAsync();
            return _mapper.Map<IEnumerable<ProdutoResponseDTO>>(produtos);
        }

        public async Task<PagedResult<ProdutoResponseDTO>> GetPaginatedAsync(int page, int size, string? categoria = null, string? nome = null, bool? ativo = null)
        {
            if (page < 1)
                throw new BadRequestException("Página deve ser maior que zero");
            
            if (size < 1 || size > 100)
                throw new BadRequestException("Tamanho da página deve estar entre 1 e 100");

            var (produtos, total) = await _repository.GetPaginatedAsync(page, size, categoria, nome, ativo);
            var produtosDto = _mapper.Map<IEnumerable<ProdutoResponseDTO>>(produtos);

            var totalPages = (int)Math.Ceiling((double)total / size);

            return new PagedResult<ProdutoResponseDTO>
            {
                Data = produtosDto,
                Page = page,
                Size = size,
                Total = total,
                TotalPages = totalPages,
                HasPrevious = page > 1,
                HasNext = page < totalPages
            };
        }

        public async Task<object> GetStatisticsAsync()
        {
            var allProdutos = await _repository.GetAllAsync();
            var produtos = allProdutos.ToList();

            if (!produtos.Any())
            {
                return new
                {
                    totalProdutos = 0,
                    produtosAtivos = 0,
                    produtosInativos = 0,
                    precoMedio = 0m,
                    precoMinimo = 0m,
                    precoMaximo = 0m,
                    categorias = new List<object>(),
                    produtoMaisCaro = (object?)null,
                    produtoMaisBarato = (object?)null
                };
            }

            var produtosAtivos = produtos.Where(p => p.Ativo).ToList();
            var precos = produtos.Select(p => p.Preco).ToList();
            
            var categorias = produtos
                .GroupBy(p => p.Categoria)
                .Select(g => new
                {
                    categoria = g.Key,
                    quantidade = g.Count(),
                    precoMedio = g.Average(p => p.Preco)
                })
                .OrderBy(c => c.categoria)
                .ToList();

            return new
            {
                totalProdutos = produtos.Count,
                produtosAtivos = produtosAtivos.Count,
                produtosInativos = produtos.Count - produtosAtivos.Count,
                precoMedio = precos.Average(),
                precoMinimo = precos.Min(),
                precoMaximo = precos.Max(),
                categorias,
                produtoMaisCaro = _mapper.Map<ProdutoResponseDTO>(produtos.OrderByDescending(p => p.Preco).First()),
                produtoMaisBarato = _mapper.Map<ProdutoResponseDTO>(produtos.OrderBy(p => p.Preco).First())
            };
        }

        private async Task ValidateBusinessRules(Produto produto)
        {
            // Regra: Nome do produto deve ser único (case-insensitive)
            var allProdutos = await _repository.GetAllAsync();
            var existingProduct = allProdutos.FirstOrDefault(p => 
                p.Nome.ToLower() == produto.Nome.ToLower() && p.Id != produto.Id);
            
            if (existingProduct != null)
                throw new BusinessException("Já existe um produto com este nome");

            // Regra: Preço deve ser positivo
            if (produto.Preco <= 0)
                throw new BusinessException("Preço deve ser maior que zero");

            // Regra: Categoria deve ter pelo menos 2 caracteres
            if (string.IsNullOrWhiteSpace(produto.Categoria) || produto.Categoria.Length < 2)
                throw new BusinessException("Categoria deve ter pelo menos 2 caracteres");
        }
    }
}