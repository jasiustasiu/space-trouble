package httpError

type HTTPError struct {
	Status  int
	Details string
}

func (e *HTTPError) Error() string {
	return e.Details
}

func NewHTTPError(status int, details string) error {
	return &HTTPError{
		Status:  status,
		Details: details,
	}
}
