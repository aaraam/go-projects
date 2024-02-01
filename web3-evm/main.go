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
	client, err := pkg.ConnectToRPC()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(client)
}
