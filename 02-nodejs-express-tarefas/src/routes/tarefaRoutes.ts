import { Router } from 'express';
import { tarefaController } from '../controllers/tarefaController';
import { validateRequest, validateParams, validateQuery, validateBody } from '../middlewares/validateRequest';
import {
  tarefaCreateSchema,
  tarefaUpdateSchema,
  tarefaQuerySchema,
  tarefaParamsSchema
} from '../schemas/tarefaSchema';

const router = Router();

/**
 * GET /estatisticas
 * Rota para estatísticas (deve vir antes da rota /:id)
 */
router.get('/estatisticas', tarefaController.obterEstatisticas);

/**
 * GET /
 * Lista todas as tarefas com filtros e paginação
 */
router.get(
  '/',
  validateQuery(tarefaQuerySchema),
  tarefaController.listarTarefas
);

/**
 * POST /
 * Cria uma nova tarefa
 */
router.post(
  '/',
  validateBody(tarefaCreateSchema),
  tarefaController.criarTarefa
);

/**
 * GET /:id
 * Busca uma tarefa por ID
 */
router.get(
  '/:id',
  validateParams(tarefaParamsSchema),
  tarefaController.buscarTarefaPorId
);

/**
 * PUT /:id
 * Atualiza uma tarefa
 */
router.put(
  '/:id',
  validateRequest({
    params: tarefaParamsSchema,
    body: tarefaUpdateSchema
  }),
  tarefaController.atualizarTarefa
);

/**
 * DELETE /:id
 * Remove uma tarefa
 */
router.delete(
  '/:id',
  validateParams(tarefaParamsSchema),
  tarefaController.removerTarefa
);

/**
 * PATCH /:id/concluir
 * Marca uma tarefa como concluída
 */
router.patch(
  '/:id/concluir',
  validateParams(tarefaParamsSchema),
  tarefaController.concluirTarefa
);

/**
 * PATCH /:id/reativar
 * Reativa uma tarefa (marca como não concluída)
 */
router.patch(
  '/:id/reativar',
  validateParams(tarefaParamsSchema),
  tarefaController.reativarTarefa
);

export default router;