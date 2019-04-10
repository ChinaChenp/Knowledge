package main

import (
	"fmt"
	"time"
)

func timeStamp2Str(s int) string {
	tm := time.Unix(int64(s), 0)
	return tm.Format("2006-01-02 03:04:05 PM")
}

func NextDaySecondsAtHour(hour int) int {
	yesTime := time.Now().AddDate(0, 0, +1)
	return int(yesTime.Unix())
}

func GetTodayBegin() int64 {
	t := time.Now()
	timeStr := fmt.Sprintf("%04d-%02d-%02d 00:00:00", t.Year(), t.Month(), t.Day())
	re, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	return re.Unix()
}

func GetNextTodayBegin() int64 {
	t := time.Now().AddDate(0, 0, 1)
	timeStr := fmt.Sprintf("%04d-%02d-%02d 00:00:00", t.Year(), t.Month(), t.Day())
	re, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	return re.Unix()
}

func main() {
	now := time.Now()
	nowDay := now.Unix() / 86400 * 86400

	t, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	timeNumber := t.Unix()

	nextDay := NextDaySecondsAtHour(3)
	fmt.Println(nowDay, nextDay)
	fmt.Println(timeStamp2Str(int(now.Unix())), timeStamp2Str(int(nowDay)), timeStamp2Str(nextDay), timeStamp2Str(int(timeNumber)))

	t = time.Now()
	year, month, day := t.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	fmt.Println(today)

	fmt.Println(GetTodayBegin())
	fmt.Println(GetNextTodayBegin())
}
