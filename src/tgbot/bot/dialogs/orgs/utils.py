from aiogram_dialog.widgets.kbd import Button
from aiogram_dialog import DialogManager
from bot.core.config import ADMINS


def is_admin(data: dict, widget: Button, manager: DialogManager) -> bool:
    try:
        return True if manager.event.from_user.username in ADMINS else False
    except Exception:
        if hasattr(manager.event, "update"):
            update = manager.event.update
            if hasattr(update, "message") and update.message:
                user = update.message.from_user
            elif hasattr(update, "callback_query") and update.callback_query:
                user = update.callback_query.from_user
            elif hasattr(update, "inline_query") and update.inline_query:
                user = update.inline_query.from_user
            else:
                return False

            return user.username in ADMINS
        return False
