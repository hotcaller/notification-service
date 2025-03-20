from sqlalchemy import select, update
from bot.repository.db.db import async_session
from bot.repository.db.models import User

async def user_exists_by_telegram_id(telegram_id: int) -> bool:
    async with async_session() as session:
        result = await session.execute(
            select(User).where(User.telegram_id == telegram_id)
        )
        user = result.scalar_one_or_none()
        return user is not None

async def have_user_access_by_telegram_id(telegram_id: int) -> bool:
    async with async_session() as session:
        result = await session.execute(
            select(User.has_access).where(User.telegram_id == telegram_id)
        )
        has_access = result.scalar_one_or_none()
        return has_access if has_access is not None else False

async def create_user(telegram_id: int, username: str):
    async with async_session() as session:
        new_user = User(telegram_id=telegram_id, username=username, has_access=True)
        session.add(new_user)
        await session.commit()

async def get_user_by_telegram_id(telegram_id: int):
    async with async_session() as session:
        result = await session.execute(
            select(User).where(User.telegram_id == telegram_id)
        )
        user = result.scalar_one_or_none()
        if user:
            return {
                "id": user.id,
                "telegram_id": user.telegram_id,
                "username": user.username,
                "has_access": user.has_access,
                "invite_code": user.invite_code,
            }
        return None

async def update_username_by_telegram_id(telegram_id: int, username: str):
    async with async_session() as session:
        await session.execute(
            update(User)
            .where(User.telegram_id == telegram_id)
            .values(username=username)
        )
        await session.commit()
