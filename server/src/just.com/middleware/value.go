package middleware
import "errors"

const (
	MLEARNING_CONTENT = "MLEARNING_CONTENT"
	RESPONSE = "RESPONSE"
)
var (
	NO_TOKEN_ERR = errors.New("NO_TOKEN_ERR")
	NO_AUTHORITATION = errors.New("NO_AUTHORITATION")
)