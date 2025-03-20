import asyncio  # noqa: D100
import logging
import sys
from aiogram import Bot, Dispatcher
from aiogram.client.default import DefaultBotProperties
from aiogram.fsm.storage.redis import RedisStorage
from aiogram_dialog import setup_dialogs
from bot.dialogs import get_dialogs
from bot.handlers import router
from bot.core.config import TELEGRAM_TOKEN, REDIS_URL
from concurrent.futures import ThreadPoolExecutor
from bot.repository.redis.redis import init_redis, close_redis
from bot.repository.db.db import init_db

executor = ThreadPoolExecutor(max_workers=1)

async def main() -> None:
    await init_redis()
    await init_db()

    storage = RedisStorage.from_url(REDIS_URL)
    storage.key_builder.with_destiny = True

    defaults = DefaultBotProperties(parse_mode="Markdown")

    bot = Bot(token=TELEGRAM_TOKEN, default=defaults)
    dp = Dispatcher(storage=storage)

    dp.include_routers(router, *get_dialogs())
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

