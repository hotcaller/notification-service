from . import start, notifications


def get_dialogs() -> list:
    return [
        *start.menu_dialogs(),
        *notifications.menu_dialogs(),
    ]
