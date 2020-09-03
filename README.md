# Golang API

<!-- Section for your links, references, etc. --->

[//]: # "References"
[logo]: https://via.placeholder.com/900x300/000000/FFFFFF/?text=project+logo
[go-official-site]: https://golang.org/doc/install
[docker-install]: https://docs.docker.com/install/
[golang-install]: https://golang.org/doc/install
[vscode-go]: https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go
[vscode-prettier]: https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode


RESTful API written in Golang using RethinkDB as database

## Table of Contents

- [Golang API](#golang-api)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
  - [Installation](#installation)
    - [Run locally](#run-local)
    - [Run with docker](#run-docker)
  - [Documentation](#documentation)
    - [Folder structure](#folder-structure)
      - [Handlers](#handlers)
      - [Input](#input)
      - [Output](#output)
      - [Rethinkdb](#rethinkdb)
      - [Core](#core)
      - [Middleware](#middleware)
      - [Utils](#utils)
      - [Scripts](#scripts)
  - [Contribution](#contribution)
  - [License](#license)

## Getting Started

- Install golang from the reference [guide][golang-install]
- Install [docker][docker-install] from the official site
- Optional: vscode addons for golang
  - [Vscode][vscode-go] Golang addon
  - [Prettier][vscode-prettier] for code formatting

## Installation

- Open a terminal and navigate to your golang working directory

  ```sh
  cd PATH/TO/YOUR/GOLANG/WORK/DIRECTORY
  ```

- Run the command on the terminal to clone the project from the gitlab repository

   ```sh
    git clone https://gitlab.com/albgarjim/go-api.git
   ```

### Run locally

1. You will find a folder named go-api, navigate inside the folder with the command

   ```sh
   cd ./go-api`
   ```

2. Run this command to install all the necessary golang packages

   ```sh
   go get -v ./...
   ```

3. After completing the 3 previous steps, the project should be configured and ready to run, in order to launch it execute the following command, which will start the server and will show the logs on the terminal

    ```sh
    go run app.go
    ```

### Run with docker

1. Create a file called `docker-compose.yml` in the same directory of the go-api folder

2. Paste the following content into the `docker-compose.yml` file:

    ```yml
    version: "3.7"
    services:

    go-api-api:
        build: ./go-api-api
        image: go-api-api:latest
        restart: always
        container_name: go-api-api
        depends_on:
        environment:
            PORT: your_port
            GO_ENV: development
            SERVER_LOGS_PATH: logs/server_info.log
            FRONT_LOGS_PATH: logs/front_info.log
            CORS_ALLOWED_ORIGINS: your_allowed_url_1, your_allowed_url_2
            ACCESS_CONTROL_ALLOW_CREDENTIALS: your_allowed_url_1, your_allowed_url_2
            CORS_ADMIN_ORIGINS: your_allowed_url_1, your_allowed_url_2
            RDB_ENV: dev
            BASIC_AUTH_USER: your_basic_auth_user_name
            BASIC_AUTH_PASS: your_basic_auth_user_pass
            ADMIN_AUTH_USER: your_basic_auth_admin_name
            ADMIN_AUTH_PASS: your_basic_auth_admin_pass
            RDB_PORT: your_rethinkdb_port
            RDB_HOST: your_rethinkdb_host
        volumes:
        - ./mounted/logs:/go/src/go-api-api/logs
        ports:
        - "8082:81"
    ```

3. Add the variables to the `docker-compose.yml` file

4. To launch the container and have the system running, execute the following command on the terminal:

    ```sh
    sudo docker-compose up --build -d
    ```

    This command will build the project, copy the resulting executable into the docker container and launch it

5. To monitor the system is running, execute the following command which will display the logs

    ```sh
    sudo docker-compose logs -t -f
    ```

## Documentation

This point covers the inner workings of the API, here you will find information about how the code is structured

The server setup and launch is contained in the core file. The initialization is done on **core_init.go**, then on **core_main.go** the router is created and the middleware added. on **routes.go** the routes are initialized and on **run.go** the server is started.

When a route is called, the server executes the middleware functions attached to that route , examples of middleware are:

- login
- rate limiter
- security
- jwt

The handlers are stored on the api/v1/handlers folder. The logic of the server is contained on them. They don't do anything by themselves, instead, they glue together the different parts of the api to perform the request, some tasks are:

- URL parameters extraction
- Database calls for CRUD operations
- Perform data validation
- Error checking
- Response formatting and return of the data

When the handler is called, it receives the data from the http requests, takes this information and uses the services in the API to perform the request. If no errors occur, the response is formatted (to json) and returned to the user. If something unexpected happens, an error is returned.


### Folder structure

The folder structure of the API is as follows:

```toml
- go-api
  - api/v1
    - filter
    - handlers
    - input
    - output
    - models
      - rethinkdb
  - core
  - middleware
  - utils
```

#### Handlers

It contains the handlers of the API that manages the logic of the routes, there is one handler function per route and all the handlers have roughly the same format:

  ```go
    func GetItem(_w http.ResponseWriter, _r *http.Request) {

        //Variables used on the handler, params stores a dictionary that contains the data
        //required to execute the request, this is done to pass data to the database and other
        //parts of the api with ease
        var err error
        var params *inp.URLParams

        //Here the extraction of the parameters from the url is performed, we pass our empty dictionary and we fill it with the request data
        if params, err = inp.ExtractParameters(_r); err != nil {
            log.Error("Incorrect url parameters: ", err)
            out.RespondWithError(_w, err)
            return
        }

        //Here the database call is performed, we pass our dictionary and we retrieve a struct
        // with the data to return to the user
        ideaProperties, err := rdb.RetrieveItem(params)
        if err != nil {
            log.Error("Error retrieving item from database: ", err)
            out.RespondWithError(_w, err)
            return
        }

        //In the last step we format our golang structure to json and we return the data to
        //the user
        log.Info("-------- Finish GetItem route --------")
        out.RespondWithJSON(_w, http.StatusOK, ideaProperties)
    }
  ```


- Url parameters extraction
- Body parameters extraction
- Optional Modification of variables
- Database call for CRUD (Create, Retrieve, Update, Delete) operations
- Return of the output as JSON or error if something unexpected happens

The used packages by the handlers folder are:

- Input and output packages to have access to the entities and data extraction functions
- Rethinkdb for database operations
- Filter to process the data of some items

#### Input

It contains the functions and entities to extract and validate data from requests. Each structure has its own `validate` function which determines if the struct has the right format and data. The URL is validated using a list of keywords and a map of functions using the following steps:

- The validation system contains a list of all the keywords to look for on the url
- Each keyword has associated a function that performs the validation of the variable. This association is done with a map. e.g.

  ```go
  map[keyword] = function_validate_keyword(keyword_value)
  ```

- When an URL is processed, all the keywords and value of each keyword are extracted and for each keyword, the respective validation function is called, if the data is correct is added to a dictionary with the following format

   ```go
   map[keyword] = keyword_value
   ```

- When the validation process is finished, the data is returned to the user

The used packages by the Input folder are:

- Utils for the user login and id extraction from the request context

#### Output

It contains the functions, entities and error codes needed to work with the output data that is returned to the user. The folder contains the following data:

- codes.go: contains the error codes together with the message to return
- response_objects: contains the structs to represent entities
- responses: contains the functions required to format the data

This package does not rely on any other packages.

#### Rethinkdb

It contains the queries to communicate with the database and realize CRUD operations. The majority of the routes receives an internal map with the parameters required to process the information. The map has the following format:

```go
map[URL_VARIABLE_OR_INTERNAL_VARIABLE] = VARIABLE_VALUE
```

#### Core

It contains the initialization config and startup of the server. Reads the .env file (if local) or docker_compose file variables (if in docker container) to configure the server. It contains the following files:

- env_load: loads environment variables
- initialize: configures the routers and adds the middleware functions
- routes_constants.go: contains the route variable naming in order to make it easier to identify the routes internally
- routes.go: contains the list of routes
- run.go: contains the final configuration of the server (cors and timeouts) and the launch

Relies on the following packages

- middleware: to have access to the middleware functions
- rethinkdb: to initialize the database

#### Middleware

It contains the middleware functions for the router. The implemented functionalities are:

- auth_admin: handles basic authentication using the admin username and password
- auth_user: handles basic authentication using the user username and password
- jwt: performs authentication with json web token
- rate_limiter: monitors the amount of requests an ip has made and puts limits if they are exceeded
- secure: contains security configuration for the route (xss protection, nosniff, allowed hosts...)

Relies on the following packages:

- output: uses this package to have access to the response codes

#### Utils

It contains various helper functions:

- Context: contains the structures and variables to store the user session
- Encrypt_URL: contains the functions required to encrypt the Imgproxy URL in various formats
- Logger: logger configuration and setup
- Set_difference: set operation to do set A - B
- Set_intersection: set operation to obtain the common elements between set A and B
- Unique_list: convert a list to a set
- URL_fetch: perform get and post operations to external urls

This package is self-contained, doesn't rely on any other package of the API

#### Scripts

The main folder (go-api) contains a few scripts:

- Repo_push_docker.sh: builds the docker container with the server and pushes it to the gitlab repo. Requires the version name to execute, example

```sh
bash repo_push_docker.sh 1.1.1
```

- Start_docker.sh: builds and launches the docker container of the server
- .env: contains the environment variables of the server
- dockerfile the container building

## Contribution

Mantainer [Alberto Garcia][alberto-mail]

## License

MIT license.
