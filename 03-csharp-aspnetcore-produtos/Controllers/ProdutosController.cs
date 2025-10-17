using Microsoft.AspNetCore.Mvc;
using ProdutosApi.DTOs;
using ProdutosApi.Services;
using ProdutosApi.Exceptions;

namespace ProdutosApi.Controllers
{
    [ApiController]
    [Route("api/[controller]")]
    [Produces("application/json")]
    public class ProdutosController : ControllerBase
    {
        private readonly IProdutoService _produtoService;
        private readonly ILogger<ProdutosController> _logger;

        public ProdutosController(IProdutoService produtoService, ILogger<ProdutosController> logger)
        {
            _produtoService = produtoService;
            _logger = logger;
        }

        /// <summary>
        /// Obtém todos os produtos
        /// </summary>
        /// <returns>Lista de produtos</returns>
        [HttpGet]
        public async Task<ActionResult<IEnumerable<ProdutoResponseDTO>>> GetAll()
        {
            _logger.LogInformation("Buscando todos os produtos");
            var produtos = await _produtoService.GetAllAsync();
            return Ok(produtos);
        }

        /// <summary>
        /// Obtém um produto por ID
        /// </summary>
        /// <param name="id">ID do produto</param>
        /// <returns>Produto encontrado</returns>
        [HttpGet("{id}")]
        public async Task<ActionResult<ProdutoResponseDTO>> GetById(int id)
        {
            _logger.LogInformation("Buscando produto com ID: {Id}", id);
            
            var produto = await _produtoService.GetByIdAsync(id);
            if (produto == null)
            {
                _logger.LogWarning("Produto com ID {Id} não encontrado", id);
                return NotFound(new { message = $"Produto com ID {id} não encontrado" });
            }

            return Ok(produto);
        }

        /// <summary>
        /// Cria um novo produto
        /// </summary>
        /// <param name="produtoRequest">Dados do produto a ser criado</param>
        /// <returns>Produto criado</returns>
        [HttpPost]
        public async Task<ActionResult<ProdutoResponseDTO>> Create([FromBody] ProdutoRequestDTO produtoRequest)
        {
            _logger.LogInformation("Criando novo produto: {Nome}", produtoRequest.Nome);

            var produto = await _produtoService.CreateAsync(produtoRequest);
            
            _logger.LogInformation("Produto criado com ID: {Id}", produto.Id);
            return CreatedAtAction(nameof(GetById), new { id = produto.Id }, produto);
        }

        /// <summary>
        /// Atualiza um produto existente
        /// </summary>
        /// <param name="id">ID do produto</param>
        /// <param name="produtoRequest">Dados atualizados do produto</param>
        /// <returns>Produto atualizado</returns>
        [HttpPut("{id}")]
        public async Task<ActionResult<ProdutoResponseDTO>> Update(int id, [FromBody] ProdutoRequestDTO produtoRequest)
        {
            _logger.LogInformation("Atualizando produto com ID: {Id}", id);

            var produto = await _produtoService.UpdateAsync(id, produtoRequest);
            if (produto == null)
            {
                _logger.LogWarning("Produto com ID {Id} não encontrado para atualização", id);
                return NotFound(new { message = $"Produto com ID {id} não encontrado" });
            }

            _logger.LogInformation("Produto com ID {Id} atualizado com sucesso", id);
            return Ok(produto);
        }

        /// <summary>
        /// Remove um produto
        /// </summary>
        /// <param name="id">ID do produto</param>
        /// <returns>Confirmação da remoção</returns>
        [HttpDelete("{id}")]
        public async Task<ActionResult> Delete(int id)
        {
            _logger.LogInformation("Removendo produto com ID: {Id}", id);

            var success = await _produtoService.DeleteAsync(id);
            if (!success)
            {
                _logger.LogWarning("Produto com ID {Id} não encontrado para remoção", id);
                return NotFound(new { message = $"Produto com ID {id} não encontrado" });
            }

            _logger.LogInformation("Produto com ID {Id} removido com sucesso", id);
            return NoContent();
        }

        /// <summary>
        /// Obtém produtos por categoria
        /// </summary>
        /// <param name="categoria">Nome da categoria</param>
        /// <returns>Lista de produtos da categoria</returns>
        [HttpGet("categoria/{categoria}")]
        public async Task<ActionResult<IEnumerable<ProdutoResponseDTO>>> GetByCategory(string categoria)
        {
            _logger.LogInformation("Buscando produtos da categoria: {Categoria}", categoria);
            
            var produtos = await _produtoService.GetByCategoryAsync(categoria);
            return Ok(produtos);
        }

        /// <summary>
        /// Obtém produtos por faixa de preço
        /// </summary>
        /// <param name="minPreco">Preço mínimo</param>
        /// <param name="maxPreco">Preço máximo</param>
        /// <returns>Lista de produtos na faixa de preço</returns>
        [HttpGet("preco")]
        public async Task<ActionResult<IEnumerable<ProdutoResponseDTO>>> GetByPriceRange(
            [FromQuery] decimal minPreco, 
            [FromQuery] decimal maxPreco)
        {
            _logger.LogInformation("Buscando produtos com preço entre {MinPreco} e {MaxPreco}", minPreco, maxPreco);
            
            var produtos = await _produtoService.GetByPriceRangeAsync(minPreco, maxPreco);
            return Ok(produtos);
        }

        /// <summary>
        /// Obtém apenas produtos ativos
        /// </summary>
        /// <returns>Lista de produtos ativos</returns>
        [HttpGet("ativos")]
        public async Task<ActionResult<IEnumerable<ProdutoResponseDTO>>> GetActive()
        {
            _logger.LogInformation("Buscando produtos ativos");
            
            var produtos = await _produtoService.GetActiveAsync();
            return Ok(produtos);
        }

        /// <summary>
        /// Obtém produtos com paginação
        /// </summary>
        /// <param name="page">Número da página (padrão: 1)</param>
        /// <param name="size">Tamanho da página (padrão: 10)</param>
        /// <param name="categoria">Filtro por categoria (opcional)</param>
        /// <param name="nome">Filtro por nome (opcional)</param>
        /// <param name="ativo">Filtro por status ativo (opcional)</param>
        /// <returns>Resultado paginado</returns>
        [HttpGet("paginado")]
        public async Task<ActionResult<PagedResult<ProdutoResponseDTO>>> GetPaginated(
            [FromQuery] int page = 1,
            [FromQuery] int size = 10,
            [FromQuery] string? categoria = null,
            [FromQuery] string? nome = null,
            [FromQuery] bool? ativo = null)
        {
            _logger.LogInformation("Buscando produtos paginados - Página: {Page}, Tamanho: {Size}", page, size);
            
            var result = await _produtoService.GetPaginatedAsync(page, size, categoria, nome, ativo);
            return Ok(result);
        }

        /// <summary>
        /// Obtém estatísticas dos produtos
        /// </summary>
        /// <returns>Estatísticas consolidadas</returns>
        [HttpGet("estatisticas")]
        public async Task<ActionResult<object>> GetStatistics()
        {
            _logger.LogInformation("Obtendo estatísticas dos produtos");
            
            var statistics = await _produtoService.GetStatisticsAsync();
            return Ok(statistics);
        }
    }
}