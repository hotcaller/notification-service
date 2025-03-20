from sqlalchemy import Column, BigInteger, String, Boolean, ForeignKey
from bot.repository.db.db import Base
from sqlalchemy.orm import relationship

class User(Base):
    __tablename__ = "users"

    id = Column(BigInteger, primary_key=True)
    telegram_id = Column(BigInteger, unique=True, nullable=False)
    username = Column(String, nullable=True)
    has_access = Column(Boolean, default=False)

    subscriptions = relationship("Subscriptions", back_populates="user")

class Orgs(Base):
    __tablename__ = "orgs"

    id = Column(BigInteger, primary_key=True)
    name = Column(String, nullable=False)
    token = Column(String, nullable=False, unique=True)

    subscriptions = relationship("Subscriptions", back_populates="org")

class Subscriptions(Base):
    __tablename__ = "subscriptions"

    id = Column(BigInteger, primary_key=True)
    telegram_id = Column(BigInteger, ForeignKey("users.telegram_id"), nullable=False)
    token = Column(String, ForeignKey("orgs.token"), nullable=False)
    patient_id = Column(BigInteger, nullable=True)

    user = relationship("User", back_populates="subscriptions")

    org = relationship("Orgs", back_populates="subscriptions")
