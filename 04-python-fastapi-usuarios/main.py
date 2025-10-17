from fastapi import FastAPI
from app.routers import users
from app.core.middleware import configure_middlewares

# Criar aplicação FastAPI
app = FastAPI(
    title="API de Usuários - FastAPI",
    description="Uma API REST completa para gerenciamento de usuários desenvolvida em Python com FastAPI",
    version="1.0.0",
    docs_url="/docs",
    redoc_url="/redoc"
)

# Configurar middlewares
configure_middlewares(app)

# Incluir routers
app.include_router(users.router)

@app.on_event("startup")
async def startup_event():
    """Evento executado na inicialização da aplicação"""
    print("🚀 Iniciando API de Usuários...")
    print("✅ Banco de dados em memória inicializado!")
    print("✅ Dados de exemplo carregados!")
    print("🎉 API de Usuários iniciada com sucesso!")

@app.get("/", tags=["Root"])
async def root():
    """Endpoint raiz da API"""
    return {
        "message": "API de Usuários - FastAPI",
        "version": "1.0.0",
        "docs": "/docs",
        "redoc": "/redoc",
        "status": "running"
    }

@app.get("/health", tags=["Health"])
async def health_check():
    """Endpoint para verificação de saúde da API"""
    return {
        "status": "healthy",
        "api": "usuarios-api",
        "framework": "FastAPI",
        "python": "3.11+"
    }

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(
        "main:app",
        host="0.0.0.0",
        port=8000,
        reload=True,
        log_level="info"
    )