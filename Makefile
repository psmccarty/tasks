build: sqlc
	go build -o tasks

add:
	go run main.go add

list:
	go run main.go list

complete:
	go run main.go complete

delete:
	go run main.go delete

sqlc: sql/query.sql sql/schema.sql
	sqlc generate -f sql/sqlc.yml