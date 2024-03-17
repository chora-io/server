FROM golang:1.21

# set working directory
WORKDIR /home

# copy source code
COPY . .

# download go modules
RUN go mod download

# install idx command
RUN go install ./cmd/idx
