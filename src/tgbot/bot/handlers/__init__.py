from aiogram import Router
from .callbacks.callback import r as callbacks_r
from .common.start import r as start_r
from .common.errors import r as errors_r

router = Router()

router.include_routers(errors_r, callbacks_r, start_r)
