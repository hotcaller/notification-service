from aiogram.types import WebAppInfo
from aiogram.utils.keyboard import ReplyKeyboardBuilder


def web_app_kb(web_app_url: str):
    kb_builder = ReplyKeyboardBuilder()
    kb_builder.button(text="Забронировать", web_app=WebAppInfo(url=web_app_url))
    keyboard = kb_builder.as_markup(resize_keyboard=True)
    return keyboard
