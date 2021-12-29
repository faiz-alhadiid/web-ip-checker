from flask import Flask, jsonify, request
import ip_lookup
import os

app = Flask(__name__)

@app.route("/")
def index():
    with open("index.html", "r") as file:
        return file.read()

@app.route("/check-ip")
def check_ip():
    url = request.args.get("url", "")
    dig_text = ip_lookup.get_dig_result(url)
    result = ip_lookup.parse_dig_result(dig_text)
    return jsonify(result)

if __name__ == "__main__":
    app.run(debug=True, host="0.0.0.0", port=int(os.environ.get("PORT", 4000)))