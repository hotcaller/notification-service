from sqlalchemy import Column, BigInteger, String, Boolean
from bot.repository.db.db import Base

class User(Base):
    __tablename__= "users"

    id = Column(BigInteger, primary_key=True)
    telegram_id = Column(BigInteger, unique=True, nullable=False)
    username = Column(String, nullable=True)
    has_access = Column(Boolean, default=False)
    invite_code = Column(String, nullable=True)
