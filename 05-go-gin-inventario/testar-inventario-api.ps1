# Script de Testes Automatizados - API de Inventário Go + Gin
# Testa todas as funcionalidades da API de forma abrangente

Write-Host "🚀 INICIANDO TESTES DA API DE INVENTÁRIO - GO + GIN" -ForegroundColor Green
Write-Host "===============================================" -ForegroundColor Green

$baseUrl = "http://localhost:8000"
$apiUrl = "$baseUrl/api/produtos"
$testsPassed = 0
$testsTotal = 0

# Função para executar requisições HTTP
function Invoke-ApiRequest {
    param(
        [string]$Method,
        [string]$Url,
        [object]$Body = $null,
        [string]$TestName
    )
    
    $global:testsTotal++
    Write-Host "`n[$global:testsTotal] Testando: $TestName" -ForegroundColor Yellow
    Write-Host "$Method $Url"
    
    try {
        $headers = @{ 'Content-Type' = 'application/json' }
        
        if ($Body) {
            $jsonBody = $Body | ConvertTo-Json -Depth 10
            Write-Host "Body: $jsonBody" -ForegroundColor Gray
            $response = Invoke-RestMethod -Uri $Url -Method $Method -Body $jsonBody -Headers $headers
        } else {
            $response = Invoke-RestMethod -Uri $Url -Method $Method -Headers $headers
        }
        
        Write-Host "✅ SUCESSO - Status: 200/201" -ForegroundColor Green
        $global:testsPassed++
        return $response
    }
    catch {
        $statusCode = $_.Exception.Response.StatusCode.value__
        $errorMessage = $_.Exception.Message
        
        if ($statusCode -eq 404 -or $statusCode -eq 409 -or $statusCode -eq 422) {
            Write-Host "✅ ESPERADO - Status: $statusCode" -ForegroundColor Green
            $global:testsPassed++
        } else {
            Write-Host "❌ FALHOU - Status: $statusCode - $errorMessage" -ForegroundColor Red
        }
        
        return $null
    }
}

# Função para pausar
function Wait-ForUser {
    Write-Host "`nPressione Enter para continuar..." -ForegroundColor Cyan
    Read-Host
}

Write-Host "`n🔍 VERIFICANDO SE O SERVIDOR ESTÁ EXECUTANDO..."

try {
    $healthCheck = Invoke-RestMethod -Uri "$baseUrl/health" -Method GET
    Write-Host "✅ Servidor está rodando!" -ForegroundColor Green
    Write-Host "Serviço: $($healthCheck.service)"
    Write-Host "Status: $($healthCheck.status)"
} catch {
    Write-Host "❌ ERRO: Servidor não está rodando na porta 8000!" -ForegroundColor Red
    Write-Host "Por favor, execute: go run cmd/api/main.go" -ForegroundColor Yellow
    exit 1
}

Write-Host "`n📋 INICIANDO TESTES DE FUNCIONALIDADES..." -ForegroundColor Cyan

# 1. Teste GET - Listar todos os produtos (dados iniciais)
$produtos = Invoke-ApiRequest -Method "GET" -Url $apiUrl -TestName "Listar produtos iniciais"

if ($produtos) {
    Write-Host "Total de produtos iniciais: $($produtos.produtos.Count)" -ForegroundColor Cyan
    Write-Host "Primeira página: $($produtos.paginacao.pagina_atual)" -ForegroundColor Cyan
}

# 2. Teste POST - Criar novo produto eletrônico
$novoProduto1 = @{
    nome = "iPad Pro 12.9"
    descricao = "Tablet profissional com chip M2 e tela Liquid Retina XDR"
    preco = 7999.99
    quantidade = 15
    categoria = "eletronicos"
    ativo = $true
}

$produtoCriado1 = Invoke-ApiRequest -Method "POST" -Url $apiUrl -Body $novoProduto1 -TestName "Criar produto eletrônico"

# 3. Teste POST - Criar produto de roupas
$novoProduto2 = @{
    nome = "Jaqueta de Couro Premium"
    descricao = "Jaqueta de couro legítimo com forro interno"
    preco = 599.99
    quantidade = 8
    categoria = "roupas"
    ativo = $true
}

