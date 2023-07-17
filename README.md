# go-url-shortener
URL Shortener implemented using Go with Clean Architecture as architecture design to structure the codebase and some adapters to switch by initiliazing command.

![System Design Image](./assets/images/System%20Design%402x.png)

## Infrastructure
* ✅ Environment settings with Viper
* ⭕ Docker Image
* 🔄 Docker Compose
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
