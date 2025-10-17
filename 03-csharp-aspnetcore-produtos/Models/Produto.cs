using System.ComponentModel.DataAnnotations;

namespace ProdutosApi.Models
{
    public class Produto
    {
        public int Id { get; set; }

        [Required]
        [StringLength(100, MinimumLength = 2)]
        public string Nome { get; set; } = string.Empty;

        [StringLength(500)]
        public string? Descricao { get; set; }

        [Required]
        [Range(0.01, double.MaxValue)]
        public decimal Preco { get; set; }

        [Required]
        [StringLength(50)]
        public string Categoria { get; set; } = string.Empty;

        public DateTime DataCriacao { get; set; }

        public bool Ativo { get; set; } = true;
    }
}