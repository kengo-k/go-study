package study01

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	// Nowで現在時刻を取得
	now := time.Now()
	// 指定日時でDateを生成
	tz, _ := time.LoadLocation("Asia/Tokyo")
	date := time.Date(2000, time.December, 31, 23, 59, 0, 0, tz)
	if now != date {
		t.Errorf("got `%v` , expected `%v` must be different", now, date)
	}
}

func TestDateSub(t *testing.T) {
	tz, _ := time.LoadLocation("Asia/Tokyo")
	date1 := time.Date(2000, time.December, 31, 23, 59, 0, 0, tz)
	date2 := time.Date(2000, time.December, 31, 23, 50, 0, 0, tz)
	result := date1.Sub(date2)
	got := result.String()
	expected := "9m0s"
	if got != expected {
		t.Errorf("got `%v` , expected `%v`", got, expected)
	}
}
