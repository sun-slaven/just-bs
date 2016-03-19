package common
import "encoding/base64"

/*std encoding*/
func Base64Encoding(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

/*std decoding*/
func Base64Decoding(value string) string {
	byte, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return ""
	}
	return string(byte)
}