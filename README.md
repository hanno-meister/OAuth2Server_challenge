# Go HTTP JWT Authentication Server

This project implements a simple HTTP server written in Go that provides JWT-based user authentication, token introspection, and JSON Web Key (JWK) management.

## Endpoints

### `POST /signup`

- **Description**: Registers a new user with an email and password.
- **Request Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "securepassword"
  }
  ```

- Response: 200 Created on successful signup or appropriate error messages otherwise.
  
### `POST /token`
- **Description**: Authenticates a user and returns a JWT token.
- **Request Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "securepassword"
  }
  ```

- **Response**:
  ```json
  {
    "token": "<JWT_token_here>"
  }
  ```

### `POST /introspect`
- **Description**: Validates and introspects the provided JWT token.
- **Request Body**:
  ```json
  {
    "token": "<JWT_token_here>"
  }
  ```

- **Response**:
  ```json
  {
      "active": true,
      "email": "user@example.com",
      "client_id": 1,
      "exp": 1746123429
  }
  ```

### `GET /signingkeys`
- **Description**: Retrieves the server's public JSON Web Keys (JWK) for token verification.
- **Example Response**:
  ```json
  {
    "key": {
      "kty": "RSA",
      "n": "<base64url_encoded_modulus>",
      "e": "<base64url_encoded_exponent>"
    }
  }
  ```

## Prerequisites
- Docker installed
- Docker Compose installed

## Installation
Clone the repository
  ```sh
git clone https://github.com/hanno-meister/go_challenge.git
  ```
  Start Docker Servers:
```sh
docker-compose up
  ```
  
## Testing

You can test endpoints using tools like Postman or `curl`.

### Example `curl` requests:

All endpoints are served on `localhost:3000` by default.

**Example Signup a user:**
```sh
curl -X POST http://localhost:3000/signup \
-H "Content-Type: application/json" \
-d '{"email":"user@example.com", "password":"securepassword"}'
```