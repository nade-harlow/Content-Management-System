migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/village-square?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/village-square?sslmode=disable" -verbose down