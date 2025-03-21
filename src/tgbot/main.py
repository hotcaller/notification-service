import asyncio  # noqa: D100
import logging
import sys

from aiogram import Bot, Dispatcher
from aiogram.client.default import DefaultBotProperties
from aiogram.fsm.storage.redis import RedisStorage
from aiogram_dialog import setup_dialogs

from bot.dialogs import get_dialogs
from bot.handlers import router as handlers_router
from bot.core.config import TELEGRAM_TOKEN, REDIS_URL
from bot.repository.redis.redis import init_redis, close_redis
from bot.repository.db.db import init_db
from bot.core.kafka.consumer import consume
from bot.handlers.callbacks.callback import start_router


async def main() -> None:
    await init_redis()
    await init_db()

    storage = RedisStorage.from_url(REDIS_URL)
    storage.key_builder.with_destiny = True

    defaults = DefaultBotProperties(parse_mode="Markdown")

    bot = Bot(token=TELEGRAM_TOKEN, default=defaults)
    dp = Dispatcher(storage=storage)

    asyncio.create_task(consume(bot))

    dp.include_routers(handlers_router, *get_dialogs())
    dp.include_router(start_router)
    setup_dialogs(dp)
    await dp.start_polling(bot)


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO, stream=sys.stdout)
    # logging.basicConfig(level=logging.INFO, stream=sys.stdout,   format='%(asctime)s - %(levelname)s:%(name)s:%(funcName)s:%(lineno)d:%(message)s',datefmt='%Y-%m-%d %H:%M:%S',)
    try:
        asyncio.run(main())
    except KeyboardInterrupt:
        close_redis()
        print("Stopped")
