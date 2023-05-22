# protoc-gen-oas [![Go Reference](https://img.shields.io/badge/go-pkg-00ADD8)](https://pkg.go.dev/github.com/ogen-go/protoc-gen-oas#section-documentation) [![codecov](https://img.shields.io/codecov/c/github/ogen-go/protoc-gen-oas?label=cover)](https://codecov.io/gh/ogen-go/protoc-gen-oas) [![experimental](https://img.shields.io/badge/-experimental-blueviolet)](https://go-faster.org/docs/projects/status#experimental)

Work in progress.

`protoc-gen-oas` is protoc plugin for generate OpenAPI v3.x.x from proto files.

# Install

```console
go install github.com/ogen-go/protoc-gen-oas/cmd/protoc-gen-oas
```

# Usage

```console
protoc --oas_out=. service.proto
```

# Features

- support [API annotations](https://github.com/googleapis/googleapis/blob/master/google/api/annotations.proto) in methods
- support [field behavior](https://github.com/googleapis/googleapis/blob/master/google/api/field_behavior.proto) in message field description

# Generate OpenAPI

## Path param

```protobuf title="service.proto"
syntax = "proto3";

package service.v1;

option go_package = "service/v1;service";

import "google/api/annotations.proto";

service Service {
  rpc GetItem(GetItemRequest) returns (Item) {
    option (google.api.http) = {
      get: "/api/v1/items/{id}"
    };
  }
}

message GetItemRequest {
  string id = 1;
}

message Item {
  string id = 1;
  string name = 2;
}
```

```yaml title="openapi.yaml"
openapi: 3.1.0
info:
    title: Generated OpenAPI specification from proto file
    version: v1.0.0
paths:
    /api/v1/items/{id}:
        get:
            operationId: getItem
            parameters:
                -   name: id
                    in: path
                    schema:
                        type: string
            responses:
                "200":
                    description: service.v1.Service.GetItem response
                    content:
                        application/json:
                            schema:
                                $ref: '#/components/schemas/Item'
components:
    schemas:
        Item:
            type: object
            properties:
                id:
                    type: string
                name:
                    type: string
```

## Query param

```protobuf title="service.proto"

```

```yaml title="openapi.yaml"

```

## Mark field as required

```protobuf title="service.proto"

```

```yaml title="openapi.yaml"

```

## Mark message/method/field as deprecated
