package main

/*
Useful reference material:

https://data-dog.github.io/godog/
https://github.com/DATA-DOG/godog/tree/master/examples/api

*/
import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cucumber/godog"
	"github.com/diversemix/accountapi-client/accountclient"
	"github.com/diversemix/accountapi-client/accountclient/entities"
	"github.com/google/uuid"
)

type apiFeatureContext struct {
	Client         accountclient.ClientInterface
	Country        string
	Bic            string
	BankID         string
	OrganisationID string
	Account        *entities.Account
	APIAccount     *entities.Account
	knownID        string
	minNumber      int
	numToPage      int
}

func isToday(t time.Time) bool {
	now := time.Now()
	return t.Year() == now.Year() && t.YearDay() == now.YearDay()
}

func (a *apiFeatureContext) theCountryIs(arg1 string) error {
	a.Country = arg1
	return nil
}

func (a *apiFeatureContext) theBicIs(arg1 string) error {
	a.Bic = arg1
	return nil
}

func (a *apiFeatureContext) theBankIDIs(arg1 string) error {
	a.BankID = arg1
	return nil
}

func (a *apiFeatureContext) theOrganisationIDIs(arg1 string) error {
	a.OrganisationID = arg1
	return nil
}

func (a *apiFeatureContext) theAccountShouldBeCreatedInMemory() error {
	account, err := entities.CreateAccount(a.OrganisationID, a.Country, a.Bic, a.BankID)
	if err != nil {
		return fmt.Errorf("Could not create the Account in memory")
	}
	a.Account = account
	return nil
}

func (a *apiFeatureContext) theAccountShouldBeCreatedOverTheAPI() error {
	account, err := a.Client.Create(a.Account)
	if err != nil {
		return fmt.Errorf("Could not create the Account over the API")
	}
	a.APIAccount = account
	return nil
}

func (a *apiFeatureContext) theCreatedOnDateShouldBeToday() error {
	if !isToday(a.APIAccount.CreatedOn) {
		return fmt.Errorf("Expected CreatedOn to be today")
	}
	return nil
}

func (a *apiFeatureContext) theModifiedOnDateShouldBeToday() error {
	if !isToday(a.APIAccount.ModifiedOn) {
		return fmt.Errorf("Expected ModifiedOn to be today")
	}
	return nil
}

func (a *apiFeatureContext) versionShouldBeZero() error {
	if a.APIAccount.Version != 0 {
		return fmt.Errorf("Expected Version to be zero")
	}
	return nil
}

func (a *apiFeatureContext) theOrganisationIDShouldBe(arg1 string) error {
	if a.APIAccount.OrganisationID != "369ace13-f19c-413f-b512-2bbacfacf5a3" {
		return fmt.Errorf("Expected OrganisationID to be 369ace13-f19c-413f-b512-2bbacfacf5a3")
	}
	return nil
}

func (a *apiFeatureContext) theAccountShouldNotBeCreatedOverTheAPI() error {
	account, err := a.Client.Create(a.Account)
	if account != nil {
		return fmt.Errorf("Not expecting the Account to have been created over the API: %+v", account)
	}
	if err == nil {
		return fmt.Errorf("Expecting the error to not be nil")
	}
	return nil
}

func (a *apiFeatureContext) theDeleteCallShouldBeSuccessfull() error {
	result, err := a.Client.Delete(a.knownID)
	if err != nil || result == false {
		return fmt.Errorf("This Account was not deleted: %s", a.knownID)
	}
	return nil
}

func (a *apiFeatureContext) theAccountExistsForKnownID() error {
	a.knownID = ""
	orgID, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("Expected to be able to create a UUID")
	}
	account, err := entities.CreateAccount(orgID.String(), "GB", "NWBKGB22", "400300")
	if err != nil {
		return fmt.Errorf("Could not create the Account in memory")
	}
	newAccount, err := a.Client.Create(account)
	if err != nil {
		return fmt.Errorf("Could not create the Account over the API")
	}
	a.knownID = newAccount.ID
	return nil
}

func (a *apiFeatureContext) theAccountDoesNotExistWithTheKnownID() error {
	_, err := a.Client.Fetch(a.knownID)
	if err == nil {
		return fmt.Errorf("This Account should not exist: %s", a.knownID)
	}
	return nil
}

func (a *apiFeatureContext) theAccountShouldBeSuccessfullyFetched() error {
	if a.knownID == "" {
		return fmt.Errorf("Should have an ID")
	}
	return nil
}

func (a *apiFeatureContext) aRandomID() error {
	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("Expected to be able to create a UUID")
	}
	a.knownID = id.String()
	return nil
}

func (a *apiFeatureContext) thereAreAtLeastAccounts(arg1 int) error {
	a.minNumber = arg1
	for index := 0; index < a.minNumber; index++ {
		a.theAccountExistsForKnownID()
	}
	return nil
}

func (a *apiFeatureContext) theAccountsShouldBeSuccessfullyListed() error {
	accounts, err := a.Client.List(nil)
	if err != nil {
		return fmt.Errorf("There was an error %+v", err)
	}
	if len(accounts) < a.minNumber {
		return fmt.Errorf("Expected more than %d got %d", a.minNumber, len(accounts))
	}
	return nil
}

