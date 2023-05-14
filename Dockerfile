FROM golang:1.19

# Download all dependencies from go.mod before copying the rest of the source files. This
# ensures that downloading the dependencies (which takes a long time) doesn't happen each
# time any of the source files change. It will only repeat if there is a change to go.mod.
COPY api/go.mod /home
COPY api/go.sum /home
RUN cd /home && go mod download

WORKDIR /home

COPY . /home

RUN cd /home/api/cmd && go build -o /bin/api
