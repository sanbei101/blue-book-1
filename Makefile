.PHONY: swagger swagger-init swagger-fmt help

swagger: swagger-fmt swagger-init

swagger-init:
	swag init -d internal/api,internal/pkg/render -g routes.go --parseInternal --v3.1 -ot yaml

swagger-fmt:
	swag fmt -d ./internal/api -g internal/api/routes.go