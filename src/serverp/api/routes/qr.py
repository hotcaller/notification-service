from fastapi import APIRouter, Depends, HTTPException
from server.db.session import get_db
from server.repositories.patient_repository import PatientRepository

router = APIRouter()

async def get_patient_repository(db_session=Depends(get_db)):
    return PatientRepository(db_session)

@router.get("/qr")
async def get_qr(
    patient_id: int,
    token: str,
    patient_repo: PatientRepository = Depends(get_patient_repository)
):
    patient = await patient_repo.get_patient_by_id_and_token(patient_id, token)
    if not patient:
        raise HTTPException(status_code=404, detail="Patient not found")
    # Здесь можно добавить генерацию QR-кода
    # Например, использовать библиотеку qrcode
    # import qrcode
    # img = qrcode.make(f"Patient ID: {patient_id}, Token: {token}")
    # Возвращаем изображение или данные
    return {"patient_id": patient_id, "token": token}
