package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"web3-evm/pkg"

	"github.com/joho/godotenv"
)

func handleBalance(balanceCmd *flag.FlagSet, address *string) {
	balanceCmd.Parse(os.Args[2:])

	if *address == "" {
		fmt.Println("address argument is required")
		balanceCmd.PrintDefaults()
		os.Exit(1)
	}

	balance, err := pkg.GetBalance(*address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Balance of %s: %s ETH\n", *address, balance)
}

func handleSend(sendCmd *flag.FlagSet, fromPrivKey *string, toAddress *string, amount *string) {
	sendCmd.Parse(os.Args[2:])

	if *fromPrivKey == "" || *toAddress == "" || *amount == "" {
		fmt.Println("all arguments (fromPrivKey, toAddress, amount) are required")
		sendCmd.PrintDefaults()
		os.Exit(1)
	}

	amountBigInt, ok := new(big.Int).SetString(strings.Trim(*amount, "\""), 10)
	if !ok {
		fmt.Println("invalid amount")
		os.Exit(1)
	}

	txHash, err := pkg.SendTransaction(*fromPrivKey, *toAddress, amountBigInt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction sent! Tx Hash: %s\n", txHash)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	balanceCmd := flag.NewFlagSet("balance", flag.ExitOnError)
	balanceAddress := balanceCmd.String("address", "", "Ethereum address to fetch the balance")

	// Parse the arguments
	if len(os.Args) < 2 {
		fmt.Println("expected 'balance' subcommand")
		os.Exit(1)
	}

	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	fromPrivKey := sendCmd.String("from", "", "Private key of the sender")
	toAddress := sendCmd.String("to", "", "Ethereum address to send to")
	amount := sendCmd.String("amount", "", "Amount of Ether to send (in Wei)")

	switch os.Args[1] {
	case "balance":
		handleBalance(balanceCmd, balanceAddress)
	case "send":
		handleSend(sendCmd, fromPrivKey, toAddress, amount)
	default:
		fmt.Println("expected 'balance' or 'send' subcommand")
		os.Exit(1)
	}
}
