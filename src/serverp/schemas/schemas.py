from pydantic import BaseModel
from datetime import datetime

class NotificationBase(BaseModel):
    title: str
    message: str

class NotificationCreate(NotificationBase):
    patient_id: int

class NotificationOut(NotificationBase):
    id: int
    timestamp: datetime
    patient_id: int

    class Config:
        orm_mode = True

class PatientBase(BaseModel):
    name: str

class PatientCreate(PatientBase):
    token: str

class PatientOut(PatientBase):
    id: int
    token: str

    class Config:
        orm_mode = True

class TelegramLogin(BaseModel):
    id: int
    first_name: str
    last_name: str = None
    username: str = None
    photo_url: str = None
    auth_date: int
    hash: str
