.PHONY: openapi
openapi-codegen:
	oapi-codegen -package api -generate gin,types -o api/v1/openapi.gen.go api/openapi-spec/api__v1_openapi.yaml
#	openapi-generator generate -g go -i internal/pkg/clients/openapi-spec/api__v2_containerapi.yaml -o internal/pkg/clients/container -p packageName=container_client -p enumClassPrefix=true --skip-validate-spec

.PHONY: install
install:
	@brew install openapi-generator
	@go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go get -d -v ./...

.PHONY: swagger
swagger:
#	@swag init -g cmd/app/main.go -o cmd/docs

.PHONY: build
build:
	@go mod tidy
	@go get -d -v ./...
	@go build -v ./...
