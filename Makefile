API_URL := http://127.0.0.1:4000/v1

.PHONY: healthcheck

healthcheck_request:
	@echo "Running healthcheck request.."
	@http $(API_URL)/healthcheck

run_dev:
	@echo "Starting the development server.."
	@go run cmd/api/*.go

migrate_up:
	@echo "Running migration up"
	@migrate -path=./migrations -database='postgresql://pratik:password@localhost/greenlight?sslmode=disable' up

migrate_down:
	@echo "Running migration down"
	@migrate -path=./migrations -database='postgresql://pratik:password@localhost/greenlight?sslmode=disable' down
