### Study project to execute local aws lambda

```
npm install serverless-offline --save-dev
sls offline --help
SLS_DEBUG=* serverless offline

serverless invoke local --function HelloWorldFunction --debug
GOARCH=amd64 GOOS=linux go build -o bootstrap main.go
GOOS=linux GOARCH=amd64 go build -o bin/main aws-golang/main.go
```