package main

import (
	"fmt"
	"github.com/daveadams/onthelambda"
	"log"
)

func main() {
	fmt.Println("Attempting authentication...")
	client, err := onthelambda.VaultClient()
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	//resp, _ := client.VaultAuth()
	fmt.Println("Authentication success!")
	fmt.Println(client)
}
