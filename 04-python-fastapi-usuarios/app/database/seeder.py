from sqlalchemy.orm import Session
from app.models.user import User, UserProfile
from datetime import datetime

def seed_users(db: Session):
    """Popula o banco com usuários de exemplo"""
    
    # Verifica se já existem usuários
    if db.query(User).count() > 0:
        return
    
    # Usuários de exemplo
    sample_users = [
        User(
            nome="João Silva",
            email="joao.silva@email.com", 
            idade=28,
            ativo=True,
            perfil=UserProfile.ADMIN
        ),
        User(
            nome="Maria Santos",
            email="maria.santos@email.com",
            idade=32,
            ativo=True,
            perfil=UserProfile.USER
        ),
        User(
            nome="Pedro Oliveira", 
            email="pedro.oliveira@email.com",
            idade=25,
            ativo=True,
            perfil=UserProfile.MODERATOR
        ),
        User(
            nome="Ana Costa",
            email="ana.costa@email.com",
            idade=29,
            ativo=False,
            perfil=UserProfile.USER
        ),
        User(
            nome="Carlos Lima",
            email="carlos.lima@email.com",
            idade=35,
            ativo=True,
            perfil=UserProfile.USER
        )
    ]
    
    # Adiciona todos os usuários
    for user in sample_users:
        db.add(user)
    
    # Confirma as alterações
    db.commit()
    
    print(f"✅ {len(sample_users)} usuários de exemplo foram criados!")
    
    # Refresh para obter os IDs gerados
    for user in sample_users:
        db.refresh(user)