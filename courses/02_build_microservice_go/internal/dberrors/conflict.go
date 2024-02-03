package dberrors

type ConflictError struct{}

func (e *ConflictError) Error() string {
	return "attemted to create a record with an existing key"
}
