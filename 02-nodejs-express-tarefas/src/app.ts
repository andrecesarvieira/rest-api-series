import express, { Express } from 'express';
import cors from 'cors';
import helmet from 'helmet';
import morgan from 'morgan';
import tarefaRoutes from './routes/tarefaRoutes';
import { errorHandler } from './middlewares/errorHandler';

const app: Express = express();
const PORT = process.env.PORT || 3000;

// Middlewares de segurança e logging
app.use(helmet());
app.use(cors());
app.use(morgan('combined'));

// Parser JSON
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// Rota de health check
app.get('/health', (req, res) => {
  res.status(200).json({
    status: 'OK',
    message: 'API de Tarefas está funcionando!',
    timestamp: new Date().toISOString(),
    environment: process.env.NODE_ENV || 'development'
  });
});

// Rotas da API
app.use('/api/tarefas', tarefaRoutes);

// Rota 404 - não encontrada
app.use('*', (req, res) => {
  res.status(404).json({
    error: 'Endpoint não encontrado',
    message: `A rota ${req.method} ${req.originalUrl} não existe`,
    availableRoutes: [
      'GET /health - Health check da API',
      'GET /api/tarefas - Listar todas as tarefas',
      'POST /api/tarefas - Criar nova tarefa',
      'GET /api/tarefas/:id - Buscar tarefa por ID',
      'PUT /api/tarefas/:id - Atualizar tarefa',
      'DELETE /api/tarefas/:id - Deletar tarefa'
    ]
  });
});

// Middleware de tratamento de erros (deve vir por último)
app.use(errorHandler);

// Inicialização do servidor
if (require.main === module) {
  app.listen(PORT, () => {
    console.log(`
🚀 Servidor Express iniciado com sucesso!
📡 API: http://localhost:${PORT}
🏥 Health Check: http://localhost:${PORT}/health
📋 Tarefas API: http://localhost:${PORT}/api/tarefas
🌍 Ambiente: ${process.env.NODE_ENV || 'development'}
⏰ Timestamp: ${new Date().toISOString()}
    `);
  });
}

export default app;
