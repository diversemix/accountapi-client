package accountclient

import "fmt"

type mockLogger struct {
	dataFatalln []string
	dataPrintln []string
	dataPrintf  []string
}

func newMockLogger() *mockLogger {
	return &mockLogger{
		dataFatalln: []string{},
		dataPrintln: []string{},
		dataPrintf:  []string{},
	}
}
func (m *mockLogger) Fatalln(v ...interface{}) {
	m.dataFatalln = append(m.dataFatalln, fmt.Sprintln(v...))
}
func (m *mockLogger) Println(v ...interface{}) {
	m.dataPrintln = append(m.dataPrintln, fmt.Sprintln(v...))
}
func (m *mockLogger) Printf(format string, v ...interface{}) {
	m.dataPrintf = append(m.dataPrintf, fmt.Sprintf(format, v...))
}

// Mock Response -----

type mockResponse struct {
	data string
}

func getGoodAccountArrayJSON() string {
	return `{
		"data": [
			{
				"attributes": {
					"alternative_bank_account_names": null,
					"bank_id": "400302",
					"bic": "NWBKGB22",
					"country": "GB"
				},
				"created_on": "2020-01-06T20:28:09.878Z",
				"id": "da12a8b7-ebfa-45ba-bbec-7e3831aa267a",
				"modified_on": "2020-01-06T20:28:09.878Z",
				"organisation_id": "826b2428-b6b4-4a21-895b-3d26c75bf342",
				"type": "accounts",
				"version": 0
			},
			{
				"attributes": {
					"alternative_bank_account_names": null,
					"bank_id": "400302",
					"bic": "NWBKGB22",
					"country": "GB"
				},
				"created_on": "2020-01-06T20:28:20.672Z",
				"id": "99825362-9fc8-4ff6-b5ce-a064854b1268",
				"modified_on": "2020-01-06T20:28:20.672Z",
				"organisation_id": "826b2428-b6b4-4a21-895b-3d26c75bf342",
				"type": "accounts",
				"version": 0
			}
		]	
	}`
}

func getGoodAccountJSON() string {
	return `{
		"data" : {
		"attributes": {
			"alternative_bank_account_names": null,
			"bank_id": "400302",
			"bic": "NWBKGB22",
			"country": "GB"
		},
		"created_on": "2020-01-06T20:28:20.672Z",
		"id": "99825362-9fc8-4ff6-b5ce-a064854b1268",
		"modified_on": "2020-01-06T20:28:20.672Z",
		"organisation_id": "826b2428-b6b4-4a21-895b-3d26c75bf342",
		"type": "accounts",
		"version": 0
	}}`
}

func (r *mockResponse) GetData() []byte {
	return []byte(r.data)
}
func newMockResponse(json string) Response {
	return &mockResponse{data: json}
}

type mockRepository struct{}

func (r *mockRepository) Get(id string) (resp Response, err error) {
	return newMockResponse(getGoodAccountJSON()), nil
}

func (r *mockRepository) GetAll(opt *PaginationOptions) (resp Response, err error) {
	return newMockResponse(getGoodAccountArrayJSON()), nil
}

func (r *mockRepository) Create(request []byte) (resp Response, err error) {
	return newMockResponse(getGoodAccountJSON()), nil
}

func (r *mockRepository) Delete(id string) (isDeleted bool, err error) {
	return true, nil
}
func newMockRepository() Repository {
	return &mockRepository{}
}
