# TSS Client Applications

This directory contains client applications for interacting with the TSS (Threshold Signature Scheme) service.

## Directory Structure

```
client/
├── keygen/
│   └── main.go          # Key generation client
├── sign/
│   └── main.go          # Signing client
└── README.md            # This file
```

## Usage

### 1. Key Generation Client

The keygen client generates distributed key shares using the TSS protocol.

```bash
cd tss/client/keygen
go run main.go
```

**What it does:**
- Connects to the TSS service
- Initiates a key generation session
- Waits for the distributed key generation to complete
- Saves the encrypted share data to a file

**Output:**
- Creates a file named `share_data_<session_id>.txt`
- Contains the encrypted key share data needed for signing

### 2. Signing Client

The sign client uses the generated key shares to sign messages.

```bash
cd tss/client/sign
go run main.go share_data_<session_id>.txt
```

**What it does:**
- Loads the encrypted share data from the specified file
- Connects to the TSS service
- Initiates a signing session
- Uses the distributed key shares to sign a message

**Parameters:**
- `share_data_<session_id>.txt`: The file containing encrypted key share data

## Example Workflow

1. **Generate Keys:**
   ```bash
   cd tss/client/keygen
   go run main.go
   # Output: share_data_abc123-def456-ghi789.txt
   ```

2. **Sign a Message:**
   ```bash
   cd tss/client/sign
   go run main.go share_data_abc123-def456-ghi789.txt
   ```

## Configuration

Both clients connect to:
- **TSS Service**: `localhost:50051` (gRPC)
- **Redis**: `localhost:6379` (for result notifications)

Make sure the TSS service is running before using the clients.

## Security Notes

- Key shares are encrypted before storage
- No single node has access to the complete private key
- Threshold signing requires cooperation between multiple nodes
- Share data files should be kept secure and not shared 