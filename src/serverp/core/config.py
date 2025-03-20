import os

class Settings:
    PROJECT_NAME: str = "FastAPI Server"
    PROJECT_VERSION: str = "1.0.0"

    DATABASE_URL: str = os.getenv(
        "PG_STRING",
        "postgresql+asyncpg://postgres:postgres@db:5432/postgres"
    )


settings = Settings()
