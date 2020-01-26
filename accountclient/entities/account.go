package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

/*
	File contents is the implementation for an Organisation-Account

	An Account represents a bank account that is registered with Form3. It is used to validate and allocate inbound payments.

	Ref: https://api-docs.form3.tech/api.html#organisation-accounts
*/

// Attributes  for an account (note minimum implementation)
type Attributes struct {
	AlternativeBankAccountNames []string `json:"alternative_bank_account_names"`
	BankID                      string   `json:"bank_id"`
	Bic                         string   `json:"bic"`
	Country                     string   `json:"country"`
	// TODO: add other fields from the specification above
}

// Account  This is referenced
type Account struct {
	Attributes     `json:"attributes"`
	CreatedOn      time.Time `json:"created_on"`
	ID             string    `json:"id"`
	ModifiedOn     time.Time `json:"modified_on"`
	OrganisationID string    `json:"organisation_id"`
	Type           string    `json:"type"`
	Version        int       `json:"version"`
}

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

// CreateAccount  static function for creating a default account entity.
func CreateAccount(OrganisationID string, country string, bic string, bankID string) (*Account, error) {
	if !isValidUUID(OrganisationID) {
		return nil, errors.New("OrganisationID must be a UUID v4")
	}
	a := Account{}
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	a.OrganisationID = OrganisationID
	a.ID = uuid.String()
	a.Type = "accounts"
	a.Version = 0

	// TODO: (if time) validate the following
	a.Country = country
	a.Bic = bic
	a.BankID = bankID
	return &a, nil
}
