package sharedutils

import "time"

const (
	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
	YYYYMMDD       = "20060102"
	YYYYMMDDslash  = "2006/01/02"
)

func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func MonthAndDayEqual(date1, date2 time.Time) bool {
	_, m1, d1 := date1.Date()
	_, m2, d2 := date2.Date()
	return m1 == m2 && d1 == d2
}

func DateParseYYYYMMDDHHmmSS(date string) *time.Time {
	val, _ := time.Parse("2006-01-02 01:01:01", date)
	return &val
}

func ParseDur(s string) time.Duration {
	t, _ := time.Parse("15:04", s)
	// sb, _ := time.Parse("15:04", time.Now().Format("15:04"))
	sb, _ := time.Parse("15:04", "00:00")
	return t.Sub(sb)
}
