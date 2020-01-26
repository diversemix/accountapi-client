package main

/*
Example of using the Client for the Account API using REST

Ref: https://api-docs.form3.tech/api.html#organisation-accounts
*/
import (
	"log"
	"os"

	"github.com/diversemix/form3interview/accountclient"
	"github.com/diversemix/form3interview/accountclient/entities"
)

func main() {
	log.Println("Create a custom logger")
	logger := log.New(log.Writer(), "[client] ", log.Flags())

	logger.Println("Create our new client, based on the logger and the type of repository transport.")
	client, err := accountclient.NewRestClient(logger, "http://localhost:8080")
	if err != nil {
		logger.Printf("Could not the client %+v\n", err)
		os.Exit(1)
	}

	logger.Println("Create an account entity...")
	orgID := "826b2428-b6b4-4a21-895b-3d26c75bf342"
	newAccount, err := entities.CreateAccount(orgID, "GB", "NWBKGB22", "400302")
	if err != nil {
		logger.Printf("Could not create entity %+v\n", err)
		os.Exit(1)
	}

	logger.Println("Create the Account in the Repository")
	a, err := client.Create(newAccount)
	if err != nil {
		logger.Printf("Could not create in the Repository %+v\n", err)
		os.Exit(1)
	}

	logger.Println("Fetch the newly created Account")
	newID := a.ID
	a, err = client.Fetch(newID)
	if err != nil {
		logger.Printf("Could not get the new Account %+v\n", err)
		os.Exit(1)
	}
	logger.Printf("%+v\n", *a)

	logger.Println("List all Accounts")

	// Example if using the PaginationOptions
	// opts := accountclient.PaginationOptions{
	// 	PageNumber: 1,
	// 	PageSize:   2,
	// }
	// aList, err := client.List(&opts)

	aList, err := client.List(nil)
	if err != nil {
		logger.Printf("Could not get the list of Accounts %+v\n", err)
		os.Exit(1)
	}

	for _, account := range aList {
		logger.Printf("AccountID: %+v\n", account.ID)
	}

	logger.Printf("Listed %d Accounts\n", len(aList))

	logger.Printf("Delete the newly created Account: %s\n", newID)
	deleted, err := client.Delete(newID)
	if err != nil {
		logger.Printf("Could not delete the new Account %+v\n", err)
		os.Exit(1)
	}
	logger.Printf("New Account deleted = %t\n", deleted)

	logger.Printf("~~~~~ Example Complete ~~~~~")
}
