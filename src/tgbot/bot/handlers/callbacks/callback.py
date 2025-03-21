from aiogram.types import CallbackQuery
from bot.core.config import TELEGRAM_CHANNEL
from aiogram import Router, types
from aiogram import F
from ...repository.db.subscription import get_subscriptions_by_user_id
import aiohttp

r = Router()

@r.callback_query(F.data == "check_subscription")
async def check_subscription(callback: CallbackQuery):
    pass


QR_CODE_URL = "http://103.88.241.21/qr"

@r.callback_query(F.data == "get_qr")
async def get_qr(callback: types.CallbackQuery):
    user_id = callback.from_user.id

    # Получаем подписки пользователя из БД
    subscriptions = await get_subscriptions_by_user_id(user_id)

    if not subscriptions:
        await callback.message.answer("У вас нет активных подписок.")
        return

    async with aiohttp.ClientSession() as session:
        for sub in subscriptions:
            params = {
                "patient_id": sub["patient_id"],
                "token": sub["token"],
            }
            async with session.get(QR_CODE_URL, params=params) as response:
                if response.status == 200:
                    qr_code_data = await response.read()
                    photo = types.BufferedInputFile(qr_code_data, filename="qrcode.png")
                    await callback.message.answer_photo(photo, caption="Ваш QR-код")
                else:
                    await callback.message.answer("Ошибка при получении QR-кода.")

    await callback.answer()
