.PHONY: openapi
openapi-codegen-server:
	oapi-codegen -package api -generate gin,types -o api/v1/openapi.gen.go api/openapi-spec/api__v1_openapi.yaml
openapi-codegen-client:
	openapi-generator generate -g go -i internal/pkg/domain/openapi-spec/api__v2_container-api.yaml -o internal/pkg/domain/container -p packageName=containerdomain -p enumClassPrefix=true --global-property skipFormModel=false,modelDocs=false,apiTests=false,apiDocs=false
	openapi-generator generate -g go -i internal/pkg/domain/openapi-spec/api__v1_aaa-api.yaml -o internal/pkg/domain/aaa -p packageName=aaadomain -p enumClassPrefix=true --global-property skipFormModel=false,modelTests=false,modelDocs=false,apiTests=false,apiDocs=false
 
.PHONY: install
install:
	@brew install openapi-generator
	@go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go get -d -v ./...

.PHONY: build
build:
	@go mod tidy
	@go get -d -v ./...
	@go build -v ./...

.PHONY: run
run:
	@go run ./cmd/app/main.go