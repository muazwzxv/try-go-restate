include .env

migrations.new:
	migrate create -ext sql -dir ./db/migrations -seq $(name)

migrations.up:
	migrate -path ./db/migrations -database "${DB_SOURCE}" -verbose up $(N_MIGRATION)

migrations.down:
	migrate -path ./db/migrations -database "${DB_SOURCE}" -verbose down $(N_MIGRATION)

proto:
	protoc --twirp_out=. --go_out=. internal/rpc/user.proto
