## Setting Up the Backend Database

This section provides instructions for configuring and connecting the backend to the database. Follow the steps below to ensure proper setup:

1. **Install Database Server**  
    Install PostgreSQL on your local system. For example, on Ubuntu, run:
    Make sure the PostgreSQL service is running before proceeding.

2. **Create Database**  
    Create a new database for the project:
    createdb sip_portal_db

3. **Configure Environment Variables**  
    Create a `.env` file in the project root if it doesn't exist, and set up your database connection details as shown below. Refer to the `env_format.txt` file in the project for the exact variable names and format.

    ```
    DB_HOST=localhost
    DB_PORT=5432
    DB_NAME=sip_portal_db
    DB_USER=your_username
    DB_PASSWORD=your_password
    ```
4. **Example Database URL**  
    When configuring your database connection string, use the following format, replacing `your_postgres_password` with your actual PostgreSQL password:

        ```
    DB='postgres://postgres:your_postgres_password@localhost:5432/sip_portal_db?sslmode=disable'
        ```

    Ensure that the username, password, host, port, and database name match your local setup.

5. **Test Database Connection**  
    Start the backend server to verify that the connection to the database is successful. If you encounter any issues, refer to the project documentation for troubleshooting and additional configuration details.

## Setting Up the Go Environment

Follow these steps to install Go, set up the environment, and install dependencies required to run the backend project:

1. **Install Go**  
    Download and install Go from the [official website](https://golang.org/dl/).  

2. **Install Project Dependencies**  
    Navigate to the project root directory and run:
    ```bash
    go mod download
    ```
    This will install all required Go modules as specified in `go.mod`.

3. **Run the Backend Project**  
    The entry point for the backend is typically the `server.go` file in the project root. Start the server with:
    ```bash
    go run server.go
    ```
    Ensure your environment variables are set up as described above before running the project.

Refer to the project documentation for additional setup or troubleshooting steps.