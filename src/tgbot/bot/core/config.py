import os
from dotenv import load_dotenv

load_dotenv()

TELEGRAM_TOKEN = os.getenv("TELEGRAM_TOKEN")
TELEGRAM_USERNAME = os.getenv("TELEGRAM_USERNAME")
TELEGRAM_CHANNEL = os.getenv("TELEGRAM_CHANNEL")

PG_STRING = os.getenv("PG_STRING")

redis_address = os.getenv("REDIS_ADDRESS")
REDIS_PORT = redis_address.split(":")[1]
REDIS_HOST = redis_address.split(":")[0]
REDIS_PASSWORD = os.getenv("REDIS_PASSWORD")
REDIS_URL = f"redis://:{REDIS_PASSWORD}@{REDIS_HOST}:{REDIS_PORT}/0"

ADMINS = os.getenv("BOT_ADMINS").split(",")

