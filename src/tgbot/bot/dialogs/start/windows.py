from aiogram_dialog import Window
from aiogram_dialog.widgets.text import Const, Format
from aiogram_dialog.widgets.kbd import Button
from bot.dialogs.start import states, callbacks


def start_window() -> Window:
    return Window(
        Format("Добро пожаловать в ZabMedBot!"),

        state=states.StartMenu.select_menu,
    )
