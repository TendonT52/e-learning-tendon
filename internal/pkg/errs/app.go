package errs

var (
	UserNotFound        = appError{kind: "user not found"}
	NodeNotFound        = appError{kind: "node not found"}
	LessonNotFound      = appError{kind: "lesson not found"}
	CourseNotFound      = appError{kind: "course not found"}
	InvalidUserID       = appError{kind: "invalid user id"}
	InvalidNodeID       = appError{kind: "invalid node id"}
	InvalidLessonID     = appError{kind: "invalid lesson id"}
	InvalidCourseID = appError{kind: "invalid course id"}

	IDShouldBeEmpty          = appError{kind: "id should be empty"}
	CreateByShouldNotBeEmpty = appError{kind: "create by should not be empty"}
	TypeShouldNotBeEmpty     = appError{kind: "type should not be empty"}
	UpdatedAtShouldBeEmpty   = appError{kind: "updated at should be empty"}
	DataShouldNotBeEmpty     = appError{kind: "data should not be empty"}

	InsertFailed = appError{kind: "insert failed"}
	UpdateFailed = appError{kind: "update failed"}
	DeleteFailed = appError{kind: "delete failed"}
	FindFailed   = appError{kind: "find failed"}

	InvalidToken  = appError{kind: "invalid token"}
	TokenExpired  = appError{kind: "token expired"}
	TokenNotfound = appError{kind: "token not found"}
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
