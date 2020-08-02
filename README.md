# Go Car API Example

Example Service to manage Cars.

Fiber is used for the REST interaction, GORM to interact with the Database.

| Endpoint                     | Description                         |
|------------------------------|-------------------------------------|
| /api/car                | Lists all carss                     |
| /api/car/1               | Returns car with id 1 if available |
| /api/car   (POST)        | Adds a new car                     |
| /api/car/1 (PUT)         | Updates car with id 1 if available |
| /api/car/1 (DELETE)      | Deletes car with id 1 if available |

## Usage

You can build this programm on your platform by yourself or simply start it with "`go run main.go.`". The application opens a socket at the port 3000.
