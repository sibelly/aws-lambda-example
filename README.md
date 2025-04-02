### AWS Lambda with Serverless Application Model (AWS SAM)

#### SAM commands

```
sam local invoke HelloWorldFunction --event event.json --debug
```

#### Start Tracing Platform

- With [Jaeger](https://www.jaegertracing.io/)

```
docker run -d -p 4317:4317 -p 16686:16686 jaegertracing/all-in-one:latest
```

