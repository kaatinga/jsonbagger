package jsonbagger

type NotFoundError byte

func (err NotFoundError) Error() string {
	return "not found"
}

func (err NotFoundError) Is(target error) bool {
	_, ok := target.(NotFoundError)
	return ok
}

const ErrNotFound NotFoundError = 0
