API_URL := http://127.0.0.1:4000/v1

.PHONY: healthcheck

healthcheck_request:
	@echo "Running healthcheck request.."
	@http $(API_URL)/healthcheck

run_dev:
	@echo "Starting the development server.."
	@go run cmd/api/*.go

# get_one_movie:
# 	@echo "Running http://localhost:4000/v1/movies/1"
# 	@http $(API_URL)/movies/1

