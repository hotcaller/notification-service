from fastapi import FastAPI
from server.api.routes import api_router
from server.core.config import settings
from server.db.session import engine
from server.models import models

app = FastAPI(title=settings.PROJECT_NAME, version=settings.PROJECT_VERSION)

app.include_router(api_router)

# Создание таблиц базы данных
async def init_models():
    async with engine.begin() as conn:
        await conn.run_sync(models.Base.metadata.create_all)

# Запуск создания моделей при старте
@app.on_event("startup")
async def on_startup():
    await init_models()
