# Container Reference App

## Local setup
### Prerequisites

- [Go](https://golang.org/doc/install)
- [Make](https://www.gnu.org/software/make/)
- [openapi-generator](https://openapi-generator.tech/) install the one which is compatible with your OS.

### To Run the App

- To install the dependencies `make install`
- To build the app `make build`
- To run the app run `export $(cat .env.local | xargs)` to export the environment variables then run `make run`

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


## Sample App Features 

### Ability to initialize the tenant with container types

The server is initialized with a list of container types that it needs to create for the service to operate. As part of this reference app, we are creating two container types `Package` and `Bag`.

Package is a container that can hold a single item. Its a container of leaf type. Bag is a container that can hold multiple items/ packages. Its a container of type non leaf.

```
curl --location --request POST 'http://localhost:3000/api/v1/initialize' \
--header 'X-COREOS-REQUEST-ID: 8f1fb9f5-19b1-490e-b298-132847e41ac9' \
--header 'X-COREOS-TENANT-ID: alpha' \
--header 'Accept: application/json'
```

### Ability to create new package 

```
curl --location --request POST 'http://localhost:3000/api/v1/packages' \
--header 'X-COREOS-REQUEST-ID: b4db4f62-5531-441a-a48d-809cb8f77935' \
--header 'X-COREOS-TENANT-ID: alpha' \
--header 'Content-Type: application/json' \
--header 'Accept: application/json' \
--data-raw '{
  "trackingId": "TR223433WD",
  "origin": "facility:317221c7-831d-4fde-9ee4-13e7898c2485",
  "destination": "facility:317221c7-831d-4fde-9ee4-13e7898c2485",
  "scannableId": "CNT4433TRD",
  "isReusable": false,
  "isHazmat": false
}'
```

### Ability to create a new bag 

```
curl --location --request POST 'http://localhost:3000/api/v1/bags' \
--header 'X-COREOS-REQUEST-ID: fcf9aa52-c9a0-4d76-b501-bfcc898fc636' \
--header 'X-COREOS-TENANT-ID: alpha' \
--header 'Content-Type: application/json' \
--header 'Accept: application/json' \
--data-raw '{
  "trackingId": "TR223433WD",
  "origin": "facility:317221c7-831d-4fde-9ee4-13e7898c2485",
  "destination": "facility:317221c7-831d-4fde-9ee4-13e7898c2485",
  "scannableId": "BAGT#23223FNT",
  "isReusable": false,
  "isHazmat": false
}'
```

### Ability to perform operations on package and bag

- Open a package 

```
curl --location --request POST 'http://localhost:3000/api/v1/packages/PackageCsa:7a635074-c9b2-5b66-9d09-0a1909b16aeb/state/open' \
--header 'X-COREOS-REQUEST-ID: e993137a-da82-4a35-afc2-c5766ab084b1' \
--header 'X-COREOS-TENANT-ID: alpha' \
--header 'Accept: application/json'
```

- Close a package

```
curl --location --request POST 'http://localhost:3000/api/v1/packages/PackageCsa:7a635074-c9b2-5b66-9d09-0a1909b16aeb/state/close' \
--header 'X-COREOS-REQUEST-ID: e993137a-da82-4a35-afc2-c5766ab084b1' \
--header 'X-COREOS-TENANT-ID: alpha' \
--header 'Accept: application/json'
```

- Open a bag 

```
curl --location --request POST 'http://localhost:3000/api/v1/bags/BagT:7a635074-c9b2-5b66-9d09-0a1909b16aeb/state/open' \
--header 'X-COREOS-REQUEST-ID: e993137a-da82-4a35-afc2-c5766ab084b1' \
--header 'X-COREOS-TENANT-ID: alpha' \
--header 'Accept: application/json'
```
- Close a bag 

```
curl --location --request POST 'http://localhost:3000/api/v1/bags/BagT:7a635074-c9b2-5b66-9d09-0a1909b16aeb/state/close' \
--header 'X-COREOS-REQUEST-ID: e993137a-da82-4a35-afc2-c5766ab084b1' \
--header 'X-COREOS-TENANT-ID: alpha' \
--header 'Accept: application/json'
```

### Ability to add a package to a bag

Before adding a package to a bag, the package should be in `open` state and the bag should be in `open` state.

```
curl --location --request POST 'http://localhost:3000/api/v1/bags/BagT:859af678-ceef-5932-92d1-831c52ca5e54/add/PackageCsa:7a635074-c9b2-5b66-9d09-0a1909b16aeb' \
--header 'X-COREOS-REQUEST-ID: 0b8d01bf-050c-45f6-a0b1-21bd702fd6a6' \
--header 'X-COREOS-TENANT-ID: alpha' \
--header 'Accept: application/json'
```

### Ability to remove a package to a bag

```
curl --location --request POST 'http://localhost:3000/api/v1/bags/BagT:859af678-ceef-5932-92d1-831c52ca5e54/remove/PackageCsa:7a635074-c9b2-5b66-9d09-0a1909b16aeb' \
--header 'X-COREOS-REQUEST-ID: 0b8d01bf-050c-45f6-a0b1-21bd702fd6a6' \
--header 'X-COREOS-TENANT-ID: alpha' \
--header 'Accept: application/json'
```

### Ability to list all packages

```
curl --location --request GET 'http://localhost:3000/api/v1/packages' \
--header 'X-COREOS-REQUEST-ID: fc199991-da25-49b4-b918-6122026f76cc' \
--header 'X-COREOS-TENANT-ID: alpha' \
--header 'Accept: application/json'
```
### Ability to list all bags 

```
curl --location --request GET 'http://localhost:3000/api/v1/bags' \
--header 'X-COREOS-REQUEST-ID: fc199991-da25-49b4-b918-6122026f76cc' \
--header 'X-COREOS-TENANT-ID: alpha' \
--header 'Accept: application/json'
```

### Ability to view the package details with package Id

```
curl --location --request GET 'http://localhost:3000/api/v1/packages/PackageCsa:7a635074-c9b2-5b66-9d09-0a1909b16aeb' \
--header 'X-COREOS-REQUEST-ID: 0fd32090-bdfe-41b4-a20a-cd262f71a66e' \
--header 'X-COREOS-TENANT-ID: alpha' \
--header 'Accept: application/json'
```
### Ability to view the contents of a bag with the bag Id

```
curl --location --request GET 'http://localhost:3000/api/v1/bags/BagT:7a635074-c9b2-5b66-9d09-0a1909b16aeb' \
--header 'X-COREOS-REQUEST-ID: 1956bd6e-0fb2-41e9-bcf4-a79fe3be55de' \
--header 'X-COREOS-TENANT-ID: alpha' \
--header 'Accept: application/json'
```
