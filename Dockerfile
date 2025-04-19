# مرحله اول: build
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy

# مرحله دوم: اجرا
FROM golang:1.21
WORKDIR /app
COPY --from=builder /app /app
CMD ["go", "run", "main.go"]