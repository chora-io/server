FROM golang:1.20

# set working directory
WORKDIR /home

# copy root go mod and sum
COPY go.mod /home
COPY go.sum /home

# copy cmd source code (api command)
COPY cmd /home/cmd

# copy db source code (local dependency)
COPY db /home/db

# create api directory
RUN mkdir api

# copy api go mod and sum
COPY api/go.mod /home/api
COPY api/go.sum /home/api

# set working directory
WORKDIR /home/api

# download go modules
RUN go mod download

# copy api source code
COPY api .

# set working directory
WORKDIR /home
