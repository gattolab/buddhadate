package buddha

import "time"

// ToBuddhistYearIn converts a Gregorian (Common Era) year to its Buddhist
// Era equivalent under the given Era convention.
//
//	buddha.ToBuddhistYearIn(2026, buddha.TheravadaEra)
//	// 2570
func ToBuddhistYearIn(year int, era Era) int {
	return year + int(era)
}

// ToGregorianYearIn converts a Buddhist Era year to its Gregorian (Common
// Era) equivalent under the given Era convention.
//
//	buddha.ToGregorianYearIn(2570, buddha.TheravadaEra)
//	// 2026
func ToGregorianYearIn(year int, era Era) int {
	return year - int(era)
}

// ToBuddhistYear converts a Gregorian (Common Era) year to its Buddhist Era
// equivalent using Thailand's convention (ThaiEra, B.E. = C.E. + 543). For
// other countries' conventions, use ToBuddhistYearIn.
//
//	buddha.ToBuddhistYear(2026)
//	// 2569
func ToBuddhistYear(year int) int {
	return ToBuddhistYearIn(year, defaultEra)
}

// ToGregorianYear converts a Buddhist Era year to its Gregorian (Common
// Era) equivalent using Thailand's convention (ThaiEra, C.E. = B.E. - 543).
// For other countries' conventions, use ToGregorianYearIn.
//
//	buddha.ToGregorianYear(2569)
//	// 2026
func ToGregorianYear(year int) int {
	return ToGregorianYearIn(year, defaultEra)
}

// ToTime returns the underlying Gregorian time.Time value of d. It is
// equivalent to d.Time().
func ToTime(d Date) time.Time {
	return d.Time()
}

// IsLeapYearIn reports whether year, given as a Buddhist Era year under the
// given Era convention, is a leap year under the Gregorian calendar rules.
//
//	buddha.IsLeapYearIn(2568, buddha.TheravadaEra)
func IsLeapYearIn(year int, era Era) bool {
	g := ToGregorianYearIn(year, era)
	return g%4 == 0 && (g%100 != 0 || g%400 == 0)
}

// IsLeapYear reports whether year, given as a Buddhist Era year using
// Thailand's convention (ThaiEra), is a leap year under the Gregorian
// calendar rules. For other countries' conventions, use IsLeapYearIn.
//
//	buddha.IsLeapYear(2569)
func IsLeapYear(year int) bool {
	return IsLeapYearIn(year, defaultEra)
}

// DaysInMonthIn returns the number of days in the given month of year,
// where year is a Buddhist Era year under the given Era convention.
//
//	buddha.DaysInMonthIn(2568, time.February, buddha.TheravadaEra)
func DaysInMonthIn(year int, month time.Month, era Era) int {
	g := ToGregorianYearIn(year, era)
	firstOfNextMonth := time.Date(g, month+1, 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstOfNextMonth.AddDate(0, 0, -1)
	return lastDay.Day()
}

// DaysInMonth returns the number of days in the given month of year, where
// year is a Buddhist Era year using Thailand's convention (ThaiEra). For
// other countries' conventions, use DaysInMonthIn.
//
//	buddha.DaysInMonth(2569, time.February)
func DaysInMonth(year int, month time.Month) int {
	return DaysInMonthIn(year, month, defaultEra)
}

// DaysInYearIn returns the number of days in year, where year is a
// Buddhist Era year under the given Era convention.
func DaysInYearIn(year int, era Era) int {
	if IsLeapYearIn(year, era) {
		return 366
	}
	return 365
}

// DaysInYear returns the number of days in year, where year is a Buddhist
// Era year using Thailand's convention (ThaiEra). For other countries'
// conventions, use DaysInYearIn.
func DaysInYear(year int) int {
	return DaysInYearIn(year, defaultEra)
}
