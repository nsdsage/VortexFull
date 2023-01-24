#!/bin/bash

./script.exp

sed -i "s/\/\/\sversion: \"0.5.1\"\,/version: \"\^0.8.0\"\,/" /tmp/geth/contracts/ens/truffle-config.js
sed -i "s/\/\/ development/development/" /tmp/geth/contracts/ens/truffle-config.js
sed -i "s/\/\/  host\: \"127.0.0.1\"/ host: \"127.0.0.1\"/" /tmp/geth/contracts/ens/truffle-config.js
sed -i "s/\/\/  port\: 8545/ port\: 8545/" /tmp/geth/contracts/ens/truffle-config.js
sed -i "s/\/\/  network_id: \"\*\"/ network_id: \"\*\",\r\n     from: \"1x1\",\r\n     gas: 1000000,\r\n    },/" /tmp/geth/contracts/ens/truffle-config.js