package pkg

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// function to connect to rpc
func ConnectToRPC() (*ethclient.Client, error) {
	uri := os.Getenv("RPC_URL")
	client, err := ethclient.Dial(uri)
	if err != nil {
		return nil, err
	}
	fmt.Println("connected to sepolia")
	return client, nil
}

// function to get balances
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
	// balance = balance.Div(balance, big.NewInt(1000000000000000000))
	return balance, nil
}

// function to send transactions
func SendTransaction(fromPrivKey string, toAddress string, amount *big.Int) (string, error) {
	client, err := ConnectToRPC()
	if err != nil {
		return "", err
	}

	// Convert the private key string to a private key object
	privateKey, err := crypto.HexToECDSA(fromPrivKey)
	if err != nil {
		return "", err
	}

	// Derive the sender's address from the private key
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// Get the nonce for the sender's account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	// Set up transaction parameters
	gasLimit := uint64(21000) // Standard gas limit for a transfer
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	// Create the transaction
	tx := types.NewTransaction(nonce, common.HexToAddress(toAddress), amount, gasLimit, gasPrice, nil)

	// Sign the transaction with the private key
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	// Send the transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}