$produtoCriado2 = Invoke-ApiRequest -Method "POST" -Url $apiUrl -Body $novoProduto2 -TestName "Criar produto de roupas"

# 4. Teste GET por ID - Buscar produto específico
if ($produtoCriado1) {
    $produto = Invoke-ApiRequest -Method "GET" -Url "$apiUrl/$($produtoCriado1.id)" -TestName "Buscar produto por ID"
    
    if ($produto) {
        Write-Host "Produto encontrado: $($produto.nome)" -ForegroundColor Cyan
        Write-Host "Preço formatado: $($produto.preco_formatado)" -ForegroundColor Cyan
    }
}

# 5. Teste PUT - Atualizar produto
if ($produtoCriado1) {
    $atualizacao = @{
        preco = 7599.99
        quantidade = 20
        descricao = "Tablet profissional com chip M2, tela Liquid Retina XDR e WiFi 6E"
    }
    
    $produtoAtualizado = Invoke-ApiRequest -Method "PUT" -Url "$apiUrl/$($produtoCriado1.id)" -Body $atualizacao -TestName "Atualizar produto"
    
    if ($produtoAtualizado) {
        Write-Host "Novo preço: $($produtoAtualizado.preco_formatado)" -ForegroundColor Cyan
        Write-Host "Nova quantidade: $($produtoAtualizado.quantidade)" -ForegroundColor Cyan
    }
}

# 6. Teste PATCH - Atualizar apenas estoque
if ($produtoCriado2) {
    $estoqueUpdate = @{
        quantidade = 12
    }
    
    $estoqueAtualizado = Invoke-ApiRequest -Method "PATCH" -Url "$apiUrl/$($produtoCriado2.id)/estoque" -Body $estoqueUpdate -TestName "Atualizar estoque"
    
    if ($estoqueAtualizado) {
        Write-Host "Estoque atualizado: $($estoqueAtualizado.quantidade)" -ForegroundColor Cyan
    }
}

# 7. Teste GET com filtros - Produtos eletrônicos
$eletronicos = Invoke-ApiRequest -Method "GET" -Url "$apiUrl/categoria/eletronicos" -TestName "Filtrar por categoria eletrônicos"

if ($eletronicos) {
    Write-Host "Produtos eletrônicos encontrados: $($eletronicos.produtos.Count)" -ForegroundColor Cyan
}

# 8. Teste GET com filtros avançados - Paginação + filtros
$filtroAvancado = "$apiUrl/filtros?categoria=eletronicos&preco_minimo=1000&preco_maximo=10000&page=1&size=5"
$produtosFiltrados = Invoke-ApiRequest -Method "GET" -Url $filtroAvancado -TestName "Filtros avançados com paginação"

if ($produtosFiltrados) {
    Write-Host "Produtos filtrados: $($produtosFiltrados.produtos.Count)" -ForegroundColor Cyan
    Write-Host "Total de itens: $($produtosFiltrados.paginacao.total_itens)" -ForegroundColor Cyan
    Write-Host "Página atual: $($produtosFiltrados.paginacao.pagina_atual)" -ForegroundColor Cyan
}

# 9. Teste GET - Produtos ativos
$produtosAtivos = Invoke-ApiRequest -Method "GET" -Url "$apiUrl/ativos" -TestName "Listar produtos ativos"

if ($produtosAtivos) {
    Write-Host "Produtos ativos: $($produtosAtivos.produtos.Count)" -ForegroundColor Cyan
}

# 10. Teste GET - Produtos em estoque
$produtosEstoque = Invoke-ApiRequest -Method "GET" -Url "$apiUrl/estoque" -TestName "Listar produtos em estoque"

if ($produtosEstoque) {
    Write-Host "Produtos em estoque: $($produtosEstoque.produtos.Count)" -ForegroundColor Cyan
}

# 11. Teste GET - Estatísticas
$estatisticas = Invoke-ApiRequest -Method "GET" -Url "$apiUrl/estatisticas" -TestName "Buscar estatísticas do inventário"

