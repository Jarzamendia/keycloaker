package keycloak

import (
	"log"

	"github.com/Nerzal/gocloak"
)

//GetUsers asd
func GetUsers(url string, username string, password string, realm string) []*gocloak.User {

	client := gocloak.NewClient(url)

	token, err := client.LoginAdmin(username, password, realm)

	if err != nil {

		log.Panic(err)

	}

	if token != nil {

	}

	users, err := client.GetUsers(token.AccessToken, realm, gocloak.GetUsersParams{
		Email: "@",
	})

	//clients, err := client.GetClient((token.AccessToken, realm, "clientID" )

	return users

}

//GetClients asd
func GetClients(url string, username string, password string, realm string) []*gocloak.Client {

	client := gocloak.NewClient(url)

	token, err := client.LoginAdmin(username, password, realm)

	if err != nil {

		log.Panic(err)

	}

	if token != nil {

	}

	clients, err := client.GetClients(token.AccessToken, realm, gocloak.GetClientsParams{
		ClientID: "",
	})

	if err != nil {

		log.Println(err)

	}

	return clients

}
