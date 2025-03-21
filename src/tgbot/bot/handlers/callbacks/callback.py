from aiogram.types import CallbackQuery
from aiogram import Router
from aiogram import F

router = Router()


@router.callback_query(F.data == "check_subscription")
async def check_subscription(callback: CallbackQuery) -> None:
    pass
