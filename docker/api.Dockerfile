FROM golang:1.22

# set working directory
WORKDIR /home

# copy source code
COPY . .

# download go modules
RUN go mod download

# install api command
RUN go install ./cmd/api
