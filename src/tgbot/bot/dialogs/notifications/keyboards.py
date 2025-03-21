from aiogram_dialog.widgets.kbd import ScrollingGroup, Select, Row
from aiogram_dialog.widgets.text import Format

def paginated_notifications(on_click, when=None):
    return ScrollingGroup(
        Select(
            Format("{item.header}"),
            id="s_notifications",
            item_id_getter=lambda x: x.id,
            items="notifications",
            on_click=on_click,
        ),
        id="sg_notifications",
        width=1,
        height=5,
        when=when,
    )