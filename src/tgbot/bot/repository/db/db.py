from sqlalchemy.ext.asyncio import create_async_engine, async_sessionmaker
from sqlalchemy.orm import declarative_base
from bot.core.config import PG_STRING

PG_STRING = PG_STRING.replace("?sslmode=disable", "")
PG_STRING = PG_STRING.replace("&sslmode=disable", "")

if PG_STRING.startswith("postgresql://"):
    PG_STRING = PG_STRING.replace("postgresql://", "postgresql+asyncpg://", 1)

print(PG_STRING)
engine = create_async_engine(
    PG_STRING,
    echo=True,
)

async_session = async_sessionmaker(
    bind=engine,
    expire_on_commit=False,
)

Base = declarative_base()


async def init_db():
    async with engine.begin() as conn:
        await conn.run_sync(Base.metadata.create_all)
