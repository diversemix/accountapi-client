package accountclient

// Response  Response expected from any Account repository
type Response interface {
	GetData() []byte
}

// PaginationOptions  holds the options for pagination
type PaginationOptions struct {
	PageNumber int
	PageSize   int
}

// Repository  The interface the Account repository supports
type Repository interface {
	Get(id string) (resp Response, err error)
	GetAll(opt *PaginationOptions) (resp Response, err error)
	Create(request []byte) (resp Response, err error)
	Delete(id string) (isDeleted bool, err error)
}
