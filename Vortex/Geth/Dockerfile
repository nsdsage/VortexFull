FROM ethereum/client-go:latest

ARG GETH_PASSWORD=${GETH_PASSWORD}

EXPOSE 8545 30303 30304 3000 8000 8545-8550 7545

RUN apk update --no-cache \
    && apk upgrade --no-cache \
    && apk add --no-cache --upgrade bash \
    && apk add go \
    && apk add jq expect nodejs npm yarn git python3 make g++

RUN npm install -g solc web3 truffle

WORKDIR /root
COPY ./assets/script.exp /root/
COPY ./assets/*.sh /root/
RUN chmod +x /root/*.sh

RUN echo ${GETH_PASSWORD} > /root/geth.password
COPY ./assets/cliqueGenesis.json /root/genesis.json

# RUN /root/builder.sh

# RUN geth --datadir /root/Vortex/node0/data init /root/genesis.json

# ENTRYPOINT ["/root/launch.sh"]
ENTRYPOINT ["tail", "-f", "/dev/null"]