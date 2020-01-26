package entities

import (
	"testing"

	"github.com/google/uuid"
)

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func createTestAccount(t *testing.T, orgID string) *Account {
	got, err := CreateAccount(orgID, "", "", "")
	if err != nil {
		t.Errorf("Not expecting error = %+v", err)
	}
	return got
}

func TestCreateAccount(t *testing.T) {
	type args struct {
		OrganisationID string
	}
	testOrgID := "32590a9f-b9d9-4121-8e80-22fb2717641e"

	t.Run("succeeds with valid UUID", func(t *testing.T) {
		got := createTestAccount(t, testOrgID)
		if !IsValidUUID(got.ID) {
			t.Errorf("Expecting ID to be a UUID, got = %+v", got.ID)
		}
	})

	t.Run("succeeds with valid type", func(t *testing.T) {
		got := createTestAccount(t, testOrgID)
		if got.Type != "accounts" {
			t.Errorf("Expecting type to be a account, got = %+v", got.Type)
		}

	})

	t.Run("succeeds with valid organisation id", func(t *testing.T) {
		got := createTestAccount(t, testOrgID)
		if got.OrganisationID != testOrgID {
			t.Errorf("Expecting OrganisationID to be a %+v, got = %+v", testOrgID, got.OrganisationID)
		}
	})

	t.Run("fails with empty OrganisationID", func(t *testing.T) {
		got, err := CreateAccount("", "", "", "")
		if err == nil {
			t.Error("Expecting an error")
		} else {
			if err.Error() != "OrganisationID must be a UUID v4" {
				t.Errorf("wrong error= %v", err)
			}
		}
		if got != nil {
			t.Error("Expecting account to be nil")
		}
	})

	t.Run("fails with bad OrganisationID", func(t *testing.T) {
		got, err := CreateAccount("this should not work invalid uuid!!!", "", "", "")
		if err == nil {
			t.Error("Expecting an error")
		} else {
			if err.Error() != "OrganisationID must be a UUID v4" {
				t.Errorf("wrong error= %v", err)
			}
		}
		if got != nil {
			t.Error("Expecting account to be nil")
		}
	})
}
