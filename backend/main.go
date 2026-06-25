package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Buku struct {
	ID      int    `json:"id"`
	Judul   string `json:"judul"`
	Penulis string `json:"penulis"`
	Harga   int    `json:"harga"`
	Stok    int    `json:"stok"`
}

var db *sql.DB

func main() {
	//load file env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env File")
	}
	//akses variabel dengan standar library
	constStr := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=disable"
	db, err = sql.Open("postgres", constStr)

	if err != nil {
		log.Fatal("Gagal membuka koneksi database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Gagal Terhubung (ping) ke Postgres Docker: %v", err)
	}

	log.Println("Berhasil terhubung ke database Postgres!")
	initDB()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))
	r.GET("/books", func(ctx *gin.Context) {
		rows, err := db.Query("SELECT id, judul, penulis, harga, stok FROM buku")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()
		var daftarBuku []Buku

		for rows.Next() {
			var b Buku
			err := rows.Scan(&b.ID, &b.Judul, &b.Penulis, &b.Harga, &b.Stok)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			daftarBuku = append(daftarBuku, b)
		}
		ctx.JSON(http.StatusOK, daftarBuku)
	})
	r.DELETE("/books/:id", func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		result, err := db.Exec(`DELETE FROM buku WHERE id = $1`, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Buku Tidak di temukan"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Buku Berhasil dihapus",
			"id":      id,
		})
	})
	r.POST("/books", func(ctx *gin.Context) {
		var bukuBaru Buku
		err := ctx.ShouldBindJSON(&bukuBaru)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		query := `INSERT INTO buku (judul, penulis, harga, stok) VALUES ($1, $2, $3, $4) RETURNING id`
		err = db.QueryRow(query, bukuBaru.Judul, bukuBaru.Penulis, bukuBaru.Harga, bukuBaru.Stok).Scan(&bukuBaru.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, bukuBaru)
	})
	r.PUT("/books/:id", func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var bukuUpdate Buku
		err = ctx.ShouldBindJSON(&bukuUpdate)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		query := `UPDATE buku SET judul=$1, penulis=$2, harga=$3, stok=$4 WHERE id=$5`
		result, err := db.Exec(query, bukuUpdate.Judul, bukuUpdate.Penulis, bukuUpdate.Harga, bukuUpdate.Stok, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Buku Tidak Ditemukan"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Buku Berhasil di Update",
			"id":      id,
		})
	})
	r.Run(":8080")
}

func initDB() {
	queryBuatTabel := `
	CREATE TABLE IF NOT EXISTS buku (
		id	SERIAL PRIMARY KEY,
		judul	VARCHAR(255)	NOT NULL,
		penulis	VARCHAR(255)	NOT NULL,
		harga	INT	NOT NULL,
		stok	INT NOT NULL DEFAULT 0
	)`
	_, err := db.Exec(queryBuatTabel)
	if err != nil {
		log.Fatalf("Gagal Membuat tabel: %v", err)
	}
	log.Println("Tabel 'buku' siap.")
	//cek apakah tabel masih kosong
	var count int
	//kita panggil queryrow untuk query yang hanya menghasilkan satu baris data
	//hasilnya kita masukkan ke variabel count menggunakan &count (pointer)
	err = db.QueryRow("SELECT COUNT(*) FROM buku").Scan(&count)
	if err != nil {
		log.Fatalf("Gagal menghitung jumlah data: %v", err)
	}

	//seeding data jika kosong (count == 0)
	if count == 0 {
		querySeed := `
		INSERT INTO buku (judul, penulis, harga, stok) VALUES
			('Bumi Manusia', 'Pramoedya Ananta Toer', 95000, 5),           
			('Laskar Pelangi', 'Andrea Hirata', 79000, 3),
			('Ronggeng Dukuh Paruk', 'Ahmad Tohari', 85000, 8),            
			('Gadis Kretek', 'Ratih Kumala', 68000, 0)`
		_, err = db.Exec(querySeed)
		if err != nil {
			log.Fatalf("Gagal Seeding Data: %v", err)
		}
		log.Println("Data awal berhasil dimasukkan")
	}
}
