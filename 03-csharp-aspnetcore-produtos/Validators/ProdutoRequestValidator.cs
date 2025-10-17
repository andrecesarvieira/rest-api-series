using FluentValidation;
using ProdutosApi.DTOs;

namespace ProdutosApi.Validators
{
    public class ProdutoRequestValidator : AbstractValidator<ProdutoRequestDTO>
    {
        public ProdutoRequestValidator()
        {
            RuleFor(x => x.Nome)
                .NotEmpty().WithMessage("Nome é obrigatório")
                .Length(2, 100).WithMessage("Nome deve ter entre 2 e 100 caracteres")
                .Matches(@"^[a-zA-ZÀ-ÿ0-9\s\-\.]+$").WithMessage("Nome contém caracteres inválidos");

            RuleFor(x => x.Descricao)
                .MaximumLength(500).WithMessage("Descrição deve ter no máximo 500 caracteres");

            RuleFor(x => x.Preco)
                .GreaterThan(0).WithMessage("Preço deve ser maior que zero")
                .LessThanOrEqualTo(999999.99m).WithMessage("Preço deve ser menor ou igual a 999.999,99");

            RuleFor(x => x.Categoria)
                .NotEmpty().WithMessage("Categoria é obrigatória")
                .Length(2, 50).WithMessage("Categoria deve ter entre 2 e 50 caracteres")
                .Matches(@"^[a-zA-ZÀ-ÿ\s\-]+$").WithMessage("Categoria deve conter apenas letras, espaços e hífens");
        }
    }
}