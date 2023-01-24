package main

import (
    "fmt"
    "os"
    "context"
    "strconv"
    "math/big"
    "crypto/ecdsa"

    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "register/api"
)
func main() {
        if (len(os.Args) >= 5) {
                domain          := os.Args[1]
                ip_addr         := os.Args[2]
                seclvl, err     := strconv.ParseInt(os.Args[3], 10, 64)
                devtype         := os.Args[4]

                client, err := ethclient.Dial("http://192.168.1.6:8545")
                if err != nil {
                        fmt.Println("Failed at Dial")
                }
                privateKey, err := crypto.HexToECDSA("08146781e60cbe2d26ee78d8c450fbbc10151394e5784ce89d02d09af115541b")
                if err != nil {
                        panic(err)
                }

                publicKey := privateKey.Public()
                publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
                if !ok {
                        panic("invalid key")
                }

                fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
                fmt.Println(fromAddress)
                nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
                if err != nil {
                        panic(err)
                }
  
                chainID, err := client.ChainID(context.Background())
                if err != nil {
                        panic(err)
                }

                auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
                if err != nil {
                        panic(err)
                }           
                auth.Nonce = big.NewInt(int64(nonce))
                auth.Value = big.NewInt(0)      // in wei
                auth.GasLimit = uint64(3000000) // in units
                auth.GasPrice = big.NewInt(1000000)

                conn, err := api.NewApi(common.HexToAddress("0x5f336BFb63868F48377AbDfBE284DA70E93Ad7A8"), client)
                if err != nil {
                        panic(err)
                }
                result, err := conn.CreateRecord(auth, domain, ip_addr, big.NewInt(seclvl), devtype)
                if err != nil {
                        panic(err)
                }
                fmt.Println(result)
                fmt.Println("Registered Device!")
        }
}
