package accountclient

import (
	"errors"

	"github.com/diversemix/accountapi-client/accountclient/entities"
	"github.com/google/uuid"
)

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

type accountclient struct {
	logger ClientLogger
	repo   Repository
}

// New  Creates a new implementation of the client and returns interface
func New(logger ClientLogger, repo Repository) (ClientInterface, error) {
	if logger == nil {
		return nil, errors.New("ClientLogger cannot be nil")
	}
	if repo == nil {
		return nil, errors.New("Repository cannot be nil")
	}
	return &accountclient{
		logger: logger,
		repo:   repo,
	}, nil
}

// NewRestClient  Creates a new Rest implementation of the client based on the URL
func NewRestClient(logger ClientLogger, baseURL string) (ClientInterface, error) {
	repo := &RestRepository{
		BaseURL: baseURL,
	}

	return New(logger, repo)
}

func (s accountclient) List(opts *PaginationOptions) ([]entities.Account, error) {
	resp, err := s.repo.GetAll(opts)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	return JSONToAccountArray(resp.GetData())
}

func (s accountclient) Delete(id string) (bool, error) {
	if !isValidUUID(id) {
		s.logger.Println("Delete() was attempted with: " + id)
		return false, errors.New("Not a valid UUID")
	}
	resp, err := s.repo.Delete(id)
	if err != nil {
		s.logger.Println(resp)
		s.logger.Println(err)
		return false, err
	}

	return true, nil
}

func (s accountclient) Create(a *entities.Account) (*entities.Account, error) {
	if a == nil {
		s.logger.Println("Create was attempts with a nil object")
		return nil, errors.New("nil passed into Create()")
	}
	payload, err := AccountToJSON(a)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	resp, err := s.repo.Create(payload)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	return JSONToAccount(resp.GetData())
}

func (s accountclient) Fetch(id string) (*entities.Account, error) {
	resp, err := s.repo.Get(id)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	return JSONToAccount(resp.GetData())
}
