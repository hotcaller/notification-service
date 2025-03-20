from fastapi import APIRouter, Depends, HTTPException
from typing import List
from server.schemas.schemas import NotificationCreate, NotificationOut
from server.services.notification_service import NotificationService
from server.db.session import get_db
from server.repositories.notification_repository import NotificationRepository
from server.repositories.patient_repository import PatientRepository

router = APIRouter()

# Зависимости для репозиториев
async def get_notification_repository(db_session=Depends(get_db)):
    return NotificationRepository(db_session)

async def get_patient_repository(db_session=Depends(get_db)):
    return PatientRepository(db_session)

@router.get("/notifications/", response_model=List[NotificationOut])
async def get_notifications(
    skip: int = 0,
    limit: int = 10,
    notification_repo: NotificationRepository = Depends(get_notification_repository),
):
    notifications = await NotificationService.get_notifications(
        skip, limit, notification_repo
    )
    return notifications

@router.get("/notifications/{id}", response_model=NotificationOut)
async def get_notification(
    id: int,
    notification_repo: NotificationRepository = Depends(get_notification_repository),
):
    notification = await NotificationService.get_notification(id, notification_repo)
    if not notification:
        raise HTTPException(status_code=404, detail="Notification not found")
    return notification

@router.post("/notifications/", response_model=NotificationOut)
async def create_notification(
    notification: NotificationCreate,
    notification_repo: NotificationRepository = Depends(get_notification_repository),
    patient_repo: PatientRepository = Depends(get_patient_repository),
):
    db_notification = await NotificationService.create_notification(
        notification, notification_repo, patient_repo
    )
    if not db_notification:
        raise HTTPException(status_code=404, detail="Patient not found")
    return db_notification
