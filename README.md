# DURIANPAY SUBMISSION
```bash
Muhammad Syafri Surya Muza`im - Fullstack Developer
```


## Clone the Repository

Clone the repository using Git:
```bash
git clone https://github.com/muzaim/payment-dashboard.git
```

## Setup Backend 
Navigate to the backend folder
```bash
cd backend
```

Configure environment variables
```bash
copy the .env.example and rewrite using your env
```

generate openapi:
```bash
make openapi-gen
```

generate JWT_SECRET:
```bash
make gen-secret
```

Run server:
```bash
cp env.sample .env
make tool-openapi
make openapi-gen
make dep
make gen-secret
make run
```

## Data Seeding

Initial data seeding will be automatically executed when you run this application.

During the first startup, the system will generate seed data and create a local database file.

### 📄 Output

After the application is running, you will find the generated database file: database.db



## Setup Frontend 

 Navigate to the frontend folder
```bash
cd frontend
```

Configure environment variables
```bash
copy the .env.example and rewrite using your env
```

Install all dependencies that required
```bash
npm install
```

Compile and Hot-Reload for Development
```bash
npm run dev
```

Type-Check, Compile and Minify for Production
```bash
npm run build
```

### 🐳 Run with Docker

Run Docker

### Build Backend Container
Navigate to backend folder
```bash
cd backend
```

Build container
```bash
docker compose up --build -d
```

### Build Frontend Container
Navigate to frontend folder
```bash
cd frontend
```

Build container
```bash
docker compose up --build -d
```

## Running The App
Login
```bash
email : cs@test.com
password : password

email : operation@test.com
password : password
```


## View OpenAPI Documentation (Without Docker)

## ✅ Prerequisites

- Docker installed
- `openapi.yaml` exists in backend folder
---

## 🚀 Steps

Go to the backend folder
```bash
cd backend
```

Run Swagger UI using docker
```bash
docker run -p 8080:8080 \
  -e SWAGGER_JSON=/foo/openapi.yaml \
  -v $(pwd):/foo \
  swaggerapi/swagger-ui
```

Open in browser
```bash
http://localhost:8080
```



