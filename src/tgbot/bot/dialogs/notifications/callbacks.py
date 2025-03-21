from aiogram.types import CallbackQuery
from aiogram_dialog.widgets.kbd import Button
from aiogram_dialog import DialogManager


async def on_chosen_notification(
    c: CallbackQuery, widget: Button, manager: DialogManager
) -> None:
    pass
