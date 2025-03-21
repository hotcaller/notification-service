from bot.repository.db.notifications import get_user_notifications
from bot.repository.db.notifications import get_notification_by_id

async def get_notifications(dialog_manager, **kwargs):
    user_id = dialog_manager.event.from_user.id
    notifications = await get_user_notifications(user_id)
    
    type_emojis = {
        "news": "📰",     
        "important": "⚠️",  
        "warning": "🚨",    
        "reminder": "🔔",  
        "default": "📌"
    }
    
    notifications_text = ""
    if notifications:
        for i, notification in enumerate(notifications, 1):
            # Get emoji based on notification type or use default
            emoji = type_emojis.get(notification.get('type', '').lower(), type_emojis["default"])
            
            notifications_text += f"{emoji} *{notification['header']}*\n"
            notifications_text += f"{notification['message'][:100]}...\n\n"
    else:
        notifications_text = "У вас нет уведомлений."
    
    return {
        "notifications_text": notifications_text,
        "notifications": notifications,
        "has_items": len(notifications) > 0
    }

async def get_notification_details(dialog_manager, **kwargs):
    notification_id = dialog_manager.dialog_data.get("notification_id")
    notification = await get_notification_by_id(notification_id)
    
    if notification:
        type_emojis = {
            "news": "📰",     
            "important": "⚠️", 
            "warning": "🚨",   
            "reminder": "🔔",  
            "default": "📌"
        }
        notification["type_emoji"] = type_emojis.get(
            notification.get('type', '').lower(), 
            type_emojis["default"]
        )
    
    return {
        "notification": notification
    }