package buddha

import "time"

// Date represents a point in time expressed using the Buddhist Era (B.E.)
// calendar year. It wraps a standard Gregorian time.Time value; no separate
// calendar system is implemented. Year() returns the Buddhist Era year while
// every other computation (leap years, days in month, duration arithmetic,
// and so on) is delegated to the standard time package.
//
// The zero value of Date corresponds to the zero value of time.Time.
type Date struct {
	t time.Time
}

// Now returns the current local date and time as a Buddhist Era Date.
func Now() Date {
	return Date{t: time.Now()}
}

// Today returns today's date at midnight in the local timezone, expressed as
// a Buddhist Era Date.
func Today() Date {
	now := time.Now()
	return Date{t: time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())}
}

// NowIn returns the current date and time in the given location, expressed
// as a Buddhist Era Date.
func NowIn(loc *time.Location) Date {
	return Date{t: time.Now().In(loc)}
}

// TodayIn returns today's date at midnight in the given location, expressed
// as a Buddhist Era Date.
func TodayIn(loc *time.Location) Date {
	now := time.Now().In(loc)
	return Date{t: time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)}
}

// New creates a Date from a Buddhist Era year, month, and day, at midnight
// in the local timezone, using Thailand's Buddhist Era convention (ThaiEra).
// year must be a Buddhist Era year. For other countries' conventions, use
// NewIn.
//
//	d := buddha.New(2569, time.August, 26)
func New(year int, month time.Month, day int) Date {
	return NewIn(defaultEra, year, month, day)
}

// NewIn creates a Date from a Buddhist Era year, month, and day, at
// midnight in the local timezone, using the given Era convention. year must
// be a Buddhist Era year under era.
//
//	d := buddha.NewIn(buddha.TheravadaEra, 2570, time.August, 26)
func NewIn(era Era, year int, month time.Month, day int) Date {
	return Date{t: time.Date(ToGregorianYearIn(year, era), month, day, 0, 0, 0, 0, time.Local)}
}

// FromTime converts a standard Gregorian time.Time into a Buddhist Era Date.
// No values are modified; only the interpretation of the year changes when
// the Date is displayed or queried through Year().
func FromTime(t time.Time) Date {
	return Date{t: t}
}

// FromBuddhistTime creates a Date from explicit Buddhist Era year, month,
// day, hour, minute, second, nanosecond, and location components, using
// Thailand's Buddhist Era convention (ThaiEra). year must be a Buddhist Era
// year. For other countries' conventions, use FromBuddhistTimeIn.
func FromBuddhistTime(year int, month time.Month, day, hour, minute, second, nanosecond int, loc *time.Location) Date {
	return FromBuddhistTimeIn(defaultEra, year, month, day, hour, minute, second, nanosecond, loc)
}

// FromBuddhistTimeIn creates a Date from explicit Buddhist Era year, month,
// day, hour, minute, second, nanosecond, and location components, using the
// given Era convention. year must be a Buddhist Era year under era.
func FromBuddhistTimeIn(era Era, year int, month time.Month, day, hour, minute, second, nanosecond int, loc *time.Location) Date {
	return Date{t: time.Date(ToGregorianYearIn(year, era), month, day, hour, minute, second, nanosecond, loc)}
}

// Year returns the Buddhist Era year of d using Thailand's Buddhist Era
// convention (ThaiEra). For other countries' conventions, use YearIn.
func (d Date) Year() int {
	return d.YearIn(defaultEra)
}

// YearIn returns the Buddhist Era year of d under the given Era convention.
//
//	d.YearIn(buddha.TheravadaEra)
func (d Date) YearIn(era Era) int {
	return ToBuddhistYearIn(d.t.Year(), era)
}

// Month returns the month of d.
func (d Date) Month() time.Month {
	return d.t.Month()
}

// Day returns the day of the month of d.
func (d Date) Day() int {
	return d.t.Day()
}

// Hour returns the hour within the day of d, in the range [0, 23].
func (d Date) Hour() int {
	return d.t.Hour()
}

