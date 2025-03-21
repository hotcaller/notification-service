from sqlalchemy import select, delete
from bot.repository.db.db import async_session
from bot.repository.db.models import Subscriptions


async def subscription_exists(user_id: int, token: str, patient_id: int) -> bool:
    async with async_session() as session:
        result = await session.execute(
            select(Subscriptions).where(
                (Subscriptions.user_id == user_id)
                & (Subscriptions.token == token)
                & (Subscriptions.patient_id == patient_id)
            )
        )
        subscription = result.scalar_one_or_none()
        return subscription is not None


async def create_subscription(user_id: int, token: str, patient_id: int):
    async with async_session() as session:
        new_subscription = Subscriptions(
            telegram_id=user_id, token=token, patient_id=patient_id
        )
        session.add(new_subscription)
        await session.commit()


async def get_subscriptions_by_user_id(user_id: int):
    async with async_session() as session:
        result = await session.execute(
            select(Subscriptions).where(Subscriptions.user_id == user_id)
        )
        subscriptions = result.scalars().all()
        return [
            {
                "id": sub.id,
                "user_id": sub.user_id,
                "token": sub.token,
                "patient_id": sub.patient_id,
            }
            for sub in subscriptions
        ]


async def get_subscription_by_patient_id_and_token(patient_id: int, token: str):
    async with async_session() as session:
        result = await session.execute(
            select(Subscriptions).where(
                (Subscriptions.patient_id == patient_id)
                & (Subscriptions.token == token)
            )
        )
        subscription = result.scalar_one_or_none()
        if subscription:
            return {
                "id": subscription.id,
                "user_id": subscription.user_id,
                "token": subscription.token,
                "patient_id": subscription.patient_id,
            }
        return None


async def get_subscription_by_id(subscription_id: int):
    async with async_session() as session:
        result = await session.execute(
            select(Subscriptions).where(Subscriptions.id == subscription_id)
        )
        subscription = result.scalar_one_or_none()
        if subscription:
            return {
                "id": subscription.id,
                "user_id": subscription.user_id,
                "token": subscription.token,
                "patient_id": subscription.patient_id,
            }
        return None


async def delete_subscription(subscription_id: int):
    async with async_session() as session:
        await session.execute(
            delete(Subscriptions).where(Subscriptions.id == subscription_id)
        )
        await session.commit()
