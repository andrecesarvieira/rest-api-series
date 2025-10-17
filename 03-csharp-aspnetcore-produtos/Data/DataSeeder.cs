using ProdutosApi.Models;

namespace ProdutosApi.Data
{
    public static class DataSeeder
    {
        public static void SeedData(ProdutoContext context)
        {
            if (context.Produtos.Any())
                return;

            var produtos = new[]
            {
                new Produto
                {
                    Nome = "Smartphone Samsung Galaxy",
                    Descricao = "Smartphone com tela de 6.1 polegadas e 128GB",
                    Preco = 1299.90m,
                    Categoria = "Eletrônicos",
                    DataCriacao = DateTime.UtcNow,
                    Ativo = true
                },
                new Produto
                {
                    Nome = "Notebook Dell Inspiron",
                    Descricao = "Notebook Intel i5 8GB RAM 256GB SSD",
                    Preco = 2499.90m,
                    Categoria = "Informática",
                    DataCriacao = DateTime.UtcNow,
                    Ativo = true
                },
                new Produto
                {
                    Nome = "Fone de Ouvido Sony",
                    Descricao = "Fone com cancelamento de ruído ativo",
                    Preco = 349.90m,
                    Categoria = "Áudio",
                    DataCriacao = DateTime.UtcNow,
                    Ativo = true
                },
                new Produto
                {
                    Nome = "Câmera Canon EOS",
                    Descricao = "Câmera DSLR 24MP com lente 18-55mm",
                    Preco = 1899.90m,
                    Categoria = "Fotografia",
                    DataCriacao = DateTime.UtcNow,
                    Ativo = true
                },
                new Produto
                {
                    Nome = "Smart TV LG 55\"",
                    Descricao = "TV 4K UHD com WebOS e HDR",
                    Preco = 2199.90m,
                    Categoria = "Eletrônicos",
                    DataCriacao = DateTime.UtcNow,
                    Ativo = true
                }
            };

            context.Produtos.AddRange(produtos);
            context.SaveChanges();
        }
    }
}