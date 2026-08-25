package buddha

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// beYear4Placeholder and beYear2Placeholder are internal sentinel strings
// substituted into a layout in place of Go's reference year tokens ("2006"
// and "06") before delegating to time.Time.Format. Because they are not
// recognized reference layout elements, time.Time.Format copies them through
// verbatim, after which they are replaced with the actual Buddhist Era year.
// The placeholders use only non-printable control characters so that they
// can never accidentally match a Go reference layout token (which are all
// ASCII letters, digits, or punctuation, for example the bare digit tokens
// "1".."5" used for months/days/hours without leading zeros).
const (
	beYear4Placeholder = "\x00\x01\x00"
	beYear2Placeholder = "\x00\x02\x00"
)

// Format formats t using layout, following the same reference layout rules
// as time.Time.Format, except that the year tokens "2006" and "06" are
// replaced with the Buddhist Era year (using Thailand's ThaiEra convention)
// instead of the Gregorian year. For other countries' conventions, use
// FormatIn.
//
//	buddha.Format(
//	    time.Date(2026, 8, 26, 0, 0, 0, 0, time.Local),
//	    "02/01/2006",
//	)
//	// "26/08/2569"
func Format(t time.Time, layout string) string {
	return FormatIn(t, layout, defaultEra)
}

// FormatIn formats t using layout, following the same reference layout
// rules as time.Time.Format, except that the year tokens "2006" and "06"
// are replaced with the Buddhist Era year under the given Era convention.
//
//	buddha.FormatIn(
//	    time.Date(2026, 8, 26, 0, 0, 0, 0, time.Local),
//	    "02/01/2006",
//	    buddha.TheravadaEra,
//	)
//	// "26/08/2570"
func FormatIn(t time.Time, layout string, era Era) string {
	beYear := ToBuddhistYearIn(t.Year(), era)

	modified := strings.ReplaceAll(layout, "2006", beYear4Placeholder)
	modified = strings.ReplaceAll(modified, "06", beYear2Placeholder)

	out := t.Format(modified)

	out = strings.ReplaceAll(out, beYear4Placeholder, fmt.Sprintf("%04d", beYear))
	out = strings.ReplaceAll(out, beYear2Placeholder, fmt.Sprintf("%02d", beYear%100))

	return out
}

// Format formats d using layout. The "2006" and "06" year tokens are
// replaced with the Buddhist Era year using Thailand's ThaiEra convention.
// See Format for details. For other countries' conventions, use
// Date.FormatIn.
func (d Date) Format(layout string) string {
	return Format(d.t, layout)
}

// FormatIn formats d using layout. The "2006" and "06" year tokens are
// replaced with the Buddhist Era year under the given Era convention. See
// FormatIn for details.
func (d Date) FormatIn(layout string, era Era) string {
	return FormatIn(d.t, layout, era)
}

// FormatDate formats t as a Buddhist Era date, "02/01/2006" style, for
// example "26/08/2569".
func FormatDate(t time.Time) string {
	return Format(t, "02/01/2006")
}

// FormatDateTime formats t as a Buddhist Era date and time, for example
// "26/08/2569 14:30:00".
func FormatDateTime(t time.Time) string {
	return Format(t, "02/01/2006 15:04:05")
}

// FormatShortDate formats t as a Buddhist Era date with a two-digit year,
// for example "26/08/69".
func FormatShortDate(t time.Time) string {
	return Format(t, "02/01/06")
}

// FormatThai formats t as a Thai-language Buddhist Era date using the full
// Thai month name and Thailand's Buddhist Era convention, for example
// "26 สิงหาคม 2569". It is equivalent to
// FormatLocale(t, ThaiLocale{}, ThaiEra). For other countries' languages
// or Era conventions, use FormatLocale directly with your own Locale
// implementation.
func FormatThai(t time.Time) string {
	return FormatLocale(t, ThaiLocale{}, defaultEra)
}

// layoutToken describes a recognized Go reference layout token and the
// regular expression fragment used to match the corresponding text in a
// formatted value.
type layoutToken struct {
	token   string
	pattern string
	isYear  bool
}

// layoutTokens lists the reference layout tokens recognized when locating
// the year field within a formatted value, ordered so that longer, more
// specific tokens are tried before shorter ones that could otherwise match
// as a prefix (for example "2006" before "06", and "01" before "1").
var layoutTokens = []layoutToken{
	{"2006", `(\d{4})`, true},
	{"January", `([A-Za-z]+)`, false},
	{"Monday", `([A-Za-z]+)`, false},
	{"-0700", `([+-]\d{4})`, false},
	{"Z0700", `(Z|[+-]\d{4})`, false},
	{"-07:00", `([+-]\d{2}:\d{2})`, false},
	{"Z07:00", `(Z|[+-]\d{2}:\d{2})`, false},
	{"_2", `( ?\d{1,2})`, false},
	{"01", `(\d{2})`, false},
	{"02", `(\d{2})`, false},
	{"03", `(\d{2})`, false},
	{"04", `(\d{2})`, false},
	{"05", `(\d{2})`, false},
	{"06", `(\d{2})`, true},
	{"15", `(\d{2})`, false},
	{"Jan", `([A-Za-z]{3})`, false},
	{"Mon", `([A-Za-z]{3})`, false},
	{"MST", `([A-Za-z]+)`, false},
	{"PM", `(AM|PM)`, false},
	{"pm", `(am|pm)`, false},
	{"1", `(\d{1,2})`, false},
	{"2", `(\d{1,2})`, false},
	{"3", `(\d{1,2})`, false},
	{"4", `(\d{1,2})`, false},
	{"5", `(\d{1,2})`, false},
}

