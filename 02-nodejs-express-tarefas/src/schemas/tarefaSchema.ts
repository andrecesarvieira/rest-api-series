import { z } from 'zod';

// Schema para validar criação de tarefa
export const tarefaCreateSchema = z.object({
  titulo: z
    .string({ required_error: 'Título é obrigatório' })
    .min(1, 'Título não pode estar vazio')
    .max(200, 'Título deve ter no máximo 200 caracteres')
    .trim(),
  
  descricao: z
    .string()
    .max(1000, 'Descrição deve ter no máximo 1000 caracteres')
    .trim()
    .optional(),
  
  prioridade: z
    .enum(['baixa', 'media', 'alta'], { 
      errorMap: () => ({ message: 'Prioridade deve ser: baixa, media ou alta' }) 
    })
    .default('media'),
  
  categoria: z
    .string()
    .max(50, 'Categoria deve ter no máximo 50 caracteres')
    .trim()
    .optional(),
  
  dataVencimento: z
    .string()
    .datetime({ message: 'Data de vencimento deve estar no formato ISO 8601' })
    .optional()
    .transform((val) => val ? new Date(val) : undefined)
});

// Schema para validar atualização de tarefa
export const tarefaUpdateSchema = z.object({
  titulo: z
    .string()
    .min(1, 'Título não pode estar vazio')
    .max(200, 'Título deve ter no máximo 200 caracteres')
    .trim()
    .optional(),
  
  descricao: z
    .string()
    .max(1000, 'Descrição deve ter no máximo 1000 caracteres')
    .trim()
    .optional(),
  
  concluida: z
    .boolean({ errorMap: () => ({ message: 'Concluída deve ser true ou false' }) })
    .optional(),
  
  prioridade: z
    .enum(['baixa', 'media', 'alta'], { 
      errorMap: () => ({ message: 'Prioridade deve ser: baixa, media ou alta' }) 
    })
    .optional(),
  
  categoria: z
    .string()
    .max(50, 'Categoria deve ter no máximo 50 caracteres')
    .trim()
    .optional(),
  
  dataVencimento: z
    .string()
    .datetime({ message: 'Data de vencimento deve estar no formato ISO 8601' })
    .optional()
    .transform((val) => val ? new Date(val) : undefined)
});

// Schema para validar parâmetros de query
export const tarefaQuerySchema = z.object({
  concluida: z
    .string()
    .optional()
    .transform((val) => {
      if (val === 'true') return true;
      if (val === 'false') return false;
      return undefined;
    }),
  
  prioridade: z
    .enum(['baixa', 'media', 'alta'])
    .optional(),
  
  categoria: z
    .string()
    .trim()
    .optional(),
  
  page: z
    .string()
    .optional()
    .transform((val) => val ? parseInt(val, 10) : 1)
    .refine((val) => val > 0, { message: 'Página deve ser maior que 0' }),
  
  limit: z
    .string()
    .optional()
    .transform((val) => val ? parseInt(val, 10) : 10)
    .refine((val) => val > 0 && val <= 100, { 
      message: 'Limite deve ser entre 1 e 100' 
    })
});

// Schema para validar parâmetros de rota (ID)
export const tarefaParamsSchema = z.object({
  id: z
    .string({ required_error: 'ID é obrigatório' })
    .uuid('ID deve ser um UUID válido')
});

export type TarefaCreateInput = z.infer<typeof tarefaCreateSchema>;
export type TarefaUpdateInput = z.infer<typeof tarefaUpdateSchema>;
export type TarefaQueryInput = z.infer<typeof tarefaQuerySchema>;
export type TarefaParamsInput = z.infer<typeof tarefaParamsSchema>;