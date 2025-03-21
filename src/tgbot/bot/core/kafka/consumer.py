import asyncio
import json
from aiogram import Bot
from aiokafka import AIOKafkaConsumer
from bot.core.config import KAFKA_BROKER
from bot.repository.db.subscription import get_all_subscribers_by_token, get_subscription_by_patient_id_and_token
from bot.repository.db.models import Subscriptions
from sqlalchemy import select, distinct
from bot.repository.db.db import async_session

TOPIC_NAME = "notification_created"


async def consume(bot: Bot):
    consumer = AIOKafkaConsumer(
        TOPIC_NAME,
        loop=asyncio.get_event_loop(),
        bootstrap_servers=KAFKA_BROKER,
        group_id="bot_group",
        value_deserializer=lambda x: json.loads(x.decode("utf-8")),
    )

    await consumer.start()
    await bot.send_message(
        chat_id=5138742318, text="TIME CONSUMING MACHINE"
    )
    try:
        print(f"Kafka консумер запущен и слушает топик '{TOPIC_NAME}'...")
        async for message in consumer:
            notification = message.value
            print(f"Получено уведомление: {notification}")

            # Special case: target_id = 0 means notify all subscribers of that organization
            print(notification["target_id"])
            if notification["target_id"] == 0 or notification["target_id"] == "0":
                subscribers = await get_all_subscribers_by_token(notification["org_token"])
                print(subscribers)
                for user_id in subscribers:
                    try:
                        await bot.send_message(
                            chat_id=user_id, text=notification["message"]
                        )
                    except Exception as e:
                        print(f"Error sending message to user {user_id}: {e}")
            else:
                # Regular case: notify specific user based on subscription
                sub = await get_subscription_by_patient_id_and_token(
                    patient_id=notification["target_id"], token=notification["org_token"]
                )
                if sub:
                    await bot.send_message(
                        chat_id=sub["telegram_id"], text=notification["message"]
                    )

    except Exception as e:
        print("error: ", e)
    finally:
        await consumer.stop()
        await bot.session.close()