export interface Tarefa {
  id: string;
  titulo: string;
  descricao?: string;
  concluida: boolean;
  prioridade: 'baixa' | 'media' | 'alta';
  categoria?: string;
  dataCriacao: Date;
  dataAtualizacao: Date;
  dataVencimento?: Date;
}

export interface TarefaCreateRequest {
  titulo: string;
  descricao?: string;
  prioridade?: 'baixa' | 'media' | 'alta';
  categoria?: string;
  dataVencimento?: string;
}

export interface TarefaUpdateRequest {
  titulo?: string;
  descricao?: string;
  concluida?: boolean;
  prioridade?: 'baixa' | 'media' | 'alta';
  categoria?: string;
  dataVencimento?: string;
}

export interface TarefaResponse {
  id: string;
  titulo: string;
  descricao?: string;
  concluida: boolean;
  prioridade: 'baixa' | 'media' | 'alta';
  categoria?: string;
  dataCriacao: string;
  dataAtualizacao: string;
  dataVencimento?: string;
}