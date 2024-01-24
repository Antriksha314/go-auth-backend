# go-auth-backend
# Authentication Project

This is a simple authentication project in Go where users can register, login, and access protected routes using JWT tokens.

## Getting Started

Before starting the project, make sure you have set up the necessary environment variables and created the PostgreSQL database.

### Prerequisites

- Go installed on your machine
- PostgreSQL installed and running
- `.env` file for storing environment variables

### Environment Variables

Create a `.env` file in the root of your project with the following variables:

```env
GIN_MODE=release
PG_HOST=localhost
PG_PORT=5432
PG_USER=your_db_user
PG_PASSWORD=your_db_password
PG_DB=your_db_name
JWT_KEY=your_jwt_secret
