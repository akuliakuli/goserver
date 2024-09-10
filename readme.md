# Simple HTTP Server with Go

This project is a basic HTTP server implemented in Go that handles `GET` and `POST` requests for managing a list of people. It uses the standard library's `net/http` package to serve requests and `encoding/json` for JSON serialization and deserialization.

## Endpoints

### GET `/people`
This endpoint retrieves a list of people.

#### Request

- **Method**: `GET`
- **URL**: `/people`
  
#### Response

- **Status**: `200 OK`
- **Body**:
  ```json
  [
    {
      "Name": "John Doe",
      "Age": 30
    },
    {
      "Name": "Jane Doe",
      "Age": 25
    }
  ]
  ```

### POST `/people`
This endpoint allows you to add a new person to the list.

#### Request

- **Method**: `POST`
- **URL**: `/people`
- **Body**:
  ```json
  {
    "Name": "John Doe",
    "Age": 30
  }
  ```

#### Response

- **Status**: `200 OK`
- **Body**:
  ```text
  post new person: '%v {Name: John Doe, Age: 30}'
  ```

## Running the Project

1. **Install Go**: If you don't have Go installed, [install it here](https://golang.org/doc/install).
   
2. **Clone the Repository**:
   ```bash
   git clone https://github.com/your-username/go-http-server.git
   ```

3. **Navigate to the Project Directory**:
   ```bash
   cd go-http-server
   ```

4. **Run the Server**:
   ```bash
   go run main.go
   ```

5. **Server will start on `localhost:8080`**:
   You can access the server at `http://localhost:8080/people`.

## Example Usage

### Fetch all people:
```bash
curl http://localhost:8080/people
```

### Add a new person:
```bash
curl -X POST http://localhost:8080/people -d '{"Name": "John Doe", "Age": 30}' -H "Content-Type: application/json"
```

## Dependencies
No external dependencies are used in this project.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
