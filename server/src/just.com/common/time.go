package common
import "time"

const (
	TIME_YYYY_MM_DD_HH_DD_SS = "2006-01-02 15:04:05"
)
func TimeFormat(t time.Time) string {
	return t.Format(TIME_YYYY_MM_DD_HH_DD_SS)
}
