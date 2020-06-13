import flask
from flask import request, jsonify
from models.d2data import D2Data

app = flask.Flask(__name__)
app.config["DEBUG"] = True

data = D2Data()

@app.route('/', methods=['GET'])
def home():
    return str(data.fileData)

app.run()