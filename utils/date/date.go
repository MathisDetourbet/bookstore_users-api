package date

import "time"

const (
	apiDateLayout = "01-02-2006T15:04:05Z"
	apiDBLayout   = "2006-01-02 15:04:05"
)

func getNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return getNow().Format(apiDateLayout)
}

func GetNowDBFormat() string {
	return getNow().Format(apiDBLayout)
}
