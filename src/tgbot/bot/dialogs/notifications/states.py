from aiogram.fsm.state import State, StatesGroup

class NotificationMenu(StatesGroup):
    notifications_menu = State()
    notification_details = State()