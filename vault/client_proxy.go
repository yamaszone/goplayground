package main

import (
        "fmt"
        "github.com/daveadams/onthelambda"
        "log"
        "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
        client, err := onthelambda.VaultClient()
        if err != nil {
                log.Fatalf("ERROR: %s", err)
        }
        fmt.Fprintf(w, client.Token())
}

func main() {
        http.HandleFunc("/", handler)
        log.Fatal(http.ListenAndServe(":8080", nil))
}
