package util

import "time"

func StrToTime(str string) (time.Time, error) {
	return time.Parse(time.RFC3339, str)
}

func TimeToStr(t time.Time) string {
	return t.Format(time.RFC3339)
}

func CalcAddTime(t1, t2 time.Time) time.Time {
	return t1.Add(t2.Sub(t1))
}

func CalcSubTime(t1, t2 time.Time) time.Duration {
	return t1.Sub(t2)
}
