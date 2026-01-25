# LintasPay Backend API

LintasPay adalah backend API sederhana untuk sistem **dompet digital (e-wallet)** yang dibangun menggunakan **Go (Golang)** dengan framework **Gin**, ORM **GORM**, dan dokumentasi API menggunakan **Swagger (swaggo)**.

Project ini dirancang dengan pendekatan **Clean Architecture** agar mudah dikembangkan, diuji, dan dipelihara.

---

## ✨ Fitur Utama

- 🔐 Authentication (Register & Login)
- 🪪 JWT-based Authorization
- 👛 Wallet Management
- 💰 Top Up Saldo
- 🔁 Transfer Antar User
- 📜 Riwayat Transaksi
- 📘 Swagger API Documentation

---

## 🛠️ Tech Stack

- **Go** >= 1.21
- **Gin** (HTTP Framework)
- **GORM** (ORM)
- **PostgreSQL** (Database)
- **JWT** (Authentication)
- **Swagger (swaggo)**

---

## 📁 Struktur Folder

```
lintas-pay/
├── cmd/
│   └── main.go
├── docs/                # Swagger output
├── internal/
│   ├── app/
│   │   ├── container/
│   │   └── middleware/
│   └── modules/
│       ├── user/
│       ├── wallet/
│       └── transactions/
├── pkg/
│   ├── jwt/
│   └── refrence/
├── go.mod
└── README.md
```

---

## ⚙️ Setup & Installation

### 1️⃣ Clone Repository

```bash
git clone https://github.com/HN721/Backend-LintasPay.git
cd lintas-pay
```

### 2️⃣ Install Dependencies

```bash
go mod tidy
```

### 3️⃣ Konfigurasi Environment

Buat file `.env`:

```env
APP_PORT=8080
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=lintaspay
JWT_SECRET=supersecretkey
```

---

## ▶️ Menjalankan Aplikasi

```bash
go run cmd/main.go
```

Server akan berjalan di:

```
http://localhost:8080
```

---

## 📘 Swagger Documentation

### Generate Swagger

```bash
swag init -g cmd/main.go --parseDependency --parseInternal
```

### Akses Swagger UI

```
http://localhost:8080/swagger/index.html
```

---

## 🔐 Authentication

Gunakan **Bearer Token** untuk endpoint yang membutuhkan autentikasi:

```
Authorization: Bearer <JWT_TOKEN>
```

---

## 📌 Endpoint Utama

### Auth

- `POST /auth/register`
- `POST /auth/login`

### Wallet

- `POST /wallet/create`

### Transactions

- `POST /trx/top-up`
- `POST /trx/transfer`
- `GET  /trx/history`

---

## 💡 Best Practice yang Digunakan

- Clean Architecture
- Repository Pattern
- Dependency Injection
- Transaction-safe wallet update
- Integer-based money handling (hindari float)

---

## 👨‍💻 Author

**Hosea Nainggolan**
Backend Developer (Go / Node.js)

---
