# gRPC API with Go and gRPC-gateway

![gRPC](https://img.shields.io/badge/gRPC-1.0-blue?style=flat-square)
![Go](https://img.shields.io/badge/Go-1.21+-blue?style=flat-square)
![Protocol Buffers](https://img.shields.io/badge/Protocol_Buffers-3+-blue?style=flat-square)

A production-ready gRPC microservice with RESTful JSON API gateway for an e-commerce order system, implemented in Go.

## Features

- **gRPC Service**: High-performance order service using HTTP/2 and Protocol Buffers
- **REST Gateway**: gRPC-gateway for JSON/HTTP compatibility
- **CRUD Operations**: Full order management (Create, Read, Update, Delete)
- **Docker Support**: Containerized deployment
- **Koyeb Ready**: Pre-configured for cloud deployment

## Architecture

```mermaid
graph LR
    Client-->|REST/JSON| Gateway
    Gateway-->|gRPC| OrderService
    OrderService-->|In-Memory| Database


###Prerequisites

- Go 1.21+

- Protocol Buffer compiler (protoc) v3+

- Docker (optional)

- Koyeb account (for deployment)

##Q#uick Start

1. Clone the repository

    git clone https://github.com/your-repo/go-grpc.git

    cd example-go-grpc-gateway

2. Install dependencies

    make deps

3. Generate protobuf code

    make protoc

4. Run services locally

    docker-compose up -d

### API Endpoints

    Method	   Path	                Description
    POST	   /v0/orders	        Create new order
    GET	       /v0/orders/{id}	    Get order by ID
    PUT	       /v0/orders	        Update order
    DELETE	   /v0/orders/{id}	    Delete order