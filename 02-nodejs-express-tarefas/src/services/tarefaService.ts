import { Tarefa, TarefaCreateRequest, TarefaUpdateRequest, TarefaResponse } from '../models/Tarefa';
import { dataStore } from '../storage/dataStore';

/**
 * Service responsável pela lógica de negócio das tarefas
 */
export class TarefaService {
  
  /**
   * Converte uma Tarefa para TarefaResponse (formatação de datas)
   */
  private toTarefaResponse(tarefa: Tarefa): TarefaResponse {
    return {
      id: tarefa.id,
      titulo: tarefa.titulo,
      descricao: tarefa.descricao,
      concluida: tarefa.concluida,
      prioridade: tarefa.prioridade,
      categoria: tarefa.categoria,
      dataCriacao: tarefa.dataCriacao.toISOString(),
      dataAtualizacao: tarefa.dataAtualizacao.toISOString(),
      dataVencimento: tarefa.dataVencimento?.toISOString()
    };
  }

  /**
   * Lista todas as tarefas com filtros e paginação
   */
  async listarTarefas(filters: {
    concluida?: boolean;
    prioridade?: 'baixa' | 'media' | 'alta';
    categoria?: string;
    page?: number;
    limit?: number;
  } = {}): Promise<{
    data: TarefaResponse[];
    pagination: {
      current: number;
      total: number;
      totalItems: number;
      hasNext: boolean;
      hasPrev: boolean;
    };
  }> {
    // Aplicar filtros
    let tarefas = dataStore.getAll();
    
    if (filters.concluida !== undefined || filters.prioridade || filters.categoria) {
      tarefas = dataStore.filter({
        concluida: filters.concluida,
        prioridade: filters.prioridade,
        categoria: filters.categoria
      });
    }

    // Ordenar por data de criação (mais recentes primeiro)
    tarefas.sort((a, b) => b.dataCriacao.getTime() - a.dataCriacao.getTime());

    // Aplicar paginação
    const page = filters.page || 1;
    const limit = filters.limit || 10;
    const paginatedResult = dataStore.paginate(tarefas, page, limit);

    return {
      data: paginatedResult.data.map(tarefa => this.toTarefaResponse(tarefa)),
      pagination: paginatedResult.pagination
    };
  }

  /**
   * Busca uma tarefa por ID
   */
  async buscarTarefaPorId(id: string): Promise<TarefaResponse | null> {
    const tarefa = dataStore.getById(id);
    
    if (!tarefa) {
      return null;
    }

    return this.toTarefaResponse(tarefa);
  }

  /**
   * Cria uma nova tarefa
   */
  async criarTarefa(dadosTarefa: TarefaCreateRequest): Promise<TarefaResponse> {
    // Validações de negócio
    await this.validarDadosTarefa(dadosTarefa);

    const novaTarefa = dataStore.create({
      titulo: dadosTarefa.titulo,
      descricao: dadosTarefa.descricao,
      concluida: false, // Sempre inicia como não concluída
      prioridade: dadosTarefa.prioridade || 'media',
      categoria: dadosTarefa.categoria,
      dataVencimento: dadosTarefa.dataVencimento ? new Date(dadosTarefa.dataVencimento) : undefined
    });

    return this.toTarefaResponse(novaTarefa);
  }

  /**
   * Atualiza uma tarefa existente
   */
  async atualizarTarefa(id: string, dadosAtualizacao: TarefaUpdateRequest): Promise<TarefaResponse | null> {
    const tarefaExistente = dataStore.getById(id);
    
    if (!tarefaExistente) {
      return null;
    }

    // Validações de negócio
    await this.validarAtualizacaoTarefa(dadosAtualizacao, tarefaExistente);

    const dadosParaAtualizar: Partial<Tarefa> = {};

    if (dadosAtualizacao.titulo !== undefined) {
      dadosParaAtualizar.titulo = dadosAtualizacao.titulo;
    }

    if (dadosAtualizacao.descricao !== undefined) {
      dadosParaAtualizar.descricao = dadosAtualizacao.descricao;
    }

    if (dadosAtualizacao.concluida !== undefined) {
      dadosParaAtualizar.concluida = dadosAtualizacao.concluida;
    }

    if (dadosAtualizacao.prioridade !== undefined) {
      dadosParaAtualizar.prioridade = dadosAtualizacao.prioridade;
    }

    if (dadosAtualizacao.categoria !== undefined) {
      dadosParaAtualizar.categoria = dadosAtualizacao.categoria;
    }

    if (dadosAtualizacao.dataVencimento !== undefined) {
      dadosParaAtualizar.dataVencimento = dadosAtualizacao.dataVencimento 
        ? new Date(dadosAtualizacao.dataVencimento) 
        : undefined;
    }

    const tarefaAtualizada = dataStore.update(id, dadosParaAtualizar);

    return tarefaAtualizada ? this.toTarefaResponse(tarefaAtualizada) : null;
  }

