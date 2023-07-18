# go-url-shortener
URL Shortener implemented using Go with Clean Architecture as architecture design to structure the codebase and some adapters to switch by initiliazing command.

![System Design Image](./assets/images/System%20Design%402x.png)

## Infrastructure
* ✅ Environment settings with [Viper](https://github.com/spf13/viper)
* ✅ Docker Image
* ✅ Docker Compose
* ✅ Dependency Injection for Adapters using Factory Pattern
* ✅ OpenAPI Documentation/Swagger with [Swag](https://github.com/swaggo/swag)
* ⭕ CI&CD/Deploy

## Web Adapters
Built-in:
* ✅ net/http [🔗](https://pkg.go.dev/net/http)

External:
* ✅ Echo [🔗](https://github.com/labstack/echo)
* ✅ Fiber [🔗](https://github.com/gofiber/fiber)

## Repository Adapters
Built-in:
* ✅ In Memory

External:
* ✅ Memcached [🔗](https://github.com/bradfitz/gomemcache)
* ✅ Redis [🔗](https://github.com/redis/go-redis)

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
                - ADAPTER=http/echo # Change here
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
docker compose -f deploy/docker-compose.yml up --build -d server
```

## Swagger
After start the server, you can access the [Swagger Documentation](http://localhost:3000/swagger/).

## Makefile Commands
Run server setting ENV and ADAPTER:
```bash
make http env=<dev|qa|prod> adapter=<echo|fiber|standard>
```

Run tests with arguments:
```bash
make test args='-v'
```

Run tests coverage:
```bash
make coverage
```

Generate swagger documentation with swaggo/swag:
```bash
make swagger bin=/home/go/bin/  # bin is optional, should be used if you have not installed swag global
```