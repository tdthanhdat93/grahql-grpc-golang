# graphql & grpc with go pratice

Pratice graphql api endpoint & grpc microservices with golang

## Setup env

**MakeFile**
install Chocolatey https://chocolatey.org/install
```sh
choco install make
make all
```

**Docker**
```sh
brew install docker --cask
```

## Development

**Starting database**

```sh
docker-compose up
```

**Proto gen**

***Install*** `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
```sh
protoc -I=proto proto/*.proto --go_out=:pb --go-grpc_out=:pb
```

## Testing

**1.Running the grpc service**

***Plane***
```sh
go run grpc/planeService/main.go
```

***Flight***
```sh
go run grpc/flightService/main.go
```

**2.Running the api**
```sh
go run clients/main.go
```

Reference graphql query command: test/graphqlQuery.txt
