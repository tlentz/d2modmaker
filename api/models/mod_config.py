import json

def set_with_default_and_type(obj, key, t, default):
        if key in obj and isinstance(obj[key], t):
            return obj[key]
        else:
            return default

class MELEE_SPLASH_CFG:
    def __init__(self, obj = {})
        self.ON_JEWELS = set_with_default_and_type(obj, "ON_JEWELS", bool, False)
        self.RANDOMIZED_PROP = set_with_default_and_type(obj, "RANDOMIZED_PROP", bool, False)
        self.ON_ALL_WEAPONS = set_with_default_and_type(obj, "ON_ALL_WEAPONS", bool, False)
        self.ON_ALL_UNIQUES = set_with_default_and_type(obj, "ON_ALL_UNIQUES", bool, False)

class ModConfig:
    def __init__(self, obj = {}):
        self.INCREASED_STACK_SIZES = set_with_default_and_type(obj, "INCREASED_STACK_SIZES", bool, False)
        self.RANDOMIZE = set_with_default_and_type(obj, "RANDOMIZE", bool, False)
        self.MELEE_SPLASH = set_with_default_and_type(obj, "MELEE_SPLASH", bool, False)

    def toJSON(self):
        return json.dumps(self, default=lambda o: o.__dict__, indent=4)