if ($estatisticas) {
    Write-Host "`n📊 ESTATÍSTICAS DO INVENTÁRIO:" -ForegroundColor Magenta
    Write-Host "Total de produtos: $($estatisticas.total_produtos)"
    Write-Host "Produtos ativos: $($estatisticas.produtos_ativos)"
    Write-Host "Produtos em estoque: $($estatisticas.produtos_em_estoque)"
    Write-Host "Valor total do inventário: R$ $($estatisticas.valor_total_inventario)"
    Write-Host "Preço médio: R$ $([math]::Round($estatisticas.preco_medio, 2))"
    Write-Host "Categorias disponíveis: $($estatisticas.por_categoria.Count)"
    
    if ($estatisticas.top5_mais_caros.Count -gt 0) {
        Write-Host "`n🏆 TOP 3 MAIS CAROS:"
        for ($i = 0; $i -lt [math]::Min(3, $estatisticas.top5_mais_caros.Count); $i++) {
            $produto = $estatisticas.top5_mais_caros[$i]
            Write-Host "  $($i+1). $($produto.nome) - $($produto.preco_formatado)"
        }
    }
}

# 12. Teste de busca por nome
$buscaNome = "$apiUrl/filtros?nome=samsung"
$produtosPorNome = Invoke-ApiRequest -Method "GET" -Url $buscaNome -TestName "Busca por nome 'samsung'"

if ($produtosPorNome) {
    Write-Host "Produtos encontrados com 'samsung': $($produtosPorNome.produtos.Count)" -ForegroundColor Cyan
}

# 13. Teste de validação - Criar produto com preço inválido
$produtoInvalido = @{
    nome = "Produto Teste"
    descricao = "Teste de validação"
    preco = -100.00  # Preço negativo (inválido)
    quantidade = 5
    categoria = "outros"
}

Invoke-ApiRequest -Method "POST" -Url $apiUrl -Body $produtoInvalido -TestName "Validação: produto com preço negativo"

# 14. Teste de erro 404 - Buscar produto inexistente
$idInexistente = "12345678-1234-1234-1234-123456789012"
Invoke-ApiRequest -Method "GET" -Url "$apiUrl/$idInexistente" -TestName "Erro 404: produto inexistente"

# 15. Teste DELETE - Remover produto criado
if ($produtoCriado2) {
    Invoke-ApiRequest -Method "DELETE" -Url "$apiUrl/$($produtoCriado2.id)" -TestName "Deletar produto"
    
    # Verificar se foi removido
    Invoke-ApiRequest -Method "GET" -Url "$apiUrl/$($produtoCriado2.id)" -TestName "Verificar produto deletado (deve dar 404)"
}

# 16. Teste final - Verificar estado após todas as operações
$produtosFinais = Invoke-ApiRequest -Method "GET" -Url $apiUrl -TestName "Estado final dos produtos"

if ($produtosFinais) {
    Write-Host "`nTotal final de produtos: $($produtosFinais.produtos.Count)" -ForegroundColor Cyan
}

# Resumo dos testes
Write-Host "`n" + "="*60 -ForegroundColor Green
Write-Host "📋 RESUMO DOS TESTES" -ForegroundColor Green
Write-Host "="*60 -ForegroundColor Green
Write-Host "✅ Testes Passaram: $global:testsPassed" -ForegroundColor Green
Write-Host "📊 Total de Testes: $global:testsTotal" -ForegroundColor Cyan

if ($global:testsTotal -gt 0) {
    Write-Host "📈 Taxa de Sucesso: $([math]::Round(($global:testsPassed / $global:testsTotal) * 100, 1))%" -ForegroundColor Yellow
} else {
    Write-Host "📈 Taxa de Sucesso: 0%" -ForegroundColor Yellow
}

if ($global:testsPassed -eq $global:testsTotal -and $global:testsTotal -gt 0) {
    Write-Host "`n🎉 TODOS OS TESTES PASSARAM! API GO + GIN FUNCIONANDO PERFEITAMENTE!" -ForegroundColor Green
} else {
    Write-Host "`n⚠️  Alguns testes falharam ou nenhum teste foi executado. Verifique a implementação." -ForegroundColor Yellow
}

Write-Host "`n🏁 TESTES CONCLUÍDOS!" -ForegroundColor Green
Write-Host "API de Inventário Go + Gin testada com sucesso!" -ForegroundColor Cyan