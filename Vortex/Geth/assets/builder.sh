#!/bin/bash

if [[ ! -d "/root/Vortex/" ]];
    then mkdir /root/Vortex/;
fi

if [[ ! -d "/root/Vortex/node0" ]]; 
    then mkdir /root/Vortex/node0; 
fi

if [[ -d "/root/Vortex/node0/data" ]];
    then 
        geth --datadir "/root/Vortex/node0/data" account new --password /root/geth.password;
        apk add go;
        go install github.com/ethereum/go-ethereum/cmd/puppeth@latest;
        /root/addSealer.sh
        /root/automatePuppeth.sh
    else
        geth account import --password /root/geth.password;
fi
