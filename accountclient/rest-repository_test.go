package accountclient

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRestResponse_GetData(t *testing.T) {
	type fields struct {
		data []byte
	}
	x := fields{
		data: []byte("hello world"),
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"it returns the data", x, []byte("hello world")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RestResponse{
				data: tt.fields.data,
			}
			if got := r.GetData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RestResponse.GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRestRepository_Get(t *testing.T) {
	statusCode := 200
	// generate a test server so we can capture and inspect the request
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(statusCode)
		res.Write([]byte("body"))
	}))
	defer func() { testServer.Close() }()

	r := &RestRepository{
		BaseURL: testServer.URL,
	}

	t.Run("succeeds with a status 200 and contains the body", func(t *testing.T) {
		statusCode = 200
		gotResp, err := r.Get("")
		if err != nil {
			t.Errorf("Not expecting the error: %+v\n", err)
		}

		if string(gotResp.GetData()) != "body" {
			t.Errorf("Not expecting the gotResp: %+v\n", gotResp)
		}
	})
	t.Run("fails with a status 201", func(t *testing.T) {
		statusCode = 201
		_, err := r.Get("")
		if err == nil {
			t.Error("Expecting an error")
		}
	})
}

func TestRestRepository_GetAll(t *testing.T) {
	statusCode := 200
	// generate a test server so we can capture and inspect the request
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(statusCode)
		res.Write([]byte("body"))
	}))
	defer func() { testServer.Close() }()

	r := &RestRepository{
		BaseURL: testServer.URL,
	}

	t.Run("succeeds with a status 200 and contains the body", func(t *testing.T) {
		statusCode = 200
		gotResp, err := r.GetAll(nil)
		if err != nil {
			t.Errorf("Not expecting the error: %+v\n", err)
		}

		if string(gotResp.GetData()) != "body" {
			t.Errorf("Not expecting the gotResp: %+v\n", gotResp)
		}
	})

	t.Run("fails with a status 201", func(t *testing.T) {
		statusCode = 201
		_, err := r.GetAll(nil)
		if err == nil {
			t.Error("Expecting an error")
		}
	})
}

func TestRestRepository_Create(t *testing.T) {
	statusCode := 201
	// generate a test server so we can capture and inspect the request
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(statusCode)
		res.Write([]byte("body"))
	}))
	defer func() { testServer.Close() }()

	r := &RestRepository{
		BaseURL: testServer.URL,
	}

	t.Run("succeeds with a status 201 and contains the body", func(t *testing.T) {
		statusCode = 201
		gotResp, err := r.Create([]byte("test"))
		if err != nil {
			t.Errorf("Not expecting the error: %+v\n", err)
		}

		if string(gotResp.GetData()) != "body" {
			t.Errorf("Not expecting the gotResp: %+v\n", gotResp)
		}
	})

	t.Run("fails with a status 200", func(t *testing.T) {
		statusCode = 200
		_, err := r.Create([]byte("test"))
		if err == nil {
			t.Error("Expecting an error")
		}
	})
}

func TestRestRepository_Delete(t *testing.T) {
	statusCode := 204
	// generate a test server so we can capture and inspect the request
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(statusCode)
		res.Write([]byte("body"))
	}))
	defer func() { testServer.Close() }()

	r := &RestRepository{
		BaseURL: testServer.URL,
	}

	t.Run("succeeds with a status 204 and contains the body", func(t *testing.T) {
		statusCode = 204
		isDeleted, err := r.Delete("")
		if err != nil {
			t.Errorf("Not expecting the error: %+v\n", err)
		}

		if !isDeleted {
			t.Error("Expecting return to be true")
		}
	})

	t.Run("fails with a status 200", func(t *testing.T) {
		statusCode = 200
		isDeleted, err := r.Delete("")
		if isDeleted {
			t.Error("Expecting return to be false")
		}
		if err == nil {
			t.Error("Expecting an error")
		}
	})
}
