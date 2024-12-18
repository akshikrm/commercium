build:
	@go build -o dist cmd/main.go 
	@go build -o file-server ./storage.go
		
build-seeder:
	rm ./seeder
	@go build -o seeder ./seed

run: build
	@clear
	@./dist

run-file: build
		@clear
		@./file-server

init-db: build-seeder
	@./seeder --init-db 

seed-db: build-seeder
	@./seeder --seed
