from aiogram.types import CallbackQuery
from bot.core.config import TELEGRAM_CHANNEL
from aiogram import Router
from aiogram import F
r = Router()

@r.callback_query(F.data == "check_subscription")
async def check_subscription(callback: CallbackQuery):
    user_id = callback.from_user.id

    chat_member = await callback.bot.get_chat_member(chat_id=TELEGRAM_CHANNEL, user_id=user_id)

    if chat_member.status in ["member", "administrator", "creator"]:
        await callback.message.answer("Спасибо за подписку! ❤️\nИспользуйте промокод `telegram` и получите скидку 5% на бронирование.")
    else:
        await callback.message.answer("Вы ещё не подписались на канал. Пожалуйста, подпишитесь и нажмите кнопку ещё раз.")