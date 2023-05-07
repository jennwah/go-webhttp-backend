db_conn = "root:pass@tcp(localhost:3306)/go-backend?tls=false"

run-server:
	go build -o build/app cmd/main.go
	build/app
run-app-with-docker-compose:
	docker-compose -f docker-compose-elk.yml up -d
	docker-compose build server
	docker-compose up -d
run-app-with-minikube:
	kubectl apply -f ./kubernetes/mysql-secret.yaml
	kubectl apply -f ./kubernetes/mysql-configmap.yaml
	kubectl apply -f ./kubernetes/mysql-pv.yaml
	kubectl apply -f ./kubernetes/mysql-deployment.yaml
	kubectl apply -f ./kubernetes/redis-deployment.yaml
	kubectl apply -f ./kubernetes/server-deployment.yaml
	minikube service backend -url
# Database migrations
migrations.up:
	migrate -path database/migrations -database mysql://$(db_conn) -verbose up
migrations.down:
	migrate -path database/migrations -database mysql://$(db_conn) -verbose down
migrations.create:
	migrate create -ext sql -dir database/migrations $(name)