migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/village-square?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/village-square?sslmode=disable" -verbose down

herodeploy:
	psql --host=ec2-63-32-7-190.eu-west-1.compute.amazonaws.com --port=5432 --username=zcxydxjqbdeeqe --password --dbname=dch1l4n831j80m

heroku pg:psql --app village-square < db/migration/000001_village-square_schema.up.sql .