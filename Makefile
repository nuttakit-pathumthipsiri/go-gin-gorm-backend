# Generate Swagger documentation
swagger:
	C:\Users\nutta\go\bin\swag.exe init -g cmd/main.go

# Run the application
run:
	go run ./cmd/main.go

# Run the application with Swagger
run dev:
	C:\Users\nutta\go\bin\swag.exe init -g cmd/main.go
	go run ./cmd/main.go

# Build the application
build:
	go build -o bin/app ./cmd/main.go
