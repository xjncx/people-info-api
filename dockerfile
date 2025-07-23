FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/api/main.go

# Устанавливаем точку входа
CMD ["./server"]
