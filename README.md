# IMP Assesent

IMP Assesent backend service using golang

## Specifications

- DB Driver: MySQL
- ORM: GORM.IO
- Context Framework: Echo

## Requirements

- Install Golang v1.20
- Install [wire](https://github.com/google/wire) for Dependency Injection setup.
- Install go-migrate for Migration Database

## Migration files

Each migration has an migration's file.

```bash
1481574547-users.sql
```

Create File Migration
 
 ```console
$ migrate create -ext sql -dir ./migrations/ {TABLE_NAME}
 ```

Run File Migration
 
 ```console
$ migrate -database "$(url)" -path ./migrations/ up $(version) 
 ```

## Instalation

We are using [Echo](https://github.com/labstack/echo) for Web Framework to build our APIs

Install all Libraries from go mod

```
go get .
```

create `.env` file

```
cp .env.example .env
```

## Running App

Just run

```
go run .
```
