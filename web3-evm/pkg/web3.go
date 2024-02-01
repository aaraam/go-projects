package pkg

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

// go function to connect to rpc
func ConnectToRPC() (*ethclient.Client, error) {
	uri := os.Getenv("RPC_URL")
	fmt.Println(uri)
	client, err := ethclient.Dial(uri)
	if err != nil {
		return nil, err
	}
	fmt.Println("we have a connection")
	return client, nil
}
