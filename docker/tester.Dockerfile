FROM golang:1.22

# install dependencies
RUN apt-get update
RUN apt-get install jq libpq-dev postgresql-client -y

# set version and chain
ENV GIT_CHECKOUT='main'

# clone chora repository
RUN git clone https://github.com/chora-io/chora /home/tester

# set working directory
WORKDIR /home/tester

# check out provided version
RUN git checkout $GIT_CHECKOUT

# build chora binary
RUN make install

# set chora binary configuration
RUN chora config chain-id chora-local
RUN chora config keyring-backend test

# add user test accounts
RUN printf "cool trust waste core unusual report duck amazing fault juice wish century across ghost cigar diary correct draw glimpse face crush rapid quit equip\n\n" | chora keys --keyring-backend test add user1 -i
RUN printf "music debris chicken erode flag law demise over fall always put bounce ring school dumb ivory spin saddle ostrich better seminar heart beach kingdom\n\n" | chora keys --keyring-backend test add user2 -i

# copy tester start script
COPY docker/scripts/tester_start.sh /home/tester/scripts/

# copy tester test scripts
COPY docker/tester/ /home/tester/scripts/

# make start script executable
RUN ["chmod", "+x", "/home/tester/scripts/tester_start.sh"]

# make test scripts executable
RUN ["chmod", "+x", "/home/tester/scripts/test_idx_group_proposals.sh"]
RUN ["chmod", "+x", "/home/tester/scripts/test_idx_group_votes.sh"]
