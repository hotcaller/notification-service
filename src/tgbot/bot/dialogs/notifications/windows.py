from aiogram_dialog import Window
from aiogram_dialog.widgets.text import Const, Format
from aiogram_dialog.widgets.kbd import Back, Button
from bot.dialogs.notifications import states, getters, callbacks
from bot.dialogs.notifications import keyboards, utils

def notifications_window() -> Window:
    return Window(
        Format(
            "📋 *Ваши уведомления*\n\n"
            "{notifications_text}"
        ),
        keyboards.paginated_notifications(
            on_click=callbacks.on_chosen_notification, when=utils.check_has_items
        ),
        Const("У вас нет уведомлений.", when=lambda data: not data.get("has_items", False)),
        Back(Const("⬅️ Назад")),
        state=states.NotificationMenu.notifications_menu,
        getter=getters.get_notifications,
    )

def notification_details_window() -> Window:
    return Window(
        Format(
            "📝 *{notification.header}*\n\n"
            "{notification.message}\n\n"
            "📅 Дата: {notification.created_at}"
        ),
        Back(Const("⬅️ Назад к списку")),
        state=states.NotificationMenu.notification_details,
        getter=getters.get_notification_details,
    )