from bot.repository.db.notifications import get_user_notifications
from bot.repository.db.notifications import get_notification_by_id

async def get_notifications(dialog_manager, **kwargs):
    user_id = dialog_manager.event.from_user.id
    notifications = await get_user_notifications(user_id)
    
    return {
        "notifications": notifications,
        "has_items": len(notifications) > 0
    }

async def get_notification_details(dialog_manager, **kwargs):
    notification_id = dialog_manager.dialog_data.get("notification_id")
    notification = await get_notification_by_id(notification_id)
    
    return {
        "notification": notification
    }