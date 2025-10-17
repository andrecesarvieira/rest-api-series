# Script para verificar links nos READMEs

$readmeFiles = Get-ChildItem -Path . -Recurse -Filter "README.md"
$issues = @()

foreach ($readme in $readmeFiles) {
    $content = Get-Content $readme.FullName -Raw
    $relPath = $readme.FullName.Replace((Get-Location).Path, ".")
    
    # Verificar links relativos para LICENSE
    if ($content -match '\[LICENSE\]\(\.\./LICENSE\)') {
        $issues += "❌ $relPath - Link incorreto para LICENSE: ../LICENSE (deveria ser ../../LICENSE ou /LICENSE)"
    }
    
    if ($content -match '\[LICENSE\]\(\.\./\.\./LICENSE\)') {
        $issues += "⚠️ $relPath - Link para LICENSE pode não funcionar no GitHub: ../../LICENSE"
    }
    
    # Verificar badges vazios
    $emptyBadges = ([regex]::Matches($content, '\]\(\)') | Measure-Object).Count
    if ($emptyBadges -gt 0) {
        $issues += "⚠️ $relPath - Contém $emptyBadges badge(s) com URL vazia"
    }
    
    # Verificar links para GitHub do usuário
    if ($content -match '\[GitHub\]\(#\)' -or $content -match '\[LinkedIn\]\(#\)') {
        $issues += "ℹ️ $relPath - Contém links placeholder (#) para GitHub/LinkedIn"
    }
}

if ($issues.Count -eq 0) {
    Write-Host "✅ Nenhum problema encontrado!" -ForegroundColor Green
} else {
    Write-Host "`n🔍 Problemas encontrados:`n" -ForegroundColor Yellow
    $issues | ForEach-Object { Write-Host $_ }
}
