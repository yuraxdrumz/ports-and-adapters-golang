# Ports And Adapters Golang

An example project of Golang + Ports and Adapters structure on top of https://github.com/golang-standards/project-layout and https://github.com/yuraxdrumz/golang-starter-kit

## Getting Started
  1. git clone `git@github.com/yuraxdrumz/ports-and-adapters-golang`
  2. Create database folder project's root for sqlite adapter
  2. go mod init
  3. go run main.go
  
### Prerequisites

Ports and Adapters divides your code to 3 parts:
  - Business-logic - These are your business rules + types, implemented without any dependency on 3rd party modules (self-contained)
  - Ports - Interfaces to speak with your business rules
  - Adapters - Implementations of the ports, There are two kinds of adapters:
    - In(Driver) - your external API to the world. For example - `internal/pkg/adapters/in/http.go`
    - Out(Driven) - what your business logic uses. For example - `internal/pkg/adapters/out/reverser/in-memory.go`

Usually, you divide `ports` and `adapters` to separate directories, but the best practice in golang is to keep structs near implementations. That is why, I decided to add `ports.go` near each adapter.

Adding a new business logic:
  1. Create appropriate ports in `ports.go` file under `internal/app/your-use-case/ports.go`
  2. Create your use-case with your application specific logic under `internal/app/your-use-case/logic.go`
  3. Create your in/out adapter, for example - `Repository(out)` or `gRPC(in)`. `internal/pkg/adapters/*`
  4. Tests!!! `your-file-name_test.go` under same directory as file `internal/app/your-use-case/logic_test.go`