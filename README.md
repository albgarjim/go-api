# AIKONE API

<!-- Section for your links, references, etc. --->

[//]: # "References"
[logo]: https://via.placeholder.com/900x300/000000/FFFFFF/?text=project+logo
[shields-badge]: https://img.shields.io/badge/make%20your%20own%20badges-on%20shields.io-brightgreen.svg
[issue-tracker]: #
[contributor-one-img]: https://via.placeholder.com/150?text=profile+image
[contributor-one-link]: #
[contributor-two-img]: https://via.placeholder.com/150?text=profile+image
[contributor-two-link]: #
[contributor-three-img]: https://via.placeholder.com/150?text=profile+image
[contributor-three-link]: #
[license]: #
[sphinx]: https://www.sphinx-doc.org/en/master/
[mkdocs]: https://www.mkdocs.org/
[gitbook]: https://www.gitbook.com/
[bibtex-wikipedia]: https://en.wikipedia.org/wiki/BibTeX
[go-official-site]: https://golang.org/doc/install
[docker-install]: https://docs.docker.com/install/
[open-api2-link]: https://swagger.io/specification/v2/
[app-api-dev-docs]: https://dev.aikone.at/api-user-docs
[app-api-mantainer-docs]: https://dev.aikone.at/api-mantainer-docs
[golang-install]: https://golang.org/doc/install
[vscode-go]: https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go
[vscode-prettier]: https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode

<!-- Your project's logo --->

![Your project's logo][logo]

<!-- Your badges --->

[![shields.io badge][shields-badge]](https://shields.io)

<!-- One liner about your project --->

One short sentence about your project goes here.

## Table of Contents

- [AIKONE API](#aikone-api)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
      - [Run locally](#run-locally)
      - [Run with docker](#run-with-docker)
    - [User flow](#user-flow)
  - [Documentation](#documentation)
    - [Folder structure](#folder-structure)
      - [Filter](#filter)
      - [Handlers](#handlers)
      - [Input](#input)
      - [Output](#output)
      - [Cache](#cache)
      - [Rethinkdb](#rethinkdb)
      - [Core](#core)
      - [Middleware](#middleware)
      - [Utils](#utils)
      - [Scripts](#scripts)
    - [Development](#development)
      - [Create route](#create-route)
        - [Get route](#get-route)
        - [Post, put, patch route](#post-put-patch-route)
      - [Modify entity](#modify-entity)
      - [Add middleware](#add-middleware)
      - [Add parameters route](#add-parameters-route)
      - [Login as user](#login-as-user)
      - [Login as admin](#login-as-admin)
    - [Usage](#usage)
      - [Callable routes](#callable-routes)
  - [Contribution](#contribution)
  - [Acknowledgement](#acknowledgement)
  - [License](#license)
  - [Citation](#citation)
  - [Contact](#contact)

## Introduction

The following documentation assumes previous knowledge of:

- AIKONE infrastructure and projects; frontend, database, gitlab repositories, aws instances, technologies used, etc
- Docker containers
- General API usage (call a route, differences between get, post, patch operations, send a request with body or with data on headers, process a response...)
- Understanding of the purpose of the API; being able to answer the questions:
  - What the API has ben built for?
  - What resources are managed by it

## Getting Started

### Prerequisites

- Install golang from the reference [guide][golang-install]
- Install [docker][docker-install] from the official site
- Optional: vscode addons for golang
  - [Vscode][vscode-go] Golang addon
  - [Prettier][vscode-prettier] for code formatting

### Installation

- Open a terminal and navigate to your golang working directory

  ```sh
  cd PATH/TO/YOUR/GOLANG/WORK/DIRECTORY
  ```

- Run the command on the terminal to clone the project from the gitlab repository

   ```sh
    git clone https://gitlab.com/aikone_fashion/goggers.git
   ```

#### Run locally

1. You will find a folder named goggers, navigate inside the folder with the command

   ```sh
   cd ./goggers`
   ```

2. Run this command to install all the necessary golang packages

   ```sh
   go get -v ./...
   ```

3. Paste the following data into the .env file inside the folder and update the variables

    ```toml
    PORT=your_db_port
    GO_ENV=development
    MC_API_KEY=your_mailchimp_api_key
    MC_USER=your_mailchimp_api_user
    MC_LIST_ID=your_mailchimp_list_id
    SERVER_LOGS_PATH=logs/server_info.log
    FRONT_LOGS_PATH=logs/front_info.log
    CORS_ALLOWED_ORIGINS=your_allowed_url_1, your_allowed_url_2
    ACCESS_CONTROL_ALLOW_CREDENTIALS=your_allowed_url_1, your_allowed_url_2
    CORS_ADMIN_ORIGINS=your_allowed_url_1, your_allowed_url_2
    RDB_ENV=dev
    BASIC_AUTH_USER=your_basic_auth_user_name
    BASIC_AUTH_PASS=your_basic_auth_user_pass
    ADMIN_AUTH_USER=your_basic_auth_admin_name
    ADMIN_AUTH_PASS=your_basic_auth_admin_pass
    RDB_PORT=your_rethinkdb_port
    RDB_HOST=your_rethinkdb_host
    IMGPROXY_KEY=your_imgproxy_key
    IMGPROXY_SALT=your_imgproxy_salt
    IMGPROXY_URL=your_imgproxy_url
    FACEBOOK_CLIENT_ID=your_facebook_client_id
    FACEBOOK_CLIENT_SECRET=your_facebook_client_secret
    REDIRECT_URL=your_facebook_redirection_url
    ```

4. After completing the 3 previous steps, the project should be configured and ready to run, in order to launch it execute the following command, which will start the server and will show the logs on the terminal

    ```sh
    go run app.go
    ```

#### Run with docker

1. Create a file called `docker-compose.yml` in the same directory of the goggers folder

2. Paste the following content into the `docker-compose.yml` file:

    ```yml
    version: "3.7"
    services:

    pair-generator-api:
        build: ./pair-generator-api
        image: pair-generator-api:latest
        restart: always
        container_name: pair-generator-api
        depends_on:
        environment:
            PORT: your_db_port
            GO_ENV: development
            MC_API_KEY: your_mailchimp_api_key
            MC_USER: your_mailchimp_api_user
            MC_LIST_ID: your_mailchimp_list_id
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
            IMGPROXY_KEY: your_imgproxy_key
            IMGPROXY_SALT: your_imgproxy_salt
            IMGPROXY_URL: your_imgproxy_url
            FACEBOOK_CLIENT_ID: your_facebook_client_id
            FACEBOOK_CLIENT_SECRET: your_facebook_client_secret
            REDIRECT_URL: your_facebook_redirection_url
        volumes:
        - ./mounted/logs:/go/src/pair-generator-api/logs
        ports:
        - "8082:81"
    ```

3. Add the aws and imgproxy keys, urls, port configuration and other variables to the `docker-compose.yml` file

4. To launch the container and have the system running, execute the following command on the terminal:

    ```sh
    sudo docker-compose up --build -d
    ```

    This command will build the project, copy the resulting executable into the docker container and launch it

5. To monitor the system is running, execute the following command which will display the logs

    ```sh
    sudo docker-compose logs -t -f
    ```

### User flow

The server setup and launch is contained in the core file. The initialization is done on **core_init.go**, then on **core_main.go** the router is created and the middleware added. on **routes.go** the routes are initialized and on **run.go** the server is started.

When a route is called, the server executes the middleware functions attached to that route , examples of middleware are:

- login
- rate limiter
- security
- jwt

If all the middleware filters are executed successfully, e.g. user is logged in, the amount of requests is on normal levels, there are no weird things on the call such as timeouts or waiting times the handler is called.

The handlers are stored on the api/v1/handlers folder. The logic of the server is contained on them. They don't do anything by themselves, instead, they glue together the different parts of the api to perform the request, some tasks are:

- URL parameters extraction
- Database calls for CRUD operations
- Call to external services (mailchimp, about you API)
- Perform data validation
- Error checking
- Response formatting and return of the data

When the handler is called, it receives the data from the http requests, takes this information and uses the services in the API to perform the request. If no errors occur, the response is formatted (to json) and returned to the user. If something unexpected happens, an error is returned.

## Documentation

This point covers the inner workings of the API, here you will find information about how the code is structured

### Folder structure

The folder structure of the API is as follows:

```toml
- goggers
  - api/v1
    - filter
    - handlers
    - input
    - output
    - models
      - cache
      - rethinkdb
  - core
  - facebook
  - middleware
  - utils
```

#### Filter

-note: will be removed

#### Handlers

It contains the handlers of the API that manages the logic of the routes, there is one handler function per route and all the handlers have roughly the same format:

  ```go
    func GetIdeaProperties(_w http.ResponseWriter, _r *http.Request) {

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
        ideaProperties, err := rdb.RetrieveStylingIdeasDB(params)
        if err != nil {
            log.Error("Error retrieving idea properties from database: ", err)
            out.RespondWithError(_w, err)
            return
        }

        //In the last step we format our golang structure to json and we return the data to
        //the user
        log.Info("-------- Finish GetIdeaProperties route --------")
        out.RespondWithJSON(_w, http.StatusOK, ideaProperties)
    }
  ```

The key functionalities of a handler are: TODO: explain

- Url parameters extraction
- Body parameters extraction
- Optional Modification of variables
- Database call for CRUD (Create, Retrieve, Update, Delete) operations
- Optional modification of the data from the database
- Return of the output as JSON or error if something unexpected happens

The used packages by the handlers folder are:

- Input and output packages to have access to the entities and data extraction functions
- Cache to have access to the list of cached items
- Rethinkdb for database operations
- Filter to process the data of some items

#### Input

It contains the functions and entities to extract and validate data from requests. Each structure has its own `validate` function which determines if the struct has the right format and data. The URL is validated using a list of keywords and a map of functions using the following steps:

- The validation system contains a list of all the keywords to look for on the url, like userName, ideaType, ideaStatus...
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

#### Cache

It contains the item cache data structures, the functions to interact with it and the formatting options for the different outputs that use items. It contains the following files:

- cache.go: contains the map of items in the format of:

  ```go
  map[ITEM_ID] = ITEM_DATA_IN_RETHINKDB_FORMAT
  ```

- filter_items: contains the functions to sort the items depending of its position (top, bottom, shoes and outerwear)
- items.go: contains the functions to retrieve items from the cache, the retrieval is done in the following steps:
  - A list of ITEM_ID is passed
  - The item data is extracted from the map of items using the ids
  - The items that are not in the cache are grouped together into a list
  - The missing items are updated by calling the database using the previously made list, the data is retrieved and added to the map
  - The cache is fully cleared every 6 hours because the data from about you changes regularly

This package relies on the following packages:

- output to have access to the output structures
- utils to have access to the image url encryption functions (required for processing the raw data from rethinkdb into the format required for the response)
- rethinkdb to have access to the database in order to update the missing items

#### Rethinkdb

It contains the queries to communicate with the database and realize CRUD operations. The majority of the routes receives an internal map with the parameters required to process the information. The map has the following format:

```go
map[URL_VARIABLE_OR_INTERNAL_VARIABLE] = VARIABLE_VALUE
```

The routes modify the resources, there are 4 types of operations:

- Create: currently there is only one create operation; the one that corresponds to the register of a new user, the current ways of adding data are:
  - data preprocessor: inserts new item data
  - by hand: the data entities for ideas, collections and picks are created by hand and inserted directly on the database.

- Retrieve receives the id of the resource and the search constrains in the map shown before. Some examples are:
  - item position
  - is on sale
  - is liked/disliked
  - collections of a particular user
  - is on stock

- Update: receives the id of the resource and the data to update, the data is passed either by the map shown before or by a custom object. This route returns the code success if everything works fine

- Delete: receives the id of the resource and performs the delete operation from the database
    -process
        -realizes the single or multiple database calls
        -formats the information
        -returns it

This package relies on the following packages:

- input and output to have access to the entities
- utils to have access to the URL_encrypt and set operation functions required to format the data

#### Core

It contains the initialization config and startup of the server. Reads the .env file (if local) or docker_compose file variables (if in docker container) to configure the server. It contains the following files:

- env_load: loads environment variables
- initialize: configures the routers and adds the middleware functions
- routes_constants.go: contains the route variable naming in order to make it easier to identify the routes internally
- routes.go: contains the list of routes
- run.go: contains the final configuration of the server (cors and timeouts) and the launch

Relies on the following packages

- middleware: to have access to the middleware functions
- facebook: to set up the login functionality
- rethinkdb: to initialize the database
- cache: to initialize the cache

#### Middleware

It contains the middleware functions for the router. The implemented functionalities are:

- auth_admin: handles basic authentication using the admin username and password
- auth_user: handles basic authentication using the user username and password
- jwt: performs authentication with json web token
- login_facebook: performs login and register using the facebook API
- rate_limiter: monitors the amount of requests an ip has made and puts limits if they are exceeded
- secure: contains security configuration for the route (xss protection, nosniff, allowed hosts...)

Relies on the following packages:

- output: uses this package to have access to the response codes
- rethinkdb: is used for the login_facebook middleware to check if the user exists on our database and is logged in

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

The main folder (goggers) contains a few scripts:

- Repo_push_docker.sh: builds the docker container with the server and pushes it to the gitlab repo. Requires the version name to execute, example

```sh
bash repo_push_docker.sh 1.1.1
```

- Start_docker.sh: builds and launches the docker container of the server
- .env: contains the environment variables of the server
- dockerfile the container building

### Development

how to work with the codebase and extend it from a developer perspective with specific use cases

#### Create route

##### Get route

    -how to extract information from url
    -how to format the data
    -how to pass the data to the database
    -how to format the database response
    -how to return the data to the user

##### Post, put, patch route

    -how to extract information from body
    -how to pass this to the database
    -how to return the response or error

#### Modify entity

    -golang and entities explanation of json tags
    -entities located in files body_objects for input and response_objects for output
    -to add an entity go to the site https://mholt.github.io/json-to-go/ and paste the json of the entity, on the right will appear the golang code for the structure
    -how to remove fields: remove the fields for the entity (either by erasing the line or commenting it out), go to the functions where the entity is used and remove all the lines of code where the removed field is used
    -how to use:
        -paste the golang structure into the input or output file depending on the intent of the structure
        -copy and paste the code into the file renaming the fiends required
        -declare the structure and initialize where it's required

#### Add middleware

    -middleware routes are located on the middleware folder
    -middleware routes have the following structure: MIDDLEWARE STRUCTURE CODE
    -how to create a new one
        -use the template
        -access the response and writer resources to get information to the request and process it as required
    -how to use
        -in the router initialization (core folder) add the middleware to the router with the Use function; e.g. EXAMPLE CODE ADD MIDDLEWARE TO ROUTER

#### Add parameters route

    -route template
    -where to add
        -declare parameter name in url_params
        -add to list of variables to check
        -implement function to validate
        -when want to use, access the params struct with the field name
    -how to call route
    -how to work with parameters in url and in body request
        -for url use extract url params
        -for body us extract body params
        -use the resulting variable as a dictionary
    -how to integrate database, middleware and entitites

#### Login as user

    -how is structured
        -third party library
        -check if user registered
        -retrieve user data
        -add to local storage
        -return cookie with session data
        -in next request extract cookie and username with it to validate
    -what the third party library does
    -how user is registered
    -how user is logged in
    -how session is persisted
    -how authentication is checked

#### Login as admin

    -how is structured
        -third party library
        -check if user registered
        -retrieve user data
        -add to local storage
        -return cookie with session data
        -in next request extract cookie and username with it to validate
    -what the third party library does
    -how user is registered
    -how user is logged in
    -how session is persisted
    -how authentication is checked


### Usage


#### Callable routes

Those are the core resources of the REST API and the routes used to modify them

| Resource    | Route                             | Method | Information                               | Params                |
| :---------- | :-------------------------------- | :----: | :---------------------------------------- | :-------------------- |
| collections | `/api/v1/GtS2Fa`                  |  GET   | Returns the collection names.             |                       |
| collections | `/api/v1/5JUSUk`                  |  GET   | Returns all items collections.            | collection_id         |
| collections | `/api/v1/5JUSUk/l9ew8v`           |  GET   | Items collections overview.               | collection_id         |
| collections | `/api/v1/5JUSUk/3fj48f`           |  GET   | Returns three items fron new collections. |                       |
| ideas       | `/api/v1/AcD9pq`                  |  GET   | Returns styng ideas of a user.            |                       |
| ideas       | `/api/v1/AcD9pq/8g73jg`           |  GET   | Returns liked styling ideas.              |                       |
| ideas       | `/api/v1/AcD9pq/j3f84f`           |  GET   | Returns disliked styling ideas.           |                       |
| ideas       | `/api/v1/AcD9pq/MtcZ29`           |  GET   | Returns styling ideas properties.         | idea                  |
| ideas       | `/api/v1/AcD9pq`                  | PATCH  | Updates idea to liked state.              | ideaName & ideaStatus |
| ideas       | `/api/v1/AcD9pq/PgwS23`           | PATCH  | Updates idea to seen state.               | ideaName              |
| ideas       | `/api/v1/AcD9pq/gmr9qj`           |  GET   | Returns items of ideas.                   |                       |
| picks       | `/api/v1/nv2sdo/82fj93`           |  GET   | Returns new picks.                        |                       |
| picks       | `/api/v1/nv2sdo/zw4w84`           |  GET   | Returns old picks.                        |                       |
| wishlist    | `/api/v1/BiT7gF/{ITEM_ID}`        | DELETE | Deletes item from wishlist.               |                       |
| wishlist    | `/api/v1/BiT7gF/{ITEM_ID}`        |  GET   | Returns wishlist.                         |                       |
| Wishlist    | `/api/v1/BiT7gF/j3f84f`           |  GET   | Returns disliked items.                   |                       |
| wishlist    | `/api/v1/BiT7gF`                  | PATCH  | Adds item to wishlist.                    | itemid struct         |
| items       | `/api/v1/5JUSUk/{ITEM_ID}`        | DELETE | Deletes item from wishlist.               |                       |
| items       | `/api/v1/5JUSUk/gmr9qj/{ITEM_ID}` |  GET   | Returns single item.                      |                       |
| user        | `/api/v1/Pq3Gm9`                  | PATCH  | Updates userData.                         |                       |
| user        | `/api/v1/Pq3Gm9/j484hs`           |  GET   | Returns  user preferences.                |                       |
| login       | `/api/v1/G35df2/Pdfe32`           |  POST  | post login.                               | code & state          |
| mailchimp   | `/api/v1/FIYnCI/t3Sf9a`           |  POST  | mailchimp subscription.                   |                       |
| survey      | `/api/v1/nZiejS`                  |  POST  | post survey.                              |                       |
| wardrobe    | `/api/v1/jfm2gw`                  |  GET   | Returns wardrobe.                         |                       |
| stylist     | `/api/v1/ZQdG8p`                  |  GET   | Returns stylist data.                     |                       |
| creator     | `/api/v1/abc`                     |  GET   | Returns arrivals from about you.          | aylimit & aypage      |

## Contribution

Describe how your team members and contributors can contribute to this project. If you have specific style guidelines which should be obeyed be sure to mention it here as well. Also, include a link to your issues tracker if you have one e.g. [github.com/project/issues][issue-tracker]

## Acknowledgement

| [![Max Mustermann][contributor-one-img]<br>Max Mustermann][contributor-one-link] | [![John Doe][contributor-two-img]<br>John Doe][contributor-two-link] | [![Zhang San][contributor-three-img]<br>Zhang San][contributor-three-link] |
| :------------------------------------------------------------------------------: | :------------------------------------------------------------------: | :------------------------------------------------------------------------: |


Now that you have included yourself, your contributors and/or team members above also make sure to list some further inspirational resources and references for your project:

- [Some nice article you have read](#)
- [An awesome GitHub repository that you have made use of](#)
- [A StackOverflow question that has helped you out](#)

## License

Provide your license information here and link to your [LICENSE.md][license] file.

## Citation

Explain how other people can cite your project.

Make sure you provide a [BibTeX formatted reference][bibtex-wikipedia] e.g.:

```
@misc{yourlabel,
 author    = {Name of the authors},
 title     = {Your project's title},
 publisher = {Your publisher},
 year      = {2019},
}
```

## Contact

Put here your contact details so contributors and team members can reach out to you if they have open questions (or want to compliment you).
