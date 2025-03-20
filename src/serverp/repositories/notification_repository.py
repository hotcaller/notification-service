from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.future import select
from server.models.models import Notification

class NotificationRepository:
    def __init__(self, db_session: AsyncSession):
        self.db_session = db_session

    async def create_notification(self, notification: Notification) -> Notification:
        self.db_session.add(notification)
        await self.db_session.commit()
        await self.db_session.refresh(notification)
        return notification

    async def get_notifications(self, skip: int = 0, limit: int = 10):
        result = await self.db_session.execute(
            select(Notification).offset(skip).limit(limit)
        )
        return result.scalars().all()

    async def get_notification_by_id(self, notification_id: int) -> Notification:
        result = await self.db_session.execute(
            select(Notification).where(Notification.id == notification_id)
        )
        return result.scalar_one_or_none()
