import { Request, Response, NextFunction } from 'express';
import { tarefaService } from '../services/tarefaService';
import { 
  asyncErrorHandler, 
  notFoundError, 
  badRequestError, 
  createError 
} from '../middlewares/errorHandler';
import { 
  TarefaCreateInput, 
  TarefaUpdateInput, 
  TarefaQueryInput, 
  TarefaParamsInput 
} from '../schemas/tarefaSchema';
import { TarefaCreateRequest, TarefaUpdateRequest } from '../models/Tarefa';

/**
 * Controller responsável por gerenciar as rotas e requisições das tarefas
 */
export class TarefaController {

  /**
   * GET /api/tarefas
   * Lista todas as tarefas com filtros e paginação
   */
  listarTarefas = asyncErrorHandler(async (req: Request, res: Response): Promise<void> => {
    const query = req.query as unknown as TarefaQueryInput;
    
    const resultado = await tarefaService.listarTarefas({
      concluida: query.concluida,
      prioridade: query.prioridade,
      categoria: query.categoria,
      page: query.page,
      limit: query.limit
    });

    res.status(200).json({
      message: 'Tarefas recuperadas com sucesso',
      data: resultado.data,
      pagination: resultado.pagination,
      timestamp: new Date().toISOString()
    });
  });

  /**
   * GET /api/tarefas/:id
   * Busca uma tarefa específica por ID
   */
  buscarTarefaPorId = asyncErrorHandler(async (req: Request, res: Response): Promise<void> => {
    const { id } = req.params as TarefaParamsInput;

    const tarefa = await tarefaService.buscarTarefaPorId(id);

    if (!tarefa) {
      throw notFoundError('Tarefa');
    }

    res.status(200).json({
      message: 'Tarefa encontrada com sucesso',
      data: tarefa,
      timestamp: new Date().toISOString()
    });
  });

  /**
   * POST /api/tarefas
   * Cria uma nova tarefa
   */
  criarTarefa = asyncErrorHandler(async (req: Request, res: Response): Promise<void> => {
    const dadosTarefa = req.body as TarefaCreateInput;

    // Converter dados do schema para o formato esperado pelo service
    const dadosParaService: TarefaCreateRequest = {
      titulo: dadosTarefa.titulo,
      descricao: dadosTarefa.descricao,
      prioridade: dadosTarefa.prioridade,
      categoria: dadosTarefa.categoria,
      dataVencimento: dadosTarefa.dataVencimento?.toISOString()
    };

    try {
      const novaTarefa = await tarefaService.criarTarefa(dadosParaService);

      res.status(201).json({
        message: 'Tarefa criada com sucesso',
        data: novaTarefa,
        timestamp: new Date().toISOString()
      });
    } catch (error) {
      if (error instanceof Error) {
        throw badRequestError(error.message);
      }
      throw error;
    }
  });

  /**
   * PUT /api/tarefas/:id
   * Atualiza uma tarefa existente
   */
  atualizarTarefa = asyncErrorHandler(async (req: Request, res: Response): Promise<void> => {
    const { id } = req.params as TarefaParamsInput;
    const dadosAtualizacao = req.body as TarefaUpdateInput;

    // Converter dados do schema para o formato esperado pelo service
    const dadosParaService: TarefaUpdateRequest = {
      titulo: dadosAtualizacao.titulo,
      descricao: dadosAtualizacao.descricao,
      concluida: dadosAtualizacao.concluida,
      prioridade: dadosAtualizacao.prioridade,
      categoria: dadosAtualizacao.categoria,
      dataVencimento: dadosAtualizacao.dataVencimento?.toISOString()
    };

    try {
      const tarefaAtualizada = await tarefaService.atualizarTarefa(id, dadosParaService);

      if (!tarefaAtualizada) {
        throw notFoundError('Tarefa');
      }

      res.status(200).json({
        message: 'Tarefa atualizada com sucesso',
        data: tarefaAtualizada,
        timestamp: new Date().toISOString()
      });
    } catch (error) {
      if (error instanceof Error && error.message !== 'Tarefa não encontrada') {
        throw badRequestError(error.message);
      }
      throw error;
    }
  });

  /**
   * DELETE /api/tarefas/:id
   * Remove uma tarefa
   */
  removerTarefa = asyncErrorHandler(async (req: Request, res: Response): Promise<void> => {
    const { id } = req.params as TarefaParamsInput;

    const removido = await tarefaService.removerTarefa(id);

    if (!removido) {
      throw notFoundError('Tarefa');
    }

    res.status(200).json({
      message: 'Tarefa removida com sucesso',
      timestamp: new Date().toISOString()
    });
  });

  /**
   * PATCH /api/tarefas/:id/concluir
   * Marca uma tarefa como concluída
   */
  concluirTarefa = asyncErrorHandler(async (req: Request, res: Response): Promise<void> => {
    const { id } = req.params as TarefaParamsInput;

    const tarefaAtualizada = await tarefaService.concluirTarefa(id);

    if (!tarefaAtualizada) {
      throw notFoundError('Tarefa');
    }

    res.status(200).json({
      message: 'Tarefa marcada como concluída',
      data: tarefaAtualizada,
      timestamp: new Date().toISOString()
    });
  });

  /**
   * PATCH /api/tarefas/:id/reativar
   * Marca uma tarefa como não concluída
   */
  reativarTarefa = asyncErrorHandler(async (req: Request, res: Response): Promise<void> => {
    const { id } = req.params as TarefaParamsInput;

    const tarefaAtualizada = await tarefaService.reativarTarefa(id);

    if (!tarefaAtualizada) {
      throw notFoundError('Tarefa');
    }

    res.status(200).json({
      message: 'Tarefa reativada com sucesso',
      data: tarefaAtualizada,
      timestamp: new Date().toISOString()
    });
  });

  /**
   * GET /api/tarefas/estatisticas
   * Obtém estatísticas das tarefas
   */
  obterEstatisticas = asyncErrorHandler(async (req: Request, res: Response): Promise<void> => {
    const estatisticas = await tarefaService.obterEstatisticas();

    res.status(200).json({
      message: 'Estatísticas recuperadas com sucesso',
      data: estatisticas,
      timestamp: new Date().toISOString()
    });
  });
}

// Singleton instance
export const tarefaController = new TarefaController();
export default tarefaController;