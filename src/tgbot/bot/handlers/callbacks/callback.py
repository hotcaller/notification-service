from aiogram.types import CallbackQuery
from aiogram.types import Message
from aiogram import Router
from aiogram import F
from bot.repository.db.subscriptions import create_subscription, subscription_exists
from bot.repository.db.users import user_exists_by_telegram_id, create_user

router = Router()
QR_CODE_URL = "http://103.88.241.21:8000/"


@router.callback_query(F.data == "check_subscription")
async def check_subscription(callback: CallbackQuery) -> None:
    pass


@router.message(F.data == "/start")
async def start_handler(message: Message):
    args = message.text.split(" ", 1)
    telegram_id = message.from_user.id
    username = message.from_user.username or "Unknown"

    if len(args) > 1:
        try:
            patient_id, token = args[1].split("|")
            patient_id = int(patient_id)
        except ValueError:
            await message.answer("Некорректный формат ссылки. Попробуйте еще раз.")
            return

        if not await user_exists_by_telegram_id(telegram_id):
            await create_user(telegram_id, username)

        if not await subscription_exists(telegram_id, token, patient_id):
            await create_subscription(telegram_id, token, patient_id)
            await message.answer("✅ Вы успешно подписались на уведомления!")
        else:
            await message.answer("ℹ️ Вы уже подписаны на эти уведомления.")

    else:
        await message.answer("Привет! Используйте специальную ссылку для подписки.")
