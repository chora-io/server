FROM golang:1.22

# set chora version
ENV GIT_CHECKOUT='main'

# install dependencies
RUN apt-get update
RUN apt-get install jq -y

# clone chora repository
RUN git clone https://github.com/chora-io/chora /home/chora

# set working directory
WORKDIR /home/chora

# check out provided version
RUN git checkout $GIT_CHECKOUT

# build chora binary
RUN make install

# initialize validator node
RUN chora init validator --chain-id chora-local

# set chora binary configuration
RUN chora config set client chain-id chora-local
RUN chora config set client keyring-backend test

# update stake to uchora
RUN sed -i "s/stake/uchora/g" /root/.chora/config/genesis.json

# add local validator and test user accounts
RUN printf "trouble alarm laptop turn call stem lend brown play planet grocery survey smooth seed describe hood praise whale smile repeat dry sauce front future\n\n" | chora keys --keyring-backend test add validator -i
RUN printf "cool trust waste core unusual report duck amazing fault juice wish century across ghost cigar diary correct draw glimpse face crush rapid quit equip\n\n" | chora keys --keyring-backend test add user1 -i
RUN printf "music debris chicken erode flag law demise over fall always put bounce ring school dumb ivory spin saddle ostrich better seminar heart beach kingdom\n\n" | chora keys --keyring-backend test add user2 -i

# add local validator account to genesis
RUN chora genesis add-genesis-account validator 1000000000uchora --keyring-backend test

# add local test user accounts to genesis
RUN chora genesis add-genesis-account user1 1000000000uchora --keyring-backend test
RUN chora genesis add-genesis-account user2 1000000000uchora --keyring-backend test

# add chora coop user accounts to genesis
RUN chora genesis add-genesis-account chora1jx34255cgvxpthkg572ma6rhq6crwl6xh7g0md 1000000000uchora
RUN chora genesis add-genesis-account chora19gayv6d9zx27hs4gfk344wxf3mdx9s2w3vqvr0 1000000000uchora
RUN chora genesis add-genesis-account chora1k50tgf3jcjjaa6dz63tml7fy8n7wl8v44rpd60 1000000000uchora
RUN chora genesis add-genesis-account chora1eute2zprd6w0y9apcm6e2gj5exkj3jgzazf7rj 1000000000uchora

# generate validator genesis transaction
RUN chora genesis gentx validator 1000000uchora

# prepare genesis file
RUN chora genesis collect-gentxs

# set minimum gas price
RUN sed -i "s/minimum-gas-prices = \"\"/minimum-gas-prices = \"0uchora\"/" /root/.chora/config/app.toml

# set cors allow all origins
RUN sed -i "s/cors_allowed_origins = \[\]/cors_allowed_origins = [\"*\"]/" /root/.chora/config/config.toml

# copy genesis state files
COPY docker/data /home/chora/data

# add authz state to genesis
RUN jq '.app_state.authz |= . + input' /root/.chora/config/genesis.json /home/chora/data/chora_authz.json > genesis-tmp-1.json

# add feegrant state to genesis
RUN jq '.app_state.feegrant |= . + input' genesis-tmp-1.json /home/chora/data/chora_feegrant.json > genesis-tmp-2.json

# add geonode state to genesis
RUN jq '.app_state.geonode |= . + input' genesis-tmp-2.json /home/chora/data/chora_geonode.json > genesis-tmp-3.json

# add group state to genesis
RUN jq '.app_state.group |= . + input' genesis-tmp-3.json /home/chora/data/chora_group.json > genesis-tmp-4.json

# add voucher state to genesis
RUN jq '.app_state.voucher |= . + input' genesis-tmp-4.json /home/chora/data/chora_voucher.json > genesis-tmp-5.json

# overwrite genesis file with updated genesis file
RUN mv -f genesis-tmp-5.json /root/.chora/config/genesis.json

# copy chora start script
COPY docker/scripts/chora_start.sh /home/chora/scripts/

# make start script executable
RUN ["chmod", "+x", "/home/chora/scripts/chora_start.sh"]
