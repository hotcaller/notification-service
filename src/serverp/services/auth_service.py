from typing import Optional
from server.schemas.schemas import TelegramLogin
from server.models.models import Patient
from server.repositories.patient_repository import PatientRepository
from hashlib import sha256
import hmac
import os

class AuthService:
    @staticmethod
    def verify_telegram_login(data: TelegramLogin) -> bool:
        # Верификация данных, полученных от Telegram
        bot_token = os.getenv("TELEGRAM_BOT_TOKEN")
        check_hash = data.hash
        data_dict = data.dict()
        data_dict.pop('hash')
        data_check_arr = [f"{k}={v}" for k, v in sorted(data_dict.items())]
        data_check_string = "\n".join(data_check_arr)
        secret_key = sha256(bot_token.encode()).digest()
        hmac_string = hmac.new(secret_key, data_check_string.encode(), sha256).hexdigest()
        return hmac_string == check_hash

    @staticmethod
    async def authenticate_user(
        telegram_data: TelegramLogin,
        patient_repo: PatientRepository
    ) -> Optional[Patient]:
        if not AuthService.verify_telegram_login(telegram_data):
            return None
        patient = await patient_repo.get_patient_by_id(telegram_data.id)
        if not patient:
            # Создаем нового пациента
            new_patient = Patient(
                id=telegram_data.id,
                name=telegram_data.first_name,
                token=os.urandom(24).hex()
            )
            patient = await patient_repo.create_patient(new_patient)
        return patient
