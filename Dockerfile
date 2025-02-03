# Using official image Go
FROM golang:1.23-alpine

# Create working directory
WORKDIR /app

# Copy code to image
COPY . .

# Build the app
RUN go build -o go-Hello-kub .

# Open port 8082
EXPOSE 8082

# Run the app
CMD ["./go-Hello-kub"]