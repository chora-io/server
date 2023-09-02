FROM golang:1.20

# install dependencies
RUN apt-get update
RUN apt-get install jq -y

# set version and chain
ENV GIT_CHECKOUT='main'

# clone chora repository
RUN git clone https://github.com/choraio/chora /home/chora

# set working directory
WORKDIR /home/chora

# check out provided version
RUN git checkout $GIT_CHECKOUT

# build chora binary
RUN make install

# initialize validator node
RUN chora --chain-id chora-local init validator

# set chora binary configuration
RUN chora config chain-id chora-local
RUN chora config keyring-backend test

# update stake to uchora
RUN sed -i "s/stake/uchora/g" /root/.chora/config/genesis.json

# add validator and user accounts
RUN printf "trouble alarm laptop turn call stem lend brown play planet grocery survey smooth seed describe hood praise whale smile repeat dry sauce front future\n\n" | chora keys --keyring-backend test add validator -i
RUN printf "cool trust waste core unusual report duck amazing fault juice wish century across ghost cigar diary correct draw glimpse face crush rapid quit equip\n\n" | chora keys --keyring-backend test add user1 -i
RUN printf "music debris chicken erode flag law demise over fall always put bounce ring school dumb ivory spin saddle ostrich better seminar heart beach kingdom\n\n" | chora keys --keyring-backend test add user2 -i

# add validator to genesis
RUN chora add-genesis-account validator 1000000000uchora --keyring-backend test
RUN chora gentx validator 1000000uchora

# add test users to genesis
RUN chora add-genesis-account user1 1000000000uchora --keyring-backend test
RUN chora add-genesis-account user2 1000000000uchora --keyring-backend test

# prepare genesis file
RUN chora collect-gentxs

# set minimum gas price
RUN sed -i "s/minimum-gas-prices = \"\"/minimum-gas-prices = \"0uchora\"/" /root/.chora/config/app.toml

# set cors allow all origins
RUN sed -i "s/cors_allowed_origins = \[\]/cors_allowed_origins = [\"*\"]/" /root/.chora/config/config.toml

# copy chora start script
COPY docker/scripts/chora_start.sh /home/chora/scripts/

# make start script executable
RUN ["chmod", "+x", "/home/chora/scripts/chora_start.sh"]