func init() {
	sort.SliceStable(layoutTokens, func(i, j int) bool {
		return len(layoutTokens[i].token) > len(layoutTokens[j].token)
	})
}

// yearCorrectedValue rewrites value by replacing the digits matched by the
// layout's year token ("2006" or "06") with the equivalent Gregorian year
// under the given Era convention, so that a subsequent call to time.Parse
// validates the date (in particular, leap-day validity) against the
// correct Gregorian calendar year instead of the raw Buddhist Era digits.
//
// If the year token cannot be located (for example because layout uses
// tokens not recognized by this best-effort scanner, such as month or
// weekday names), value is returned unchanged and the subsequent
// time.Parse call handles validation using the raw digits.
func yearCorrectedValue(layout, value string, era Era) string {
	var pattern strings.Builder
	yearGroup := -1
	groupIndex := 0
	fourDigitYear := true

	i := 0
	for i < len(layout) {
		matched := false
		for _, tk := range layoutTokens {
			if strings.HasPrefix(layout[i:], tk.token) {
				pattern.WriteString(tk.pattern)
				groupIndex++
				if tk.isYear && yearGroup == -1 {
					yearGroup = groupIndex
					fourDigitYear = tk.token == "2006"
				}
				i += len(tk.token)
				matched = true
				break
			}
		}
		if !matched {
			pattern.WriteString(regexp.QuoteMeta(string(layout[i])))
			i++
		}
	}

	if yearGroup == -1 {
		return value
	}

	re, err := regexp.Compile("^" + pattern.String() + "$")
	if err != nil {
		return value
	}

	loc := re.FindStringSubmatchIndex(value)
	if loc == nil {
		return value
	}

	start, end := loc[2*yearGroup], loc[2*yearGroup+1]
	if start < 0 || end < 0 {
		return value
	}

	rawYear, err := strconv.Atoi(value[start:end])
	if err != nil {
		return value
	}

	var gregorian int
	var replacement string
	if fourDigitYear {
		gregorian = ToGregorianYearIn(rawYear, era)
		replacement = fmt.Sprintf("%04d", gregorian)
	} else {
		// A 2-digit Buddhist Era year is interpreted literally as
		// 2500+yy (for example "69" -> 2569 B.E.), rather than
		// mirroring time.Package's ambiguous pivot-year heuristic for
		// 2-digit Gregorian years. This keeps the interpretation
		// explicit and documented instead of guessing a century.
		gregorian = ToGregorianYearIn(2500+rawYear, era)
		replacement = fmt.Sprintf("%02d", gregorian%100)
	}

	return value[:start] + replacement + value[end:]
}

// parseStrict parses value using layout with time.Parse and additionally
// verifies that reformatting the result with the same layout reproduces
// value exactly. This rejects dates that time.Parse would otherwise
// silently normalize, such as 29 February in a non-leap year.
func parseStrict(layout, value string) (time.Time, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return time.Time{}, fmt.Errorf("%w: %v", ErrInvalidFormat, err)
	}
	if t.Format(layout) != value {
		return time.Time{}, fmt.Errorf("%w: %q is not a valid date for layout %q", ErrInvalidFormat, value, layout)
	}
	return t, nil
}

// ParseGregorian parses value using layout, treating the year in value as a
// Gregorian (Common Era) year. Returns an error wrapping ErrInvalidFormat if
// value does not match layout or does not represent a valid date.
//
//	buddha.ParseGregorian("2006-01-02", "2026-08-26")
func ParseGregorian(layout, value string) (Date, error) {
	t, err := parseStrict(layout, value)
	if err != nil {
		return Date{}, err
	}
	return Date{t: t}, nil
}

// ParseBuddhist parses value using layout, treating the year in value as a
// Buddhist Era year under Thailand's convention (ThaiEra). Returns an error
// wrapping ErrInvalidFormat if value does not match layout or does not
// represent a valid date. For other countries' conventions, use ParseIn.
//
//	buddha.ParseBuddhist("02/01/2006", "26/08/2569")
func ParseBuddhist(layout, value string) (Date, error) {
	return ParseIn(layout, value, defaultEra)
}

// ParseIn parses value using layout, treating the year in value as a
// Buddhist Era year under the given Era convention. Returns an error
// wrapping ErrInvalidFormat if value does not match layout or does not
// represent a valid date.
//
//	buddha.ParseIn("02/01/2006", "26/08/2570", buddha.TheravadaEra)
func ParseIn(layout, value string, era Era) (Date, error) {
	corrected := yearCorrectedValue(layout, value, era)
	t, err := parseStrict(layout, corrected)
	if err != nil {
		return Date{}, fmt.Errorf("%w: %q is not a valid Buddhist Era date for layout %q", ErrInvalidFormat, value, layout)
	}
	return Date{t: t}, nil
}

// Parse parses value using layout, treating the year in value as a Buddhist
// Era year under Thailand's convention (ThaiEra). It is equivalent to
// ParseBuddhist and is provided for symmetry with Format, whose output year
// is also Buddhist Era.
func Parse(layout, value string) (Date, error) {
	return ParseBuddhist(layout, value)
}
