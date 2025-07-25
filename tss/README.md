# TSS Implementation

This project implements a Threshold Signature Scheme (TSS) using gRPC and Redis.

## Project Structure

```
tss
├── client                  # Client code calling keygen and sign APIs, interacts with server via gRPC
│ └── keygen_client.go      # Client calls key generations (keygen) API
│ └── sign_client.go        # Client calls signing (sign) API
├── config                  # System configuration
├── db                      # Database management
│ └── migrations            # SQL scripts for creating and modifying tables/schema
│ └── queries               # SQL files declaring queries
│ └── sqlc                  # Go code auto-generated by sqlc based on SQL files
│ └── db.go                 # DB connection initialization and common helpers
├── party                   # Logic for each node in MPC (multi-party computation) algorithm
├── proto                   # gRPC definitions, auto-generated protobuf code
├── server                  # Core server logic
│ └── helper.go             # Utility helper functions
│ └── keygen.go             # MPC key generation (keygen) logic
│ └── sign.go               # MPC signing (sign) logic
│ └── server.go             # gRPC server initialization and running
├── utils                   # Common utility functions (encryption, compression, file save, conversion, etc.)
├── .env.example            # Environment configuration file
├── docker-compose.yml      # Docker compose configuration
├── Dockerfile              # Docker image build configuration
├── go.mod / go.sum         # Go dependencies management files
├── main.go                 # Server entry point
├── Makefile                # Script for automating common tasks
└── README.md               # Project documentation
├── sqlc.yaml               # Configuration for sqlc tool
```

## Prerequisites

- Go 1.22
- PostgreSQL
- Redis

## Setup

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/tss-impl.git
   cd tss
   ```

2. Copy the example environment file and update it with your configuration:

   ```sh
   cp .env.example .env
   ```

3. Install dependencies:

   ```sh
   go mod tidy
   ```

4. Start Services (PostgreSQL, Redis, App):

   ```sh
   docker-compose up -d
   ```

5. Run Database Migrations with Goose:

First, install Goose if you haven't already:

   ```sh
   go install github.com/pressly/goose/v3/cmd/goose@latest
   ```

Then run the migrations:

   ```sh
   goose -dir db/migrations postgres "postgres://vaulta:123@localhost:5432/mpc_key?sslmode=disable" up
   ```

   ✅ Note: If the command above fails, try replacing localhost with 127.0.0.1.

## Running the Project

1. Start the server:

   ```sh
   make run
   ```

2. Run the client to generate the key:
   ```sh
   go run client/keygen_client.go
   ```

3. Run the client to sign:
   ```sh
   go run client/sign_client.go
   ```

## Project Components

### Client

The client initiates key generation and signing processes. It communicates with the server using gRPC.

#### Key Generation

The client sends a `NotifyAction` request to the server to initiate the key generation process. The server responds with the generated key shares, which are then saved to a file.

#### Signing

The client loads the key shares from the file and sends a `NotifyAction` request to the server to initiate the signing process. The server responds with the signature.

### Server

The server handles key generation and signing requests. It uses Redis for inter-node communication and PostgreSQL for storing encrypted shares.

#### Key Generation

The server receives a `NotifyAction` request from the client to initiate the key generation process. It generates the key shares and publishes them to a Redis channel.

#### Signing

The server receives a `NotifyAction` request from the client to initiate the signing process. It retrieves the key shares from the database and performs the signing operation.

### Configuration

Configuration is managed using environment variables. See [.env.example](http://_vscodecontentref_/2) for the required variables.

### Database

The database schema is managed using SQL migrations. See the [migrations](http://_vscodecontentref_/3) directory for the migration files.

### Utilities

Utility functions for file operations, encryption, and compression are located in the [utils](http://_vscodecontentref_/4) directory.

## License

This project is licensed under the MIT License.
