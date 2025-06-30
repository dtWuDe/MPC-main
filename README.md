# MPC Wallet Project - Complete Setup Guide

A comprehensive Multi-Party Computation (MPC) wallet system with B2B wallet-as-a-service capabilities, featuring distributed key generation, threshold signatures, and Ethereum integration.

## 🏗️ Project Architecture

This project consists of three main components:

1. **Backend API** (`backend/`) - Go-based REST API with MPC integration
2. **TSS Service** (`tss/`) - Threshold Signature Scheme service for distributed signing
3. **Frontend Client** (`client/`) - React/TypeScript web application

## 📋 Prerequisites

### Required Software
- **Go** 1.22.2 or higher
- **Node.js** 18+ and npm/yarn
- **PostgreSQL** 13+
- **Redis** 6+
- **Docker** and Docker Compose (optional, for easy setup)

### Optional
- **Ethereum Node** access (for blockchain integration)
- **Goose** (for database migrations)

## 🚀 Quick Start

### Option 1: Using Docker Compose (Recommended)

1. **Clone and navigate to the project:**
   ```bash
   git clone <repository-url>
   cd MPC-main
   ```

2. **Start all services:**
   ```bash
   # Start PostgreSQL, Redis, and other dependencies
   docker-compose up -d
   ```

3. **Run database migrations:**
   ```bash
   # Backend migrations
   cd backend
   make migrate
   
   # TSS migrations
   cd ../tss
   make migrate
   ```

4. **Start the services:**
   ```bash
   # Terminal 1: Start TSS service
   cd tss
   make run
   
   # Terminal 2: Start Backend API
   cd backend
   make run
   
   # Terminal 3: Start Frontend
   cd client
   npm install
   npm run dev
   ```

### Option 2: Manual Setup

#### 1. Database Setup

**PostgreSQL:**
```bash
# Create databases
createdb mpc_db
createdb mpc_key

# Or using Docker
docker run --name postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=mpc_db -p 5432:5432 -d postgres:13
```

**Redis:**
```bash
# Using Docker
docker run --name redis -p 6379:6379 -d redis:6-alpine

# Or install locally
# Follow Redis installation guide for your OS
```

#### 2. Environment Configuration

**Backend Environment** (`backend/.env`):
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=mpc_db
REDIS_URL=localhost:6379
ETH_NODE_URL=wss://ethereum-sepolia-rpc.publicnode.com
JWT_SECRET=your-secret-key
API_KEY_SECRET=your-api-key-secret
```

**TSS Environment** (`tss/.env`):
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=mpc_key
REDIS_URL=localhost:6379
NODE_ID=1
ENCRYPT_KEY=your-encryption-key
```

**Frontend Environment** (`client/.env`):
```env
VITE_API_BASE_URL=http://localhost:5001
VITE_TSS_BASE_URL=http://localhost:50051
```

#### 3. Install Dependencies

**Backend:**
```bash
cd backend
go mod download
```

**TSS Service:**
```bash
cd tss
go mod download
```

**Frontend:**
```bash
cd client
npm install
```

#### 4. Run Database Migrations

**Backend:**
```bash
cd backend
make migrate
```

**TSS:**
```bash
cd tss
make migrate
```

#### 5. Start Services

**Terminal 1 - TSS Service:**
```bash
cd tss
make run
```

**Terminal 2 - Backend API:**
```bash
cd backend
make run
```

**Terminal 3 - Frontend:**
```bash
cd client
npm run dev
```

## 🧪 Testing the System

### 1. Test TSS Key Generation

```bash
cd tss/client/keygen
go run main.go
```

This will:
- Connect to the TSS service
- Generate distributed key shares
- Save encrypted share data to a file

### 2. Test TSS Signing

```bash
cd tss/client/sign
go run main.go share_data_<session_id>.txt
```

This will:
- Load the key shares from the file
- Sign a test message using distributed signing

### 3. Test Backend API

The API will be available at `http://localhost:5001`

**Health Check:**
```bash
curl http://localhost:5001/health
```

**Swagger Documentation:**
```
http://localhost:5001/swagger/index.html
```

### 4. Test Frontend

The frontend will be available at `http://localhost:5173`

## 🔧 Development Commands

### Backend Commands
```bash
cd backend
make run          # Start the API server
make build        # Build the application
make test         # Run tests
make migrate      # Run database migrations
make generate     # Generate SQL code
```

### TSS Commands
```bash
cd tss
make run          # Start the TSS service
make build        # Build the application
make test         # Run tests
make migrate      # Run database migrations
make generate     # Generate SQL code
```

### Frontend Commands
```bash
cd client
npm run dev       # Start development server
npm run build     # Build for production
npm run preview   # Preview production build
```

## 📊 Service Ports

| Service | Port | Description |
|---------|------|-------------|
| Frontend | 5173 | React development server |
| Backend API | 5001 | REST API server |
| TSS Service | 50051 | gRPC TSS service |
| PostgreSQL | 5432 | Database |
| Redis | 6379 | Cache/Message broker |

## 🔐 B2B Features

The system includes B2B wallet-as-a-service features:

- **Organization Management**: Multi-tenant organization support
- **API Key Authentication**: Secure API key-based authentication
- **Usage Tracking**: Monitor API usage and transactions
- **Webhook Support**: Real-time notifications for events

### Creating an Organization

1. Register an organization via the API
2. Generate API keys for the organization
3. Use API keys to authenticate requests
4. Monitor usage and transactions

## 🐛 Troubleshooting

### Common Issues

1. **Database Connection Errors:**
   - Ensure PostgreSQL is running
   - Check database credentials in `.env` files
   - Verify database exists

2. **Redis Connection Errors:**
   - Ensure Redis is running
   - Check Redis URL in configuration

3. **TSS Service Issues:**
   - Verify all TSS nodes are running
   - Check node configuration
   - Ensure proper network connectivity between nodes

4. **Frontend Build Issues:**
   - Clear node_modules and reinstall: `rm -rf node_modules && npm install`
   - Check Node.js version compatibility

### Logs

Check logs for each service:
- **Backend**: Console output or log files
- **TSS**: Console output with debug information
- **Frontend**: Browser developer tools console

## 📚 Documentation

- **Backend API**: `http://localhost:5001/swagger/index.html`
- **B2B Implementation**: See `B2B_IMPLEMENTATION.md`
- **TSS Client Usage**: See `tss/client/README.md`

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.