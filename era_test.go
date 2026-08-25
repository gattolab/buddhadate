package buddha

import (
	"testing"
	"time"
)

func TestToBuddhistYearIn(t *testing.T) {
	if got, want := ToBuddhistYearIn(2026, ThaiEra), 2569; got != want {
		t.Errorf("ToBuddhistYearIn(2026, ThaiEra) = %d, want %d", got, want)
	}
	if got, want := ToBuddhistYearIn(2026, TheravadaEra), 2570; got != want {
		t.Errorf("ToBuddhistYearIn(2026, TheravadaEra) = %d, want %d", got, want)
	}
}

func TestToGregorianYearIn(t *testing.T) {
	if got, want := ToGregorianYearIn(2569, ThaiEra), 2026; got != want {
		t.Errorf("ToGregorianYearIn(2569, ThaiEra) = %d, want %d", got, want)
	}
	if got, want := ToGregorianYearIn(2570, TheravadaEra), 2026; got != want {
		t.Errorf("ToGregorianYearIn(2570, TheravadaEra) = %d, want %d", got, want)
	}
}

func TestDefaultEraMatchesThaiEra(t *testing.T) {
	if got, want := ToBuddhistYear(2026), ToBuddhistYearIn(2026, ThaiEra); got != want {
		t.Errorf("ToBuddhistYear default = %d, want %d (ThaiEra)", got, want)
	}
}

func TestIsLeapYearIn(t *testing.T) {
	// 2024 CE is a leap year: ThaiEra 2567, TheravadaEra 2568.
	if !IsLeapYearIn(2567, ThaiEra) {
		t.Error("expected 2567 (ThaiEra) to be a leap year")
	}
	if !IsLeapYearIn(2568, TheravadaEra) {
		t.Error("expected 2568 (TheravadaEra) to be a leap year")
	}
	if IsLeapYearIn(2568, ThaiEra) {
		t.Error("did not expect 2568 (ThaiEra) to be a leap year")
	}
}

func TestDaysInMonthInAndDaysInYearIn(t *testing.T) {
	if got := DaysInMonthIn(2568, time.February, TheravadaEra); got != 29 {
		t.Errorf("DaysInMonthIn(2568, Feb, TheravadaEra) = %d, want 29", got)
	}
	if got := DaysInYearIn(2568, TheravadaEra); got != 366 {
		t.Errorf("DaysInYearIn(2568, TheravadaEra) = %d, want 366", got)
	}
}

func TestNewInAndYearIn(t *testing.T) {
	d := NewIn(TheravadaEra, 2570, time.August, 26)
	if d.Time().Year() != 2026 {
		t.Errorf("underlying Gregorian year = %d, want 2026", d.Time().Year())
	}
	if got := d.YearIn(TheravadaEra); got != 2570 {
		t.Errorf("YearIn(TheravadaEra) = %d, want 2570", got)
	}
	if got := d.YearIn(ThaiEra); got != 2569 {
		t.Errorf("YearIn(ThaiEra) = %d, want 2569", got)
	}
	// Year() defaults to ThaiEra.
	if got := d.Year(); got != 2569 {
		t.Errorf("Year() = %d, want 2569 (ThaiEra default)", got)
	}
}

func TestFromBuddhistTimeIn(t *testing.T) {
	d := FromBuddhistTimeIn(TheravadaEra, 2570, time.August, 26, 12, 0, 0, 0, time.UTC)
	if d.Time().Year() != 2026 {
		t.Errorf("underlying Gregorian year = %d, want 2026", d.Time().Year())
	}
}

func TestFormatIn(t *testing.T) {
	gt := time.Date(2026, time.August, 26, 0, 0, 0, 0, time.UTC)
	if got, want := FormatIn(gt, "02/01/2006", TheravadaEra), "26/08/2570"; got != want {
		t.Errorf("FormatIn = %q, want %q", got, want)
	}
	d := FromTime(gt)
	if got, want := d.FormatIn("02/01/2006", TheravadaEra), "26/08/2570"; got != want {
		t.Errorf("Date.FormatIn = %q, want %q", got, want)
	}
}

func TestParseIn(t *testing.T) {
	d, err := ParseIn("02/01/2006", "26/08/2570", TheravadaEra)
	if err != nil {
		t.Fatalf("ParseIn error: %v", err)
	}
	if d.Time().Year() != 2026 {
		t.Errorf("underlying Gregorian year = %d, want 2026", d.Time().Year())
	}

	// Leap-day validation should use the correct Gregorian year for the
	// given era: 2568 TheravadaEra = 2024 CE, a leap year.
	leap, err := ParseIn("02/01/2006", "29/02/2568", TheravadaEra)
	if err != nil {
		t.Fatalf("unexpected error parsing valid TheravadaEra leap date: %v", err)
	}
	if leap.Day() != 29 || leap.Month() != time.February {
		t.Errorf("got %v, want 29 February", leap)
	}
}

func TestFormatLocale(t *testing.T) {
	gt := time.Date(2026, time.August, 26, 0, 0, 0, 0, time.UTC)
	if got, want := FormatLocale(gt, ThaiLocale{}, ThaiEra), "26 สิงหาคม 2569"; got != want {
		t.Errorf("FormatLocale = %q, want %q", got, want)
	}
	if got, want := FormatLocale(gt, ThaiLocale{}, TheravadaEra), "26 สิงหาคม 2570"; got != want {
		t.Errorf("FormatLocale (TheravadaEra) = %q, want %q", got, want)
	}
}

// customLocale is a minimal Locale implementation used to verify that
// FormatLocale works with locales other than ThaiLocale.
type customLocale struct{}

func (customLocale) MonthName(month time.Month) string        { return "MONTH-" + month.String() }
func (customLocale) MonthShortName(month time.Month) string   { return "M" }
func (customLocale) WeekdayName(day time.Weekday) string      { return "DAY-" + day.String() }
func (customLocale) WeekdayShortName(day time.Weekday) string { return "D" }

func TestFormatLocaleCustomLocale(t *testing.T) {
	gt := time.Date(2026, time.August, 26, 0, 0, 0, 0, time.UTC)
	if got, want := FormatLocale(gt, customLocale{}, ThaiEra), "26 MONTH-August 2569"; got != want {
		t.Errorf("FormatLocale custom = %q, want %q", got, want)
	}
}
