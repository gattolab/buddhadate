package buddha

// Era represents a Buddhist Era numbering convention: the number of years
// added to a Gregorian (Common Era) year to obtain the corresponding
// Buddhist Era year under that convention.
//
// "The" Buddhist Era is not a single universal reckoning. Thailand counts
// B.E. 1 as the year after the Buddha's Parinibbana (death), giving the
// widely known B.E. = C.E. + 543 rule. Sri Lanka, Myanmar, Cambodia, and
// Laos instead count the year of the Parinibbana itself as year 1, one
// year earlier, giving B.E. = C.E. + 544. This package defaults to the
// Thai convention (ThaiEra) everywhere an Era is not specified explicitly,
// since that is the behavior documented throughout this package. Callers
// working with another country's convention should use the "*In" family
// of functions (for example ToBuddhistYearIn, NewIn, FormatIn) together
// with TheravadaEra, or a custom Era value, instead of assuming the
// default.
type Era int

const (
	// ThaiEra is Thailand's Buddhist Era numbering convention:
	// B.E. = C.E. + 543. This is the convention used by every function in
	// this package that does not take an explicit Era.
	ThaiEra Era = 543

	// TheravadaEra is the Buddhist Era numbering convention commonly used
	// by Sri Lanka, Myanmar, Cambodia, and Laos: B.E. = C.E. + 544. It
	// begins counting one year earlier than ThaiEra.
	TheravadaEra Era = 544
)

// defaultEra is the convention used by every package function and Date
// method that does not take an explicit Era, matching this package's
// documented Thai-first defaults.
const defaultEra = ThaiEra
