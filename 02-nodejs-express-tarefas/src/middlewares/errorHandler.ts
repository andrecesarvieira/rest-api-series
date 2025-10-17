import { Request, Response, NextFunction } from 'express';
import { ZodError } from 'zod';

/**
 * Interface para erros customizados da aplicação
 */
export interface AppError extends Error {
  statusCode?: number;
  isOperational?: boolean;
}

/**
 * Classe para erros customizados da aplicação
 */
export class CustomError extends Error implements AppError {
  statusCode: number;
  isOperational: boolean;

  constructor(message: string, statusCode: number = 500, isOperational: boolean = true) {
    super(message);
    this.name = this.constructor.name;
    this.statusCode = statusCode;
    this.isOperational = isOperational;

    Error.captureStackTrace(this, this.constructor);
  }
}

/**
 * Middleware global de tratamento de erros
 */
export function errorHandler(
  error: Error | AppError | ZodError,
  req: Request,
  res: Response,
  next: NextFunction
): void {
  console.error('🚨 Error caught by middleware:', {
    name: error.name,
    message: error.message,
    stack: error.stack,
    url: req.originalUrl,
    method: req.method,
    timestamp: new Date().toISOString()
  });

  // Erro de validação do Zod (não deveria chegar aqui se o middleware de validação estiver funcionando)
  if (error instanceof ZodError) {
    const formattedErrors = error.errors.map(err => ({
      field: err.path.join('.'),
      message: err.message,
      code: err.code
    }));

    res.status(400).json({
      error: 'Erro de validação',
      message: 'Os dados fornecidos são inválidos',
      details: formattedErrors,
      timestamp: new Date().toISOString()
    });
    return;
  }

  // Erro customizado da aplicação
  if (error instanceof CustomError || (error as AppError).isOperational) {
    const statusCode = (error as AppError).statusCode || 500;
    
    res.status(statusCode).json({
      error: getErrorTypeByStatusCode(statusCode),
      message: error.message,
      timestamp: new Date().toISOString(),
      path: req.originalUrl,
      method: req.method
    });
    return;
  }

  // Erro interno do servidor (não operacional)
  const isDevelopment = process.env.NODE_ENV === 'development';
  
  res.status(500).json({
    error: 'Erro interno do servidor',
    message: isDevelopment 
      ? error.message 
      : 'Algo deu errado. Tente novamente mais tarde.',
    timestamp: new Date().toISOString(),
    path: req.originalUrl,
    method: req.method,
    ...(isDevelopment && { stack: error.stack })
  });
}

/**
 * Middleware para capturar erros assíncronos
 */
export function asyncErrorHandler(
  fn: (req: Request, res: Response, next: NextFunction) => Promise<any>
) {
  return (req: Request, res: Response, next: NextFunction) => {
    Promise.resolve(fn(req, res, next)).catch(next);
  };
}

/**
 * Função utilitária para criar erros específicos
 */
export function createError(message: string, statusCode: number = 500): CustomError {
  return new CustomError(message, statusCode);
}

/**
 * Funções utilitárias para erros comuns
 */
export function notFoundError(resource: string = 'Recurso'): CustomError {
  return new CustomError(`${resource} não encontrado`, 404);
}

export function badRequestError(message: string): CustomError {
  return new CustomError(message, 400);
}

export function unauthorizedError(message: string = 'Não autorizado'): CustomError {
  return new CustomError(message, 401);
}

export function forbiddenError(message: string = 'Acesso negado'): CustomError {
  return new CustomError(message, 403);
}

export function conflictError(message: string): CustomError {
  return new CustomError(message, 409);
}

export function internalServerError(message: string = 'Erro interno do servidor'): CustomError {
  return new CustomError(message, 500, false);
}

/**
 * Função auxiliar para obter o tipo de erro baseado no status code
 */
function getErrorTypeByStatusCode(statusCode: number): string {
  switch (statusCode) {
    case 400: return 'Requisição inválida';
    case 401: return 'Não autorizado';
    case 403: return 'Acesso negado';
    case 404: return 'Não encontrado';
    case 409: return 'Conflito';
    case 422: return 'Entidade não processável';
    case 429: return 'Muitas requisições';
    case 500: return 'Erro interno do servidor';
    case 502: return 'Gateway inválido';
    case 503: return 'Serviço indisponível';
    default: return 'Erro';
  }
}

/**
 * Middleware para lidar com rotas não encontradas
 */
export function notFoundHandler(req: Request, res: Response, next: NextFunction): void {
  const error = new CustomError(
    `A rota ${req.method} ${req.originalUrl} não foi encontrada`,
    404
  );
  next(error);
}