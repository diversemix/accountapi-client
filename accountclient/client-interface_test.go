package accountclient

import (
	"testing"

	"github.com/diversemix/accountapi-client/accountclient/entities"
)

type goodClientInterface struct{}

func (c *goodClientInterface) Create(account *entities.Account) (*entities.Account, error) {
	return nil, nil
}
func (c *goodClientInterface) Fetch(string) (*entities.Account, error) {
	return nil, nil
}
func (c *goodClientInterface) List(opts *PaginationOptions) ([]entities.Account, error) {
	return nil, nil
}
func (c *goodClientInterface) Delete(string) (bool, error) {
	return false, nil
}

func isClientInterface(c *goodClientInterface) ClientInterface {
	return c
}

func TestClientInterface(t *testing.T) {

	t.Run("interface has required members", func(t *testing.T) {
		x := goodClientInterface{}
		if isClientInterface(&x) == nil {
			t.Error("Expecting ClientLogger")
		}
		return
	})
}
