using Microsoft.EntityFrameworkCore;
using ProdutosApi.Data;
using ProdutosApi.Models;

namespace ProdutosApi.Repositories
{
    public class ProdutoRepository : IProdutoRepository
    {
        private readonly ProdutoContext _context;

        public ProdutoRepository(ProdutoContext context)
        {
            _context = context;
        }

        public async Task<IEnumerable<Produto>> GetAllAsync()
        {
            return await _context.Produtos
                .OrderBy(p => p.Nome)
                .ToListAsync();
        }

        public async Task<Produto?> GetByIdAsync(int id)
        {
            return await _context.Produtos
                .FirstOrDefaultAsync(p => p.Id == id);
        }

        public async Task<Produto> CreateAsync(Produto produto)
        {
            produto.DataCriacao = DateTime.UtcNow;
            _context.Produtos.Add(produto);
            await _context.SaveChangesAsync();
            return produto;
        }

        public async Task<Produto?> UpdateAsync(int id, Produto produto)
        {
            var existingProduto = await GetByIdAsync(id);
            if (existingProduto == null)
                return null;

            existingProduto.Nome = produto.Nome;
            existingProduto.Descricao = produto.Descricao;
            existingProduto.Preco = produto.Preco;
            existingProduto.Categoria = produto.Categoria;
            existingProduto.Ativo = produto.Ativo;

            await _context.SaveChangesAsync();
            return existingProduto;
        }

        public async Task<bool> DeleteAsync(int id)
        {
            var produto = await GetByIdAsync(id);
            if (produto == null)
                return false;

            _context.Produtos.Remove(produto);
            await _context.SaveChangesAsync();
            return true;
        }

        public async Task<IEnumerable<Produto>> GetByCategoryAsync(string categoria)
        {
            return await _context.Produtos
                .Where(p => p.Categoria.ToLower() == categoria.ToLower())
                .OrderBy(p => p.Nome)
                .ToListAsync();
        }

        public async Task<IEnumerable<Produto>> GetByPriceRangeAsync(decimal minPreco, decimal maxPreco)
        {
            return await _context.Produtos
                .Where(p => p.Preco >= minPreco && p.Preco <= maxPreco)
                .OrderBy(p => p.Preco)
                .ToListAsync();
        }

        public async Task<IEnumerable<Produto>> GetActiveAsync()
        {
            return await _context.Produtos
                .Where(p => p.Ativo)
                .OrderBy(p => p.Nome)
                .ToListAsync();
        }

        public async Task<(IEnumerable<Produto> produtos, int total)> GetPaginatedAsync(
            int page, int size, string? categoria = null, string? nome = null, bool? ativo = null)
        {
            var query = _context.Produtos.AsQueryable();

            // Aplicar filtros
            if (!string.IsNullOrWhiteSpace(categoria))
                query = query.Where(p => p.Categoria.ToLower().Contains(categoria.ToLower()));

            if (!string.IsNullOrWhiteSpace(nome))
                query = query.Where(p => p.Nome.ToLower().Contains(nome.ToLower()));

            if (ativo.HasValue)
                query = query.Where(p => p.Ativo == ativo.Value);

            var total = await query.CountAsync();

            var produtos = await query
                .OrderBy(p => p.Nome)
                .Skip((page - 1) * size)
                .Take(size)
                .ToListAsync();

            return (produtos, total);
        }
    }
}