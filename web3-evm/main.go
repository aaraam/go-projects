package main

import (
	"fmt"
	"log"

	"web3-evm/pkg"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	balance, err := pkg.GetBalance("0xC4bFccB1668d6E464F33a76baDD8C8D7D341e04A")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance, "ETH")
}
