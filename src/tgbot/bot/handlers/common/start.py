from aiogram.types import Message
from aiogram.filters import CommandStart, CommandObject
from aiogram_dialog import DialogManager, StartMode
from aiogram import Router

from bot.dialogs.start.states import StartMenu
from bot.repository.db.user import (
    user_exists_by_telegram_id,
    have_user_access_by_telegram_id,
    get_user_by_telegram_id,
    create_user,
    update_username_by_telegram_id,
)
from bot.repository.db.subscription import create_subscription

r = Router()


async def handle_start_with_invite_code(
    message: Message, invite_code: str, dialog_manager: DialogManager
) -> None:
    params = invite_code.split("|")

    if len(params) == 2:
        id = params[0]
        token = params[1]

    id = params
    token = 123

    await create_user(message.from_user.id, message.from_user.username)
    await create_subscription(message.from_user.id, token, id)


@r.message(CommandStart())
async def start_handler(
    message: Message, command: CommandObject, dialog_manager: DialogManager
) -> None:
    user_id = int(message.from_user.id)
    invite_code = command.args
    await message.answer(f"{invite_code} {command}")

    if not await user_exists_by_telegram_id(user_id):
        if invite_code:
            await handle_start_with_invite_code(message, invite_code, dialog_manager)
        await create_user(user_id, message.from_user.username)

    if not await have_user_access_by_telegram_id(user_id):
        await message.answer("*Упс...* Пока что у вас нет доступа к этому боту ⏳")
        return

    user = await get_user_by_telegram_id(user_id)

    if user["username"] != message.from_user.username:
        await update_username_by_telegram_id(user_id, message.from_user.username)

    await dialog_manager.start(StartMenu.select_menu, mode=StartMode.RESET_STACK)
