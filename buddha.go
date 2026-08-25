// Package buddha provides utilities for working with Buddhist Era (B.E./พ.ศ.)
// dates.
//
// The package does not implement a separate calendar system. Internally it
// relies entirely on Go's standard time package and the Gregorian calendar;
// a Buddhist Era year is always simply the Gregorian year plus a fixed
// offset. buddha only changes how years are interpreted and displayed,
// never how dates are computed.
//
//	now := buddha.Now()
//	fmt.Println(now)
//	// 26/08/2569
//
// # A note on scope
//
// Every function in this package that does not take an explicit Era or
// Locale parameter defaults to Thailand's conventions: the +543 year
// offset (ThaiEra) and Thai month/weekday names (ThaiLocale). "The"
// Buddhist Era is not actually a single universal system — Sri Lanka,
// Myanmar, Cambodia, and Laos also use a Buddhist Era, but commonly count
// one year earlier (+544, see TheravadaEra), and each has its own
// language for month and weekday names. Use the "*In" functions (for
// example ToBuddhistYearIn, NewIn, FormatIn, ParseIn) with an explicit
// Era, and implement the Locale interface for languages other than Thai,
// rather than assuming this package's Thai-first defaults apply
// everywhere.
//
// # Design principles
//
// Buddha is intentionally small:
//
//   - Zero external dependencies.
//   - No custom calendar or leap-year logic; everything is delegated to
//     time.Time and time.Date.
//   - Explicit, unambiguous naming. Functions that accept or return a
//     Buddhist Era year are named accordingly (for example ToBuddhistYear,
//     ParseBuddhist), and functions dealing with Gregorian years are named
//     ToGregorianYear and ParseGregorian. There is no function that silently
//     guesses which calendar era or country convention applies.
package buddha
