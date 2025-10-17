# Script de Testes para API de Usuários - Python FastAPI

# URL base da API
$baseUrl = "http://localhost:8000/api/usuarios"

Write-Host "=== TESTE DA API DE USUÁRIOS - Python FastAPI ===" -ForegroundColor Green
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
    # 1. Teste GET - Listar todos os usuários
    Write-Host "1. Testando GET /api/usuarios (Listar todos)" -ForegroundColor Yellow
    $usuarios = Invoke-RestMethod -Uri $baseUrl -Method GET
    Show-TestResult "Listar todos os usuários" "Total: $($usuarios.Count) usuários"

    # 2. Teste GET - Obter usuário por ID
    if ($usuarios.Count -gt 0) {
        $primeiroId = $usuarios[0].id
        Write-Host "2. Testando GET /api/usuarios/{id} (Usuário por ID)" -ForegroundColor Yellow
        $usuario = Invoke-RestMethod -Uri "$baseUrl/$primeiroId" -Method GET
        Show-TestResult "Obter usuário por ID $primeiroId" $usuario.nome
    }

    # 3. Teste POST - Criar novo usuário
    Write-Host "3. Testando POST /api/usuarios (Criar usuário)" -ForegroundColor Yellow
    $novoUsuario = @{
        nome = "Usuário Teste API"
        email = "teste.api@exemplo.com"
        idade = 30
        perfil = "user"
        ativo = $true
    } | ConvertTo-Json

    $headers = @{ "Content-Type" = "application/json" }
    $usuarioCriado = Invoke-RestMethod -Uri $baseUrl -Method POST -Body $novoUsuario -Headers $headers
    $novoId = $usuarioCriado.id
    Show-TestResult "Criar novo usuário" "ID: $novoId - $($usuarioCriado.nome)"

    # 4. Teste PUT - Atualizar usuário
    Write-Host "4. Testando PUT /api/usuarios/{id} (Atualizar usuário)" -ForegroundColor Yellow
    $usuarioAtualizado = @{
        nome = "Usuário Teste ATUALIZADO"
        email = "teste.atualizado@exemplo.com"
        idade = 35
        perfil = "moderator"
        ativo = $true
    } | ConvertTo-Json

    $usuarioEditado = Invoke-RestMethod -Uri "$baseUrl/$novoId" -Method PUT -Body $usuarioAtualizado -Headers $headers
    Show-TestResult "Atualizar usuário ID $novoId" $usuarioEditado.nome

    # 5. Teste GET - Filtros por perfil
    Write-Host "5. Testando GET /api/usuarios/perfil/{perfil}" -ForegroundColor Yellow
    $usuariosPerfil = Invoke-RestMethod -Uri "$baseUrl/perfil/user" -Method GET
    Show-TestResult "Filtrar por perfil 'user'" "Total: $($usuariosPerfil.Count) usuários"

    # 6. Teste GET - Usuários ativos
    Write-Host "6. Testando GET /api/usuarios/status/ativos" -ForegroundColor Yellow
    $usuariosAtivos = Invoke-RestMethod -Uri "$baseUrl/status/ativos" -Method GET
    Show-TestResult "Listar usuários ativos" "Total: $($usuariosAtivos.Count) usuários ativos"

    # 7. Teste GET - Filtro por faixa de idade
    Write-Host "7. Testando GET /api/usuarios/idade/faixa?min_idade=25&max_idade=35" -ForegroundColor Yellow
    $usuariosIdade = Invoke-RestMethod -Uri "$baseUrl/idade/faixa?min_idade=25&max_idade=35" -Method GET
    Show-TestResult "Filtrar por faixa de idade (25-35)" "Total: $($usuariosIdade.Count) usuários"

    # 8. Teste GET - Paginação
    Write-Host "8. Testando GET /api/usuarios/paginado/lista?page=1&size=3" -ForegroundColor Yellow
    $usuariosPaginados = Invoke-RestMethod -Uri "$baseUrl/paginado/lista?page=1&size=3" -Method GET
    Show-TestResult "Paginação (página 1, 3 itens)" "Página: $($usuariosPaginados.page), Total: $($usuariosPaginados.total)"

    # 9. Teste GET - Paginação com filtros
    Write-Host "9. Testando paginação com filtros (nome='João')" -ForegroundColor Yellow
    $usuariosFiltrados = Invoke-RestMethod -Uri "$baseUrl/paginado/lista?page=1&size=10&nome=João" -Method GET
    Show-TestResult "Paginação com filtro por nome" "Encontrados: $($usuariosFiltrados.users.Count) usuários"

    # 10. Teste GET - Estatísticas
    Write-Host "10. Testando GET /api/usuarios/estatisticas/dados" -ForegroundColor Yellow
    $estatisticas = Invoke-RestMethod -Uri "$baseUrl/estatisticas/dados" -Method GET
    Show-TestResult "Obter estatísticas" "Total usuários: $($estatisticas.total_usuarios), Idade média: $($estatisticas.idade_media)"

    # 11. Teste DELETE - Remover usuário
    Write-Host "11. Testando DELETE /api/usuarios/{id}" -ForegroundColor Yellow
    Invoke-RestMethod -Uri "$baseUrl/$novoId" -Method DELETE
    Show-TestResult "Remover usuário ID $novoId" "Usuário removido com sucesso"

    # 12. Teste de validação - Usuário inválido (email duplicado)
    Write-Host "12. Testando validação (email duplicado)" -ForegroundColor Yellow
    $usuarioInvalido = @{
        nome = "Teste Duplicado"
        email = "joao.silva@email.com"  # Email já existe
        idade = 25
        perfil = "user"
    } | ConvertTo-Json

    try {
        Invoke-RestMethod -Uri $baseUrl -Method POST -Body $usuarioInvalido -Headers $headers
        Show-TestResult "Validação de email duplicado" "ERRO: Deveria ter falhado" $false
    } catch {
        Show-TestResult "Validação de email duplicado" "Erro capturado corretamente: $($_.Exception.Message)"
    }

    # 13. Teste de validação - Dados inválidos
    Write-Host "13. Testando validação (dados inválidos)" -ForegroundColor Yellow
    $dadosInvalidos = @{
        nome = ""
        email = "email-invalido"
        idade = -5
    } | ConvertTo-Json

    try {
        Invoke-RestMethod -Uri $baseUrl -Method POST -Body $dadosInvalidos -Headers $headers
        Show-TestResult "Validação de dados inválidos" "ERRO: Deveria ter falhado" $false
    } catch {
        Show-TestResult "Validação de dados inválidos" "Erro capturado corretamente: $($_.Exception.Message)"
    }

    # 14. Teste 404 - Usuário inexistente
    Write-Host "14. Testando erro 404 (usuário inexistente)" -ForegroundColor Yellow
    try {
        Invoke-RestMethod -Uri "$baseUrl/99999" -Method GET
        Show-TestResult "Teste 404" "ERRO: Deveria ter retornado 404" $false
    } catch {
        Show-TestResult "Teste 404" "Erro 404 capturado corretamente"
    }

    Write-Host ""
    Write-Host "=== TODOS OS TESTES CONCLUÍDOS! ===" -ForegroundColor Green
    Write-Host "API Python FastAPI funcionando corretamente! ✅" -ForegroundColor Green

} catch {
    Write-Host "❌ Erro durante os testes: $($_.Exception.Message)" -ForegroundColor Red
    Write-Host "Verifique se a API está rodando em http://localhost:8000" -ForegroundColor Yellow
}