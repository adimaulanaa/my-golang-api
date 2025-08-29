package main

import (
    "log"
    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
    connStr := "postgres://postgres:mypassword@localhost/my_go_api?sslmode=disable" // Ganti mypassword
    m, err := migrate.New("file://migrations", connStr)
    if err != nil {
        log.Fatal("Gagal init migration:", err)
    }
    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatal("Gagal jalankan migration:", err)
    }
    log.Println("Migration selesai")
}