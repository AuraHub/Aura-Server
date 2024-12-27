# Step 1: Use an official Golang image as a builder
FROM golang:1.23 as builder

# Step 2: Set the working directory in the container
WORKDIR /app

# Step 3: Copy go mod and sum files
COPY go.mod go.sum ./

# Step 4: Download dependencies
RUN go mod tidy && go mod download

# Step 5: Copy the source code
COPY . .

# Step 6: Build the application
RUN go build -o main.go .

# Step 7: Use a minimal image for the final container
FROM alpine:latest

# Step 8: Install certificates
RUN apk --no-cache add ca-certificates

# Step 9: Set the working directory
WORKDIR /root/

# Step 10: Copy the binary from the builder
COPY --from=builder /app/main .

# Step 11: Expose port 8080
EXPOSE 3000

# Step 12: Run the binary
CMD ["./main"]
