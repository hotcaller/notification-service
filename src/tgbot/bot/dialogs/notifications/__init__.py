from aiogram_dialog import Dialog

from bot.dialogs.notifications import windows

def menu_dialogs() -> list:
    return [
        Dialog(
            windows.notifications_window(),
            windows.notification_details_window(),
        ),
    ]