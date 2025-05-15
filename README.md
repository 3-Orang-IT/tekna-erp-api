# ERP Tekna API

This document provides instructions to set up the ERP Tekna API project from scratch.

## Prerequisites

1. Install [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/).
2. Install [Git](https://git-scm.com/) for version control.

> **Note**: PostgreSQL is included in the Docker setup, so you do not need to install it manually.

## Setup Instructions

### 1. Clone the Repository

```bash
git clone <repository-url>
cd erp-tekna-api
```

### 2. Configure Environment Variables

Copy the example environment file and adjust the values as needed:

```bash
cp .env.example .env
```

### 3. Build and Start the Docker Containers

```bash
docker-compose up --build
```

This will start the application and the PostgreSQL database.

### 4. Access the Application

The application will be available at `http://localhost:8080`.

## Additional Docker Commands

-   **Stop containers**: `docker-compose down`
-   **Rebuild containers**: `docker-compose up --build`
-   **View logs**: `docker-compose logs -f`

## Troubleshooting

-   Ensure Docker is running.
-   Check the `.env` file for correct configurations.
-   Use `docker-compose logs` to debug issues.

## License

This project is licensed under the MIT License.
