# Web Service

Simple Web Server 

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

*  GoLang >= 1.12
*  Modules: See go.mod file

### Installing Go Module

Install modules (modules may auto-download when running 'go run *.go')
```
go mod download
```

Verify modules
```
go mod verify
```

### Installing redis and run (Mac)
```
brew install redis

brew services start redis
```

## Running the tests
```
cd /Testing
go test -v
```

## Running on localhost
```
go run app.go
```
