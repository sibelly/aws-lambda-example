
https://github.com/open-telemetry/opentelemetry-lambda/tree/main/go
https://pkg.go.dev/go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda#section-readme

### Invoke local

```
serverless invoke local --function HelloWorldFunction --debug
SLS_DEBUG=* serverless offline
 DISABLE_TRACING=true STAGE=dev serverless offline start
```

```
curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{}'
```