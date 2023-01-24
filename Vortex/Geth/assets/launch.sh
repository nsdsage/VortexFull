#!/bin/bash
geth \
--networkid 999 \
--datadir /root/Vortex/node0/data \
--port 30303 \
--ipcdisable \
--syncmode full \
--http \
--allow-insecure-unlock \
--http.corsdomain "*" \
--http.port 8545 \
--http.addr 0.0.0.0 \
--unlock "0x0" \
--password /root/geth.password \
--mine \
--http.api personal,admin,db,eth,net,web3,miner,shh,txpool,debug,clique \
--ws \
--ws.addr 0.0.0.0 \
--ws.port 8545 \
--ws.origins "*" \
--ws.api personal,admin,db,eth,net,web3,miner,shh,txpool,debug,clique \
--maxpeers 25 \
--miner.etherbase 0x0 \
--miner.gasprice 0