package buddha

import (
	"errors"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	gt := time.Date(2026, time.August, 26, 0, 0, 0, 0, time.UTC)
	if got, want := Format(gt, "02/01/2006"), "26/08/2569"; got != want {
		t.Errorf("Format = %q, want %q", got, want)
	}
}

func TestDateFormat(t *testing.T) {
	d := New(2569, time.August, 26)
	if got, want := d.Format("02/01/2006"), "26/08/2569"; got != want {
		t.Errorf("Date.Format = %q, want %q", got, want)
	}
}

func TestFormatDate(t *testing.T) {
	gt := time.Date(2026, time.August, 26, 0, 0, 0, 0, time.UTC)
	if got, want := FormatDate(gt), "26/08/2569"; got != want {
		t.Errorf("FormatDate = %q, want %q", got, want)
	}
}

func TestFormatDateTime(t *testing.T) {
	gt := time.Date(2026, time.August, 26, 14, 30, 0, 0, time.UTC)
	if got, want := FormatDateTime(gt), "26/08/2569 14:30:00"; got != want {
		t.Errorf("FormatDateTime = %q, want %q", got, want)
	}
}

func TestFormatShortDate(t *testing.T) {
	gt := time.Date(2026, time.August, 26, 0, 0, 0, 0, time.UTC)
	if got, want := FormatShortDate(gt), "26/08/69"; got != want {
		t.Errorf("FormatShortDate = %q, want %q", got, want)
	}
}

func TestFormatThai(t *testing.T) {
	gt := time.Date(2026, time.August, 26, 0, 0, 0, 0, time.UTC)
	if got, want := FormatThai(gt), "26 สิงหาคม 2569"; got != want {
		t.Errorf("FormatThai = %q, want %q", got, want)
	}
}

func TestParseBuddhist(t *testing.T) {
	d, err := ParseBuddhist("02/01/2006", "26/08/2569")
	if err != nil {
		t.Fatalf("ParseBuddhist error: %v", err)
	}
	if d.Year() != 2569 || d.Month() != time.August || d.Day() != 26 {
		t.Errorf("ParseBuddhist result = %v", d)
	}
	if d.Time().Year() != 2026 {
		t.Errorf("underlying Gregorian year = %d, want 2026", d.Time().Year())
	}
}

func TestParseGregorian(t *testing.T) {
	d, err := ParseGregorian("2006-01-02", "2026-08-26")
	if err != nil {
		t.Fatalf("ParseGregorian error: %v", err)
	}
	if d.Time().Year() != 2026 {
		t.Errorf("Gregorian year = %d, want 2026", d.Time().Year())
	}
	if d.Year() != 2569 {
		t.Errorf("Year() = %d, want 2569", d.Year())
	}
}

func TestParseIsParseBuddhist(t *testing.T) {
	d1, err1 := Parse("02/01/2006", "26/08/2569")
	d2, err2 := ParseBuddhist("02/01/2006", "26/08/2569")
	if err1 != nil || err2 != nil {
		t.Fatalf("unexpected errors: %v %v", err1, err2)
	}
	if !d1.Equal(d2) {
		t.Errorf("Parse and ParseBuddhist should agree: %v vs %v", d1, d2)
	}
}

func TestParseInvalidLayout(t *testing.T) {
	_, err := ParseBuddhist("02-01-2006", "26/08/2569")
	if !errors.Is(err, ErrInvalidFormat) {
		t.Errorf("expected ErrInvalidFormat, got %v", err)
	}
}

func TestParseInvalidDate(t *testing.T) {
	_, err := ParseBuddhist("02/01/2006", "31/02/2569")
	if !errors.Is(err, ErrInvalidFormat) {
		t.Errorf("expected ErrInvalidFormat for invalid date, got %v", err)
	}
}

func TestParseInvalidLeapDate(t *testing.T) {
	// 2569 B.E. = 2026 C.E., not a leap year, so Feb 29 is invalid.
	_, err := ParseBuddhist("02/01/2006", "29/02/2569")
	if !errors.Is(err, ErrInvalidFormat) {
		t.Errorf("expected ErrInvalidFormat for invalid leap date, got %v", err)
	}
}

func TestParseValidLeapDate(t *testing.T) {
	// 2567 B.E. = 2024 C.E., a leap year.
	d, err := ParseBuddhist("02/01/2006", "29/02/2567")
	if err != nil {
		t.Fatalf("unexpected error for valid leap date: %v", err)
	}
	if d.Day() != 29 || d.Month() != time.February {
		t.Errorf("got %v, want 29 February", d)
	}
}
