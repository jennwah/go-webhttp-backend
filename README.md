# Go template for web http API service

- [Gin framework](https://github.com/gin-gonic/gin)
- Clean Architecture
- MySQL Database integration
- Docker compose development
- ELK integration for application's logs


## To bootstrap service

1. go get ./...
2. docker compose -f docker-compose-elk.yml up -d
3. docker compose up -d
4. make migrations.up
5. try CURL 
```bash
curl --location --request POST 'http://localhost:8080/api/v1/task' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "new task"
}'
```
