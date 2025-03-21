from aiogram_dialog import Window
from aiogram_dialog.widgets.text import Const, Format
from aiogram_dialog.widgets.kbd import Back
from bot.dialogs.notifications import states, getters, callbacks
from bot.dialogs.notifications import keyboards, utils


def notifications_window() -> Window:
    return Window(
        Format(
            "*Ваши брони в OneDayDom*👋\n\nНажмите на бронь, чтобы посмотреть подробности."
        ),
        keyboards.paginated_bookings(
            on_click=callbacks.on_chosen_notification, when=utils.check_has_items
        ),
        Back(Const("⬅️ Назад")),
        state=states.NotificationMenu.notifications_menu,
        getter=getters.get_bookings,
    )
