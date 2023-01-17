package utls

import "time"

func GetCurrentDateTime() string {
	CurrentTime := time.Now()
	return CurrentTime.Format(time.UnixDate)
}

func GetCurrentDate()string{
	
		date:=time.Now().Format("02-01-2006")
	return date
}