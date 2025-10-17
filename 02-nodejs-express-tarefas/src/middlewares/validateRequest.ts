import { Request, Response, NextFunction } from 'express';
import { ZodSchema, ZodError } from 'zod';

/**
 * Middleware genérico para validação usando Zod
 */
export function validateRequest(schema: {
  body?: ZodSchema;
  params?: ZodSchema;
  query?: ZodSchema;
}) {
  return (req: Request, res: Response, next: NextFunction) => {
    try {
      // Validar corpo da requisição
      if (schema.body) {
        req.body = schema.body.parse(req.body);
      }

      // Validar parâmetros da URL
      if (schema.params) {
        req.params = schema.params.parse(req.params);
      }

      // Validar query parameters
      if (schema.query) {
        req.query = schema.query.parse(req.query);
      }

      next();
    } catch (error) {
      if (error instanceof ZodError) {
        const formattedErrors = error.errors.map(err => ({
          field: err.path.join('.'),
          message: err.message,
          code: err.code
        }));

        return res.status(400).json({
          error: 'Dados inválidos',
          message: 'Os dados fornecidos não atendem aos requisitos',
          details: formattedErrors,
          timestamp: new Date().toISOString()
        });
      }

      // Erro não esperado
      next(error);
    }
  };
}

/**
 * Middleware específico para validação de corpo da requisição
 */
export function validateBody(schema: ZodSchema) {
  return validateRequest({ body: schema });
}

/**
 * Middleware específico para validação de parâmetros
 */
export function validateParams(schema: ZodSchema) {
  return validateRequest({ params: schema });
}

/**
 * Middleware específico para validação de query parameters
 */
export function validateQuery(schema: ZodSchema) {
  return validateRequest({ query: schema });
}