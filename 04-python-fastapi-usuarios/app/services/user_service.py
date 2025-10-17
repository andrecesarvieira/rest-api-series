from app.models.user import User, UserProfile
from app.schemas.user import UserCreate, UserUpdate, UserListResponse
from app.database.database import InMemoryDatabase
from fastapi import HTTPException, status
from typing import Optional, List
import math

class UserService:
    """Serviço para lógica de negócio dos usuários"""
    
    def __init__(self, db: InMemoryDatabase):
        self.db = db
    
    def get_all_users(self) -> List[User]:
        """Retorna todos os usuários"""
        users = self.db.get_all_users()
        return sorted(users, key=lambda x: x.nome)
    
    def get_user_by_id(self, user_id: int) -> Optional[User]:
        """Busca usuário por ID"""
        if user_id <= 0:
            raise HTTPException(
                status_code=status.HTTP_400_BAD_REQUEST,
                detail="ID deve ser um número positivo"
            )
        
        return self.db.get_user_by_id(user_id)
    
    def get_user_by_email(self, email: str) -> Optional[User]:
        """Busca usuário por email"""
        return self.db.get_user_by_email(email)
    
    def create_user(self, user_data: UserCreate) -> User:
        """Cria um novo usuário"""
        
        # Verifica se email já existe
        existing_user = self.get_user_by_email(user_data.email)
        if existing_user:
            raise HTTPException(
                status_code=status.HTTP_409_CONFLICT,
                detail=f"Email '{user_data.email}' já está em uso"
            )
        
        # Cria o usuário
        user_dict = user_data.model_dump()
        user_dict['email'] = user_dict['email'].lower()
        
        return self.db.create_user(user_dict)
    
    def update_user(self, user_id: int, user_data: UserUpdate) -> Optional[User]:
        """Atualiza um usuário existente"""
        
        existing_user = self.get_user_by_id(user_id)
        if not existing_user:
            return None
        
        # Verifica se novo email já existe (se fornecido)
        if user_data.email and user_data.email.lower() != existing_user.email:
            email_exists = self.get_user_by_email(user_data.email)
            if email_exists:
                raise HTTPException(
                    status_code=status.HTTP_409_CONFLICT,
                    detail=f"Email '{user_data.email}' já está em uso"
                )
        
        # Atualiza apenas campos fornecidos
        update_dict = user_data.model_dump(exclude_unset=True)
        if 'email' in update_dict and update_dict['email']:
            update_dict['email'] = update_dict['email'].lower()
        
        return self.db.update_user(user_id, update_dict)
    
    def delete_user(self, user_id: int) -> bool:
        """Remove um usuário"""
        
        if user_id <= 0:
            raise HTTPException(
                status_code=status.HTTP_400_BAD_REQUEST,
                detail="ID deve ser um número positivo"
            )
        
        return self.db.delete_user(user_id)
    
    def get_users_by_profile(self, profile: UserProfile) -> List[User]:
        """Busca usuários por perfil"""
        users = self.db.get_all_users()
        filtered_users = [u for u in users if u.perfil == profile]
        return sorted(filtered_users, key=lambda x: x.nome)
    
    def get_active_users(self) -> List[User]:
        """Retorna apenas usuários ativos"""
        users = self.db.get_all_users()
        active_users = [u for u in users if u.ativo]
        return sorted(active_users, key=lambda x: x.nome)
    
    def get_users_by_age_range(self, min_age: int, max_age: int) -> List[User]:
        """Busca usuários por faixa etária"""
        if min_age < 0:
            raise HTTPException(
                status_code=status.HTTP_400_BAD_REQUEST,
                detail="Idade mínima não pode ser negativa"
            )
        
        if max_age < min_age:
            raise HTTPException(
                status_code=status.HTTP_400_BAD_REQUEST,
                detail="Idade máxima deve ser maior que a idade mínima"
            )
        
        users = self.db.get_all_users()
        filtered_users = [u for u in users if min_age <= u.idade <= max_age]
        return sorted(filtered_users, key=lambda x: x.idade)
    
    def get_users_paginated(
        self,
        page: int = 1,
        size: int = 10,
        nome: Optional[str] = None,
        email: Optional[str] = None,
        perfil: Optional[UserProfile] = None,
        ativo: Optional[bool] = None
    ) -> UserListResponse:
        """Retorna usuários com paginação e filtros"""
        
        if page < 1:
            raise HTTPException(
                status_code=status.HTTP_400_BAD_REQUEST,
                detail="Página deve ser maior que zero"
            )
        
        if size < 1 or size > 100:
            raise HTTPException(
                status_code=status.HTTP_400_BAD_REQUEST,
                detail="Tamanho da página deve estar entre 1 e 100"
            )
        
        # Obter todos os usuários
        all_users = self.db.get_all_users()
        
        # Aplicar filtros
        filtered_users = all_users
        
        if nome:
            filtered_users = [u for u in filtered_users if nome.lower() in u.nome.lower()]
        
        if email:
            filtered_users = [u for u in filtered_users if email.lower() in u.email.lower()]
        
        if perfil:
            filtered_users = [u for u in filtered_users if u.perfil == perfil]
        
        if ativo is not None:
            filtered_users = [u for u in filtered_users if u.ativo == ativo]
        
        # Ordenar
        filtered_users = sorted(filtered_users, key=lambda x: x.nome)
        
        # Total de registros
        total = len(filtered_users)
        
        # Paginação
        start_index = (page - 1) * size
        end_index = start_index + size
        paginated_users = filtered_users[start_index:end_index]
        
        # Calcular metadados da paginação
        total_pages = math.ceil(total / size) if total > 0 else 0
        has_previous = page > 1
        has_next = page < total_pages
        
        return UserListResponse(
            users=paginated_users,
            page=page,
            size=size,
            total=total,
            total_pages=total_pages,
            has_previous=has_previous,
            has_next=has_next
        )
    
    def get_statistics(self) -> dict:
        """Retorna estatísticas dos usuários"""
        
        all_users = self.db.get_all_users()
        total_users = len(all_users)
        
        if total_users == 0:
            return {
                "total_usuarios": 0,
                "usuarios_ativos": 0,
                "usuarios_inativos": 0,
                "idade_media": 0.0,
                "idade_minima": 0,
                "idade_maxima": 0,
                "perfis": [],
                "usuario_mais_velho": None,
                "usuario_mais_novo": None
            }
        
        # Usuários ativos/inativos
        active_users = [u for u in all_users if u.ativo]
        active_count = len(active_users)
        inactive_count = total_users - active_count
        
        # Estatísticas de idade
        ages = [u.idade for u in all_users]
        avg_age = sum(ages) / len(ages)
        min_age = min(ages)
        max_age = max(ages)
        
        # Estatísticas por perfil
        profiles_stats = []
        for profile in UserProfile:
            profile_users = [u for u in all_users if u.perfil == profile]
            count = len(profile_users)
            
            if count > 0:
                profile_avg_age = sum(u.idade for u in profile_users) / count
                profiles_stats.append({
                    "perfil": profile.value,
                    "quantidade": count,
                    "idade_media": round(profile_avg_age, 1)
                })
        
        # Usuários extremos
        oldest_user = max(all_users, key=lambda u: u.idade)
        youngest_user = min(all_users, key=lambda u: u.idade)
        
        return {
            "total_usuarios": total_users,
            "usuarios_ativos": active_count,
            "usuarios_inativos": inactive_count,
            "idade_media": round(avg_age, 1),
            "idade_minima": min_age,
            "idade_maxima": max_age,
            "perfis": profiles_stats,
            "usuario_mais_velho": {
                "id": oldest_user.id,
                "nome": oldest_user.nome,
                "email": oldest_user.email,
                "idade": oldest_user.idade
            },
            "usuario_mais_novo": {
                "id": youngest_user.id,
                "nome": youngest_user.nome,
                "email": youngest_user.email,
                "idade": youngest_user.idade
            }
        }