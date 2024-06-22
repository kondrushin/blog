package response

type HttpError struct {
	error
	StatusCode int
}

func SetHttpStatusCode(err error, code int) error {
	return &HttpError{error: err, StatusCode: code}
}
