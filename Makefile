db_conn = "root:pass@tcp(localhost:3307)/go-backend?tls=false"

run-server:
	go build -o build/app cmd/main.go
	build/app

# Database migrations
migrations.up:
	migrate -path database/migrations -database mysql://$(db_conn) -verbose up
migrations.down:
	migrate -path database/migrations -database mysql://$(db_conn) -verbose down
migrations.create:
	migrate create -ext sql -dir database/migrations $(name)