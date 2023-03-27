init:
	swag init
	@open "http://localhost:7000/swagger/index.html"
	go run main.go