package pkg

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// go function to connect to rpc
func ConnectToRPC() (*ethclient.Client, error) {
	uri := os.Getenv("RPC_URL")
	client, err := ethclient.Dial(uri)
	if err != nil {
		return nil, err
	}
	fmt.Println("connected to sepolia")
	return client, nil
}

func GetBalance(address string) (*big.Int, error) {
	client, err := ConnectToRPC()
	if err != nil {
		return nil, err
	}
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		return nil, err
	}
	// convert balance to ether
	balance = balance.Div(balance, big.NewInt(1000000000000000000))
	return balance, nil
}
