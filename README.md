# API Documentation

## Overview

This project is a Go backend service. Below are the available API endpoints and their corresponding controllers.

---

## Endpoints

### 1. `GET /users`
- **Description:** Retrieve a list of users.
- **Controller:** `UserController.GetUsers`
- **Response:**  
    ```json
    [
        {
            "id": 1,
            "name": "John Doe",
            "email": "john@example.com"
        }
    ]
    ```

### 2. `GET /users/{id}`
- **Description:** Retrieve a user by ID.
- **Controller:** `UserController.GetUserByID`
- **Response:**  
    ```json
    {
        "id": 1,
        "name": "John Doe",
        "email": "john@example.com"
    }
    ```

### 3. `POST /users`
- **Description:** Create a new user.
- **Controller:** `UserController.CreateUser`
- **Request Body:**  
    ```json
    {
        "name": "Jane Doe",
        "email": "jane@example.com"
    }
    ```
- **Response:**  
    ```json
    {
        "id": 2,
        "name": "Jane Doe",
        "email": "jane@example.com"
    }
    ```

### 4. `PUT /users/{id}`
- **Description:** Update an existing user.
- **Controller:** `UserController.UpdateUser`
- **Request Body:**  
    ```json
    {
        "name": "Jane Smith",
        "email": "jane.smith@example.com"
    }
    ```
- **Response:**  
    ```json
    {
        "id": 2,
        "name": "Jane Smith",
        "email": "jane.smith@example.com"
    }
    ```

### 5. `DELETE /users/{id}`
- **Description:** Delete a user by ID.
- **Controller:** `UserController.DeleteUser`
- **Response:**  
    ```json
    {
        "message": "User deleted successfully"
    }
    ```

---

## Controllers

- **UserController**
    - `GetUsers`
    - `GetUserByID`
    - `CreateUser`
    - `UpdateUser`
    - `DeleteUser`

---

> **Note:**  
> Update the documentation as you add more endpoints or controllers.
