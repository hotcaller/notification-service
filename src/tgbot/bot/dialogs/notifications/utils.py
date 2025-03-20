
def check_has_items(data:dict, widget, manager) -> bool:
    return True if len(data["notifications"]) > 0 else False
