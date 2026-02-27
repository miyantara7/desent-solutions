# 🚀 Desent Solutions API

Production-ready REST API built with **Golang + Gin + Clean Architecture + DI (sarulabs/di)**.

## 🚀 Quick Start (Local)

### 1️⃣ Install deps

```bash
go mod tidy
```

### 2️⃣ Run app

```bash
make run
```

### 3️⃣ Test

```bash
curl http://localhost:8080/ping
```

---

## 🐳 Run with Docker

### Build image

```bash
make docker-build
```

### Run container

```bash
make docker-run
```

Test:

```bash
curl http://localhost:8080/ping
```

---

## 🧪 Run Speed Test (Local)

```bash
make speedrun
```

---

One-shot end-to-end verification.

```bash
curl -X POST https://desent-solutions-production.up.railway.app/speedrun
```

---

## 🔌 API Endpoints

### Health

```
GET /ping
POST /echo
```

### Books

```
POST   /books
GET    /books
GET    /books/:id
PUT    /books/:id
DELETE /books/:id
```

### Auth

```
POST /auth/token
GET  /protected/books
```

### Speedrun

```
POST /speedrun
```

---

## 🧪 Example cURL

### Create Book

```bash
curl -X POST https://desent-solutions-production.up.railway.app/books \
  -H "Content-Type: application/json" \
  -d '{"title":"Golang","author":"Arya"}'
```

### Protected Route

```bash
curl -X GET https://desent-solutions-production.up.railway.app/protected/books \
  -H "Authorization: Bearer secret-token"
```

---

## 👨‍💻 Author

Miyaantara.

---
