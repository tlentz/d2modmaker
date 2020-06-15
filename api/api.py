import flask
from flask import request, jsonify
from models.d2data import D2Data
import json

app = flask.Flask(__name__)
app.config["DEBUG"] = True

data = D2Data()

@app.route('/', methods=['GET'])
def home():
    return f"<pre>{data.toJSON()}</pre>"

app.run()