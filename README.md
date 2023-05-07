# Go template for web http API service

- [Gin framework](https://github.com/gin-gonic/gin)
- Clean Architecture
- MySQL Database integration
- Redis Cache integration
- Docker compose development
- minikube local development for kubernetes
- ELK integration for application's logs


## To bootstrap service with docker compose

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

## To bootstrap service with minikube 

1. `brew install minikube`
2. `minikube start`
3. `kubectl apply -f {all files in ./kubernetes dir}` 
4. `kubectl logs -l app=backend/mysql/redis` to debug. `kubectl get pods`, `kubectl describe pods {pod_name}`, `kubectl get deployments`
5. `minikube service backend -url` to get backend server URL. 


## Reference Gin's [Quick Start Guide] (https://github.com/gin-gonic/gin/blob/master/docs/doc.md)

- for web development stuff such as middlewares, custom validators, handling errors, routes, request binding and etc
