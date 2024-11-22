# go\_JWT\_api

This project is a Go-based API that uses JWT for authentication and authorization. It includes user registration, login, and logout functionalities, as well as user information retrieval and role-based access control.

## Getting Started

### Prerequisites

- Go 1.23.3 or later
- PostgreSQL
- Git

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/RikiLaNeko/go-postgres-jwt-auth-api.git
    cd go-postgres-jwt-auth-api
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Set up environment variables:
   Create a `.env` file in the root directory and add the following variables:
    ```env
    POSTGRES_HOST=your_postgres_host
    POSTGRES_USER=your_postgres_user
    POSTGRES_PASSWORD=your_postgres_password
    POSTGRES_DB=your_postgres_db
    POSTGRES_PORT=your_postgres_port
    JWT_SECRET=your_jwt_secret
    JWT_EXPIRED_IN=your_jwt_expiration_duration
    JWT_MAXAGE=your_jwt_max_age
    CLIENT_ORIGIN=your_client_origin
    ```

4. Run the application:
    ```sh
    go run main.go
    ```

### Usage

- Register a new user: `POST /api/auth/register`
- Login: `POST /api/auth/login`
- Logout: `GET /api/auth/logout`
- Get authenticated user info: `GET /api/users/me`
- Get all users (admin/moderator only): `GET /api/users/`

### Documentation

You can generate and view the documentation using `godoc`.

1. Install `godoc` if you haven't already:
    ```sh
    go install golang.org/x/tools/cmd/godoc@latest
    ```

2. Run `godoc`:
    ```sh
    godoc -http=:6060
    ```

3. Open your web browser and navigate to `http://localhost:6060/pkg/github.com/RikiLaNeko/go-postgres-jwt-auth-api/` to view the documentation.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.