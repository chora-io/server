FROM golang:1.19

WORKDIR /home

RUN mkdir idx

COPY idx/go.mod /home/idx
COPY idx/go.sum /home/idx

COPY db /home/db

WORKDIR /home/idx

RUN go mod download

COPY . /home/idx
