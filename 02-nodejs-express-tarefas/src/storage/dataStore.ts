import { Tarefa } from '../models/Tarefa';
import { v4 as uuidv4 } from 'uuid';

/**
 * Armazenamento de dados em memória para as tarefas
 * Em um ambiente de produção, isso seria substituído por um banco de dados
 */
class DataStore {
  private tarefas: Map<string, Tarefa> = new Map();
  
  constructor() {
    this.initializeWithSampleData();
  }

  /**
   * Inicializa o store com dados de exemplo
   */
  private initializeWithSampleData(): void {
    const sampleTarefas: Omit<Tarefa, 'id'>[] = [
      {
        titulo: 'Implementar API REST',
        descricao: 'Desenvolver API completa usando Node.js e Express',
        concluida: false,
        prioridade: 'alta',
        categoria: 'Desenvolvimento',
        dataCriacao: new Date('2025-10-16T10:00:00Z'),
        dataAtualizacao: new Date('2025-10-16T10:00:00Z'),
        dataVencimento: new Date('2025-10-20T23:59:59Z')
      },
      {
        titulo: 'Escrever documentação',
        descricao: 'Criar documentação técnica da API',
        concluida: false,
        prioridade: 'media',
        categoria: 'Documentação',
        dataCriacao: new Date('2025-10-17T09:00:00Z'),
        dataAtualizacao: new Date('2025-10-17T09:00:00Z'),
        dataVencimento: new Date('2025-10-25T17:00:00Z')
      },
      {
        titulo: 'Configurar CI/CD',
        descricao: 'Implementar pipeline de integração e deploy contínuo',
        concluida: true,
        prioridade: 'alta',
        categoria: 'DevOps',
        dataCriacao: new Date('2025-10-15T14:30:00Z'),
        dataAtualizacao: new Date('2025-10-16T16:45:00Z')
      }
    ];

    sampleTarefas.forEach(tarefaData => {
      const id = uuidv4();
      this.tarefas.set(id, { ...tarefaData, id });
    });

    console.log(`📋 DataStore inicializado com ${this.tarefas.size} tarefas de exemplo`);
  }

  /**
   * Busca todas as tarefas
   */
  getAll(): Tarefa[] {
    return Array.from(this.tarefas.values());
  }

  /**
   * Busca tarefa por ID
   */
  getById(id: string): Tarefa | undefined {
    return this.tarefas.get(id);
  }

  /**
   * Cria nova tarefa
   */
  create(tarefaData: Omit<Tarefa, 'id' | 'dataCriacao' | 'dataAtualizacao'>): Tarefa {
    const id = uuidv4();
    const now = new Date();
    
    const novaTarefa: Tarefa = {
      ...tarefaData,
      id,
      dataCriacao: now,
      dataAtualizacao: now
    };

    this.tarefas.set(id, novaTarefa);
    return novaTarefa;
  }

  /**
   * Atualiza tarefa existente
   */
  update(id: string, updateData: Partial<Omit<Tarefa, 'id' | 'dataCriacao' | 'dataAtualizacao'>>): Tarefa | undefined {
    const tarefaExistente = this.tarefas.get(id);
    
    if (!tarefaExistente) {
      return undefined;
    }

    const tarefaAtualizada: Tarefa = {
      ...tarefaExistente,
      ...updateData,
      dataAtualizacao: new Date()
    };

    this.tarefas.set(id, tarefaAtualizada);
    return tarefaAtualizada;
  }

  /**
   * Remove tarefa por ID
   */
  delete(id: string): boolean {
    return this.tarefas.delete(id);
  }

  /**
   * Filtra tarefas com base nos critérios fornecidos
   */
  filter(criteria: {
    concluida?: boolean;
    prioridade?: 'baixa' | 'media' | 'alta';
    categoria?: string;
  }): Tarefa[] {
    const tarefas = this.getAll();

    return tarefas.filter(tarefa => {
      if (criteria.concluida !== undefined && tarefa.concluida !== criteria.concluida) {
        return false;
      }
      
      if (criteria.prioridade && tarefa.prioridade !== criteria.prioridade) {
        return false;
      }
      
      if (criteria.categoria && tarefa.categoria !== criteria.categoria) {
        return false;
      }
      
      return true;
    });
  }

  /**
   * Paginação de tarefas
   */
  paginate(tarefas: Tarefa[], page: number = 1, limit: number = 10): {
    data: Tarefa[];
    pagination: {
      current: number;
      total: number;
      totalItems: number;
      hasNext: boolean;
      hasPrev: boolean;
    };
  } {
    const startIndex = (page - 1) * limit;
    const endIndex = startIndex + limit;
    const paginatedTarefas = tarefas.slice(startIndex, endIndex);
    const totalPages = Math.ceil(tarefas.length / limit);

    return {
      data: paginatedTarefas,
      pagination: {
        current: page,
        total: totalPages,
        totalItems: tarefas.length,
        hasNext: page < totalPages,
        hasPrev: page > 1
      }
    };
  }

  /**
   * Retorna estatísticas das tarefas
   */
  getStats(): {
    total: number;
    concluidas: number;
    pendentes: number;
    porPrioridade: Record<'baixa' | 'media' | 'alta', number>;
    porCategoria: Record<string, number>;
  } {
    const tarefas = this.getAll();
    const stats = {
      total: tarefas.length,
      concluidas: tarefas.filter(t => t.concluida).length,
      pendentes: tarefas.filter(t => !t.concluida).length,
      porPrioridade: { baixa: 0, media: 0, alta: 0 },
      porCategoria: {} as Record<string, number>
    };

    tarefas.forEach(tarefa => {
      stats.porPrioridade[tarefa.prioridade]++;
      
      if (tarefa.categoria) {
        stats.porCategoria[tarefa.categoria] = (stats.porCategoria[tarefa.categoria] || 0) + 1;
      }
    });

    return stats;
  }

  /**
   * Limpa todas as tarefas (útil para testes)
   */
  clear(): void {
    this.tarefas.clear();
  }

  /**
   * Retorna o número total de tarefas
   */
  count(): number {
    return this.tarefas.size;
  }
}

// Singleton instance
export const dataStore = new DataStore();
export default dataStore;