build:
	@go build -o temp cmd/main.go 
	@go build -o file-server ./storage.go
	@rm ./dist && mv ./temp ./dist
		
build-seeder:
	@go build -o seeder ./seed

run: build
	@ls dist | entr -cr ./dist

run-file: build
		@clear
		@./file-server

init-db: build-seeder
	@./seeder --init-db 

nuke-db: build-seeder
	@./seeder --nuke-db 

seed-db: build-seeder
	@./seeder --seed

refresh-db: build-seeder
	@./seeder --refresh-db
	
	
