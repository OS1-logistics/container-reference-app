.PHONY: openapi
openapi-codegen:
	oapi-codegen -package api -generate gin,types -o api/v1/openapi.gen.go api/openapi-spec/api__v1_openapi.yaml	
	oapi-codegen -package container_client -generate client,types -o internal/pkg/clients/container/openapi.container.types.gen.go internal/pkg/clients/openapi-spec/api__v2_containerapi.yaml
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
