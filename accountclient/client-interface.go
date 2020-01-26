package accountclient

import "github.com/diversemix/accountapi-client/accountclient/entities"

// ClientInterface  The interface for the client
type ClientInterface interface {
	Create(account *entities.Account) (*entities.Account, error)
	Fetch(string) (*entities.Account, error)
	List(opts *PaginationOptions) ([]entities.Account, error)
	Delete(string) (bool, error)
}
