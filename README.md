## Mini Backend

### Requirement

-   go
-   docker(optional, for dev server)
-   docker-compose(optional, for dev server)

#### How to run dev server with docker-compose

`$ docker-compose up `

#### Available Make command

-   Make test
-   Make build
-   Make clean

#### Available Endpoint

-   POST /users
-   GET /users (must login)
-   DELETE /users/{id} (must login)
-   POST /users/auth
-   POST /charts (must login)
-   GET /charts (must login)
-   POST /charts/favorite (must login)

#### POSTMAN collection

you can use postman.json file to test all of the available endpoint with postman(https://www.postman.com/)
