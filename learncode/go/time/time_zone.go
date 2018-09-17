package main

import (
	"fmt"
	"time"
)

var gCountryId = map[string]string{
	"Hungary":     "Europe/Budapest",
	"Egypt":       "Africa/Cairo",
	"Los_Angeles": "America/Los_Angeles",
	"chongqing":   "Asia/Chongqing",
}

func timeIn(name string) time.Time {
	loc, err := time.LoadLocation(gCountryId[name])
	if err != nil {
		panic(err)
	}
	return time.Now().In(loc)
}

var StringTimeFormat = "2006-01-02 15:04:05"
func ConvertTimeString2Timestamp(timeStr string, Locale *time.Location) (int64, error) {
	t, err := time.ParseInLocation(StringTimeFormat, timeStr, Locale)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func main() {
	utc := time.Now().UTC().Format("15:04")
	hun := timeIn("Hungary")
	eg := timeIn("Egypt")
	al := timeIn("Los_Angeles")
	cq := timeIn("chongqing")
	fmt.Println(utc, hun, eg, al, cq)
	fmt.Println(eg.Weekday(), al.Weekday(), cq.Weekday())

	fmt.Println()

	t := time.Now()
	fmt.Println(t.Unix(), t.UTC().Unix())

	t = time.Now()
	y, w := t.ISOWeek()
	fmt.Println(t.Weekday(), y, w)


	str := "2018-05-05 05:05:05"
	locale := "Asia/Chongqing"
	t1, _ := ConvertTimeString2Timestamp(str, locale)
	fmt.Println(t1)
}
