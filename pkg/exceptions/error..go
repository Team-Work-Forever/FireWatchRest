package exec

type (
	Error struct {
		Status int
		Title  string
		Detail string
	}
)

func NewError(status int, title string, detail string) *Error {
	return &Error{
		Status: status,
		Title:  title,
		Detail: detail,
	}
}

func (e *Error) Error() string {
	return e.Detail
}
