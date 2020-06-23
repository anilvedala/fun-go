package utils

import "time"

func ParseCalendarDate(dateString string) (time.Time, error)  {
	date, err :=  time.Parse("2006-01-02", dateString)

	if err != nil {
		return time.Now(), err
	}
	return date, nil
}
