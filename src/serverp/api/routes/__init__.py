from fastapi import APIRouter
from server.api.routes import login, qr, notifications

api_router = APIRouter()
api_router.include_router(login.router, tags=["login"])
api_router.include_router(qr.router, tags=["qr"])
api_router.include_router(notifications.router, tags=["notifications"])
