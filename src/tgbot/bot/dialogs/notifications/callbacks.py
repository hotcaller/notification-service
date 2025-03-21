from aiogram.types import CallbackQuery
from aiogram_dialog import DialogManager
from bot.dialogs.notifications.states import NotificationMenu

async def on_chosen_notification(callback: CallbackQuery, widget, dialog_manager: DialogManager, item_id: str):
    dialog_manager.dialog_data["notification_id"] = int(item_id)
    await dialog_manager.switch_to(NotificationMenu.notification_details)