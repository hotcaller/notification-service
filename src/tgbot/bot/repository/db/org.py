from sqlalchemy import select
from bot.repository.db.db import async_session
from bot.repository.db.models import Orgs


async def org_exists_by_id(id: int) -> bool:
    async with async_session() as session:
        result = await session.execute(select(Orgs).where(Orgs.id == id))
        org = result.scalar_one_or_none()
        return org is not None


async def create_org(name: str, token: str):
    async with async_session() as session:
        new_org = Orgs(name=name, token=token)
        session.add(new_org)
        await session.commit()


async def get_org_by_id(id: int):
    async with async_session() as session:
        result = await session.execute(select(Orgs).where(Orgs.id == id))
        org = result.scalar_one_or_none()
        if org:
            return {
                "id": org.id,
                "name": org.name,
                "token": org.token,
            }
        return None
