# Movie API

A simple REST API for managing a collection of movies. The API allows users to perform CRUD (Create, Read, Update, Delete) operations on movie data stored in an in-memory slice.

---

## **Features**

- Fetch all movies
- Fetch a single movie by ID
- Add a new movie
- Update an existing movie
- Delete a movie

---

## **Technologies Used**

- **Language:** Go (Golang)
- **Router:** Gorilla Mux
- **Data Storage:** In-memory slice (suitable for small projects or testing)

---

## **Project Structure**

```plaintext
movie-api/
├── main.go              # Entry point of the application
├── router/              
│   └── router.go        # Defines routes
├── handlers/            
│   └── movie_handlers.go # Contains handler functions for CRUD operations
├── models/              
│   └── movie.go         # Defines Movie and Director structs
├── data/                
│   └── seed.go          # Seeds initial movie data
├── utils/               
│   └── response.go      # Helper functions for responses (e.g., setting headers)
├── go.mod               # Go module file
└── go.sum               # Dependencies file
```

---

## **Endpoints**

### **Base URL**
`http://localhost:8000`

### **1. Get All Movies**
**Endpoint:** `GET /movies`

**Response:**
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

---

### **2. Get a Movie by ID**
**Endpoint:** `GET /movies/{id}`

**Response (Success):**
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

**Response (Not Found):**
```json
"Movie not found"
```

---

### **3. Create a Movie**
**Endpoint:** `POST /movies`

**Request Body:**
```json
{
  "isbn": "9876543210",
  "title": "New Movie",
  "director": {
    "firstname": "Jane",
    "lastname": "Doe"
  }
}
```

**Response:**
```json
{
  "id": "123456",
  "isbn": "9876543210",
  "title": "New Movie",
  "director": {
    "firstname": "Jane",
    "lastname": "Doe"
  }
}
```

---

### **4. Update a Movie**
**Endpoint:** `PUT /movies/{id}`

**Request Body:**
```json
{
  "isbn": "9876543210",
  "title": "Updated Movie",
  "director": {
    "firstname": "Jane",
    "lastname": "Doe"
  }
}
```

**Response:**
```json
{
  "id": "1",
  "isbn": "9876543210",
  "title": "Updated Movie",
  "director": {
    "firstname": "Jane",
    "lastname": "Doe"
  }
}
```

---

### **5. Delete a Movie**
**Endpoint:** `DELETE /movies/{id}`

**Response:**
```json
[
  {
    "id": "2",
    "isbn": "2345678901",
    "title": "Mystery at Dawn",
    "director": {
      "firstname": "Jane",
      "lastname": "Smith"
    }
  }
]
```

---

## **Setup and Run**

### **Prerequisites**
- Go installed (version 1.19 or higher)

### **Steps to Run**

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/movie-api.git
   cd movie-api
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the server:
   ```bash
   go run main.go
   ```

4. Test the API:
   - Use tools like Postman or curl to test endpoints.
