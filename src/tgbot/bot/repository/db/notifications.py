# bot/repository/db/notifications.py
from sqlalchemy import select, or_
from bot.repository.db.db import async_session
from bot.repository.db.models import Notifications, Subscriptions

async def get_user_notifications(user_id: int):

    async with async_session() as session:
        # First, get all user's subscriptions
        subscriptions_result = await session.execute(
            select(Subscriptions).where(Subscriptions.telegram_id == user_id)
        )
        subscriptions = subscriptions_result.scalars().all()
        
        if not subscriptions:
            return []
        
        conditions = []
        org_tokens = set()
        
        for sub in subscriptions:
            conditions.append(
                (Notifications.target_id == sub.patient_id) & 
                (Notifications.org_token == sub.token)
            )
            org_tokens.add(sub.token)
        
        for token in org_tokens:
            conditions.append(
                (Notifications.target_id == 0) & 
                (Notifications.org_token == token)
            )
        
        if conditions:
            notifications_result = await session.execute(
                select(Notifications)
                .where(or_(*conditions))
                .order_by(Notifications.created_at.desc())
            )
            notifications = notifications_result.scalars().all()
            
            return [
                {
                    "id": notification.id,
                    "message": notification.message,
                    "target_id": notification.target_id,
                    "org_token": notification.org_token, 
                    "header": notification.header,
                    "type": notification.type,
                    "created_at": notification.created_at
                }
                for notification in notifications
            ]
        
        return []
    

async def get_notification_by_id(notification_id: int):

    async with async_session() as session:
        result = await session.execute(
            select(Notifications).where(Notifications.id == notification_id)
        )
        notification = result.scalar_one_or_none()
        
        if notification:
            return {
                "id": notification.id,
                "message": notification.message,
                "target_id": notification.target_id,
                "org_token": notification.org_token,
                "header": notification.header,
                "type": notification.type,
                "created_at": notification.created_at
            }
        return None