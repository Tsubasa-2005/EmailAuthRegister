.PHONY: run
run:
	go run main.go serve

.PHONY: ogen
ogen:
	ogen -target ui/api -clean docs/bundle.yaml

.PHONY: redocly
redocly:
	redocly lint ./docs/openapi.yaml
	redocly bundle ./docs/openapi.yaml -o ./docs/bundle.yaml

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: docker
docker:
	docker compose up -d