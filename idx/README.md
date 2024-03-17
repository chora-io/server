# idx

## Prerequisites

- [Go >= 1.21](https://golang.org/doc/install)

## Installation

```
go install github.com/chora-io/server/cmd/idx@latest
```

## Usage

```
idx help
```

## Example

```
DATABASE_URL=postgres://postgres:password@localhost:5432/server?sslmode=disable idx localhost:9090 chora-local --start-block 1
```
