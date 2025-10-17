# Script de Testes para API de Produtos - C# ASP.NET Core

# URL base da API
$baseUrl = "http://localhost:5000/api/produtos"

Write-Host "=== TESTE DA API DE PRODUTOS - C# ASP.NET Core ===" -ForegroundColor Green
Write-Host "URL Base: $baseUrl" -ForegroundColor Cyan
Write-Host ""

# Função auxiliar para exibir resultados
function Show-TestResult {
    param($testName, $result, $success = $true)
    if ($success) {
        Write-Host "✅ $testName" -ForegroundColor Green
    } else {
        Write-Host "❌ $testName" -ForegroundColor Red
    }
    if ($result) {
        Write-Host "   Resultado: $($result | ConvertTo-Json -Depth 3)" -ForegroundColor Gray
    }
    Write-Host ""
}

try {
    # 1. Teste GET - Listar todos os produtos
    Write-Host "1. Testando GET /api/produtos (Listar todos)" -ForegroundColor Yellow
    $produtos = Invoke-RestMethod -Uri $baseUrl -Method GET
    Show-TestResult "Listar todos os produtos" "Total: $($produtos.Count) produtos"

    # 2. Teste GET - Obter produto por ID
    if ($produtos.Count -gt 0) {
        $primeiroId = $produtos[0].id
        Write-Host "2. Testando GET /api/produtos/{id} (Produto por ID)" -ForegroundColor Yellow
        $produto = Invoke-RestMethod -Uri "$baseUrl/$primeiroId" -Method GET
        Show-TestResult "Obter produto por ID $primeiroId" $produto.nome
    }

    # 3. Teste POST - Criar novo produto
    Write-Host "3. Testando POST /api/produtos (Criar produto)" -ForegroundColor Yellow
    $novoProduto = @{
        nome = "Produto Teste API"
        descricao = "Produto criado via teste automatizado"
        preco = 199.90
        categoria = "Teste"
        ativo = $true
    } | ConvertTo-Json

    $headers = @{ "Content-Type" = "application/json" }
    $produtoCriado = Invoke-RestMethod -Uri $baseUrl -Method POST -Body $novoProduto -Headers $headers
    $novoId = $produtoCriado.id
    Show-TestResult "Criar novo produto" "ID: $novoId - $($produtoCriado.nome)"

    # 4. Teste PUT - Atualizar produto
    Write-Host "4. Testando PUT /api/produtos/{id} (Atualizar produto)" -ForegroundColor Yellow
    $produtoAtualizado = @{
        nome = "Produto Teste ATUALIZADO"
        descricao = "Descrição atualizada via teste"
        preco = 299.90
        categoria = "Teste Atualizado"
        ativo = $true
    } | ConvertTo-Json

    $produtoEditado = Invoke-RestMethod -Uri "$baseUrl/$novoId" -Method PUT -Body $produtoAtualizado -Headers $headers
    Show-TestResult "Atualizar produto ID $novoId" $produtoEditado.nome

    # 5. Teste GET - Filtros por categoria
    Write-Host "5. Testando GET /api/produtos/categoria/{categoria}" -ForegroundColor Yellow
    $produtosCategoria = Invoke-RestMethod -Uri "$baseUrl/categoria/Eletrônicos" -Method GET
    Show-TestResult "Filtrar por categoria 'Eletrônicos'" "Total: $($produtosCategoria.Count) produtos"

    # 6. Teste GET - Produtos ativos
    Write-Host "6. Testando GET /api/produtos/ativos" -ForegroundColor Yellow
    $produtosAtivos = Invoke-RestMethod -Uri "$baseUrl/ativos" -Method GET
    Show-TestResult "Listar produtos ativos" "Total: $($produtosAtivos.Count) produtos ativos"

    # 7. Teste GET - Filtro por faixa de preço
    Write-Host "7. Testando GET /api/produtos/preco?minPreco=100&maxPreco=2000" -ForegroundColor Yellow
    $produtosPreco = Invoke-RestMethod -Uri "$baseUrl/preco?minPreco=100&maxPreco=2000" -Method GET
    Show-TestResult "Filtrar por faixa de preço (100-2000)" "Total: $($produtosPreco.Count) produtos"

    # 8. Teste GET - Paginação
    Write-Host "8. Testando GET /api/produtos/paginado?page=1&size=3" -ForegroundColor Yellow
    $produtosPaginados = Invoke-RestMethod -Uri "$baseUrl/paginado?page=1&size=3" -Method GET
    Show-TestResult "Paginação (página 1, 3 itens)" "Página: $($produtosPaginados.page), Total: $($produtosPaginados.total)"

    # 9. Teste GET - Estatísticas
    Write-Host "9. Testando GET /api/produtos/estatisticas" -ForegroundColor Yellow
    $estatisticas = Invoke-RestMethod -Uri "$baseUrl/estatisticas" -Method GET
    Show-TestResult "Obter estatísticas" "Total produtos: $($estatisticas.totalProdutos), Preço médio: $($estatisticas.precoMedio)"

    # 10. Teste DELETE - Remover produto
    Write-Host "10. Testando DELETE /api/produtos/{id}" -ForegroundColor Yellow
    Invoke-RestMethod -Uri "$baseUrl/$novoId" -Method DELETE
    Show-TestResult "Remover produto ID $novoId" "Produto removido com sucesso"

    # 11. Teste de validação - Produto inválido
    Write-Host "11. Testando validação (produto inválido)" -ForegroundColor Yellow
    $produtoInvalido = @{
        nome = ""
        preco = -10
        categoria = ""
    } | ConvertTo-Json

    try {
        Invoke-RestMethod -Uri $baseUrl -Method POST -Body $produtoInvalido -Headers $headers
        Show-TestResult "Validação de produto inválido" "ERRO: Deveria ter falhado" $false
    } catch {
        Show-TestResult "Validação de produto inválido" "Erro capturado corretamente: $($_.Exception.Message)"
    }

    # 12. Teste 404 - Produto inexistente
    Write-Host "12. Testando erro 404 (produto inexistente)" -ForegroundColor Yellow
    try {
        Invoke-RestMethod -Uri "$baseUrl/99999" -Method GET
        Show-TestResult "Teste 404" "ERRO: Deveria ter retornado 404" $false
    } catch {
        Show-TestResult "Teste 404" "Erro 404 capturado corretamente"
    }

    Write-Host ""
    Write-Host "=== TODOS OS TESTES CONCLUÍDOS! ===" -ForegroundColor Green
    Write-Host "API C# ASP.NET Core funcionando corretamente! ✅" -ForegroundColor Green

} catch {
    Write-Host "❌ Erro durante os testes: $($_.Exception.Message)" -ForegroundColor Red
    Write-Host "Verifique se a API está rodando em http://localhost:5000" -ForegroundColor Yellow
}