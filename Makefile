setup:
	#golang-migrate
	brew install golang-migrate

# run as make create-migration-file table_name="create_exercise_table"
create-migration-file:
	migrate create -ext sql -dir db/migrations -seq $(table_name)

migrations-up:
	migrate -source file://db/migrations -database mysql://'root:root@123'@'tcp(localhost:3306)'/gym_streak_app up

migrations-down:

