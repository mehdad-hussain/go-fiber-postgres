# Contacts Application API

## Table of Contents
- [Overview](#overview)
- [Project Setup](#project-setup)
- [Running the Application](#running-the-application)
- [API Consumption](#api-consumption)

## Overview
This Contacts Application API is built using Go with the Fiber framework and GORM for ORM. The API provides functionalities to manage contacts, including creating and retrieving contacts. It retrieves user information based on email for contact creation.

## Project Setup

### Prerequisites
Before you begin, ensure you have the following installed on your machine:
- [Go](https://golang.org/doc/install) (version 1.17 or later)
- [PostgreSQL](https://www.postgresql.org/download/)
- [GORM](https://gorm.io/index.html) library

### Clone the Repository
First, clone the repository to your local machine:

```bash
git clone https://github.com/mehdad-hussain/go-fiber-postgres.git
cd go-fiber-postgres
```

### Configure Database
1. Create a PostgreSQL database for the application.
2. Update the database connection settings in the codebase. You can typically find this in the `config` or `database` package.

### Install Dependencies
Navigate to the project directory and install the required Go modules:

```bash
go mod tidy
```

## Running the Application

To run the application, use the following command:

```bash
go run cmd/app/main.go
```

This command will start the server on the default port `3000`. You can change the port by modifying the configuration in the main application file if needed.

## API Consumption

### Base URL
The base URL for the API is:

```
http://localhost:3000/api/v1
```

### Endpoints

1. **Get Contacts**
   - **Endpoint**: `GET /contacts`
   - **Description**: Retrieve all contacts for the authenticated user.
   - **Headers**:
     - `Authorization`: Bearer token for authentication (required).
   - **Query Parameters**:
     - `page`: The page number for pagination (optional, default: 1).
     - `limit`: The number of contacts per page (optional, default: 10).
   - **Example Request**:
     ```http
     GET http://localhost:3000/api/v1/contacts?page=1&limit=10
     Authorization: Bearer <your_jwt_token>
     ```
   - **Success Response**:
     ```json
     [
       {
         "id": 1,
         "user_id": 1,
         "name": "John Doe",
         "email": "john@example.com",
         "phone": "123-456-7890"
       },
       ...
     ]
     ```
   - **Error Response**:
     ```json
     {
       "error": "Unable to retrieve contacts"
     }
     ```

2. **Create Contact**
   - **Endpoint**: `POST /contacts`
   - **Description**: Create a new contact.
   - **Headers**:
     - `Authorization`: Bearer token for authentication (required).
     - `Content-Type`: application/json
   - **Request Body**:
     ```json
     {
       "name": "John Doe",
       "email": "john@example.com",
     }
     ```
   - **Example Request**:
     ```http
     POST http://localhost:3000/api/v1/contacts
     Authorization: Bearer <your_jwt_token>
     Content-Type: application/json
     ```
   - **Success Response**:
     ```json
     {
       "id": 1,
       "name": "John Doe",
       "email": "john@example.com",
     }
     ```
   - **Error Response**:
     ```json
     {
       "error": "Failed to create contact"
     }
     ```

### Testing the API
You can test the API using tools like [Postman](https://www.postman.com/) or [cURL](https://curl.se/) to send requests to the endpoints mentioned above. Make sure to include the JWT token in the `Authorization` header for all requests that require authentication.
