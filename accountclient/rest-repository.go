package accountclient

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"

	"net/http"
)

// RestResponse  Implementation of a response from the Rest Account repository
type RestResponse struct {
	data []byte
}

// GetData  Returns the data from the Rest Account repository
func (r RestResponse) GetData() []byte {
	return r.data
}

// RestRepository  This behaves like a Repository
type RestRepository struct {
	BaseURL string
}

func returnResponse(restResponse *http.Response, expectedStatus int) (resp Response, err error) {

	body, err := ioutil.ReadAll(restResponse.Body)
	if err != nil {
		return nil, err
	}
	if restResponse.StatusCode != expectedStatus {
		msg := fmt.Sprintf("Not expecting StatusCode %d : Error: %s", restResponse.StatusCode, string(body))
		return nil, errors.New(msg)
	}

	return RestResponse{
		data: body,
	}, err
}

// Get  Gets an Account over a Rest API
func (r *RestRepository) Get(id string) (resp Response, err error) {
	restResponse, err := http.Get(r.BaseURL + "/v1/organisation/accounts/" + id)

	if err != nil {
		return nil, err
	}

	return returnResponse(restResponse, 200)
}

// GetAll  Gets all Accounts over a Rest API
func (r *RestRepository) GetAll(opts *PaginationOptions) (resp Response, err error) {
	url := r.BaseURL + "/v1/organisation/accounts"
	if opts != nil {
		// TODO find a library to do this encoding for us
		url += fmt.Sprintf("?page[number]=%d&page[size]=%d", opts.PageNumber, opts.PageSize)
	}
	restResponse, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	return returnResponse(restResponse, 200)
}

// Create  Send data over Rest
func (r *RestRepository) Create(request []byte) (resp Response, err error) {
	restResponse, err := http.Post(r.BaseURL+"/v1/organisation/accounts",
		"application/vnd.api+json",
		bytes.NewBuffer(request))

	if err != nil {
		return nil, err
	}

	return returnResponse(restResponse, 201)
}

// Delete  Delete an Account over a Rest API
func (r *RestRepository) Delete(id string) (isDeleted bool, err error) {
	url := r.BaseURL + "/v1/organisation/accounts/" + id + "?version=0"
	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return false, err
	}

	restResponse, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}

	if restResponse.StatusCode == 204 {
		return true, nil
	}

	msg := fmt.Sprintf("DELETE returned unexpected status code: %d", restResponse.StatusCode)
	return false, errors.New(msg)

}
