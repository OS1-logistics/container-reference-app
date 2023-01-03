.PHONY: openapi
openapi-codegen:
	oapi-codegen -package api -generate gin,types -o api/v1/openapi.gen.go api/openapi-spec/api__v1_openapi.yaml
	openapi-generator generate -g go -i internal/pkg/clients/openapi-spec/api__v2_container-api.yaml -o internal/pkg/clients/container -p packageName=container_client -p enumClassPrefix=true --global-property skipFormModel=false
	openapi-generator generate -g go -i internal/pkg/clients/openapi-spec/api__v1_aaa-api.yaml -o internal/pkg/clients/aaa -p packageName=aaa_client -p enumClassPrefix=true --global-property skipFormModel=false
#	openapi-generator generate -g go -i internal/pkg/clients/openapi-spec/api__v1_aaa-api.yaml -o internal/pkg/clients/aaa -p packageName=aaa_client -p enumClassPrefix=true -p prependFormOrBodyParameters=true -p structPrefix=true -p useOneOfDiscriminatorLookup=true --global-property generateAliasAsModel=true
#	openapi-generator generate -g go -i internal/pkg/clients/openapi-spec/api__v1_aaa-api.yaml -o internal/pkg/clients/aaa --additional-properties packageName=aaa_client --additional-properties enumClassPrefix=true --additional-properties prependFormOrBodyParameters=true --additional-properties structPrefix=true --additional-properties useOneOfDiscriminatorLookup=true  --global-property skipFormModel=false

openapi-generator generate -g go -i internal/pkg/clients/openapi-spec/api__v1_aaa-api.yaml -o internal/pkg/clients/aaa --additional-properties packageName=aaa_client --additional-properties enumClassPrefix=true --additional-properties useOneOfDiscriminatorLookup=true  --global-property skipFormModel=false
 
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
