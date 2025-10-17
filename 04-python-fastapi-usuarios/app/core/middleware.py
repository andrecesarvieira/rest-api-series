from fastapi import Request, HTTPException
from fastapi.responses import JSONResponse
from fastapi.middleware.cors import CORSMiddleware
from starlette.middleware.base import BaseHTTPMiddleware
import logging
import time
from typing import Callable

# Configurar logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

class ExceptionHandlerMiddleware(BaseHTTPMiddleware):
    """Middleware global para tratamento de exceções"""
    
    async def dispatch(self, request: Request, call_next: Callable):
        try:
            response = await call_next(request)
            return response
        except HTTPException as exc:
            # HTTPException já tratada pelo FastAPI
            raise exc
        except Exception as exc:
            logger.error(f"Erro não tratado: {str(exc)}")
            return JSONResponse(
                status_code=500,
                content={
                    "error": "Erro interno do servidor",
                    "message": "Ocorreu um erro inesperado. Tente novamente mais tarde.",
                    "timestamp": time.time()
                }
            )

class LoggingMiddleware(BaseHTTPMiddleware):
    """Middleware para logging de requisições"""
    
    async def dispatch(self, request: Request, call_next: Callable):
        start_time = time.time()
        
        # Log da requisição
        logger.info(f"🔄 {request.method} {request.url}")
        
        response = await call_next(request)
        
        # Log da resposta
        process_time = time.time() - start_time
        logger.info(f"✅ {request.method} {request.url} - {response.status_code} - {process_time:.3f}s")
        
        return response

def configure_cors(app):
    """Configura CORS para a aplicação"""
    app.add_middleware(
        CORSMiddleware,
        allow_origins=["*"],  # Em produção, especificar origens
        allow_credentials=True,
        allow_methods=["*"],
        allow_headers=["*"],
    )

def configure_middlewares(app):
    """Configura todos os middlewares"""
    
    # Middleware de logging
    app.add_middleware(LoggingMiddleware)
    
    # Middleware de exceções
    app.add_middleware(ExceptionHandlerMiddleware)
    
    # CORS
    configure_cors(app)