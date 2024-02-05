package i18nd

import (
	"time"
)

// Day parses a YYYY-MM-DD formatted text and panics on error.
func Day(text string) (t time.Time) {
	var err error
	t, err = DayParse(text, "en")
	if err != nil { panic(err) }
	return
}

// DayFormat print a good date format for a specific language.
func DayFormat(language string) (format string) {
	switch language {
	case "es": return "02/01/2006"
	default:   return "2006-01-02"
	}
}

// DayParse parses a YYYY-MM-DD formatted text.
func DayParse(text, language string) (t time.Time, err error) {
	t, err = time.Parse("2006-01-02", text)
	return
}

// DayShow shows a date in YYYY-MM-DD format.
func DayShow(t time.Time, language string) string {
	return t.Format(DayFormat(language))
}
