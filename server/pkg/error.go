package pkg

type ServiceError struct {
	Code   int
	ErrMsg string
}

func (s ServiceError) Error() string {
	return s.ErrMsg
}

func (s ServiceError) ErrCode() int {
	return s.Code
}

func NewCustomError(code int, errMsg string) ServiceError {
	return ServiceError{
		Code:   code,
		ErrMsg: errMsg,
	}
}

func ParseError(err error) ServiceError {
	if serviceError, ok := err.(ServiceError); ok {
		return serviceError
	}
	return GeneralFailure
}

var (
	GeneralFailure = ServiceError{Code: 10000, ErrMsg: "general failure"}
	BindingFailure = ServiceError{Code: 10001, ErrMsg: "binding failure"}
	NotExist       = ServiceError{Code: 10002, ErrMsg: "not exist"}
	NotHaveRight   = ServiceError{Code: 10003, ErrMsg: "not have right"}
	InvalidProfile = ServiceError{Code: 10004, ErrMsg: "invalid profile"}
)
