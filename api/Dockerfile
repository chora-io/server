FROM golang:1.20

WORKDIR /home

RUN mkdir api

COPY api/go.mod /home/api
COPY api/go.sum /home/api

COPY db /home/db

WORKDIR /home/api

RUN go mod download

COPY . /home/api
