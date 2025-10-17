using ProdutosApi.DTOs;

namespace ProdutosApi.Services
{
    public interface IProdutoService
    {
        Task<IEnumerable<ProdutoResponseDTO>> GetAllAsync();
        Task<ProdutoResponseDTO?> GetByIdAsync(int id);
        Task<ProdutoResponseDTO> CreateAsync(ProdutoRequestDTO produtoRequest);
        Task<ProdutoResponseDTO?> UpdateAsync(int id, ProdutoRequestDTO produtoRequest);
        Task<bool> DeleteAsync(int id);
        Task<IEnumerable<ProdutoResponseDTO>> GetByCategoryAsync(string categoria);
        Task<IEnumerable<ProdutoResponseDTO>> GetByPriceRangeAsync(decimal minPreco, decimal maxPreco);
        Task<IEnumerable<ProdutoResponseDTO>> GetActiveAsync();
        Task<PagedResult<ProdutoResponseDTO>> GetPaginatedAsync(int page, int size, string? categoria = null, string? nome = null, bool? ativo = null);
        Task<object> GetStatisticsAsync();
    }

    public class PagedResult<T>
    {
        public IEnumerable<T> Data { get; set; } = new List<T>();
        public int Page { get; set; }
        public int Size { get; set; }
        public int Total { get; set; }
        public int TotalPages { get; set; }
        public bool HasPrevious { get; set; }
        public bool HasNext { get; set; }
    }
}