# A pluggable, simple and fast API for your CRM.

## Installation

### Setup environment variables

Before running the app, you must first need to setup the environment variables. 
It is provided a `.env.example` file under the `api/` directory, to indicate which variables must be set. 

The file holds some pre-defined values that work with the current setup of Docker. If you wish to change them, do not forget to tweak them in the Docker files accordingly.


* For production use:
    ```bash
    $ cp api/.env.example api/.env
    ```

* For development use:
    ```bash
    $ cp api/.env.example api/.env.dev
    ```

### Run the app

To run the up we simply use the power of Docker and Docker Compose v2.

**Production**

To spin up the app for production use:

```bash
$ cd docker/
$ docker compose build
$ docker compose up -d
```

The above command does the following:

1. Spins up a MySQL database
2. Builds and runs the Go application

Verify it is running by making a request

```bash
$ curl localhost:8080
```

It should return:

```text
A pluggable, simple and fast CRM API
```

**Development**

There are also provided a `docker-compose.dev.yml` with a `Dockerfile.dev` for use of local development.

```bash
$ cd docker/
$ docker compose -f docker-compose.yml -f docker-compose.dev.yml build
$ docker compose -f docker-compose.yml -f docker-compose.dev.yml up -d
```

## API

This API provides a simple and easy-to-use management of customers. 

### Usage
The usage of API is described in detail in the [API documentation](swagger/openapi.yaml) following the OpenAPI Specification 3.0.

For this purpose, you can explore the API and its' capabilities at:

`http://localhost:8512/swagger/`

where an instance of the Swagger UI application is provided through Docker. The Swagger UI provides a user-friendly interface for exploring and interacting with your API, allowing you to easily test API requests and responses, view documentation, and more.

### Info
The API currently supports the management of:

1. Customers
2. Contacts (of Customers)
3. Addresses (of Customers or Contacts)

However, this API implements the following business logic for the above entities, as every CRM is supposed to:

* A Customer can have multiple Contacts and Addresses
* Only 3 types of Addresses are supported: "legal" or "branch" for Customers or "contact" for Contacts.
* A Customer can have only one "legal" Address. Only then it is considered as active.
* A Customer can have multiple "branch" Addresses.
* A Contact may or may not have an Address.
