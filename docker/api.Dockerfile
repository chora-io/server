FROM golang:1.20

# set working directory
WORKDIR /home

# copy source code
COPY . .

# install api command
RUN go install ./cmd/api
