migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/village-square?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/village-square?sslmode=disable" -verbose down

herodeploy:
	psql --host=ec2-54-170-163-224.eu-west-1.compute.amazonaws.com --port=5432 --username=vysjvbplqhhnkq --password --dbname=dfi5c1hhmnm8k4

heroku pg:psql --app village-square < db/migration/000001_village-square_schema.up.sql