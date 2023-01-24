#!/bin/bash

accounts="$(jq -r .address /root/Vortex/node0/data/keystore/*)"
read -ra sealer <<< ${accounts}
#sealer=${accounts[0]}
sed -i "s/0000000000000000000000000000000000000000/$sealer/g" /root/automatePuppeth.sh
echo $sealer > /root/gethSealer.txt
#echo $sealer
sed -i "s/0x0/0x$(cat /root/gethSealer.txt)/g" /root/launch.sh
echo "Added Sealer"
