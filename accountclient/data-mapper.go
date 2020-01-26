package accountclient

import (
	"encoding/json"

	"github.com/diversemix/form3interview/accountclient/entities"
)

type payloadAccount struct {
	Data entities.Account `json:"data"`
}

type payloadAccountArray struct {
	Data []entities.Account `json:"data"`
}

// AccountToJSON  Converts an Account entity to Json
func AccountToJSON(a *entities.Account) ([]byte, error) {
	p := payloadAccount{}
	p.Data = *a

	bytes, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// JSONToAccountArray  Converts a byte slice to an Array of Accounts
func JSONToAccountArray(data []byte) ([]entities.Account, error) {
	a := &payloadAccountArray{}
	jsonErr := json.Unmarshal(data, a)

	if jsonErr != nil {
		return nil, jsonErr
	}
	return a.Data, nil
}

// JSONToAccount  Converts a byte slice to an Account
func JSONToAccount(data []byte) (*entities.Account, error) {
	a := &payloadAccount{}
	jsonErr := json.Unmarshal(data, a)

	if jsonErr != nil {
		return nil, jsonErr
	}
	return &a.Data, nil
}