  /**
   * Remove uma tarefa
   */
  async removerTarefa(id: string): Promise<boolean> {
    return dataStore.delete(id);
  }

  /**
   * Marca uma tarefa como concluída
   */
  async concluirTarefa(id: string): Promise<TarefaResponse | null> {
    return this.atualizarTarefa(id, { concluida: true });
  }

  /**
   * Marca uma tarefa como não concluída
   */
  async reativarTarefa(id: string): Promise<TarefaResponse | null> {
    return this.atualizarTarefa(id, { concluida: false });
  }

  /**
   * Obtém estatísticas das tarefas
   */
  async obterEstatisticas(): Promise<{
    total: number;
    concluidas: number;
    pendentes: number;
    porPrioridade: Record<'baixa' | 'media' | 'alta', number>;
    porCategoria: Record<string, number>;
    proximosVencimentos: TarefaResponse[];
  }> {
    const stats = dataStore.getStats();
    
    // Buscar tarefas que vencem nos próximos 7 dias
    const agora = new Date();
    const seteDiasDepois = new Date(agora.getTime() + 7 * 24 * 60 * 60 * 1000);
    
    const proximosVencimentos = dataStore.getAll()
      .filter(tarefa => 
        !tarefa.concluida &&
        tarefa.dataVencimento &&
        tarefa.dataVencimento >= agora &&
        tarefa.dataVencimento <= seteDiasDepois
      )
      .sort((a, b) => (a.dataVencimento!.getTime() - b.dataVencimento!.getTime()))
      .slice(0, 5) // Máximo 5 tarefas
      .map(tarefa => this.toTarefaResponse(tarefa));

    return {
      ...stats,
      proximosVencimentos
    };
  }

  /**
   * Validações de negócio para criação de tarefa
   */
  private async validarDadosTarefa(dados: TarefaCreateRequest): Promise<void> {
    // Verificar se já existe tarefa com o mesmo título
    const tarefasExistentes = dataStore.getAll();
    const tituloJaExiste = tarefasExistentes.some(
      tarefa => tarefa.titulo.toLowerCase() === dados.titulo.toLowerCase()
    );

    if (tituloJaExiste) {
      throw new Error('Já existe uma tarefa com este título');
    }

    // Validar data de vencimento
    if (dados.dataVencimento) {
      const dataVencimento = new Date(dados.dataVencimento);
      const agora = new Date();
      
      if (dataVencimento < agora) {
        throw new Error('Data de vencimento não pode ser no passado');
      }
    }
  }

  /**
   * Validações de negócio para atualização de tarefa
   */
  private async validarAtualizacaoTarefa(dados: TarefaUpdateRequest, tarefaExistente: Tarefa): Promise<void> {
    // Verificar se título já existe (exceto na própria tarefa)
    if (dados.titulo) {
      const tarefasExistentes = dataStore.getAll();
      const tituloJaExiste = tarefasExistentes.some(
        tarefa => 
          tarefa.id !== tarefaExistente.id &&
          tarefa.titulo.toLowerCase() === dados.titulo!.toLowerCase()
      );

      if (tituloJaExiste) {
        throw new Error('Já existe uma tarefa com este título');
      }
    }

    // Validar data de vencimento
    if (dados.dataVencimento) {
      const dataVencimento = new Date(dados.dataVencimento);
      const agora = new Date();
      
      if (dataVencimento < agora) {
        throw new Error('Data de vencimento não pode ser no passado');
      }
    }
  }
}

// Singleton instance
export const tarefaService = new TarefaService();
export default tarefaService;