from aiogram.types import CallbackQuery
from aiogram_dialog.widgets.kbd import Button
from aiogram_dialog import DialogManager


async def on_chosen_notifications(
    c: CallbackQuery, widget: Button, manager: DialogManager
) -> None:
    pass


async def on_chosen_orgs(
    c: CallbackQuery, widget: Button, manager: DialogManager
) -> None:
    pass
