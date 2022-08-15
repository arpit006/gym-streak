setup:
	#golang-migrate
	brew install golang-migrate

create-migration-file:
	migrate create -ext sql -dir db/migrations -seq create_users_table

migrations-up:
	migrate -source file://db/migrations -database mysql://'root:root@123'@'tcp(localhost:3306)'/gym_streak_app up

migrations-down:

