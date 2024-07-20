docs:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag i -g handler.go -dir internal/controller/http/v1 --instanceName v1 --parseDependency


.PHONY: goose
goose:
	goose -dir migrations create $(name) sql
