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

## Setup Frontend 

 Navigate to the frontend folder
```bash
cd frontend
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

## Running The App
Login
```bash
email : cs@test.com
password : password

email : operation@test.com
password : password
```


### 🐳 Run with Docker

Run Docker

---

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




