
# 🍃 DURIANPAY SUBMISSION

**Muhammad Syafri Surya Muza`im — Fullstack Developer**

---

## 📌 Project Overview

This repository contains a fullstack payment dashboard application consisting of:

- ⚙️ Backend Service  
- 🎨 Frontend Web App  
- 📘 OpenAPI Documentation  
- 🐳 Docker Support  
- 🗄 Automatic Data Seeding  

---

## 📦 Clone Repository

```bash
git clone https://github.com/muzaim/payment-dashboard.git
cd payment-dashboard
````

---

## ✅ Requirements

Make sure you have:

* Node.js >= 18
* Docker & Docker Compose
* Make
* Git

---

# 🛠 Backend Setup

Navigate to backend folder:

```bash
cd backend
```

### 🔐 Environment Configuration

Copy environment template and adjust values:

```bash
cp env.sample .env
```

---

### ⚙️ Generate OpenAPI

```bash
make openapi-gen
```

---

### 🔑 Generate JWT Secret

```bash
make gen-secret
```

---

### 🚀 Run Backend Server

```bash
cp env.sample .env
make tool-openapi
make openapi-gen
make dep
make gen-secret
make run
```

---

## 🗄 Data Seeding

Initial data seeding runs automatically when the application starts for the first time.

A local SQLite database will be created.

### 📄 Output

```bash
database.db
```

### 🧠 Notes

* Seeding runs automatically
* No manual command required
* Deleting `database.db` will regenerate data on next run

---

# 🎨 Frontend Setup

Navigate to frontend folder:

```bash
cd frontend
```

### 🔐 Environment Configuration

```bash
cp .env.example .env
```

---

### 📦 Install Dependencies

```bash
npm install
```

---

### 🔥 Development Mode (Hot Reload)

```bash
npm run dev
```

---

### 🏗 Production Build

```bash
npm run build
```

---

# 🐳 Run with Docker

Ensure Docker is running.

---

## ▶️ Backend Container

```bash
cd backend
docker compose up --build -d
```

---

## ▶️ Frontend Container

Open new terminal:

```bash
cd frontend
docker compose up --build -d
```

---

# 🔐 Login Credentials

```bash
email: cs@test.com
password: password

email: operation@test.com
password: password
```

---

# 📘 View OpenAPI Documentation (Swagger UI)

### ✅ Prerequisites

* Docker installed
* `openapi.yaml` exists in backend folder

---

### 🚀 Steps

```bash
cd backend
```

Run Swagger UI:

```bash
docker run -p 8080:8080 \
  -e SWAGGER_JSON=/foo/openapi.yaml \
  -v $(pwd):/foo \
  swaggerapi/swagger-ui
```

---

### 🌐 Open in Browser

```bash
http://localhost:8080
```

You can:

✅ View all API endpoints
✅ Inspect request & response
✅ See schemas and models
✅ Test APIs via Swagger UI

---

## ✨ Quick Access

* Backend → [http://localhost:3000](http://localhost:8000)
* Frontend → [http://localhost:5173](http://localhost:5173)
* Swagger → [http://localhost:8080](http://localhost:8080)

---

Thank you for reviewing my technical test 🚀


tinggal bilang 😄
```
