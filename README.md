# go-url-shortener
URL Shortener implemented using Go with Clean Architecture as architecture design to structure the codebase and some adapters to switch by initiliazing command.

![System Design Image](./assets/images/System%20Design%402x.png)

## Infrastructure
* ✅ Environment settings with Viper
* ✅ Docker Image
* ✅ Docker Compose
* ✅ Dependency Injection for Adapters using Factory Pattern
* ⭕ OpenAPI Documentation/Swagger
* ⭕ CI&CD/Deploy

## Web Adapters
Built-in:
* ✅ net/http

External:
* ✅ Echo
* ✅ Fiber

## Repository Adapters
Built-in:
* ✅ In Memory

External:
* ✅ Memcached
* ✅ Redis

## To Do
* ⭕ Add tests for HTTP Adapters
* ⭕ Add tests for Presenters

## Setting up with Docker
This project was built to easily change vendors, so you can configure it according your desire.  

For Dockerfile and Docker Compose, by default the _HTTP Adapter_ configured is _Fiber_, but you can easily change it on:
* **Dockerfile**: When build the docker image you need to pass `--build-arg="ADAPTER=http/echo"` or `--build-arg="ADAPTER=http/stantard"` as argument. For example:
    ```bash
    docker build --build-arg="ADAPTER=http/echo" -f build/packages/Dockerfile -t url-shortener .
    ```
* **Docker Compose**: You need change in [docker-compose.yml](deploy/docker-compose.yml) file the `services.server.build.args.ADAPTER` configuration to `http/echo` or `http/standard`. For example:
    ```yml
    server:
        build:
        context: ../.
        dockerfile: build/packages/Dockerfile
        args:
            - GO_VERSION=1.20
            - ADAPTER=http/echo # Here is the change
        network: "host"
    ```

### Running with Docker
To setup the server with docker you need build the image first:
```bash
docker build --build-arg="ADAPTER=http/fiber" -f build/packages/Dockerfile -t url-shortener .
```

Then, run a new container with image built:
```bash
docker run --name url-shortener -v "$PWD":/usr/app/ -p 3000:3000 -d url-shortener
```

### Running with Docker Compose
To setup the server with docker compose you need run:
```bash
sudo docker compose -f deploy/docker-compose.yml up --build -d server
```
