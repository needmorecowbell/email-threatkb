# Yara Processor

YARA Processor Server. Just takes in the raw eml and checks it against the yara ruleset, returning a json response. 

## Usage
- `go run main.go`

spins up the server, opens an index route with welcome banner and a /scan route which accepts post requests, with the raw eml as byte data.

curl:

```
curl -X POST --data-binary "@/home/user/Downloads/test.eml"  http://127.0.0.1:8080/scan     
{"matches":["DetectMalicious"],"status":"malicious"}
```