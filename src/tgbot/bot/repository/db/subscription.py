from sqlalchemy import select, delete, distinct
from bot.repository.db.db import async_session
from bot.repository.db.models import Subscriptions



async def get_all_subscribers_by_token(token: str):
    async with async_session() as session:
        result = await session.execute(
            select(distinct(Subscriptions.telegram_id)).where(
                Subscriptions.token == token
            )
        )
        return result.scalars().all()
    
async def subscription_exists(user_id: int, token: str, patient_id: int) -> bool:
    async with async_session() as session:
        result = await session.execute(
            select(Subscriptions).where(
                (Subscriptions.telegram_id == user_id)
                & (Subscriptions.token == token)
                & (Subscriptions.patient_id == patient_id)
            )
        )
        subscription = result.scalar_one_or_none()
        return subscription is not None


async def create_subscription(user_id: int, token: str, patient_id: int):

    async with async_session() as session:
        result = await session.execute(
            select(Subscriptions).where(
                (Subscriptions.patient_id == patient_id) & 
                (Subscriptions.token == token)
            )
        )
        existing_subscription = result.scalar_one_or_none()
        print(existing_subscription, "LOW SKILLLL")
        if existing_subscription:
            # Update existing subscription
            existing_subscription.telegram_id = user_id
        else:
            # Create new subscription
            new_subscription = Subscriptions(
                telegram_id=user_id, token=token, patient_id=patient_id
            )
            session.add(new_subscription)
            
        await session.commit()  


async def get_subscriptions_by_user_id(user_id: int):
    async with async_session() as session:
        result = await session.execute(
            select(Subscriptions).where(Subscriptions.telegram_id == user_id)
        )
        subscriptions = result.scalars().all()
        return [
            {
                "id": sub.id,
                "telegram_id": sub.telegram_id,
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
                "telegram_id": subscription.telegram_id,
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
                "telegram_id": subscription.telegram_id,
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
