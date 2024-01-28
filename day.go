package i18nd

import (
	"time"
)

// DayParse parses a YYYY-MM-DD formatted text.
func DayParse(text, language string) (t time.Time, err error) {
	t, err = time.Parse("2006-01-02", text)
	return
}

// DayShow shows a date in YYYY-MM-DD format.
func DayShow(t time.Time, language string) string {
	return t.Format("2006-01-02")
}
