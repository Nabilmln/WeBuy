# Architecture Diagram & Monorepo Folder Structure

## 1. Diagram Arsitektur Sistem

```mermaid
flowchart TB
    Client["React Frontend (Vite)"]

    subgraph Gateway["Entry Point (opsional, fase lanjut)"]
        GW["Reverse Proxy / API Gateway"]
    end

    subgraph Services["Microservices (Go)"]
        US["User Service<br/>(Auth, Profil)"]
        PS["Product Catalog Service<br/>(Produk, Kategori, Stok)"]
        CS["Cart Service<br/>(Keranjang)"]
        OS["Order Service<br/>(Order, Status)"]
        PayS["Payment Service<br/>(Integrasi Midtrans)"]
        NS["Notification Service<br/>(Email/SMS - fase lanjut)"]
    end

    subgraph Databases["PostgreSQL per Service (Docker)"]
        DBUser[("user_db")]
        DBProduct[("product_db")]
        DBCart[("cart_db")]
        DBOrder[("order_db")]
        DBPayment[("payment_db")]
    end

    Midtrans["Midtrans<br/>(Payment Gateway - Sandbox)"]

    Client -->|REST API + JWT| GW
    GW --> US
    GW --> PS
    GW --> CS
    GW --> OS
    GW --> PayS

    US --> DBUser
    PS --> DBProduct
    CS --> DBCart
    OS --> DBOrder
    PayS --> DBPayment

    OS -->|create order| PayS
    PayS -->|request Snap token| Midtrans
    Midtrans -->|webhook notification| PayS
    PayS -->|update status| OS
    OS -->|trigger notif, fase lanjut| NS
```

**Catatan alur:**
- Semua request dari Client melewati JWT yang diterbitkan oleh **User Service** dan diverifikasi oleh service lain.
- **API Gateway** bersifat opsional di awal — bisa langsung expose tiap service lewat port berbeda selama development, baru ditambahkan reverse proxy saat mendekati deployment.
- Tiap service punya database PostgreSQL sendiri (database-per-service), dijalankan sebagai container terpisah lewat Docker Compose.
- Payment Service adalah satu-satunya service yang berkomunikasi langsung dengan Midtrans (request Snap token & menerima webhook).

---

## 2. Struktur Folder Monorepo

```
ecommerce-platform/
├── apps/
│   └── web/                        # Frontend React (Vite)
│       ├── src/
│       │   ├── components/
│       │   ├── pages/
│       │   ├── context/             # Auth context, dsb
│       │   ├── hooks/
│       │   ├── services/            # axios/fetch wrapper ke tiap microservice
│       │   ├── store/                # Zustand store (jika dipakai)
│       │   ├── routes/
│       │   └── App.tsx
│       ├── index.html
│       ├── package.json
│       └── vite.config.ts
│
├── services/
│   ├── user-service/
│   │   ├── cmd/
│   │   │   └── main.go
│   │   ├── internal/
│   │   │   ├── handler/
│   │   │   ├── service/
│   │   │   ├── repository/
│   │   │   └── model/
│   │   ├── migrations/
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   ├── product-service/
│   │   ├── cmd/main.go
│   │   ├── internal/{handler,service,repository,model}/
│   │   ├── migrations/
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   ├── cart-service/
│   │   ├── cmd/main.go
│   │   ├── internal/{handler,service,repository,model}/
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   ├── order-service/
│   │   ├── cmd/main.go
│   │   ├── internal/{handler,service,repository,model}/
│   │   ├── migrations/
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   ├── payment-service/
│   │   ├── cmd/main.go
│   │   ├── internal/
│   │   │   ├── handler/
│   │   │   ├── service/
│   │   │   ├── repository/
│   │   │   ├── model/
│   │   │   └── midtrans/            # wrapper client Midtrans
│   │   ├── migrations/
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   └── notification-service/        # fase lanjut
│       ├── cmd/main.go
│       ├── internal/{handler,service}/
│       ├── Dockerfile
│       └── go.mod
│
├── docker-compose.yml               # orkestrasi semua service + database
├── .env.example
├── docs/
│   ├── 01-project-requirements.md
│   └── 02-architecture-and-folder-structure.md
└── README.md
```

**Prinsip struktur:**
- Setiap service Go punya `go.mod` sendiri (independent module), bukan satu `go.mod` besar — supaya benar-benar independen sesuai prinsip microservices.
- Folder `internal/` di tiap service mengikuti pola `handler → service → repository → model` yang konsisten di semua service, sehingga begitu paham satu service, service lain lebih mudah diikuti.
- `docker-compose.yml` di root menjadi satu titik untuk menjalankan seluruh sistem (`docker compose up`) saat development.
- Frontend berada di `apps/web` — dipisah jelas dari `services/` agar monorepo tetap terorganisir walau nanti bertambah service lain.
