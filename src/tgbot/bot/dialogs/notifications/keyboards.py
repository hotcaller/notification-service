import operator
from aiogram_dialog.widgets.kbd import ScrollingGroup, Select
from aiogram_dialog.widgets.text import Format

def paginated_bookings(on_click: callable, when: callable) -> ScrollingGroup:
    return ScrollingGroup(
        Select(
            Format("📅 {item[id]}"),
            id="s_scroll_notifications",
            item_id_getter=operator.itemgetter("id"),
            items="notifications",
            on_click=on_click,
            when=when
        ),
        id="notifications_id",
        width=1,
        height=6,
    )