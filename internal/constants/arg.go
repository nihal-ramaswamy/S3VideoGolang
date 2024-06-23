package constants

const (
	DELIMITER = "="
	PREFIX    = "-"
)

type ArgParseError struct {
	Message string
}

func NewArgParseError(message string) *ArgParseError {
	return &ArgParseError{
		Message: message,
	}
}

func (a *ArgParseError) Error() string {
	return a.Message
}
