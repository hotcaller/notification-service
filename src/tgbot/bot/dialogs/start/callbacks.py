from aiogram.types import CallbackQuery
from aiogram_dialog.widgets.kbd import Button
from aiogram_dialog import DialogManager
from bot.dialogs.notifications.states import NotificationMenu


async def on_chosen_notifications(
    c: CallbackQuery, widget: Button, manager: DialogManager
) -> None:
    await manager.start(NotificationMenu.notifications_menu)


async def on_chosen_orgs(
    c: CallbackQuery, widget: Button, manager: DialogManager
) -> None:
    pass
