package buddha

import (
	"testing"
	"time"
)

func TestMonthName(t *testing.T) {
	if got, want := MonthName(time.August), "สิงหาคม"; got != want {
		t.Errorf("MonthName(August) = %q, want %q", got, want)
	}
}

func TestMonthShortName(t *testing.T) {
	if got, want := MonthShortName(time.August), "ส.ค."; got != want {
		t.Errorf("MonthShortName(August) = %q, want %q", got, want)
	}
}

func TestWeekdayName(t *testing.T) {
	if got, want := WeekdayName(time.Monday), "วันจันทร์"; got != want {
		t.Errorf("WeekdayName(Monday) = %q, want %q", got, want)
	}
}

func TestWeekdayShortName(t *testing.T) {
	if got, want := WeekdayShortName(time.Monday), "จ."; got != want {
		t.Errorf("WeekdayShortName(Monday) = %q, want %q", got, want)
	}
}

func TestMonthNamesAllMonths(t *testing.T) {
	for m := time.January; m <= time.December; m++ {
		if MonthName(m) == "" {
			t.Errorf("MonthName(%v) is empty", m)
		}
		if MonthShortName(m) == "" {
			t.Errorf("MonthShortName(%v) is empty", m)
		}
	}
}

func TestWeekdayNamesAllDays(t *testing.T) {
	for d := time.Sunday; d <= time.Saturday; d++ {
		if WeekdayName(d) == "" {
			t.Errorf("WeekdayName(%v) is empty", d)
		}
		if WeekdayShortName(d) == "" {
			t.Errorf("WeekdayShortName(%v) is empty", d)
		}
	}
}

func TestThaiLocaleImplementsLocale(t *testing.T) {
	var loc Locale = ThaiLocale{}
	if got, want := loc.MonthName(time.August), "สิงหาคม"; got != want {
		t.Errorf("ThaiLocale.MonthName(August) = %q, want %q", got, want)
	}
	if got, want := loc.MonthShortName(time.August), "ส.ค."; got != want {
		t.Errorf("ThaiLocale.MonthShortName(August) = %q, want %q", got, want)
	}
	if got, want := loc.WeekdayName(time.Monday), "วันจันทร์"; got != want {
		t.Errorf("ThaiLocale.WeekdayName(Monday) = %q, want %q", got, want)
	}
	if got, want := loc.WeekdayShortName(time.Monday), "จ."; got != want {
		t.Errorf("ThaiLocale.WeekdayShortName(Monday) = %q, want %q", got, want)
	}
}
