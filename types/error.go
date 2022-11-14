package types

type Error struct {
	Message string `json:"message"`
}

func CreateErrorMessage(message string) Error {
	return Error{
		Message: message,
	}
}
