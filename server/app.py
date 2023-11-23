from flask import Flask, request, jsonify
import yara
import tempfile
import os
from datetime import datetime

app = Flask(__name__)

# Define YARA rules
yara_rules = """
rule DetectMalicious {
    strings:
        $malicious_string = "malicious_pattern"
    condition:
        $malicious_string
}
"""

# Load YARA rules
compiled_rules = yara.compile(source=yara_rules)

@app.route('/')
def index():
    return jsonify({"status": "success","message":"Welcome to the eml processor"})

@app.route('/scan', methods=['POST'])
def scan_eml():
    try:
        # Get the raw EML file from the POST request
        raw_eml: bytes = request.get_data()
        
        # Create a temporary file to save the raw EML content
        with tempfile.NamedTemporaryFile(delete=False) as temp_file:
            temp_filename = temp_file.name
            temp_file.write(raw_eml)

        # Scan the EML file with YARA rules
        matches = compiled_rules.match(temp_filename)
        # Delete the temporary file
        os.remove(temp_filename)

        # Check if any matches were found
        if matches:
            result = {'status': 'malicious', 'matches': [str(match) for match in matches]}
            
        else:
            result = {'status': 'clean', 'matches': []}

        return jsonify(result)

    except Exception as e:
        return jsonify({'status': 'error', 'message': str(e)}), 500

if __name__ == '__main__':
    app.run(debug=True, host="0.0.0.0",port=6000)

