#!/bin/bash

./script.exp

sed -i "s/\/\/\sversion: \"0.5.1\"\,/version: \"\^0.8.0\"\,/" /opt/geth/contracts/ens/truffle-config.js
sed -i "s/\/\/ development/development/" /opt/geth/contracts/ens/truffle-config.js
sed -i "s/\/\/  host\: \"127.0.0.1\"/ host: \"127.0.0.1\"/" /opt/geth/contracts/ens/truffle-config.js
sed -i "s/\/\/  port\: 8545/ port\: 8545/" /opt/geth/contracts/ens/truffle-config.js
sed -i "s/\/\/  network_id: \"\*\"/ network_id: \"\*\",\r\n     from: \"1x1\",\r\n     gas: 1000000,\r\n    },/" /opt/geth/contracts/ens/truffle-config.js