package timer

import (
	"time"
)

func GetNowTime() time.Time {
	return time.Now()
}

func GetCalculateTime(currentTimeer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTimeer.Add(duration), nil
}

func ConvertTimeStamp(timeStamp int64) string {
	return time.Unix(timeStamp, 0).Format("2006-01-02 15-04-05")
}
