# sql-swagger-generator
Generate swagger file from CREATE TABLE statement.
This results in a boilerplate swaggerfile (YML) that should probably be edited to completely suit your needs.
By combining the generated swagger with the [openapi generator](https://github.com/OpenAPITools/openapi-generator), you can implement your application in a fast way.

## Swagger version
Swagen generates Swagger/OpenAPI 2.0 configuration.

## Installation
Swagen can be acquired and installed by running

```
go get -v -u github.com/toshi1127/sql-swagger-generator
```

## How to build
Or, alternatively, the binary can be manually built:

```
go build -o ~/bin/sql-swagger-generator main.go
```

## Usage

### Creating a config file
You will need to create a small config file that Swagen uses to generate YML. For example:

```
service:
  name: MyService
  host: 127.0.0.1:1234
  prefix: /v1

resources:
  table_name:
    # The written name of the resource, e.g. "product". Will be used in summaries.
    title: written name of resource
    definition:
      # This is how the resource will be named in the definitions list.
      name: ResourceDefinition
  # Continue listing all the tables you want to generate Swagger YAML for.
```

### Running Swagen
Swagger YML can then be generated using the following command:

```
sql-swagger-generator conf=./example/conf.yml -sql=./example/queries.sql -outputDir=./example/swagger.yml
```

### Example
Products table:

```sql
CREATE TABLE products1 (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    PRIMARY KEY (id)
); 
CREATE TABLE products2 (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    PRIMARY KEY (id)
); 
```

Config:

```
service:
  name: TestService
  host: 1.1.1.1:1234
  prefix: /v1

resources:
  products1:
    title: product1
    definition:
      name: Product1
  products2:
    title: product2
    definition:
      name: Product2
```

**Entity generation results:**

```models/product1.yml
title: Product1
type: object
properties:
  id:
    type: integer
    format: int64
  name:
    type: string
  price:
    type: integer
    format: int64
required:
  - name
  - price
```

```models/product2.yml
title: Product2
type: object
properties:
  id:
    type: integer
    format: int64
  name:
    type: string
  price:
    type: integer
    format: int64
required:
  - name
  - price
```

**Endpoint generation results:**
<details><summary>show:</summary><div>

```
swagger: '2.0'
info:
  title: TestService
  version: 1.0.0
host: 1.1.1.1:1234
basePath: /v1
schemes:
  - http
consumes:
  - application/json
produces:
  - application/json
paths:
  /health-check:
    get:
      operationId: HealthCheck
      description: Returns 200 if the service is healthy.
      responses:
        200:
          description: Healthy
        500:
          description: Not healthy
  /products1:
    get:
      operationId: GetProduct1s
      summary: Get product1s
      description: Returns all product1 resources.
      parameters:
        - in: query
          name: limit
          type: integer
        - in: query
          name: offset
          type: integer
      responses:
        200:
          description: List of product1 resources.
          schema:
            type: array
            items:
              $ref: './models/Product1.yml'
        500:
          description: Internal server error
    post:
      operationId: CreateProduct1
      summary: Create product1
      description: Creates a product1.
      parameters:
        - name: resource
          in: body
          required: true
          schema:
            $ref: './models/Product1.yml'
      responses:
        201:
          description: Created
          schema:
            $ref: './models/Product1.yml'
        400:
          description: Bad request
          schema:
            $ref: '#/definitions/Error'
        422:
          description: Unprocessable entity
          schema:
            $ref: '#/definitions/Error'
        500:
          description: Internal server error
  /products1/batch:
    get:
      operationId: GetProduct1sByID
      summary: Get product1s by ID
      description: Returns the product1 resources with the given IDs.
      parameters:
        - in: query
          name: ids
          type: array
          items:
            type: integer
      responses:
        200:
          description: List of product1 resources
          schema:
            type: array
            items:
              $ref: './models/Product1.yml'
        500:
          description: Internal server error
  /products1/{id}:
    get:
      operationId: GetProduct1
      summary: Get product1 by ID
      description: Returns the product1 with the given ID.
      parameters:
        - in: path
          name: id
          type: integer
          required: true
      responses:
        200:
          description: Single product1
          schema:
            $ref: './models/Product1.yml'
        404:
          description: Not found
        500:
          description: Internal server error
    patch:
      operationId: PatchProduct1
      summary: Patch product1
      description: Patches the product1 with the given ID.
      parameters:
        - name: id
          in: path
          type: integer
          required: true
        - name: patch
          in: body
          required: true
          schema:
            $ref: '#/definitions/Patch'
      responses:
        200:
          description: Success
          schema:
            $ref: './models/Product1.yml'
        400:
          description: Bad request
          schema:
            $ref: '#/definitions/Error'
        404:
          description: Not found
        422:
          description: Unprocessable entity
          schema:
            $ref: '#/definitions/Error'
        500:
          description: Internal server error
    put:
      operationId: PutProduct1
      summary: Put product1
      description: Replaces the product1 with the given ID.
      parameters:
        - name: id
          in: path
          type: integer
          required: true
        - name: resource
          in: body
          required: true
          schema:
            $ref: './models/Product1.yml'
      responses:
        200:
          description: Success
        400:
          description: Bad request
          schema:
            $ref: '#/definitions/Error'
        404:
          description: Not found
        422:
          description: Unprocessable entity
          schema:
            $ref: '#/definitions/Error'
        500:
          description: Internal server error
    delete:
      operationId: DeleteProduct1
      summary: Delete product1
      description: Deletes the product1 with the given ID.
      parameters:
        - name: id
          in: path
          type: integer
          required: true
      responses:
        200:
          description: Success
        404:
          description: Not found
        500:
          description: Internal server error
  /products2:
    get:
      operationId: GetProduct2s
      summary: Get product2s
      description: Returns all product2 resources.
      parameters:
        - in: query
          name: limit
          type: integer
        - in: query
          name: offset
          type: integer
      responses:
        200:
          description: List of product2 resources.
          schema:
            type: array
            items:
              $ref: './models/Product2.yml'
        500:
          description: Internal server error
    post:
      operationId: CreateProduct2
      summary: Create product2
      description: Creates a product2.
      parameters:
        - name: resource
          in: body
          required: true
          schema:
            $ref: './models/Product2.yml'
      responses:
        201:
          description: Created
          schema:
            $ref: './models/Product2.yml'
        400:
          description: Bad request
          schema:
            $ref: '#/definitions/Error'
        422:
          description: Unprocessable entity
          schema:
            $ref: '#/definitions/Error'
        500:
          description: Internal server error
  /products2/batch:
    get:
      operationId: GetProduct2sByID
      summary: Get product2s by ID
      description: Returns the product2 resources with the given IDs.
      parameters:
        - in: query
          name: ids
          type: array
          items:
            type: integer
      responses:
        200:
          description: List of product2 resources
          schema:
            type: array
            items:
              $ref: './models/Product2.yml'
        500:
          description: Internal server error
  /products2/{id}:
    get:
      operationId: GetProduct2
      summary: Get product2 by ID
      description: Returns the product2 with the given ID.
      parameters:
        - in: path
          name: id
          type: integer
          required: true
      responses:
        200:
          description: Single product2
          schema:
            $ref: './models/Product2.yml'
        404:
          description: Not found
        500:
          description: Internal server error
    patch:
      operationId: PatchProduct2
      summary: Patch product2
      description: Patches the product2 with the given ID.
      parameters:
        - name: id
          in: path
          type: integer
          required: true
        - name: patch
          in: body
          required: true
          schema:
            $ref: '#/definitions/Patch'
      responses:
        200:
          description: Success
          schema:
            $ref: './models/Product2.yml'
        400:
          description: Bad request
          schema:
            $ref: '#/definitions/Error'
        404:
          description: Not found
        422:
          description: Unprocessable entity
          schema:
            $ref: '#/definitions/Error'
        500:
          description: Internal server error
    put:
      operationId: PutProduct2
      summary: Put product2
      description: Replaces the product2 with the given ID.
      parameters:
        - name: id
          in: path
          type: integer
          required: true
        - name: resource
          in: body
          required: true
          schema:
            $ref: './models/Product2.yml'
      responses:
        200:
          description: Success
        400:
          description: Bad request
          schema:
            $ref: '#/definitions/Error'
        404:
          description: Not found
        422:
          description: Unprocessable entity
          schema:
            $ref: '#/definitions/Error'
        500:
          description: Internal server error
    delete:
      operationId: DeleteProduct2
      summary: Delete product2
      description: Deletes the product2 with the given ID.
      parameters:
        - name: id
          in: path
          type: integer
          required: true
      responses:
        200:
          description: Success
        404:
          description: Not found
        500:
          description: Internal server error
definitions:
  Patch:
    type: array
    description: Patch instructions
    items:
      type: object
      required:
        - op
        - path
      properties:
        op:
          type: string
          description: Operation
        path:
          type: string
          description: Path to field to operate on
        value:
          $ref: '#/definitions/AnyValue'
  AnyValue:
    description: Any type of value
  Error:
    type: object
    properties:
      code:
        type: integer
        format: int64
        x-nullable: true
      message:
        type: string
  Principal:
    type: object
    description: Security principal for validating that a user is authorized to execute certain actions
    properties:
      userId:
        type: string
      permissions:
        type: array
        items:
          type: string

```
</summary><div>