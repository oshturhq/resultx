package resultx

type Error struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
}

type Result[T any] struct {
	Success bool      `json:"success"`
	Data    *T        `json:"data,omitempty"`
	Message string    `json:"message,omitempty"`
	Error   *Error    `json:"error,omitempty"`
	Meta    *Metadata `json:"meta,omitempty"`
}

func Ok[T any](val T, message string, opts ...MetaOption) *Result[T] {
	result := &Result[T]{
		Success: true,
		Data:    &val,
		Message: message,
	}

	if len(opts) > 0 {
		result.Meta = &Metadata{}
		for _, opt := range opts {
			opt(result.Meta)
		}
	}

	return result
}

func Fail[T any](code string, err error, opts ...MetaOption) *Result[T] {
	result := &Result[T]{
		Success: false,
		Data:    nil,
		Error: &Error{
			Code:    code,
			Message: err.Error(),
		},
	}

	if len(opts) > 0 {
		result.Meta = &Metadata{}
		for _, opt := range opts {
			opt(result.Meta)
		}
	}

	return result
}
