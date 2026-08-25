package buddha

import "time"

// This file provides a small set of "fun" APIs representing the
// traditional dates associated with the life of the Buddha, as commonly
// cited in Thai Theravada Buddhist tradition.
//
// Important: none of these dates are historically verified Gregorian
// dates. Ancient chronology before reliable written records is inherently
// uncertain, and different Buddhist traditions (Thai, Sri Lankan, Burmese,
// and various academic estimates) do not agree on exact years. The values
// returned here follow one commonly cited traditional Thai narrative:
//
//   - The Buddha is traditionally said to have been born, to have attained
//     enlightenment, and to have died (reached Parinibbana) on the same day
//     of the year: the full moon day of the sixth lunar month, observed in
//     Thailand as Visakha Bucha day and conventionally associated with 15
//     May on the solar calendar.
//   - Thai tradition places the Buddha's death (and the start of the
//     Buddhist Era) 543 years before 1 CE, at age 80.
//   - Working backwards from that convention: enlightenment at age 35, and
//     birth at age 0, 80 years before death.
//
// Years before 1 CE are represented using Go's astronomical year numbering
// (as used by time.Date), where 1 BCE is year 0, 2 BCE is year -1, and so
// on. These functions exist purely for the package's playful surface API;
// they should not be used as a source of historical fact.

// BirthDate returns the traditional/conventional date associated with the
// birth of the Buddha, following the Thai Buddhist tradition described in
// this file. This is not a historically verified date.
func BirthDate() Date {
	return Date{t: time.Date(-622, time.May, 15, 0, 0, 0, 0, time.UTC)}
}

// EnlightenmentDate returns the traditional/conventional date associated
// with the Buddha's enlightenment (at the traditional age of 35),
// following the Thai Buddhist tradition described in this file. This is not
// a historically verified date.
func EnlightenmentDate() Date {
	return Date{t: time.Date(-587, time.May, 15, 0, 0, 0, 0, time.UTC)}
}

// DeathDate returns the traditional/conventional date associated with the
// Buddha's death, or Parinibbana (at the traditional age of 80), following
// the Thai Buddhist tradition described in this file. This date marks the
// beginning of the Buddhist Era, so DeathDate().Year() returns 1. This is
// not a historically verified date.
func DeathDate() Date {
	return Date{t: time.Date(-542, time.May, 15, 0, 0, 0, 0, time.UTC)}
}

// VisakhaBucha returns this year's conventional observance date for
// Visakha Bucha day (Vesak), the day traditionally associated with the
// Buddha's birth, enlightenment, and death. In actual practice the
// observed date shifts every year according to the Thai lunar calendar,
// which this package does not implement; VisakhaBucha instead returns 15
// May of the current year as a fixed, documented approximation commonly
// cited on the solar calendar.
func VisakhaBucha() Date {
	year, _, _ := time.Now().Date()
	return Date{t: time.Date(year, time.May, 15, 0, 0, 0, 0, time.Local)}
}
