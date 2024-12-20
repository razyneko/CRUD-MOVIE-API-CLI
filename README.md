```markdown
# Movie API

A simple RESTful API built with Go (Golang) for managing a collection of movies. This project demonstrates CRUD operations (Create, Read, Update, Delete) on movie data using the Gorilla Mux router.

## Features
- Retrieve all movies
- Retrieve a specific movie by ID
- Create a new movie
- Update an existing movie by ID
- Delete a movie by ID

## Technologies Used
- **Go (Golang)**: Main programming language
- **Gorilla Mux**: HTTP router and URL matcher for building RESTful services
- **JSON**: Data format for request and response

## Getting Started

### Prerequisites
- [Go](https://golang.org/doc/install) installed on your system

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/movie-api.git
   ```
2. Navigate to the project directory:
   ```bash
   cd movie-api
   ```
3. Initialize or ensure Go modules are set up:
   ```bash
   go mod tidy
   ```

### Running the Server
1. Run the application:
   ```bash
   go run main.go
   ```
2. The server will start at `http://localhost:8000`.

### API Endpoints

#### 1. Get All Movies
- **Endpoint**: `GET /movies`
- **Description**: Retrieves all movies in the database.
- **Response**:
  ```json
  [
    {
      "id": "1",
      "isbn": "1234567890",
      "title": "The Great Adventure",
      "director": {
        "firstname": "John",
        "lastname": "Doe"
      }
    }
  ]
  ```

#### 2. Get a Single Movie
- **Endpoint**: `GET /movies/{id}`
- **Description**: Retrieves a specific movie by its ID.
- **Response**:
  ```json
  {
    "id": "1",
    "isbn": "1234567890",
    "title": "The Great Adventure",
    "director": {
      "firstname": "John",
      "lastname": "Doe"
    }
  }
  ```

#### 3. Create a Movie
- **Endpoint**: `POST /movies`
- **Description**: Adds a new movie to the database.
- **Request Body**:
  ```json
  {
    "isbn": "7894561230",
    "title": "New Movie Title",
    "director": {
      "firstname": "Director Firstname",
      "lastname": "Director Lastname"
    }
  }
  ```
- **Response**:
  ```json
  {
    "id": "generated-id",
    "isbn": "7894561230",
    "title": "New Movie Title",
    "director": {
      "firstname": "Director Firstname",
      "lastname": "Director Lastname"
    }
  }
  ```

#### 4. Update a Movie
- **Endpoint**: `PUT /movies/{id}`
- **Description**: Updates the details of an existing movie by its ID.
- **Request Body**:
  ```json
  {
    "isbn": "updated-isbn",
    "title": "Updated Title",
    "director": {
      "firstname": "Updated Firstname",
      "lastname": "Updated Lastname"
    }
  }
  ```
- **Response**:
  ```json
  {
    "id": "1",
    "isbn": "updated-isbn",
    "title": "Updated Title",
    "director": {
      "firstname": "Updated Firstname",
      "lastname": "Updated Lastname"
    }
  }
  ```

#### 5. Delete a Movie
- **Endpoint**: `DELETE /movies/{id}`
- **Description**: Deletes a movie by its ID.
- **Response**:
  ```json
  [
    // Updated list of movies
  ]
  ```


## Contribution
Contributions are welcome! Please fork this repository, make your changes, and submit a pull request.
Feel free to customize and extend this project as needed!
```
