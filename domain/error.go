package domain

const (
	NotFound         = 404
	InvalidArguments = 400
	AlreadyExists    = 409
	PermissionDenied = 403
	Unauthorized     = 401
	Internal         = 500
)

type Error struct {
	Code    int32
	Message string
}

func NewError(code int32, message string) error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func (e Error) Error() string {
	return e.Message
}
