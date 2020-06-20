import flask
from flask import request, jsonify
from models.d2data import D2Data
from models.mod_config import ModConfig
import json

app = flask.Flask(__name__)
app.config["DEBUG"] = True

# data = D2Data()
cfg_dict = {
    'INCREASED_STACK_SIZES': True,
    'MELEE_SPLASH': {
        'ON_JEWELS': True
    }
}
cfg = ModConfig(cfg_dict)

@app.route('/', methods=['GET'])
def home():
    return f"<pre>{cfg.toJSON()}</pre>"

app.run()