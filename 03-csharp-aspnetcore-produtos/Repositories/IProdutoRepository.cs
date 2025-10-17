using ProdutosApi.Models;

namespace ProdutosApi.Repositories
{
    public interface IProdutoRepository
    {
        Task<IEnumerable<Produto>> GetAllAsync();
        Task<Produto?> GetByIdAsync(int id);
        Task<Produto> CreateAsync(Produto produto);
        Task<Produto?> UpdateAsync(int id, Produto produto);
        Task<bool> DeleteAsync(int id);
        Task<IEnumerable<Produto>> GetByCategoryAsync(string categoria);
        Task<IEnumerable<Produto>> GetByPriceRangeAsync(decimal minPreco, decimal maxPreco);
        Task<IEnumerable<Produto>> GetActiveAsync();
        Task<(IEnumerable<Produto> produtos, int total)> GetPaginatedAsync(int page, int size, string? categoria = null, string? nome = null, bool? ativo = null);
    }
}