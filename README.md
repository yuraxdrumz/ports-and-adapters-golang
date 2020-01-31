# Golang Starter Kit

A starter kit written in Golang + Ports and Adapters structure on top of https://github.com/golang-standards/project-layout

## Getting Started
  1. git clone `git@github.com/yuraxdrumz/golang-starter-kit`
  2. chmod +x `./scripts/hooks.sh`
  3. chmod +x `./hooks/pre-commit.sh`
  4. run `./scripts/hooks.sh` to add `pre-commit` hook with linter and import checker
  5. run `go mod init` to initialize your module
  6. run `go get -u golang.org/x/lint/golint`
  7. run `go get -u golang.org/x/tools/cmd/goimports`
  
### Prerequisites

Ports and Adapters divides your code to 3 parts:
  - Business-logic - These are your business rules + types, implemented without any dependency on 3rd party modules (self-contained)
  - Ports - Interfaces to speak with your business rules
  - Adapters - Implementations of the ports, There are two kinds of adapters:
    - In(Driver) - your external API to the world. For example - `internal/pkg/adapters/in/http.go`
    - Out(Driven) - what your business logic uses. For example - `internal/pkg/adapters/out/reverser/in-memory.go`

Usually, you divide `ports` and `adapters` to separate directories, but the best practice in golang is to keep structs near implementations. That is why, I decided to add `ports.go` near each adapter.

Adding a new business logic:
  1. Create appropriate structs in `ports.go` file under `internal/app/your-use-case/ports.go`
  2. Create your use-case with your application specific logic under `internal/app/your-use-case/logic.go`
  3. Create your in/out adapter, for example - `Repository(out)` or `gRPC(in)`. `internal/pkg/adapters/*`
  4. Tests!!! `your-file-name_test.go` under same directory as file `internal/app/your-use-case/logic_test.go`

### Installing Packages
```golang
go get -u <repository-name>
```
### Running the service
```golang
go run main.go
```

### Environment Variables
To read environement variables, this service uses `envconfig` library, which allows defining a struct of environment variables.

The default is:
```golang
type Specification struct {
	LogzioToken string `envconfig:"LOGZIO_TOKEN"`
	AppName     string `envconfig:"APP_NAME" default:"example-app"`
	LogLevel    string `envconfig:"LOG_LEVEL" default:"info"`
	Port        string `envConfig:"PORT" default:"8080"`
}
```

### Logging

The repo uses `logrus` logger to write to stdout and has and option to write to `logz.io`. In order to send logs to `logz.io` you will need an environment variable called `LOGZIO_TOKEN` which is your `logz.io` token

To change log level change the `LOG_LEVEL` environment variable

Possible log levels:
  * info - default
  * debug
  * error
  * warn
  * fatal

### Building the service
The repo compiles to all 3 major os's - `windows`, `linux`, `darwin(macOS)`

To compile the project, run `./scripts/build.sh`.
This will create a build directory with the binaries both zipped and unzipped for all os's

### Running the service + watch mode
```golang
go get -u github.com/cespare/reflex
reflex -d none -s -R vendor. -r \.go$ -- go run main.go
```

### Pre-Commit
Will run on every pre commit and check if all `.go` files are valid
