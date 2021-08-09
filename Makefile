API_NAME=kafka-ms-poc

build: 
	@echo "Creando Binario ..."
	@go build -o bin/main main.go
	@echo "Binario generado en /bin/${API_NAME}"
run:
	@echo "Creando Binario ..."
	@go build -o bin/main main.go
	@go run main.go

vendor:
	@echo "Vendoring..."
	@go mod vendor

.PHONY: test build