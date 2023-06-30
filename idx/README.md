# idx

## Prerequisites

- [Go >= 1.20](https://golang.org/doc/install)

## Installation

```
go install github.com/choraio/server/cmd/idx@latest
```

## Usage

```
idx help
```

## Example

```
DATABASE_URL=postgres://postgres:password@localhost:5432/postgres?sslmode=disable idx localhost:9090 chora-local 1
```
