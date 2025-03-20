from aiogram import Router
from aiogram.types.error_event import ErrorEvent
import logging
import traceback
from aiogram_dialog import DialogManager

r = Router()


@r.error()
async def error_handler(event: ErrorEvent, dialog_manager: DialogManager) -> None:
    logging.error("An error occurred:", exc_info=event.exception)

    exception = event.exception
    tb_exception = traceback.TracebackException.from_exception(exception)
    traceback_text = "".join(tb_exception.format())

    user_id = None

    if event.update:
        if hasattr(event.update, "message") and event.update.message:
            user_id = event.update.message.from_user.id
        elif hasattr(event.update, "callback_query") and event.update.callback_query:
            user_id = event.update.callback_query.from_user.id
        elif hasattr(event.update, "inline_query") and event.update.inline_query:
            user_id = event.update.inline_query.from_user.id



    await dialog_manager.event.bot.send_message(chat_id=user_id, text=traceback_text)

