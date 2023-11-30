import openai
import os
from dotenv import load_dotenv
from flask import Flask, request, jsonify, render_template, send_file, Response
from typing import Union

app = Flask(__name__)

# Load environment variables
load_dotenv()

# Set OpenAI API Key
openai.api_key = os.getenv('OPENAI_API_KEY')


@app.route('/', methods=['GET'])
def index() -> str:
    return render_template('index.html')


@app.route('/upload', methods=['POST'])
def upload_file() -> Union[tuple[Response, int], Response]:
    if 'file' not in request.files or request.files['file'].filename == '':
        return jsonify({"error": "No file part or no selected file"}), 400

    file = request.files['file']
    original_content = file.read().decode('utf-8')
    specification = generate_specification(original_content)

    return jsonify({
        "original_content": original_content,
        "specification": specification
    })


def generate_specification(code) -> str:
    response = openai.Completion.create(
        model="text-davinci-003",
        prompt=f"""
            Please create a specification for the following code.
            Below is the template.
            #template
            Language: enter programming language here
            Library: Put your library here
            Package: put the package her
            Class: Put your class here
            Function: Here are local functions specific to this file.
            ・Here, the details of each local function are listed for each local function for this file only.
            Summary:
            This is where I put all the information from the file analysis.
            list format:\n\n{code}",
            """,
        temperature=0.3,
        max_tokens=200
    )
    return response.choices[0].text.strip()


@app.route('/download/<filename>', methods=['GET'])
def download_file(file_id) -> Response:
    filename = f"{file_id}.txt"
    response = send_file(filename, as_attachment=True)
    os.remove(filename)
    return response


if __name__ == '__main__':
    app.run(debug=True)
