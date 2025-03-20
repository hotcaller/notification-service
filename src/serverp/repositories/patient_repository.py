from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.future import select
from server.models.models import Patient

class PatientRepository:
    def __init__(self, db_session: AsyncSession):
        self.db_session = db_session

    async def get_patient_by_id_and_token(self, patient_id: int, token: str) -> Patient:
        result = await self.db_session.execute(
            select(Patient).where(Patient.id == patient_id, Patient.token == token)
        )
        return result.scalar_one_or_none()

    async def get_patient_by_id(self, patient_id: int) -> Patient:
        result = await self.db_session.execute(
            select(Patient).where(Patient.id == patient_id)
        )
        return result.scalar_one_or_none()

    async def create_patient(self, patient: Patient) -> Patient:
        self.db_session.add(patient)
        await self.db_session.commit()
        await self.db_session.refresh(patient)
        return patient
