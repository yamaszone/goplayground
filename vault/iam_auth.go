package main

import (
        "fmt"
        "github.com/daveadams/onthelambda"
        "log"
)

func main() {
        client, err := onthelambda.VaultClient()
        if err != nil {
                log.Fatalf("ERROR: %s", err)
        }

        fmt.Println(client.Token())
}
