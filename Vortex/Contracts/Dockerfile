FROM ubuntu:latest

RUN apt-get -y update && \
        apt-get -y dist-upgrade && \
        apt-get -y upgrade

RUN apt-get install -y software-properties-common && \
      add-apt-repository -y ppa:ethereum/ethereum && \
      # add-apt-repository ppa:ethereum/ethereum-dev && \
      # apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 2A518C819BE37D2C2031944D1C52189C923F6CA9 && \
      apt-get -y update

RUN apt-get install -y jq expect nodejs yarn git solc npm golang

RUN npm install -g web3 truffle

RUN mkdir -p /opt/geth/contracts/ens/ && \
      mkdir -p /opt/geth/contracts/ens/build

COPY ./assets/* /opt/geth/contracts/ens/

COPY ./assets/truffleBuild.sh /opt/geth/contracts/ens/truffleBuild.sh
RUN chmod +x /opt/geth/contracts/ens/truffleBuild.sh

RUN go get -u github.com/ethereum/go-ethereum && \
      cd /root/go/src/github.com/ethereum/go-ethereum/cmd/abigen/ && \
      go build . && \
      ln -s /root/go/src/github.com/ethereum/go-ethereum/cmd/abigen/abigen /usr/bin/abigen

RUN cd /opt/geth/contracts/ens/contracts/registry/ && \
      solc --abi ENSRegistry.sol -o /opt/geth/contracts/ens/build/ && \
      solc --bin ENSRegistry.sol -o /opt/geth/contracts/ens/build/

RUN cd /opt/geth/contracts/ens/build/ && \
      abigen --abi=ENSRegistry.abi --pkg=ens --out=ENSRegistry.sol && \
      abigen --bin=ENSRegistry.bin --abi=ENSRegistry.abi --pkg=ens --out=ENSRegistry.go && \
      mkdir -p /tmp/geth/contracts/ens/

RUN cp /opt/geth/contracts/ens/build/* /tmp/geth/contracts/ens/

ENTRYPOINT ["tail", "-f", "/dev/null"]