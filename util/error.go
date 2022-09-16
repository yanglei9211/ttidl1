package util

type Error struct {
	StatusCode int    `json:"statusCode"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

func (e *Error) Error() string {
	return e.Msg
}

func NewError(statusCode, code int, msg string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       code,
		Msg:        msg,
	}
}
