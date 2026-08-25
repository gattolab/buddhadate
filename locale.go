package buddha

import (
	"fmt"
	"time"
)

// Locale supplies the human-readable month and weekday names used when
// rendering a date in a particular language. buddha ships only ThaiLocale
// built in.
//
// Other Buddhist-Era-using countries (Cambodia, Laos, Myanmar, Sri Lanka,
// ...) use their own scripts and languages for month and weekday names.
// This package intentionally does not bundle translations for those
// languages, since shipping unverified translations in a published
// library risks being subtly wrong. Instead, implement Locale with the
// correct names for the target language and pass it to FormatLocale, or
// call its methods directly.
type Locale interface {
	// MonthName returns the full name of month in this locale's language.
	MonthName(month time.Month) string
	// MonthShortName returns the abbreviated name of month in this
	// locale's language.
	MonthShortName(month time.Month) string
	// WeekdayName returns the full name of day in this locale's language.
	WeekdayName(day time.Weekday) string
	// WeekdayShortName returns the abbreviated name of day in this
	// locale's language.
	WeekdayShortName(day time.Weekday) string
}

// FormatLocale formats t as "<day> <month name> <year>" using loc for the
// month name and era for the Buddhist Era year numbering convention.
//
//	buddha.FormatLocale(t, buddha.ThaiLocale{}, buddha.ThaiEra)
//	// "26 สิงหาคม 2569"
func FormatLocale(t time.Time, loc Locale, era Era) string {
	return fmt.Sprintf("%d %s %d", t.Day(), loc.MonthName(t.Month()), ToBuddhistYearIn(t.Year(), era))
}
