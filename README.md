# User-Service

This is a Go application for managing user service.The User Service is responsible for user-related operations such as registration, authentication, and profile management. It provides endpoints for creating, updating, and deleting user accounts, as well as handling authentication and session management
## Features

The Rider Service offers comprehensive features tailored for efficient rider management and delivery operations. Key functionalities include:

-User Registration: Allows users to create new accounts by providing necessary information such as username, email, and password.

-Authentication: Provides secure authentication mechanisms such as token-based authentication or session management to verify user identity during login attempts.

-Profile Management: Enables users to update their profile information, including personal details, contact information, and preferences.

## Getting Started with User-Service

### Requirements

- A working Go environment - [https://go.dev/dl/](https://go.dev/dl/)
- Check the go version with command: go version.
- One should also be familiar with the Golang syntax. [Golang Tour](https://tour.golang.org/) has an excellent guided tour and highly recommended.

### Installation

## GOFR as dependency used for migrations

- To get the GOFR as a dependency, use the command:
  `go get gofr.dev`

- Then use the command `go mod tidy`, to download the necessary packages.


### To Run Server

Run `go run main.go` command in CLI.

## Usage

The application provides the following RESTful endpoints:

- `POST /users/register`: Register a new user.
- `POST /users/login`: Login User
- `GET /users/profile"`: Get the user profile.


## Dependencies

The application uses the following dependencies:

- `gofr.dev/pkg/gofr`: A Go web framework used for handling HTTP requests.
- `User-service/handlers`: Handlers package for handling HTTP requests related to users.
- `User-service/services`: Services package for business logic related to users.
- `User-service/stores`:Store package for handling db operations related to users.

For any information please reach out to me via rishimanudeepg@gmail.com