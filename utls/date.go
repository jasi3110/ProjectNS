package utls

import "time"

func GetCurrentDateTime() string {
	CurrentTime := time.Now()
	return CurrentTime.Format(time.UnixDate)
}
