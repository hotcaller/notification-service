from aiogram_dialog import Dialog

from bot.dialogs.start import windows

def menu_dialogs() -> list:
    return [
        Dialog(
            windows.start_window(),
        ),
    ]
