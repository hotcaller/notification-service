from aiogram.types import CallbackQuery
from bot.core.config import TELEGRAM_CHANNEL
from aiogram import Router, types
from aiogram import F
r = Router()

@r.callback_query(F.data == "check_subscription")
async def check_subscription(callback: CallbackQuery):
    user_id = callback.from_user.id

    chat_member = await callback.bot.get_chat_member(chat_id=TELEGRAM_CHANNEL, user_id=user_id)

    # if chat_member.status in ["member", "administrator", "creator"]:
    #     await callback.message.answer("Спасибо за подписку! ❤️\nИспользуйте промокод `telegram` и получите скидку 5% на бронирование.")
    # else:
    #     await callback.message.answer("Вы ещё не подписались на канал. Пожалуйста, подпишитесь и нажмите кнопку ещё раз.")


import aiohttp


QR_CODE_URL = "http://103.88.241.21/qr"  

@r.callback_query(lambda c: c.data == "get_qr")
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
