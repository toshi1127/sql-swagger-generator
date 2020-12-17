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
sql-swagger-generator conf=./example/conf.yml -sql=./example/queries.sql -outputDir=./example/swagger/swagger.yml
```

### Example

Products table:

```sql
CREATE TABLE products (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users(
    user_id varchar(255) PRIMARY KEY,
    nick_name varchar(255) NOT NULL,
    profile_image_uri varchar(255),
    email varchar(255) NOT NULL,
    description varchar(255),
    social_link varchar(255),
    gender enum('male', 'female', 'other'),
    identify_status varchar(255),
    customer_id varchar(255),
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    deleted_at timestamp
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

Config:

```
service:
  name: TestService
  host: 1.1.1.1:1234
  prefix: /v1

resources:
  users:
    title: user
    definition:
      name: User
  products:
    title: product
    definition:
      name: Product
```

**Entity generation results:**

```models/user.yml
title: User
type: object
properties:
  user_id:
    type: string
    x-nullable: true
  nick_name:
    type: string
  profile_image_uri:
    type: string
    x-nullable: true
  email:
    type: string
  description:
    type: string
    x-nullable: true
  social_link:
    type: string
    x-nullable: true
  gender:
    type: string
    enum:
      - male
      - female
      - other
    x-nullable: true
  identify_status:
    type: string
    x-nullable: true
  customer_id:
    type: string
    x-nullable: true
  created_at:
    type: string
    format: date-time
  updated_at:
    type: string
    format: date-time
  deleted_at:
    type: string
    format: date-time
    x-nullable: true
required:
  - nick_name
  - email
  - created_at
  - updated_at
```

```models/product.yml
title: Product
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
  /users:
    get:
      operationId: GetUsers
      summary: Get users
      description: Returns all user resources.
      parameters:
        - in: query
          name: limit
          type: integer
        - in: query
          name: offset
          type: integer
      responses:
        200:
          description: List of user resources.
          schema:
            type: array
            items:
              $ref: './models/User.yml'
        500:
          description: Internal server error
    post:
      operationId: CreateUser
      summary: Create user
      description: Creates a user.
      parameters:
        - name: resource
          in: body
          required: true
          schema:
            $ref: './models/User.yml'
      responses:
        201:
          description: Created
          schema:
            $ref: './models/User.yml'
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
  /users/batch:
    get:
      operationId: GetUsersByID
      summary: Get users by ID
      description: Returns the user resources with the given IDs.
      parameters:
        - in: query
          name: ids
          type: array
          items:
            type: integer
      responses:
        200:
          description: List of user resources
          schema:
            type: array
            items:
              $ref: './models/User.yml'
        500:
          description: Internal server error
  /users/{id}:
    get:
      operationId: GetUser
      summary: Get user by ID
      description: Returns the user with the given ID.
      parameters:
        - in: path
          name: id
          type: integer
          required: true
      responses:
        200:
          description: Single user
          schema:
            $ref: './models/User.yml'
        404:
          description: Not found
        500:
          description: Internal server error
    patch:
      operationId: PatchUser
      summary: Patch user
      description: Patches the user with the given ID.
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
            $ref: './models/User.yml'
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
      operationId: PutUser
      summary: Put user
      description: Replaces the user with the given ID.
      parameters:
        - name: id
          in: path
          type: integer
          required: true
        - name: resource
          in: body
          required: true
          schema:
            $ref: './models/User.yml'
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
      operationId: DeleteUser
      summary: Delete user
      description: Deletes the user with the given ID.
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
  /products:
    get:
      operationId: GetProducts
      summary: Get products
      description: Returns all product resources.
      parameters:
        - in: query
          name: limit
          type: integer
        - in: query
          name: offset
          type: integer
      responses:
        200:
          description: List of product resources.
          schema:
            type: array
            items:
              $ref: './models/Product.yml'
        500:
          description: Internal server error
    post:
      operationId: CreateProduct
      summary: Create product
      description: Creates a product.
      parameters:
        - name: resource
          in: body
          required: true
          schema:
            $ref: './models/Product.yml'
      responses:
        201:
          description: Created
          schema:
            $ref: './models/Product.yml'
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
  /products/batch:
    get:
      operationId: GetProductsByID
      summary: Get products by ID
      description: Returns the product resources with the given IDs.
      parameters:
        - in: query
          name: ids
          type: array
          items:
            type: integer
      responses:
        200:
          description: List of product resources
          schema:
            type: array
            items:
              $ref: './models/Product.yml'
        500:
          description: Internal server error
  /products/{id}:
    get:
      operationId: GetProduct
      summary: Get product by ID
      description: Returns the product with the given ID.
      parameters:
        - in: path
          name: id
          type: integer
          required: true
      responses:
        200:
          description: Single product
          schema:
            $ref: './models/Product.yml'
        404:
          description: Not found
        500:
          description: Internal server error
    patch:
      operationId: PatchProduct
      summary: Patch product
      description: Patches the product with the given ID.
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
            $ref: './models/Product.yml'
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
      operationId: PutProduct
      summary: Put product
      description: Replaces the product with the given ID.
      parameters:
        - name: id
          in: path
          type: integer
          required: true
        - name: resource
          in: body
          required: true
          schema:
            $ref: './models/Product.yml'
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
      operationId: DeleteProduct
      summary: Delete product
      description: Deletes the product with the given ID.
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

<div></details>
