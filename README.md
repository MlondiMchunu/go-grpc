## gRPC Order Service with Go and gRPC-Gateway

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)
![gRPC](https://img.shields.io/badge/gRPC-1.0-4285F4?style=for-the-badge&logo=google)
![Protocol Buffers](https://img.shields.io/badge/Protocol_Buffers-3+-3178C6?style=for-the-badge&logo=protobuf)

A production-ready e-commerce order service with:
- gRPC microservice (HTTP/2 + Protobuf)
- RESTful JSON API gateway
- Dockerized deployment
- Koyeb cloud-ready configuration

#### 📂 Project Structure

    .
    ├── cmd/
    │   ├── server/        # gRPC server main
    │   └── client/        # Gateway server main
    ├── internal/
    │   ├── orderservice.go # Service implementation
    │   └── db.go          # Database layer
    ├── proto/             # Protocol Buffer definitions
    ├── protogen/          # Generated code
    └── docker-compose.yml # Local development


### Architecture

```mermaid
graph LR
    A[Client] -->|REST/JSON| B[Gateway]
    B -->|gRPC| C[OrderService]
    C -->|In-Memory| D[Database]
```

### 🚀 Quick Start


#### Clone the repository

    git clone https://github.com/MlondiMchunu/go-grpc.git

    cd go-grpc

#### Install dependencies

    make deps

#### Generate protobuf code

    make protoc

#### Run services

    docker-compose up -d

#### 🌐 API Endpoints

    Method	            Endpoint	        Description
    POST	            /v0/orders	        Create a new order
    GET	                /v0/orders/{id}	    Get order details
    PUT	                /v0/orders	        Update an existing order
    DELETE	            /v0/orders/{id}	    Remove an order


#### 🛠️ Development

- Build binaries

      make build
    
- Run tests

      make test
    
- Regenerate protobuf files

      make protoc


#### 📦 Deployment

1. Order Service

        docker build --target orders-service -t orders-service .


2. Gateway Service

        docker build --target gateway-service -t gateway-service .

