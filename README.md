# HCF (Highest Common Factor) 

This Service just to run search Highest Common Factor

## Clone Project

```bash
https://github.com/bayusm0506/hcf-or-gcd.git
cd /hcf-or-gcd
```

## Structure Folder

```
├── app
│   ├── app.go
│   ├── controllers
│   │   ├── hcf.go
│   │   └── response.go
│   └── models
│       └── hcf.go
├── config
│   └── config.go
├── go.mod
├── go.sum
├── main.go
```

## Installation

```bash
# Install mux router
go get -u github.com/gorilla/mux

# Install godotenv
go get -u github.com/joho/godotenv
```

## Usage

```go
# create file .env in root folder and add variable
APPS_PORT=2021

# running project
go run main.go

# or if you wanna run project with .exe file
go build
./hcf-or-gcd
```

## API

#### CREATE A TEAM
```sh
POST : /api/hcf
```
```json
JSON : 
{
    "number1": 20,
    "number2": 25
}
```
```
#### APP INFO

#### Author originally
Bayu Setra Maulana [bayusm.com]

Postman Collection : https://www.getpostman.com/collections/e6ac7bc82c04bdd3688d
```
