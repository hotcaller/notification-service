from aiogram_dialog import Window
from aiogram_dialog.widgets.text import Const, Format
from aiogram_dialog.widgets.kbd import Button
from bot.dialogs.start import states, callbacks


def start_window() -> Window:
    return Window(
        Format("Добро пожаловать в ZabMedBot!"),
        Button(
            Const("📑 Список уведомлений"),
            "notifications_button",
            callbacks.on_chosen_notifications,
        ),
        Button(Const("📆 Организации"), "orgs_button", callbacks.on_chosen_orgs),
        state=states.StartMenu.select_menu,
    )
