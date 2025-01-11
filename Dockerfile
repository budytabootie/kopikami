# Menggunakan Go versi terbaru yang sesuai dengan go.mod
FROM golang:1.23.1-alpine

# Set working directory
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Menginstal dependensi
RUN go mod tidy

# Build aplikasi
RUN go build -o main .

# Expose port yang digunakan oleh aplikasi
EXPOSE 8080

# Menjalankan aplikasi
CMD ["./main"]
