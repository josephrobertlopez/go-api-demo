# Use an official Go runtime as a parent image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary files into the container
COPY ./cmd /app/cmd
COPY go.mod /app/
# If importing from outside of stl a go.sum file will be generated
COPY go.sum /app/

# Build the Go app
RUN go mod download
RUN go build -o main ./cmd

# Expose port 8080
EXPOSE 8080

# Run the Go app
CMD ["./main"]

