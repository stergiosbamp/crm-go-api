# A pluggable, simple and fast API for your CRM. ⚡

## 🛠 Installation

### Environment Variables 

Before running the app, you must first need to setup the environment variables. 
It is provided a `.env.example` file to indicate which variables must be set. 

The file holds some pre-defined values that work with the current setup of Docker. If you wish to change them, do not forget to tweak them in the Docker files accordingly.


* For production use:
    ```bash
    $ cp .env.example .env
    ```

### 🚀 Run the Application

The whole application and its dependencies are containarized, utilizing Docker. 🐋

**Steps**

Start the application's infrastructure services using Docker Compose.

```bash
$ cd docker/
$ docker compose build
$ docker compose up -d
```

**Dependencies**

- **MySQL** database.
- **Swagger** UI application for API documentation and specification.

**Test it**

```bash
$ curl localhost:8080  # Ping health check
```
```
It should return:

A pluggable, simple and fast CRM API
```

## 💡 Features of the API

This API provides a comprehensive set of features for a CRM system, and implemented business logic for the management of its entities.

### 📖 OpenAPI Specification
- The definition of the API is described in detail in the [API documentation](swagger/openapi.yaml) following the OpenAPI Specification 3.0.
- Explore the API and its' capabilities at: `http://localhost:8512/swagger/`

### 🔒 Authentication

- API supports authentication using JSON Web Tokens (JWT).
- Protect your endpoints and control access to your CRM data.
- Generate tokens for your users and invalidate them using server-side blacklisting.

### 🌐 Entities

The API supports full CRUD (Create, Read, Update, Delete) operations for the following entities:

- **Customers**: Manage customer data efficiently.
- **Contacts**: Keep track of your contacts and interactions.
- **Addresses**: Handle address information seamlessly.
