# Toko Buku - Fullstack App (React + Go + PostgreSQL)

Aplikasi Toko Buku sederhana yang dibangun menggunakan arsitektur fullstack modern. Data buku tersimpan secara permanen di database PostgreSQL yang berjalan di dalam Docker, diakses menggunakan API Backend berbahasa Go (Gin), dan ditampilkan menggunakan Frontend React (Vite + TypeScript).

---

## 🌟 Fitur Utama (CRUD)
* **Read (Tampil Buku)**: Menampilkan kartu buku lengkap dengan judul, penulis, harga, dan stok yang diambil langsung dari database.
* **Create (Tambah Buku)**: Menambahkan buku baru ke database melalui formulir input dinamis secara real-time.
* **Update (Beli Buku)**: Mengurangi stok buku secara langsung di database saat tombol "Beli" diklik (data persisten saat refresh).
* **Delete (Hapus Buku)**: Menghapus buku secara permanen dari database dan langsung memperbarui tampilan tanpa reload halaman.

---

## 🛠️ Tech Stack
* **Frontend**: React 19, TypeScript, Vite, React Compiler, Vanilla CSS.
* **Backend**: Go (Golang) 1.22+, Gin Web Framework, `database/sql` standard library, `lib/pq` (driver Postgres).
* **Database**: PostgreSQL 16 (berjalan di dalam Docker Compose).
* **Configuration**: `godotenv` untuk mengelola environment variables secara aman.

---

## 📊 Alur Kerja Aplikasi

```mermaid
graph TD
    subgraph Frontend (React - Port 5173/5174)
        App[App.tsx] -->|Render| BookCard[BookCard.tsx]
        App -->|Submit Form| Form[Form Input]
        App -->|API Request| Fetch[Fetch API]
    end

    subgraph Backend (Go - Port 8080)
        Fetch -->|GET/POST/PUT/DELETE| Gin[Gin Router]
        Gin -->|CORS Check| Cors[CORS Middleware]
        Cors -->|SQL Query| DB[Postgres Docker]
    end
```

---

## 🚀 Panduan Menjalankan Proyek

Pastikan laptop Anda sudah terinstall **Docker**, **Go**, dan **Node.js**.

### LINGKUNGAN DATABASE (Docker)
1. Masuk ke PostgreSQL lokal kamu dan buat databasenya sekali saja:
   ```bash
   sudo -u postgres psql
   # Di dalam psql run:
   CREATE DATABASE toko_buku;
   CREATE USER awi WITH PASSWORD 'awi';
   GRANT ALL PRIVILEGES ON DATABASE toko_buku TO awi;
   \q
   ```
2. Pastikan servis Postgres lokal laptop mati agar tidak bentrok port:
   ```bash
   sudo systemctl stop postgresql
   ```

### 1. Setup & Jalankan Backend (Go)
1. Masuk ke folder backend:
   ```bash
   cd backend
   ```
2. Buat file `.env` di dalam folder `backend/` dan sesuaikan nilainya:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=awi
   DB_PASSWORD=awi
   DB_NAME=toko_buku
   ```
3. Jalankan Postgres di Docker:
   ```bash
   docker compose up -d
   ```
4. Jalankan server Go:
   ```bash
   go run main.go
   ```
   *Server Go akan berjalan di `http://localhost:8080` dan otomatis melakukan inisialisasi tabel serta memasukkan data awal jika database masih kosong.*

### 2. Setup & Jalankan Frontend (React)
1. Buka terminal baru di root folder proyek (`toko_buku/`).
2. Install dependensi package Node:
   ```bash
   npm install
   ```
3. Jalankan aplikasi Vite:
   ```bash
   npm run dev
   ```
4. Buka alamat localhost yang tertera di terminal (biasanya `http://localhost:5173`) di browser Anda.

---

## 🔒 Catatan Keamanan
File `backend/.env` berisi password rahasia database kamu dan **tidak akan ikut ter-upload ke GitHub** karena sudah didaftarkan di dalam file `.gitignore`.
