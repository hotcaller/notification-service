from sqlalchemy.ext.asyncio import create_async_engine, AsyncSession
from sqlalchemy.orm import sessionmaker
from server.core.config import settings

engine = create_async_engine(settings.DATABASE_URL, future=True, echo=True)

AsyncSessionLocal = sessionmaker(
    bind=engine, class_=AsyncSession, expire_on_commit=False
)

# Dependency
async def get_db():
    async with AsyncSessionLocal() as session:
        yield session
