import redis.asyncio as redis
import logging
from bot.core.config import REDIS_HOST, REDIS_PORT, REDIS_PASSWORD

RDB = None

# Настройка логирования
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


async def init_redis():
    """Инициализирует подключение к Redis."""
    global RDB
    RDB = redis.Redis(host=REDIS_HOST, port=REDIS_PORT, password=REDIS_PASSWORD, db=0)
    try:
        await RDB.ping()
        logger.info("Подключение к Redis установлено.")
    except Exception as e:
        logger.error(f"Не удалось подключиться к Redis: {e}")
        raise e


async def get(key: str) -> str:
    """Получает значение по ключу из Redis.

    :param key: Ключ для поиска.
    :return: Значение, связанное с ключом, или None, если ключ не найден.
    """
    try:
        value = await RDB.get(key)
        if value:
            return value.decode("utf-8")
        return None
    except Exception as e:
        logger.error(f"Ошибка при получении из Redis: {e}")
        return None


async def set(key: str, value: str):
    """Устанавливает значение по ключу в Redis.

    :param key: Ключ для сохранения.
    :param value: Значение для сохранения.
    """
    try:
        await RDB.set(key, value)
        logger.info(f"Значение для ключа '{key}' успешно сохранено в Redis.")
    except Exception as e:
        logger.error(f"Ошибка при сохранении в Redis: {e}")


async def close_redis():
    """Закрывает подключение к Redis."""
    if RDB:
        await RDB.close()  # Закрываем подключение
        logger.info("Подключение к Redis закрыто.")
