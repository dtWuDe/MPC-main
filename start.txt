@echo off
echo 🚀 Starting MPC Wallet Project...

REM Check if Docker is running
docker info >nul 2>&1
if errorlevel 1 (
    echo ❌ Docker is not running. Please start Docker and try again.
    pause
    exit /b 1
)

REM Start services with Docker Compose
echo 📦 Starting PostgreSQL and Redis...
docker-compose up -d

REM Wait for services to be ready
echo ⏳ Waiting for services to be ready...
timeout /t 10 /nobreak >nul

echo ✅ Services started successfully!
echo.
echo 📋 Next steps:
echo 1. Set up environment files (see README.md for details)
echo 2. Run database migrations:
echo    cd backend ^&^& make migrate
echo    cd ../tss ^&^& make migrate
echo 3. Start the applications:
echo    Terminal 1: cd tss ^&^& make run
echo    Terminal 2: cd backend ^&^& make run
echo    Terminal 3: cd client ^&^& npm install ^&^& npm run dev
echo.
echo 🌐 Services will be available at:
echo    Frontend: http://localhost:5173
echo    Backend API: http://localhost:5001
echo    TSS Service: localhost:50051
echo    Swagger Docs: http://localhost:5001/swagger/index.html
pause 