.PHONY: all run hot-reload

# Default target
all: run

# 🔥 Jalankan dengan hot reload (butuh nodemon)
hot-reload:
	nodemon --exec "go run cmd/main.go" --signal SIGTERM --ext go,mod --watch .

# 🚀 Jalankan aplikasi biasa
run:
	go run cmd/main.go

