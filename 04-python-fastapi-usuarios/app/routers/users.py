from fastapi import APIRouter, Depends, HTTPException, status, Query
from typing import Optional, List
from app.database.database import get_database, InMemoryDatabase
from app.schemas.user import UserCreate, UserUpdate, UserResponse, UserListResponse
from app.services.user_service import UserService
from app.models.user import UserProfile

router = APIRouter(prefix="/api/usuarios", tags=["Usuários"])

def get_user_service(db: InMemoryDatabase = Depends(get_database)) -> UserService:
    """Dependency para obter o serviço de usuários"""
    return UserService(db)

@router.get("/", response_model=List[UserResponse], summary="Listar todos os usuários")
async def get_all_users(service: UserService = Depends(get_user_service)):
    """
    Retorna todos os usuários cadastrados.
    """
    users = service.get_all_users()
    return users

@router.get("/{user_id}", response_model=UserResponse, summary="Obter usuário por ID")
async def get_user_by_id(
    user_id: int,
    service: UserService = Depends(get_user_service)
):
    """
    Retorna um usuário específico pelo ID.
    
    - **user_id**: ID único do usuário
    """
    user = service.get_user_by_id(user_id)
    if not user:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail=f"Usuário com ID {user_id} não encontrado"
        )
    return user

@router.post("/", response_model=UserResponse, status_code=status.HTTP_201_CREATED, summary="Criar novo usuário")
async def create_user(
    user_data: UserCreate,
    service: UserService = Depends(get_user_service)
):
    """
    Cria um novo usuário.
    
    - **nome**: Nome completo do usuário (2-100 caracteres)
    - **email**: Email único e válido
    - **idade**: Idade entre 0 e 120 anos
    - **perfil**: Perfil do usuário (user, admin, moderator)
    - **ativo**: Status ativo do usuário (padrão: true)
    """
    return service.create_user(user_data)

@router.put("/{user_id}", response_model=UserResponse, summary="Atualizar usuário")
async def update_user(
    user_id: int,
    user_data: UserUpdate,
    service: UserService = Depends(get_user_service)
):
    """
    Atualiza um usuário existente.
    
    - **user_id**: ID do usuário a ser atualizado
    - Apenas campos fornecidos serão atualizados
    """
    user = service.update_user(user_id, user_data)
    if not user:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail=f"Usuário com ID {user_id} não encontrado"
        )
    return user

@router.delete("/{user_id}", status_code=status.HTTP_204_NO_CONTENT, summary="Remover usuário")
async def delete_user(
    user_id: int,
    service: UserService = Depends(get_user_service)
):
    """
    Remove um usuário do sistema.
    
    - **user_id**: ID do usuário a ser removido
    """
    success = service.delete_user(user_id)
    if not success:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail=f"Usuário com ID {user_id} não encontrado"
        )

@router.get("/perfil/{profile}", response_model=List[UserResponse], summary="Listar usuários por perfil")
async def get_users_by_profile(
    profile: UserProfile,
    service: UserService = Depends(get_user_service)
):
    """
    Retorna usuários filtrados por perfil.
    
    - **profile**: Perfil do usuário (user, admin, moderator)
    """
    users = service.get_users_by_profile(profile)
    return users

@router.get("/status/ativos", response_model=List[UserResponse], summary="Listar usuários ativos")
async def get_active_users(service: UserService = Depends(get_user_service)):
    """
    Retorna apenas usuários com status ativo.
    """
    users = service.get_active_users()
    return users

@router.get("/idade/faixa", response_model=List[UserResponse], summary="Filtrar usuários por idade")
async def get_users_by_age_range(
    min_idade: int = Query(..., ge=0, description="Idade mínima"),
    max_idade: int = Query(..., le=120, description="Idade máxima"),
    service: UserService = Depends(get_user_service)
):
    """
    Retorna usuários dentro de uma faixa etária.
    
    - **min_idade**: Idade mínima (>= 0)
    - **max_idade**: Idade máxima (<= 120)
    """
    users = service.get_users_by_age_range(min_idade, max_idade)
    return users

@router.get("/paginado/lista", response_model=UserListResponse, summary="Listar usuários com paginação")
async def get_users_paginated(
    page: int = Query(1, ge=1, description="Número da página"),
    size: int = Query(10, ge=1, le=100, description="Itens por página"),
    nome: Optional[str] = Query(None, description="Filtro por nome (busca parcial)"),
    email: Optional[str] = Query(None, description="Filtro por email (busca parcial)"),
    perfil: Optional[UserProfile] = Query(None, description="Filtro por perfil"),
    ativo: Optional[bool] = Query(None, description="Filtro por status ativo"),
    service: UserService = Depends(get_user_service)
):
    """
    Retorna usuários com paginação e filtros opcionais.
    
    - **page**: Número da página (padrão: 1)
    - **size**: Itens por página (padrão: 10, máx: 100)
    - **nome**: Filtro por nome (busca parcial, opcional)
    - **email**: Filtro por email (busca parcial, opcional)
    - **perfil**: Filtro por perfil (opcional)
    - **ativo**: Filtro por status ativo (opcional)
    """
    return service.get_users_paginated(
        page=page,
        size=size,
        nome=nome,
        email=email,
        perfil=perfil,
        ativo=ativo
    )

@router.get("/estatisticas/dados", summary="Obter estatísticas dos usuários")
async def get_statistics(service: UserService = Depends(get_user_service)):
    """
    Retorna estatísticas consolidadas dos usuários.
    
    Inclui:
    - Total de usuários
    - Usuários ativos/inativos
    - Estatísticas de idade (média, mín, máx)
    - Distribuição por perfis
    - Usuário mais velho/novo
    """
    return service.get_statistics()