package date

import "time"

const (
	apiDateLayout = "01-02-2006T15:04:05Z"
)

func getNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return getNow().Format(apiDateLayout)
}
