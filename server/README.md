# Yara Processor

Proof of concept yara processor. Just takes in the raw eml and checks it against the yara ruleset, returning a json response. 

## Installation

- `apt install python3 python3-pip yara python3-flask python3-yara gunicorn` (or using a virtualenv)

## Usage
- `gunicorn app:app --bind=127.0.0.1:6000`

spins up the flask server, opens an index route with welcome banner and a /scan route which accepts post requests, with the raw eml as byte data.

curl:

```
curl -X POST --data-binary "@/home/user/Downloads/test.eml"  http://127.0.0.1:6000/scan     
{"matches":["DetectMalicious"],"status":"malicious"}
```