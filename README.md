# Go template for web http API service

- [Gin framework](https://github.com/gin-gonic/gin)
- Clean Architecture
- MySQL Database integration
- Redis Cache integration
- Docker compose development
- ELK integration for application's logs


## To bootstrap service

1. install go dependencies `go get ./...`
2. run `make run-app` to bootstrap app with docker compose
3. `make migrations.up` for database migrations
4. try CURL 
```bash
curl --location --request POST 'http://localhost:8080/api/v1/task' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "new task"
}'
```
5. try CURL 
```
curl --location --request GET 'http://localhost:8080/api/v1/task/1' \
--header 'Content-Type: application/json' \
```
6. Check application's logs at <b>http://localhost:5601/app/logs/stream</b> with Kibana UI
7. If you're interested in caching strategy used in this app, (read aside strategy), see on `Get task by ID usecase` on its implementation

`docker exec -it {redis-container-id} redis-cli` 

`HGET key fields..`


## Reference Gin's [Quick Start Guide] (https://github.com/gin-gonic/gin/blob/master/docs/doc.md)

- for web development stuff such as middlewares, custom validators, handling errors, routes, request binding and etc
