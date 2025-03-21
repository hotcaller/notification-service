import asyncio
import json
from aiogram import Bot
from aiokafka import AIOKafkaConsumer
from bot.core.config import KAFKA_BROKER
from bot.repository.db.subscription import get_all_subscribers_by_token, get_subscription_by_patient_id_and_token
from bot.repository.db.db import async_session
from bot.repository.db.models import Subscriptions

TOPIC_NAME = "notification_created"


async def consume(bot: Bot):
    consumer = None
    try:
        consumer = AIOKafkaConsumer(
            TOPIC_NAME,
            loop=asyncio.get_event_loop(),
            bootstrap_servers=KAFKA_BROKER,
            group_id="bot_group",
            value_deserializer=lambda x: json.loads(x.decode("utf-8")),
            auto_offset_reset="earliest"
        )

        await consumer.start()
        print(f"Kafka consumer started for topic '{TOPIC_NAME}'")
        
        await bot.send_message(
            chat_id=5138742318, text="TIME CONSUMING MACHINE"
        )
        
        async for message in consumer:
            try:
                notification = message.value
                print(f"Received notification: {notification}")

                notification_type = notification.get("type", "default")
                emoji_map = {
                    "news": "📰",
                    "reminder": "⏰",
                    "warning": "⚠️",
                    "important": "❗",
                    "default": "ℹ️"  # Default emoji if type not specified
                }
                emoji = emoji_map.get(notification_type, emoji_map["default"])
                formatted_message = f"{emoji} *{notification['header']}*\n\n{notification['message']}"
                # Special case: target_id = 0 means notify all subscribers of that organization
                if notification["target_id"] == 0:
                    subscribers = await get_all_subscribers_by_token(notification["org_token"])
                    print(f"Broadcasting to {len(subscribers)} subscribers")
                    
                    for user_id in subscribers:
                        try:
                            await bot.send_message(
                                chat_id=user_id, text=formatted_message, parse_mode="Markdown"
                            )
                        except Exception as e:
                            print(f"Failed to send to user {user_id}: {e}")
                else:
                    # Regular case: notify specific user based on subscription
                    sub = await get_subscription_by_patient_id_and_token(
                        patient_id=notification["target_id"], token=notification["org_token"]
                    )
                    if sub:
                        print(f"Sending targeted notification to user {sub['telegram_id']}")
                        await bot.send_message(
                            chat_id=sub["telegram_id"], text=notification["message"]
                        )
                    else:
                        print(f"No subscription found for patient_id={notification['target_id']}")
                
            except Exception as e:
                print(f"Error processing message: {e}")
                import traceback
                print(traceback.format_exc())

    except Exception as e:
        print(f"Kafka consumer error: {e}")
        import traceback
        print(traceback.format_exc())
    finally:
        if consumer:
            await consumer.stop()
        await bot.session.close()