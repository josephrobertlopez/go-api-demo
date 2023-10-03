# Golang API Presentation Demo

This Golang API presentation demo showcases a simple RESTful API built using the Gorilla Mux router. The API provides three endpoints for health checking, retrieving data, and posting data.

## Getting Started

### Prerequisites

- Golang installed on your system
- Docker installed (if you plan to use the provided Dockerfile)

### Installation

1. Clone this repository to your local machine.

```bash
git clone <repository_url>
cd <repository_directory>
```

2. Build and run the API:
```bash
go run cmd/main.go
```

The API will start running on http://localhost:8080.

## Endpoints

### Health Check

- Endpoint: /health
- Method: GET
- Description: Check the health of the API.
- Response: Returns a 200 OK response with the message "OK".
- Get Data

### GET Data

- Endpoint: /get
- Method: GET
- Description: Retrieve stored data.
- Response: Returns a 200 OK response with the stored data.

### Post Data

- Endpoint: /post
- Method: POST
- Description: Store data in the API.
- Request Body: JSON data to be stored.
- Response: Returns a 200 OK response with the message "Data received and stored".

## Docker Support
You can also run the API using Docker. A Docker image is provided for your convenience.

1. Build the Docker img:
```bash
docker build -t api-demo .
```
2. Run the Docker container:
```bash
docker run -p 8080:8080 api-demo
```

The API will be accessible at http://localhost:8080.

## OpenAPI Documentation
The API documentation can be found in the swagger.yaml file. You can use this Swagger file to generate API documentation or test the API endpoints using tools like [Swagger UI](https://editor.swagger.io/) or Postman.

## Server Scaffolding
If you want to generate a server scaffolding for your API based on the Swagger documentation, you can use [go-swagger](https://github.com/go-swagger/go-swagger).
```bash
swagger generate server -f swagger.yaml