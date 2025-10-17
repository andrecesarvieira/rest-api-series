from datetime import datetime
from typing import List, Optional, Dict
from enum import Enum

class UserProfile(str, Enum):
    ADMIN = "admin"
    USER = "user"
    MODERATOR = "moderator"

class User:
    """Modelo simples de usuário para armazenamento em memória"""
    
    def __init__(self, **kwargs):
        self.id = kwargs.get('id')
        self.nome = kwargs.get('nome', '')
        self.email = kwargs.get('email', '')
        self.idade = kwargs.get('idade', 0)
        self.ativo = kwargs.get('ativo', True)
        self.data_criacao = kwargs.get('data_criacao', datetime.utcnow())
        self.perfil = kwargs.get('perfil', UserProfile.USER)
    
    def to_dict(self) -> Dict:
        """Converte para dicionário"""
        return {
            'id': self.id,
            'nome': self.nome,
            'email': self.email,
            'idade': self.idade,
            'ativo': self.ativo,
            'data_criacao': self.data_criacao,
            'perfil': self.perfil
        }
    
    def __repr__(self):
        return f"<User(id={self.id}, nome='{self.nome}', email='{self.email}')>"