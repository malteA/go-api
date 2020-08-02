# Go Car API Example

Example Service to manage Cars.

Fiber is used for the REST interaction, GORM to interact with the Database.

API Documentation is done with the help of SWAGGER

| Endpoint                     | Description                         |
|------------------------------|-------------------------------------|
| /api/car                | Lists all carss                     |
| /api/car/1               | Returns car with id 1 if available |
| /api/car   (POST)        | Adds a new car                     |
| /api/car/1 (PUT)         | Updates car with id 1 if available |
| /api/car/1 (DELETE)      | Deletes car with id 1 if available |

## Usage

1. Install swag `go get -u github.com/swaggo/swag/cmd/swag`
2. Add GOPATH to yout PATH to be able to run `swag`

To update the swagger documentation run `swag init`.

You can build this programm on your platform by yourself or simply start it with "`go run main.go.`". The application opens a socket at the port 3000.

Swagger will be accessible at [localhost:3000/swagger](localhost:3000/swagger)
