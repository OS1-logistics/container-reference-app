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

### OpenAPI

The OpenAPI spec is generated from the `api/openapi-spec/v1/api__v1_openapi.yaml` file. To update the spec, run:

```
make openapi-codegen
```

