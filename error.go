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

type NestingOverflowError byte

func (err NestingOverflowError) Error() string {
	return "nesting level overflow"
}

func (err NestingOverflowError) Is(target error) bool {
	_, ok := target.(NestingOverflowError)
	return ok
}

const ErrNestingOverflow NestingOverflowError = 0
