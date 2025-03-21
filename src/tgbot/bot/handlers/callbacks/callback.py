from aiogram.types import CallbackQuery
from aiogram import Router
from aiogram import F

r = Router()

@r.callback_query(F.data == "check_subscription")
async def check_subscription(callback: CallbackQuery) -> None:
    pass
