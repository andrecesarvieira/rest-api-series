from pydantic import BaseModel, EmailStr, Field, field_validator
from typing import Optional
from datetime import datetime
from enum import Enum

class UserProfile(str, Enum):
    ADMIN = "admin"
    USER = "user"
    MODERATOR = "moderator"

class UserBase(BaseModel):
    nome: str = Field(..., min_length=2, max_length=100, description="Nome do usuário")
    email: EmailStr = Field(..., description="Email único do usuário")
    idade: int = Field(..., ge=0, le=120, description="Idade do usuário (0-120 anos)")
    perfil: UserProfile = Field(default=UserProfile.USER, description="Perfil do usuário")
    ativo: bool = Field(default=True, description="Status ativo do usuário")

class UserCreate(UserBase):
    """Schema para criação de usuário"""
    
    @field_validator('nome')
    def validate_nome(cls, v):
        if not v.strip():
            raise ValueError('Nome não pode estar vazio')
        if len(v.strip()) < 2:
            raise ValueError('Nome deve ter pelo menos 2 caracteres')
        return v.strip()
    
    @field_validator('email')
    def validate_email(cls, v):
        # EmailStr já valida o formato, aqui fazemos validações adicionais
        if not v:
            raise ValueError('Email é obrigatório')
        return v.lower()
    
    @field_validator('idade')
    def validate_idade(cls, v):
        if v < 0:
            raise ValueError('Idade não pode ser negativa')
        if v > 120:
            raise ValueError('Idade não pode ser maior que 120 anos')
        return v

class UserUpdate(BaseModel):
    """Schema para atualização de usuário"""
    nome: Optional[str] = Field(None, min_length=2, max_length=100)
    email: Optional[EmailStr] = None
    idade: Optional[int] = Field(None, ge=0, le=120)
    perfil: Optional[UserProfile] = None
    ativo: Optional[bool] = None
    
    @field_validator('nome')
    def validate_nome(cls, v):
        if v is not None:
            if not v.strip():
                raise ValueError('Nome não pode estar vazio')
            if len(v.strip()) < 2:
                raise ValueError('Nome deve ter pelo menos 2 caracteres')
            return v.strip()
        return v
    
    @field_validator('email')
    def validate_email(cls, v):
        if v is not None:
            return v.lower()
        return v

class UserResponse(UserBase):
    """Schema para resposta de usuário"""
    id: int
    data_criacao: datetime
    
    class Config:
        from_attributes = True
        json_encoders = {
            datetime: lambda v: v.isoformat()
        }

class UserListResponse(BaseModel):
    """Schema para lista paginada de usuários"""
    users: list[UserResponse]
    page: int = Field(..., ge=1, description="Número da página atual")
    size: int = Field(..., ge=1, le=100, description="Tamanho da página")
    total: int = Field(..., ge=0, description="Total de usuários")
    total_pages: int = Field(..., ge=0, description="Total de páginas")
    has_previous: bool = Field(..., description="Tem página anterior")
    has_next: bool = Field(..., description="Tem próxima página")