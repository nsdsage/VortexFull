# CREATE NEW ACCOUNT WITH PASSWORD

'''
geth --datadir "/opt/ethPoA/node0/data" account new --password
'''

### ATTACH TO GET

'''
geth attach http://127.0.0.1:8545
'''


## GET ACCOUNTS

'''
eth.accounts
'''

## GET BALANCE

'''
eth.getBalance("8b78Bf87fa46D75317718F2c9AbB761031e2445E")
'''

## UNLOCK ACCOUNT

'''
personal.unlockAccount("0x8b78bf87fa46d75317718f2c9abb761031e2445e")
'''

## SEND ETHER

'''
eth.sendTransaction({from:"0x8b78bf87fa46d75317718f2c9abb761031e2445e" to:"0x07860e21e57e677907330e0f7cedb52d0c02438d", value: web3.toWei(3.0., "ether")})
'''

## CHECK BALANCE FROM CLI

'''
curl -s -X POST --data '{"jsonrpc":"2.0", "method":"eth_getBalance", "params":["0x09059c234934215ac27a7d7e2790e155cc6b93ed", "latest"], "id":1}' http://localhost:8545 -H 'Content-Type: application/json'
'''
