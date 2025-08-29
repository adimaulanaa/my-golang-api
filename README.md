# My Go API

API sederhana dengan Go untuk belajar backend menggunakan clean architecture.

## Struktur Proyek
```
my-go-api/
├── cmd/
│   └── api/
│       └── main.go        # Entry point server
├── internal/
│   ├── handlers/
│   │   └── tasks.go       # Handler untuk /tasks
│   ├── models/
│   │   └── task.go        # Struct Task
│   └── services/
│       └── task_service.go # Logika bisnis task
├── go.mod                 # Module definition
└── README.md              # Dokumentasi
```

## Prasyarat
- Go 1.25 (`go version` untuk cek).
- Dependency: `github.com/gorilla/mux` (`go get github.com/gorilla/mux`).
- (Opsional) Air untuk auto-reload: `go install github.com/air-verse/air@latest`.

## Cara Menjalankan
1. Pastikan Go terinstall: `go version`.
2. Masuk ke folder proyek: `cd ~/Documents/my-go-api`.
3. Jalankan `go mod tidy` untuk cek dependency.
4. Jalankan server:
   - Manual: `go run cmd/api/main.go`
   - Auto-reload: `air` (pastikan Air terinstall dan `.air.toml` ada).
5. Test endpoint di browser (`http://localhost:8080`) atau pakai `curl`.

## File Konfigurasi Air
Untuk auto-reload, buat `.air.toml`:
```toml
root = "."
tmp_dir = "tmp"
[build]
cmd = "go build -o ./tmp/main ./cmd/api/main.go"
bin = "./tmp/main"
include_ext = ["go"]
exclude_dir = ["tmp"]
delay = 1000
```

## Model
- `Task`: Struct untuk data task, dengan field:
  - `ID` (int): ID unik task.
  - `Name` (string): Nama task.
  - `IsCompleted` (bool, opsional): Status penyelesaian task.

## Service
- `TaskService`: Menangani logika bisnis task, seperti:
  - Mengambil semua task.
  - Menambah task baru.
  - Memperbarui task berdasarkan ID.
  - Menghapus task berdasarkan ID.

## Handler
- `GetTasks`: Menangani `GET /tasks`.
- `CreateTask`: Menangani `POST /tasks`.
- `UpdateTask`: Menangani `PUT /tasks/{id}`.
- `DeleteTask`: Menangani `DELETE /tasks/{id}`.

## Endpoint
- `GET /`: Menampilkan pesan selamat datang.
  ```bash
  curl http://localhost:8080
  ```
  Output: `"Halo, API Go pertamaku!"`
- `GET /tasks`: Menampilkan daftar task dalam JSON.
  ```bash
  curl http://localhost:8080/tasks
  ```
  Output: `[{"id":1,"name":"Belajar Go"},{"id":2,"name":"Bikin API"}]`
- `POST /tasks`: Menambah task baru (kirim JSON).
  ```bash
  curl -X POST -H "Content-Type: application/json" -d '{"name":"Test Task"}' http://localhost:8080/tasks
  ```
  Output: `{"id":3,"name":"Test Task"}`
- `PUT /tasks/{id}`: Memperbarui task berdasarkan ID.
  ```bash
  curl -X PUT -H "Content-Type: application/json" -d '{"name":"Belajar Go Updated"}' http://localhost:8080/tasks/1
  ```
  Output: `{"id":1,"name":"Belajar Go Updated"}`
- `DELETE /tasks/{id}`: Menghapus task berdasarkan ID.
  ```bash
  curl -X DELETE http://localhost:8080/tasks/1
  ```
  Output: (kosong, status 204)