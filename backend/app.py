from flask import Flask, request, jsonify
from flask_cors import CORS
import cohere
import os
from dotenv import load_dotenv

# 🔄 Load environment variables from .env
load_dotenv()
api_key = os.getenv("COHERE_API_KEY")

# ✅ Initialize Flask app
app = Flask(__name__)
CORS(app)  # Enable Cross-Origin Resource Sharing

# ✅ Initialize Cohere client
co = cohere.Client(api_key)

@app.route('/')
def home():
    return "🚀 Flask summarizer server is running!"

@app.route('/summarize', methods=['POST'])
def summarize():
    data = request.get_json()
    text = data.get("text", "")

    # 🚫 Check for empty input
    if not text:
        return jsonify({"error": "No text provided"}), 400

    # 🧠 Call Cohere API
    try:
        print("🧠 Calling Cohere to summarize...")
        response = co.summarize(
            text=text,
            length='medium',       # short / medium / long
            format='paragraph',    # sentence / paragraph / bulleted
            model='summarize-xlarge',
            temperature=0.3,
        )
        summary = response.summary
        print("✅ Summary:", summary)
        return jsonify({"summary": summary})

    except Exception as e:
        print("❌ Error:", str(e))
        return jsonify({"error": "Summarization failed"}), 500

# 🚀 Run Flask server on port 5001
if __name__ == '__main__':
    app.run(port=5001, debug=True)
