package main

import (
	"fmt"

	"github.com/jarzamendia/keycloaker/keycloak"
)

func main() {

	clients := keycloak.GetClients("http://127.0.0.1:8080", "Jarzamendia", "senha123", "Ms")

	for _, client := range clients {

		fmt.Println(client.ClientID)

	}

}
