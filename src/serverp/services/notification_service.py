from typing import List, Optional
from server.schemas.schemas import NotificationCreate
from server.models.models import Notification
from server.repositories.notification_repository import NotificationRepository
from server.repositories.patient_repository import PatientRepository

class NotificationService:
    @staticmethod
    async def create_notification(
        notification_data: NotificationCreate,
        notification_repo: NotificationRepository,
        patient_repo: PatientRepository
    ) -> Optional[Notification]:
        db_patient = await patient_repo.get_patient_by_id(notification_data.patient_id)
        if not db_patient:
            return None
        notification = Notification(
            title=notification_data.title,
            message=notification_data.message,
            patient_id=notification_data.patient_id
        )
        return await notification_repo.create_notification(notification)

    @staticmethod
    async def get_notifications(
        skip: int,
        limit: int,
        notification_repo: NotificationRepository
    ) -> List[Notification]:
        return await notification_repo.get_notifications(skip=skip, limit=limit)

    @staticmethod
    async def get_notification(
        notification_id: int,
        notification_repo: NotificationRepository
    ) -> Optional[Notification]:
        return await notification_repo.get_notification_by_id(notification_id)
