package service
import "strings"

func IsEmpty(args...string) bool {
	for _, arg := range args {
		if strings.TrimSpace(arg) != "" {
			return true
		}
	}
	return false
}