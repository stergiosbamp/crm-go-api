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

# {TODO this runs by default the override file which is for dev mode}
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
