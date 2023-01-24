# Compilation

## ABI & BIN Generations

'''
solc --optimize --abi bdnsRegistry.sol -o build
solc --optimize --bin bdnsRegistry.sol -o build
'''

## Generate GOLANG API

'''
abigen --abi=./build/bdnsRegistry.abi --bin=./build/bdnsRegistry.bin --pkg=api --out=./api/bdnsRegistry.go
'''

curl -X POST --data '{"id":1,"jsonrpc":2.0,"method":"eth_blockNumber","params":["test.vortex.mil", "127.0.0.100", 2, "test"]}' http://127.0.0.1:8545/08146781e60cbe2d26ee78d8c450fbbc10151394e5784ce89d02d09af115541b


curl -v -d '{"jsonrpc": "2.0","method": "invokefunction","params": ["<scripthash>","Main",[{"type":"String","value":"createRecord"},{"type":"Array", "value":[{"type":"String","value":"example123"},{"type":"String","value":"00"},{"type":"String","value":"00"},{"type":"String","value":"00"}]}]],"id": 1}' http://127.0.0.1:8545 -H 'Content-Type: application/json'