package common

func IsNil(args...interface{}) bool {
	for _, arg := range args {
		if arg == nil {
			return true
		}
	}
	return false
}