func (a *apiFeatureContext) onlyAccountsShouldBeListedWithPageSizeOf(arg1, arg2 int) error {
	opts := accountclient.PaginationOptions{
		PageNumber: 0,
		PageSize:   arg2,
	}
	accounts, err := a.Client.List(&opts)
	if err != nil {
		return fmt.Errorf("There was an error %+v", err)
	}
	if len(accounts) != arg1 {
		return fmt.Errorf("Expected exactly %d got %d", arg1, len(accounts))
	}
	return nil
}

func (a *apiFeatureContext) thereAreAnOddNumberOfAccounts() error {
	accounts, err := a.Client.List(nil)
	if err != nil {
		return fmt.Errorf("There was an error %+v", err)
	}
	a.numToPage = len(accounts)
	if a.numToPage%2 == 0 {
		a.theAccountExistsForKnownID()
		a.numToPage++
	}
	return nil
}

func (a *apiFeatureContext) eachPageOfContainsTheExpectedNumber(arg1 int) error {
	opts := accountclient.PaginationOptions{
		PageNumber: 0,
		PageSize:   arg1,
	}

	expectedPagesOfPageSize := a.numToPage / opts.PageSize
	expectedlastPageSize := a.numToPage - expectedPagesOfPageSize*opts.PageSize

	for index := 0; index < expectedPagesOfPageSize; index++ {
		accounts, err := a.Client.List(&opts)
		if err != nil {
			return fmt.Errorf("There was an error %+v", err)
		}
		if len(accounts) != arg1 {
			return fmt.Errorf("Expected exactly %d got %d", arg1, len(accounts))
		}
		opts.PageNumber++
	}
	fmt.Printf(">>>> Paged %d Accounts with a PageSize=%d into %d pages\n", a.numToPage, opts.PageSize, expectedPagesOfPageSize)

	accounts, err := a.Client.List(&opts)
	if err != nil {
		return fmt.Errorf("There was an error %+v on last page", err)
	}

	if len(accounts) != expectedlastPageSize {
		return fmt.Errorf("Expected last page to be %d got %d", expectedlastPageSize, len(accounts))
	}
	fmt.Printf(">>>> Last Page contained %d Accounts\n", expectedlastPageSize)

	return nil

}

func CreateFeatureContext(s *godog.Suite) {
	api := &apiFeatureContext{}
	log.Println("Create a CreateFeatureContext logger")
	logger := log.New(log.Writer(), "[CreateFeatureContext] ", log.Flags())

	logger.Println("Create our new client, based on the logger and the type of repository transport.")
	var err error
	api.Client, err = accountclient.NewRestClient(logger, os.Getenv("URL_UNDER_TEST"))
	if err != nil {
		logger.Println("Could not create the Client, exiting...")
		os.Exit(1)
	}

	// For Create
	s.Step(`^the country is "([^"]*)"$`, api.theCountryIs)
	s.Step(`^the Bic is "([^"]*)"$`, api.theBicIs)
	s.Step(`^the BankID is "([^"]*)"$`, api.theBankIDIs)
	s.Step(`^the OrganisationID is "([^"]*)"$`, api.theOrganisationIDIs)
	s.Step(`^the Account should be created in memory$`, api.theAccountShouldBeCreatedInMemory)
	s.Step(`^the Account should be created over the API$`, api.theAccountShouldBeCreatedOverTheAPI)
	s.Step(`^the CreatedOn date should be today$`, api.theCreatedOnDateShouldBeToday)
	s.Step(`^the ModifiedOn date should be today$`, api.theModifiedOnDateShouldBeToday)
	s.Step(`^Version should be zero$`, api.versionShouldBeZero)
	s.Step(`^the OrganisationID should be "([^"]*)"$`, api.theOrganisationIDShouldBe)
	s.Step(`^the Account should not be created over the API$`, api.theAccountShouldNotBeCreatedOverTheAPI)

	// For Delete
	s.Step(`^the Account exists for known ID$`, api.theAccountExistsForKnownID)
	s.Step(`^the Account does not exist with the known ID$`, api.theAccountDoesNotExistWithTheKnownID)
	s.Step(`^the Delete call should be successfull$`, api.theDeleteCallShouldBeSuccessfull)

	// For Fetch
	s.Step(`^a random ID$`, api.aRandomID)
	s.Step(`^the Account should be successfully fetched$`, api.theAccountShouldBeSuccessfullyFetched)

	// For List
	s.Step(`^there are at least (\d+) Accounts$`, api.thereAreAtLeastAccounts)
	s.Step(`^the Accounts should be successfully listed$`, api.theAccountsShouldBeSuccessfullyListed)
	s.Step(`^only (\d+) Accounts should be listed with page size of (\d+)$`, api.onlyAccountsShouldBeListedWithPageSizeOf)
	s.Step(`^there are an odd number of Accounts$`, api.thereAreAnOddNumberOfAccounts)
	s.Step(`^each page of (\d+) contains the expected number$`, api.eachPageOfContainsTheExpectedNumber)
}
