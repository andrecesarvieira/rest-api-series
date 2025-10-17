using System.ComponentModel.DataAnnotations;

namespace ProdutosApi.DTOs
{
    public class ProdutoRequestDTO
    {
        [Required(ErrorMessage = "Nome é obrigatório")]
        [StringLength(100, MinimumLength = 2, ErrorMessage = "Nome deve ter entre 2 e 100 caracteres")]
        public string Nome { get; set; } = string.Empty;

        [StringLength(500, ErrorMessage = "Descrição deve ter no máximo 500 caracteres")]
        public string? Descricao { get; set; }

        [Required(ErrorMessage = "Preço é obrigatório")]
        [Range(0.01, double.MaxValue, ErrorMessage = "Preço deve ser maior que zero")]
        public decimal Preco { get; set; }

        [Required(ErrorMessage = "Categoria é obrigatória")]
        [StringLength(50, MinimumLength = 2, ErrorMessage = "Categoria deve ter entre 2 e 50 caracteres")]
        public string Categoria { get; set; } = string.Empty;

        public bool Ativo { get; set; } = true;
    }
}