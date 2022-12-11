package errs

var (
	ErrNotFound    = appError{kind: "not found"}
	ErrWrongFormat = appError{kind: "wrong format"}
	ErrDatabase    = appError{kind: "database error"}

	ErrInvalidToken    = appError{kind: "invalid token"}
	ErrTokenExpired    = appError{kind: "token expired"}
)

type appError struct {
	kind string
	err  error
}

func NewaAppError(kind string, err error) appError {
	return appError{
		kind: kind,
		err:  err,
	}
}

func (e appError) Error() string {
	return e.kind
}

func (e appError) From(err error) appError {
	e.err = err
	return e
}

func (e appError) Is(err error) bool {
	t, ok := err.(appError)
	if !ok {
		return false
	}
	return e.kind == t.kind
}

func (e appError) Unwrap() error {
	return e.err
}
