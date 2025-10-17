from typing import Dict, List, Optional
from app.models.user import User, UserProfile
from datetime import datetime

class InMemoryDatabase:
    """Banco de dados simples em memória"""
    
    def __init__(self):
        self.users: Dict[int, User] = {}
        self.next_id = 1
        self._populate_initial_data()
    
    def _populate_initial_data(self):
        """Popula dados iniciais"""
        initial_users = [
            {
                'nome': 'João Silva',
                'email': 'joao.silva@email.com',
                'idade': 28,
                'ativo': True,
                'perfil': UserProfile.ADMIN
            },
            {
                'nome': 'Maria Santos',
                'email': 'maria.santos@email.com',
                'idade': 32,
                'ativo': True,
                'perfil': UserProfile.USER
            },
            {
                'nome': 'Pedro Oliveira',
                'email': 'pedro.oliveira@email.com',
                'idade': 25,
                'ativo': True,
                'perfil': UserProfile.MODERATOR
            },
            {
                'nome': 'Ana Costa',
                'email': 'ana.costa@email.com',
                'idade': 29,
                'ativo': False,
                'perfil': UserProfile.USER
            },
            {
                'nome': 'Carlos Lima',
                'email': 'carlos.lima@email.com',
                'idade': 35,
                'ativo': True,
                'perfil': UserProfile.USER
            }
        ]
        
        for user_data in initial_users:
            user = User(**user_data)
            user.id = self.next_id
            user.data_criacao = datetime.utcnow()
            self.users[self.next_id] = user
            self.next_id += 1
    
    def get_all_users(self) -> List[User]:
        """Retorna todos os usuários"""
        return list(self.users.values())
    
    def get_user_by_id(self, user_id: int) -> Optional[User]:
        """Busca usuário por ID"""
        return self.users.get(user_id)
    
    def get_user_by_email(self, email: str) -> Optional[User]:
        """Busca usuário por email"""
        for user in self.users.values():
            if user.email.lower() == email.lower():
                return user
        return None
    
    def create_user(self, user_data: dict) -> User:
        """Cria um novo usuário"""
        user = User(**user_data)
        user.id = self.next_id
        user.data_criacao = datetime.utcnow()
        self.users[self.next_id] = user
        self.next_id += 1
        return user
    
    def update_user(self, user_id: int, user_data: dict) -> Optional[User]:
        """Atualiza um usuário"""
        user = self.users.get(user_id)
        if not user:
            return None
        
        for field, value in user_data.items():
            if hasattr(user, field) and value is not None:
                setattr(user, field, value)
        
        return user
    
    def delete_user(self, user_id: int) -> bool:
        """Remove um usuário"""
        if user_id in self.users:
            del self.users[user_id]
            return True
        return False

# Instância global do banco
database = InMemoryDatabase()

def get_database():
    """Dependency para obter o banco de dados"""
    return database