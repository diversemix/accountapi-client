package accountclient

import (
	"strings"
	"testing"

	"github.com/diversemix/accountapi-client/accountclient/entities"
)

func TestNew(t *testing.T) {

	t.Run("succeeds with logger and repository set", func(t *testing.T) {
		l := newMockLogger()
		rr := &RestRepository{}
		result, err := New(l, rr)
		if err != nil {
			t.Error("New() should not fail")
		}
		if result == nil {
			t.Error("New() should return a value")
		}
	})

	t.Run("rejects nil logger", func(t *testing.T) {
		rr := &RestRepository{}
		_, err := New(nil, rr)
		if err == nil {
			t.Error("New() should fail and result should be nil")
		}

		if strings.Compare(err.Error(), "ClientLogger cannot be nil") != 0 {
			t.Errorf("Unexpected error: %+v", err.Error())
		}
	})

	t.Run("rejects nil repo", func(t *testing.T) {
		l := newMockLogger()
		_, err := New(l, nil)
		if err == nil {
			t.Error("New() should fail and result should be nil")
		}

		if strings.Compare(err.Error(), "Repository cannot be nil") != 0 {
			t.Errorf("Unexpected error: %+v", err.Error())
		}
	})
}

func TestNewRestClient(t *testing.T) {
	t.Run("succeeds with logger set", func(t *testing.T) {
		l := newMockLogger()
		result, err := NewRestClient(l, "test-url")
		if err != nil {
			t.Error("TestNewRestClient() should not fail")
		}
		if result == nil {
			t.Error("TestNewRestClient() should return a value")
		}
	})

	t.Run("rejects nil logger", func(t *testing.T) {
		_, err := NewRestClient(nil, "")
		if err == nil {
			t.Error("TestNewRestClient() should fail and result should be nil")
		}

		if strings.Compare(err.Error(), "ClientLogger cannot be nil") != 0 {
			t.Errorf("Unexpected error: %+v", err.Error())
		}
	})
}

func getClient(t *testing.T) ClientInterface {
	l := newMockLogger()
	r := newMockRepository()

	c, err := New(l, r)
	if err != nil {
		t.Error("TestNewRestClient() should not fail")
	}
	return c
}

func Test_accountclient_List(t *testing.T) {
	t.Run("succeeds with nil options", func(t *testing.T) {
		c := getClient(t)
		result, err := c.List(nil)
		if err != nil {
			t.Errorf("List() should not fail, failed with: %+v", err)
		}

		if result == nil {
			t.Error("result returned from List() is nil")
		}
	})
}

func Test_accountclient_Delete(t *testing.T) {
	t.Run("succeeds with a valid UUID", func(t *testing.T) {
		c := getClient(t)
		result, err := c.Delete("213e6bd9-74e8-4ec8-bf78-512a0bdca080")
		if err != nil {
			t.Errorf("Delete() should not fail, failed with: %+v", err)
		}

		if result != true {
			t.Error("expecting result from Delete() to be true")
		}
	})
}

func Test_accountclient_Create(t *testing.T) {
	t.Run("succeeds with a valid Account", func(t *testing.T) {
		c := getClient(t)
		a, err := entities.CreateAccount("99825362-9fc8-4ff6-b5ce-a064854b1268", "", "", "")
		if err != nil {
			t.Errorf("CreateAccount() should not fail, failed with: %+v", err)
		}

		if a == nil {
			t.Error("result returned from CreateAccount() is nil")
		}

		result, err := c.Create(a)

		if err != nil {
			t.Errorf("Create() should not fail, failed with: %+v", err)
		}

		if result.ID != "99825362-9fc8-4ff6-b5ce-a064854b1268" {
			t.Errorf("result returned from Create() is not expected %+v\n", result)
		}
	})
}

func Test_accountclient_Fetch(t *testing.T) {
	t.Run("succeeds with a valid UUID", func(t *testing.T) {
		c := getClient(t)
		result, err := c.Fetch("99825362-9fc8-4ff6-b5ce-a064854b1268")
		if err != nil {
			t.Errorf("Fetch() should not fail, failed with: %+v", err)
		}

		if result.ID != "99825362-9fc8-4ff6-b5ce-a064854b1268" {
			t.Errorf("result returned from Create() is not expected %+v\n", result)
		}
	})
}