// Minute returns the minute offset within the hour of d, in the range
// [0, 59].
func (d Date) Minute() int {
	return d.t.Minute()
}

// Second returns the second offset within the minute of d, in the range
// [0, 59].
func (d Date) Second() int {
	return d.t.Second()
}

// Location returns the time zone information associated with d.
func (d Date) Location() *time.Location {
	return d.t.Location()
}

// Time returns the underlying Gregorian time.Time value of d.
func (d Date) Time() time.Time {
	return d.t
}

// Before reports whether d occurs before other.
func (d Date) Before(other Date) bool {
	return d.t.Before(other.t)
}

// After reports whether d occurs after other.
func (d Date) After(other Date) bool {
	return d.t.After(other.t)
}

// Equal reports whether d and other represent the same instant in time,
// regardless of location.
func (d Date) Equal(other Date) bool {
	return d.t.Equal(other.t)
}

// Add returns the Date d+dur.
func (d Date) Add(dur time.Duration) Date {
	return Date{t: d.t.Add(dur)}
}

// AddDate returns the Date corresponding to adding the given number of
// years, months, and days to d. See time.Time.AddDate for the exact
// normalization rules that are applied.
func (d Date) AddDate(years, months, days int) Date {
	return Date{t: d.t.AddDate(years, months, days)}
}

// String returns d formatted as "02/01/2006" using the Buddhist Era year,
// for example "26/08/2569".
func (d Date) String() string {
	return d.Format("02/01/2006")
}

// StartOfDay returns d at midnight (00:00:00) on the same day.
func (d Date) StartOfDay() Date {
	return Date{t: time.Date(d.t.Year(), d.t.Month(), d.t.Day(), 0, 0, 0, 0, d.t.Location())}
}

// EndOfDay returns d at the last nanosecond (23:59:59.999999999) of the same
// day.
func (d Date) EndOfDay() Date {
	return Date{t: time.Date(d.t.Year(), d.t.Month(), d.t.Day(), 23, 59, 59, 999999999, d.t.Location())}
}

// StartOfMonth returns d set to the first day of its month at midnight.
func (d Date) StartOfMonth() Date {
	return Date{t: time.Date(d.t.Year(), d.t.Month(), 1, 0, 0, 0, 0, d.t.Location())}
}

// EndOfMonth returns d set to the last day of its month at the last
// nanosecond of the day.
func (d Date) EndOfMonth() Date {
	firstOfNextMonth := time.Date(d.t.Year(), d.t.Month()+1, 1, 0, 0, 0, 0, d.t.Location())
	lastDay := firstOfNextMonth.AddDate(0, 0, -1)
	return Date{t: time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(), 23, 59, 59, 999999999, d.t.Location())}
}

// StartOfYear returns d set to January 1st of its year at midnight.
func (d Date) StartOfYear() Date {
	return Date{t: time.Date(d.t.Year(), time.January, 1, 0, 0, 0, 0, d.t.Location())}
}

// EndOfYear returns d set to December 31st of its year at the last
// nanosecond of the day.
func (d Date) EndOfYear() Date {
	return Date{t: time.Date(d.t.Year(), time.December, 31, 23, 59, 59, 999999999, d.t.Location())}
}

// StartOfDay returns d at midnight (00:00:00) on the same day.
func StartOfDay(d Date) Date {
	return d.StartOfDay()
}

// EndOfDay returns d at the last nanosecond of the same day.
func EndOfDay(d Date) Date {
	return d.EndOfDay()
}

// StartOfMonth returns d set to the first day of its month at midnight.
func StartOfMonth(d Date) Date {
	return d.StartOfMonth()
}

// EndOfMonth returns d set to the last day of its month at the last
// nanosecond of the day.
func EndOfMonth(d Date) Date {
	return d.EndOfMonth()
}

// StartOfYear returns d set to January 1st of its year at midnight.
func StartOfYear(d Date) Date {
	return d.StartOfYear()
}

// EndOfYear returns d set to December 31st of its year at the last
// nanosecond of the day.
func EndOfYear(d Date) Date {
	return d.EndOfYear()
}
