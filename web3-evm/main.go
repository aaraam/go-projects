package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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

	switch os.Args[1] {
	case "balance":
		handleBalance(balanceCmd, balanceAddress)
	default:
		fmt.Println("expected 'balance' subcommand")
		os.Exit(1)
	}
}
