# GoFiber-Login

A simple web application built using Go (Golang) and the Fiber framework. This application allows users to register, log in, and manage their sessions with JWT authentication. It uses MongoDB as the database to store user details securely.

## Features

- **User Registration:** Users can register by providing a name, email, and password. Passwords are securely hashed before being stored.
- **User Login:** Users can log in with their email and password. Upon successful login, a JWT token is generated for session management.
- **JWT Authentication:** A JWT token is issued upon login for secure authentication and session handling.
- **MongoDB Database:** All user data is stored in MongoDB, ensuring scalability and performance.

## Tech Stack

- **Go (Golang):** The programming language used for backend development.
- **Fiber Framework:** A fast and minimal web framework for Go.
- **MongoDB:** NoSQL database used to store user data.
- **JWT:** JSON Web Token for user authentication.
- **Bcrypt:** Password hashing for secure storage.

## Prerequisites

Before running the application, ensure that you have the following installed:

- [Go (Golang)](https://golang.org/dl/)
- [MongoDB](https://www.mongodb.com/)
- [Go Fiber](https://github.com/gofiber/fiber) (installed with `go get`)
- [Bcrypt](https://github.com/golang/crypto) (installed with `go get`)

Additionally, create a `.env` file for environment variables.

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/GoFiber-Login.git
   cd GoFiber-Login
2. Install dependencies:

   ```bash
   go mod tidy
3. Create a .env file in the root directory and add your environment variables:
   ```makefile
   MONGO_URI=mongodb://localhost:27017
   JWT_SECRET=your_jwt_secret_key
4. Run the application:

   ```bash
   go run main.go


## API Routes

### User Endpoints

- **GET** `"/register"`:  
  Route to display the user registration page.  
  **Request Parameters**: None  
  **Response**: HTML page for user registration.

- **POST** `"/register"`:  
  Route to register a new user by submitting the registration form.  
  **Request Parameters**:  
    - `name` (string): User's name  
    - `email` (string): User's email  
    - `password` (string): User's password  
  **Response**: Redirects to the login page upon successful registration.  
  **Error Handling**: Returns an error if any required field is missing or there is a database issue.

- **GET** `"/login"`:  
  Route to display the login page.  
  **Request Parameters**: None  
  **Response**: HTML page for user login.

- **POST** `"/login"`:  
  Route to log in a user and generate a JWT token.  
  **Request Parameters**:  
    - `email` (string): User's email  
    - `password` (string): User's password  
  **Response**:  
    - **Success**: Redirects to the dashboard (or returns a JWT token if implemented as an API endpoint).
    - **Error Handling**: Returns an error if the credentials are incorrect or if there's a database issue.

### Example Requests

1. **POST** `/register`

   Example request body:
   ```json
   {
     "name": "Abhijeet Srivastava",
     "email": "abhijeetsrivastava2189@gmail.com",
     "password": "securepassword123"
   }



2. **POST** `/login`

   Example request body:
   ```json
   {
     "email": "abhijeetsrivastava2189@gmail.com",
     "password": "securepassword123"
   }

