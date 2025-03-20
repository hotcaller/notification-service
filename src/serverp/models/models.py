from sqlalchemy import Column, Integer, String, ForeignKey, DateTime, Text
from sqlalchemy.orm import relationship
from server.db.session import Base
from datetime import datetime

class Patient(Base):
    __tablename__= "patients"
    id = Column(Integer, primary_key=True, index=True)
    name = Column(String)
    token = Column(String)
    notifications = relationship("Notification", back_populates="patient")

class Notification(Base):
    __tablename__= "notifications"
    id = Column(Integer, primary_key=True, index=True)
    title = Column(String)
    message = Column(Text)
    timestamp = Column(DateTime, default=datetime.utcnow)
    patient_id = Column(Integer, ForeignKey("patients.id"))
    patient = relationship("Patient", back_populates="notifications")
