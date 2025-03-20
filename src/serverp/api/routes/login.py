from fastapi import APIRouter, Request, Depends, HTTPException
from server.schemas.schemas import TelegramLogin, PatientOut
from server.services.auth_service import AuthService
from server.repositories.patient_repository import PatientRepository
from server.db.session import get_db

router = APIRouter()

async def get_patient_repository(db_session=Depends(get_db)):
    return PatientRepository(db_session)

@router.post("/login/callback", response_model=PatientOut)
async def telegram_login_callback(
    data: TelegramLogin,
    patient_repo: PatientRepository = Depends(get_patient_repository)
):
    patient = await AuthService.authenticate_user(data, patient_repo)
    if not patient:
        raise HTTPException(status_code=400, detail="Invalid Telegram data")
    return PatientOut(
        id=patient.id,
        name=patient.name,
        token=patient.token
    )
