### AWS Lambda with Serverless Application Model (AWS SAM)

#### SAM commands

```
sam local invoke HelloWorldFunction --event event.json --debug

 sam local invoke --debug 
```

#### Start Tracing Platform

- With [Jaeger](https://www.jaegertracing.io/)

```
docker run -d -p 4317:4317 -p 16686:16686 jaegertracing/all-in-one:latest
```

### Known Problems
- Error: Running AWS SAM projects locally requires Docker. Have you got it installed and running?
```
➜  quoteOrder git:(develop) ✗ docker context ls
NAME              DESCRIPTION                               DOCKER ENDPOINT                                        ERROR
default           Current DOCKER_HOST based configuration   unix:///var/run/docker.sock                            
desktop-linux *   Docker Desktop                            unix:///Users/sibellysanches/.docker/run/docker.sock   
➜  quoteOrder git:(develop) ✗ export DOCKER_HOST=unix:///Users/sibellysanches/.docker/run/docker.sock 
```
