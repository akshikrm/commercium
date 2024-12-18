build:
	@go build -o dist cmd/main.go 
	@go build -o file-server ./storage.go
		
run: build
	@clear
	@./dist

run-file: build
		@clear
		@./file-server

handler:
	rm ./seeder
	@go build -o seeder ./seed

init-db: seeder
	@./seeder --init-db 

seed-db: seeder
	@./seeder --seed

seed-users: seeder
	@./seeder --seed-users

seed-groups: seeder
	@./seeder --seed-groups

seed-roles: seeder
	@./seeder --seed-roles

seed-resources: seeder
	@./seeder --seed-resources

seed-products: seeder
	@./seeder --seed-products


seed-product-categories: seeder
	@./seeder --seed-product-categories

seed-permission: seeder
		@./seeder --seed-permission

nuke-db: seeder
	@./seeder --nuke-db

refresh-db: seeder 
	@./seeder --refresh-db
