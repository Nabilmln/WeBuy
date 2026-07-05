# Project Requirements — Scalable E-Commerce Platform (React + Golang)

Dokumen ini berisi seluruh komponen dan teknologi yang disepakati untuk pengembangan project e-commerce berbasis microservices, sebagai referensi selama development.

---

## 1. Ringkasan Project

Platform e-commerce berbasis **microservices architecture**, menggunakan **Go** untuk backend dan **React** untuk frontend, dengan **Docker** sebagai environment containerization. Payment gateway menggunakan **Midtrans** (khusus market Indonesia).

Tujuan ganda:
- Membangun produk e-commerce fungsional (untuk portofolio)
- Sarana belajar React dan Golang secara bersamaan

---

## 2. Frontend

| Komponen | Pilihan | Alasan |
|---|---|---|
| Build tool | **Vite + React** | Lebih ringan dari Next.js, fokus murni ke konsep React tanpa SSR/App Router yang mengalihkan fokus belajar |
| Styling | **Tailwind CSS** | Konsisten dengan stack yang sudah dikuasai (thesis project) |
| Routing | **React Router v6/v7** | Standar routing untuk SPA |
| State management (lokal/global sederhana) | **useState / useContext** | Untuk auth state & cart state di tahap awal |
| State management (lanjutan, jika diperlukan) | **Zustand** | Alternatif ringan dari Redux, dipakai jika Context API mulai terasa berat (prop drilling) |
| Data fetching & caching | **TanStack Query (React Query)** | Penting untuk konsumsi banyak microservices — caching, loading state, auto refetch |
| Form handling | **React Hook Form** | Digunakan pada form checkout, register, login |

### Urutan belajar React yang disarankan
1. Component & Props
2. useState / useEffect
3. Context API (auth global)
4. React Query (integrasi ke backend Go)
5. Zustand (opsional, jika diperlukan)
6. React Hook Form (form checkout)

---

## 3. Backend

| Komponen | Pilihan | Alasan |
|---|---|---|
| Bahasa | **Go (Golang)** | Ringan, cepat, concurrency native (goroutines) cocok untuk microservices |
| Web framework | **Gin** atau **Fiber** | REST API ringan dan populer di ekosistem Go |
| ORM / DB access | **GORM** (mirip Sequelize/Prisma) atau **sqlc** (lebih type-safe, kontrol query manual) | Sesuai preferensi tingkat kontrol terhadap query |
| Arsitektur kode per service | `handler → service → repository` | Clean architecture sederhana, konsisten di semua service |
| Komunikasi antar service | **REST API** (gRPC opsional di tahap lanjut) | Lebih sederhana untuk tahap awal |
| Autentikasi | **JWT**, diterbitkan oleh User Service, diverifikasi oleh service lain | Standar auth terdesentralisasi untuk microservices |

### Strategi belajar
Stabilkan dulu satu service (contoh: User Service) sampai terhubung end-to-end dengan frontend React, baru replikasi pattern-nya ke service lain.

---

## 4. Core Microservices (MVP Scope)

| Service | Tanggung Jawab |
|---|---|
| **User Service** | Registrasi, login, profil, role (admin/customer) |
| **Product Catalog Service** | CRUD produk, kategori, stok |
| **Shopping Cart Service** | Tambah/hapus/update item keranjang |
| **Order Service** | Membuat order, tracking status, riwayat order |
| **Payment Service** | Integrasi Midtrans, webhook status pembayaran |
| **Notification Service** *(fase lanjutan)* | Email/SMS notifikasi status order |

> Payment & Notification Service dapat masuk ke fase 2 pengembangan, tidak wajib di MVP awal.

---

## 5. Database

| Komponen | Pilihan | Alasan |
|---|---|---|
| DBMS | **PostgreSQL** | Relational, mendukung constraint kuat untuk data transaksional (order, payment) |
| Strategi | **Database-per-service** | Standar microservices — tiap service punya database sendiri, dijalankan sebagai container Docker terpisah |
| Environment | **Docker container** (bukan Laragon) | Supaya environment local sama persis dengan environment production; Laragon lebih cocok untuk stack PHP tanpa Docker |

---

## 6. Payment Gateway — Midtrans

- Gunakan **Snap API** (checkout page siap pakai) untuk mempercepat development, dibanding Core API yang butuh UI pembayaran custom.
- Gunakan **Sandbox environment** Midtrans selama development/testing.
- Flow dasar:
  1. Order Service membuat order
  2. Payment Service request Snap token ke Midtrans
  3. Frontend redirect/embed Snap
  4. Midtrans mengirim **webhook notification** ke Payment Service saat status berubah
  5. Payment Service update status, notify Order Service
- SDK Go: `midtrans/midtrans-go` (community/semi-official) atau call langsung REST API Midtrans dengan `net/http`.

---

## 7. Infrastruktur & DevOps

| Komponen | Pilihan | Catatan |
|---|---|---|
| Containerization | **Docker** | Dockerfile per microservice |
| Orkestrasi lokal | **Docker Compose** | Menjalankan seluruh service + database sekaligus di local |
| Docker environment (Windows) | **Docker Desktop dengan WSL2 backend** | Gratis untuk personal/education use; pastikan WSL2 aktif (bukan Hyper-V) untuk performa lebih ringan |
| Resource management | Batasi CPU/RAM Docker via `.wslconfig` | Mencegah laptop kehabisan resource saat banyak container jalan bersamaan |
| API Gateway | *(opsional, fase lanjut)* Kong/Traefik/NGINX | Bisa di-skip di awal, langsung expose service atau reverse proxy sederhana |
| Service Discovery | *(skip di awal)* Consul/Eureka | Overengineering untuk skala kecil, hanya relevan jika jumlah service sudah banyak |
| Monitoring/Logging | *(opsional, fase lanjut)* Prometheus + Grafana, ELK stack | Tidak wajib untuk MVP |
| CI/CD | *(fase lanjut)* GitHub Actions | Otomasi build/test/deploy setelah MVP stabil |

---

## 8. Deployment (Rencana)

| Komponen | Platform |
|---|---|
| Frontend | Vercel |
| Backend services | Railway / Render |
| Database | Managed PostgreSQL (Railway/Render addon) atau Supabase |

---

## 9. Roadmap Fase Pengembangan

1. **Fase 1**: Auth (User Service) + Product Catalog (CRUD) + Frontend dasar (Vite + React + Tailwind)
2. **Fase 2**: Cart Service + Order Service
3. **Fase 3**: Payment integration (Midtrans sandbox)
4. **Fase 4**: Notification Service, deployment, polish (README, demo)
