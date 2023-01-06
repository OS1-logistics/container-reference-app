# Container Reference App

## Local setup
### Prerequisites

- [Go](https://golang.org/doc/install)
- [Make](https://www.gnu.org/software/make/)

Run make to install the dependencies:

```
make install
```

## Codegen 

Client and server code can be generated using the `openapi-generator` tool. The tool is installed as part of the `make install` command. 

- Server code is generated from the `api/openapi-spec/api__v1_openapi.yaml` file using command `make openapi-codegen-server`. The generated code is stored in `api/v1/` directory.
- OS1 platform client code is generated from the `internal/pkg/domain/openapi-spec/api__v1_<domainname>-api.yaml` file using command `make openapi-codegen-client`. The generated code is stored in `internal/pkg/domain/<domainname>` directory.

### Add a new API to the service 

- Update `api/openapi-spec/api__v1_openapi.yaml` file to add the new API
- Run `make openapi-codegen-server` to generate the server code and its associated types.

### Add a new OS1 service 

- Download the OpenAPI spec from the OS1 documentation site https://swagger.preprod.fxtrt.io/
- Download openapi yaml file to internal/package/domain/openapi-spec/api__v1_<domainname>-api.yaml
- Add openapi spec generation command to Makefile. Example: Replace domainname with the actual domain name. 
```
openapi-codegen-client:	
	openapi-generator generate -g go -i internal/pkg/domain/openapi-spec/api__v2_<domainname>-api.yaml -o internal/pkg/domain/<domainname> -p packageName=<domainname>domain -p enumClassPrefix=true --global-property skipFormModel=false,modelDocs=false,apiTests=false,apiDocs=false
```
- Run `make openapi-codegen-client` to generate the client code.

### To Run the App

To run the app in local for development, run the following command: 

```
make run
```