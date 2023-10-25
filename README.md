# Hub-Connect Documentation
### Requirements: Implement API with golang and postgres db

**Hub**, **Team**, and **Users** are in the order of hierarchy. A short description is as below:

- A Hub is an entity that associates Team depending on their geological location.
- A Team is an entity that associates Users based on their types.
- A User is an entity that holds the information of the human users.

**Todo**
- [ ] Implement a Create for each _hub, team, and user_.
- [ ] Implement a Search which will return _team and hub information_.
- [ ] Implement a Join for user into team, and team into hub (for simplicity: one user belongs to one team, one team belongs to one hub).
- [ ] Write the test cases
- [ ] Provide a SQL script which creates tables needed for the API.
- [ ] Good to use `docker/docker-compose` for local development setup(not mandatory)
- [ ] Good to provide the solution with security concern

The Hub Connect API allows you to manage hubs, teams, and users within the Hub Connect application. This documentation provides information about the available endpoints, request and response formats, and their functionalities.


## Running the Application
- Ensure that `docker & docker compose` compose are installed on your workstation.
- Some of the variables given below are set up as environment variables, so you may simply change them to suit your needs, modifying the .env file
```
PORTS = 80
DB_PORT = 5432
SWAGGER_DOMAIN = 127.0.0.1
```
- `PORT` is the port on which the web application will execute
- `DB_PORT` is the port on which the database application will execute
- `SWAGGER_DOMAIN` is the swagger domain that will be used for API testing.
To run the hub-connect, execute the following command:
```sh
git clone https://github.com/pdhoang91/hub-connect.git
```
- Please stand at project root folder and run:
```sh
`docker-compose up` or `docker-compose --env-file ./.env up`
```
# Testing API Endpoints with Swagger and Postman
To increase API security, to access system resources as well as our APIs. You need to add Basic token below to the header:
`Basic cGRob2FuZzkxQGdtYWlsLmNvbTpjdWJpY2FzYQ==`

##### Using Swagger
Follow these steps to test API endpoints using Swagger:
Swagger documentation may be found at `http://127.0.0.1:{PORT}/swagger/index.html`.

#####  Using Postman
To test API endpoints using Postman, follow these steps:

1. This API allows you to create a new User: POST /v1/users
```shell
curl --location 'http://localhost:80/v1/users' \
--header 'Content-Type: application/json' \
--header 'Authorization: Basic cGRob2FuZzkxQGdtYWlsLmNvbTpjdWJpY2FzYQ==' \
--data-raw '{"name": "User 1", "email": "user1@example.com"}'
```

2. This API returns User information: GET /v1/users/:userID
```sh
curl --location 'http://localhost:80/v1/users/1' \
--header 'Authorization: Basic cGRob2FuZzkxQGdtYWlsLmNvbTpjdWJpY2FzYQ=='
```
3. This API allows you to assign a User to a Team: GET /v1/users/:userID/teams/:teamID
```sh
curl --location 'http://localhost:80/v1/users/1/teams/1' \
--header 'Authorization: Basic cGRob2FuZzkxQGdtYWlsLmNvbTpjdWJpY2FzYQ=='
```

4.  This API allows you to create a new Team: POST /v1/teams
```sh
curl --location 'http://localhost:80/v1/teams' \
--header 'Content-Type: application/json' \
--header 'Authorization: Basic cGRob2FuZzkxQGdtYWlsLmNvbTpjdWJpY2FzYQ==' \
--data '{"name": "Team 1", "type": "Type A"}'
```
5.  This API allows you to search for Teams by name: POST /v1/teams/search
```sh
curl --location 'http://localhost:80/v1/teams/search' \
--header 'Content-Type: application/json' \
--header 'Authorization: Basic cGRob2FuZzkxQGdtYWlsLmNvbTpjdWJpY2FzYQ==' \
--data '{"search_name": ""}'
```
6.  This API returns Team information: GET /v1/teams/:teamID
```sh
curl --location 'http://localhost:80/v1/teams/1' \
--header 'Authorization: Basic cGRob2FuZzkxQGdtYWlsLmNvbTpjdWJpY2FzYQ=='
```
7. This API allows a Team to connect to a Hub: GET /v1/teams/:teamID/hubs/:hubID
```sh
curl --location 'http://localhost:80/v1/teams/1/hubs/1' \
--header 'Authorization: Basic cGRob2FuZzkxQGdtYWlsLmNvbTpjdWJpY2FzYQ=='
```

8. This API allows you to create a new Hub: POST /v1/hubs
```sh
curl --location 'http://localhost:80/v1/hubs' \
--header 'Content-Type: application/json' \
--header 'Authorization: Basic cGRob2FuZzkxQGdtYWlsLmNvbTpjdWJpY2FzYQ==' \
--data '{"name": "Hub 1", "location": "Location 1"}'
```
9. This API allows you to search for Hubs by name: POST /v1/hubs/search
```sh
curl --location 'http://localhost:80/v1/hubs/search' \
--header 'Content-Type: application/json' \
--header 'Authorization: Basic cGRob2FuZzkxQGdtYWlsLmNvbTpjdWJpY2FzYQ==' \
--data '{"search_name": ""}'
```
10. This API returns Hub information: GET /v1/hubs/:hubID
```sh
curl --location 'http://localhost:80/v1/hubs/1' \
--header 'Authorization: Basic cGRob2FuZzkxQGdtYWlsLmNvbTpjdWJpY2FzYQ=='
```

# Run Unittest
- Please stand at project root folder and run `go test ./...` or `go test --cover ./...`