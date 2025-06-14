DB_URL=mysql://root:root@tcp(127.0.0.1:3307)/Bank

migrate_up:
	migrate -path ./migrations -database "$(DB_URL)" up
migrate_down:
	migrate -path ./migrations -database "$(DB_URL)" down