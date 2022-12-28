.PHONY: openapi
openapi-codegen:
	oapi-codegen -generate types -o api/v1/types/openapi.types.gen.go -package api api/openapi-spec/v1/api__v1_openapi.yaml
	oapi-codegen -generate gin -o api/v1/server/openapi.server.gen.go -package api api/openapi-spec/v1/api__v1_openapi.yaml	
	
.PHONY: install
install:
	@go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	@go install github.com/swaggo/swag/cmd/swag@v1.7.8

.PHONY: swagger
swagger:
#	@swag init -g cmd/app/main.go -o cmd/docs

.PHONY: build
build:
	@go mod tidy
	@go build -v ./...
