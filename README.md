# Book Store Management API

This is a RESTful API built using Golang for managing books in a bookstore. It supports CRUD operations and uses MySQL as the database for persistent storage. The API is developed using the Gorilla Mux router for handling routes efficiently.

## Features
- **Create** a new book record with details such as title, author, and publication year.
- **Retrieve** all books available in the database.
- **Retrieve** a specific book by its unique ID.
- **Update** book details such as title, author, or publication year.
- **Delete** a book from the inventory.

## Tech Stack
- **Golang** - A statically typed, compiled programming language designed for efficiency and reliability.
- **Gorilla Mux** - A powerful router and dispatcher for handling API routes in Go.
- **MySQL** - A relational database management system for data persistence.
- **Jinzhu Inflection** - A package for handling pluralization of words, useful in REST API design.

## Project Structure
```
bookStoreManagementAPI/
├── cmd/main/main.go               # Entry point of the application
├── pkg/config/app.go              # Database configuration and connection setup
├── pkg/controllers/book-controller.go  # Business logic for book-related CRUD operations
├── pkg/models/book.go             # Book model definition with struct mapping to database
├── pkg/routes/bookStore-Route.go  # API routes definition and mapping to controllers
├── pkg/utils/utils.go             # Utility functions to support application logic
├── go.mod                         # Go module file for dependency management
├── go.sum                         # Dependency lock file
```

## Setup and Installation
### Prerequisites
- Install [Go](https://go.dev/dl/) (version 1.18 or later recommended)
- Install [MySQL](https://dev.mysql.com/downloads/) and create a database

### Steps to Run the Project
1. **Clone the repository:**
   ```sh
   git clone <repo_url>
   cd bookStoreManagementAPI
   ```
2. **Install dependencies:**
   ```sh
   go mod tidy
   ```
3. **Configure MySQL database connection:**
   Open `pkg/config/app.go` and update the connection string to match your database credentials:
   ```go
   import (
       "database/sql"
       _ "github.com/go-sql-driver/mysql"
   )

   var DB *sql.DB

   func Connect() {
       dsn := "user:password@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
       db, err := sql.Open("mysql", dsn)
       if err != nil {
           panic("Failed to connect to database")
       }
       DB = db
   }
   ```
4. **Run database migrations** (if required):
   ```sh
   go run cmd/main/main.go
   ```
5. **Start the API server:**
   ```sh
   go run cmd/main/main.go
   ```
6. The API will be running at `http://localhost:8080/`

## API Endpoints
| Method | Endpoint          | Description |
|--------|------------------|-------------|
| GET    | /books           | Get all books available in the bookstore |
| GET    | /books/:id       | Get a book by its ID |
| POST   | /books           | Add a new book to the inventory |
| PUT    | /books/:id       | Update book details |
| DELETE | /books/:id       | Remove a book from the store |

## Example API Request
### Create a New Book
#### Request:
```sh
POST /books
Content-Type: application/json

{
  "title": "The Golang Guide",
  "author": "John Doe",
  "publication_year": 2023
}
```
#### Response:
```json
{
  "id": 1,
  "title": "The Golang Guide",
  "author": "John Doe",
  "publication_year": 2023
}
```

## Future Enhancements
- Implement authentication and authorization for API security.
- Add pagination for large datasets.
- Improve error handling and logging.
- Deploy the application using Docker and Kubernetes.

## License
This project is licensed under the MIT License.

