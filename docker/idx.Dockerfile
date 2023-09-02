FROM golang:1.20

# set working directory
WORKDIR /home

# copy root go mod and sum
COPY go.mod /home
COPY go.sum /home

# copy cmd source code (idx command)
COPY cmd /home/cmd

# copy db source code (local dependency)
COPY db /home/db

# create idx directory
RUN mkdir idx

# copy idx go mod and sum
COPY idx/go.mod /home/idx
COPY idx/go.sum /home/idx

# set working directory
WORKDIR /home/idx

# download go modules
RUN go mod download

# copy idx source code
COPY idx .

# set working directory
WORKDIR /home
