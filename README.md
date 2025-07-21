# Public Library API

## 📚 Description

The **Public Library API** streamlines the retrieval, creation, updating, and deletion of book records, providing an efficient and organized way to manage book data.

---

## 🚀 Run Service

To run the project using Docker:

1. **Clone the repository**

   ```bash
   git clone https://github.com/your-username/your-repo-name.git
   cd your-repo-name

2. **Start the services and access the API**

  Run the following command to build and start the containers:
  
  ```bash
  docker-compose up --build
  ```
  
  Once the containers are running, the API will be accessible at: http://localhost:8082

## 📖 Endpoints Description

  ### Get All Books
  
  **GET** `/api/v1/books?offset=0&limit=10`  
  Returns a paginated list of books.
  
  ---
  
  ### Get Book By ID
  
  **GET** `/api/v1/books/:id`  
  Fetch a specific book by its ID.
  
  ---
  
  ### Create a New Book
  
  **POST** `/api/v1/books`  
  Creates a new book entry.
  
  **Request Body:**
  
  ```json
  {
    "author": "John Henry",
    "title": "The return of king",
    "year": 1957
  }
  ```
  ### Update an Existing Book
  
  **PUT** `/api/v1/books/:id`  
  Updates book details by ID.
  
  **Request Body:**
  
  ```json
  {
    "author": "John Henry",
    "title": "The return of king",
    "year": 1958
  }
  ```
  ### 🔹 Delete a Book
  
  **DELETE** `/api/v1/books/:id`  
  Deletes the book with the given ID.

Import postman collection from postmanCollection folder to get started!
    
