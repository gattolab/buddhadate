package buddha

import (
	"testing"
	"time"
)

func TestConversionRoundTrip(t *testing.T) {
	cases := []struct {
		gregorian int
		buddhist  int
	}{
		{2026, 2569},
		{2000, 2543},
		{1900, 2443},
	}
	for _, c := range cases {
		if got := ToBuddhistYear(c.gregorian); got != c.buddhist {
			t.Errorf("ToBuddhistYear(%d) = %d, want %d", c.gregorian, got, c.buddhist)
		}
		if got := ToGregorianYear(c.buddhist); got != c.gregorian {
			t.Errorf("ToGregorianYear(%d) = %d, want %d", c.buddhist, got, c.gregorian)
		}
	}
}

func TestIsLeapYear(t *testing.T) {
	cases := []struct {
		gregorian int
		leap      bool
	}{
		{2000, true},
		{1900, false},
		{2024, true},
		{2026, false},
	}
	for _, c := range cases {
		be := ToBuddhistYear(c.gregorian)
		if got := IsLeapYear(be); got != c.leap {
			t.Errorf("IsLeapYear(%d) [gregorian %d] = %v, want %v", be, c.gregorian, got, c.leap)
		}
	}
}

func TestDaysInMonth(t *testing.T) {
	if got := DaysInMonth(ToBuddhistYear(2024), time.February); got != 29 {
		t.Errorf("DaysInMonth(2567, Feb) = %d, want 29", got)
	}
	if got := DaysInMonth(ToBuddhistYear(2026), time.February); got != 28 {
		t.Errorf("DaysInMonth(2569, Feb) = %d, want 28", got)
	}
	if got := DaysInMonth(ToBuddhistYear(2026), time.January); got != 31 {
		t.Errorf("DaysInMonth(2569, Jan) = %d, want 31", got)
	}
}

func TestDaysInYear(t *testing.T) {
	if got := DaysInYear(ToBuddhistYear(2024)); got != 366 {
		t.Errorf("DaysInYear(2567) = %d, want 366", got)
	}
	if got := DaysInYear(ToBuddhistYear(2026)); got != 365 {
		t.Errorf("DaysInYear(2569) = %d, want 365", got)
	}
}
