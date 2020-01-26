package accountclient

import (
	"strings"
	"testing"

	"github.com/diversemix/form3interview/accountclient/entities"
)

func TestAccountToJSON(t *testing.T) {
	t.Run("contains expected fields", func(t *testing.T) {
		a, err := entities.CreateAccount("ee067e61-b965-4c59-9c2e-16c22afb4192", "GB", "bic", "bankID")
		if err != nil {
			t.Errorf("Not expecting CreateAccount() to error: %+v\n", err)
		}
		result, err := AccountToJSON(a)
		if err != nil {
			t.Errorf("Not expecting AccountToJSON() to error: %+v\n", err)
		}
		str := string(result)

		if !strings.Contains(str, `"organisation_id":"ee067e61-b965-4c59-9c2e-16c22afb4192"`) {
			t.Errorf("Could not find OrgID in output: %s", str)
		}
		if !strings.Contains(str, `"country":"GB"`) {
			t.Errorf("Could not find country in output: %s", str)
		}
		if !strings.Contains(str, `"bic":"bic"`) {
			t.Errorf("Could not find bic in output: %s", str)
		}
		if !strings.Contains(str, `"bank_id":"bankID"`) {
			t.Errorf("Could not find bank_id in output: %s", str)
		}
	})
	// TODO: Add more failure cases
}

func TestJSONToAccountArray(t *testing.T) {
	t.Run("contains expected Account objects", func(t *testing.T) {
		aList, err := JSONToAccountArray([]byte(getGoodAccountArrayJSON()))
		if err != nil {
			t.Errorf("Not expecting JSONToAccountArray() to error: %+v\n", err)
		}
		if len(aList) != 2 {
			t.Errorf("Expecting JSONToAccountArray() length to be 2 not %d\n", len(aList))
		}

		if aList[0].ID != "da12a8b7-ebfa-45ba-bbec-7e3831aa267a" || aList[1].ID != "99825362-9fc8-4ff6-b5ce-a064854b1268" {
			t.Errorf("Not expecting JSONToAccountArray() to return those IDs: %s, %s\n", aList[0].ID, aList[1].ID)
		}
	})
	// TODO: Add more failure cases
}

func TestJSONToAccount(t *testing.T) {
	t.Run("contains expected Account", func(t *testing.T) {
		a, err := JSONToAccount([]byte(getGoodAccountJSON()))
		if err != nil {
			t.Errorf("Not expecting JSONToAccount() to error: %+v\n", err)
		}
		if a.ID != "99825362-9fc8-4ff6-b5ce-a064854b1268" {
			t.Errorf("Not expecting JSONToAccount() to return that ID: %+v\n", a)
		}
	})
	// TODO: Add more failure cases
}
