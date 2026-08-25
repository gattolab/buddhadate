package buddha

import (
	"testing"
	"time"
)

func TestNewConstructsExpectedDate(t *testing.T) {
	d := New(2569, time.August, 26)
	if d.Year() != 2569 {
		t.Errorf("Year() = %d, want 2569", d.Year())
	}
	if d.Month() != time.August {
		t.Errorf("Month() = %v, want August", d.Month())
	}
	if d.Day() != 26 {
		t.Errorf("Day() = %d, want 26", d.Day())
	}
	if d.Time().Year() != 2026 {
		t.Errorf("underlying Gregorian year = %d, want 2026", d.Time().Year())
	}
}

func TestDateConstruction(t *testing.T) {
	cases := []struct {
		name        string
		year, day   int
		month       time.Month
		wantInvalid bool
	}{
		{"start of year", 2569, 1, time.January, false},
		{"end of year", 2569, 31, time.December, false},
		{"leap day leap year", ToBuddhistYear(2024), 29, time.February, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			d := New(c.year, c.month, c.day)
			if d.Day() != c.day || d.Month() != c.month {
				t.Errorf("got %02d/%02d, want %02d/%v", d.Day(), int(d.Month()), c.day, c.month)
			}
		})
	}
}

func TestFromTimeAndToTime(t *testing.T) {
	gt := time.Date(2026, time.August, 26, 14, 30, 0, 0, time.UTC)
	d := FromTime(gt)
	if d.Year() != 2569 {
		t.Errorf("Year() = %d, want 2569", d.Year())
	}
	if got := ToTime(d); !got.Equal(gt) {
		t.Errorf("ToTime(d) = %v, want %v", got, gt)
	}
}

func TestFromBuddhistTime(t *testing.T) {
	d := FromBuddhistTime(2569, time.August, 26, 14, 30, 15, 0, time.UTC)
	if d.Year() != 2569 || d.Month() != time.August || d.Day() != 26 {
		t.Fatalf("unexpected date: %v", d)
	}
	if d.Hour() != 14 || d.Minute() != 30 || d.Second() != 15 {
		t.Errorf("unexpected time: %02d:%02d:%02d", d.Hour(), d.Minute(), d.Second())
	}
}

func TestDateAccessors(t *testing.T) {
	loc := time.FixedZone("TEST", 7*3600)
	d := FromBuddhistTime(2569, time.August, 26, 1, 2, 3, 0, loc)
	if d.Location() != loc {
		t.Errorf("Location() = %v, want %v", d.Location(), loc)
	}
}

func TestBeforeAfterEqual(t *testing.T) {
	a := New(2569, time.August, 26)
	b := New(2569, time.August, 27)
	if !a.Before(b) {
		t.Error("expected a.Before(b)")
	}
	if !b.After(a) {
		t.Error("expected b.After(a)")
	}
	if a.Equal(b) {
		t.Error("did not expect a.Equal(b)")
	}
	c := New(2569, time.August, 26)
	if !a.Equal(c) {
		t.Error("expected a.Equal(c)")
	}
}

func TestAddAndAddDate(t *testing.T) {
	d := New(2569, time.August, 26)
	added := d.Add(24 * time.Hour)
	if added.Day() != 27 {
		t.Errorf("Add(24h).Day() = %d, want 27", added.Day())
	}

	addedYear := d.AddDate(1, 0, 0)
	if addedYear.Year() != 2570 {
		t.Errorf("AddDate(1,0,0).Year() = %d, want 2570", addedYear.Year())
	}
}

func TestNowTodayNotZero(t *testing.T) {
	if Now().Time().IsZero() {
		t.Error("Now() returned zero time")
	}
	today := Today()
	if today.Hour() != 0 || today.Minute() != 0 || today.Second() != 0 {
		t.Errorf("Today() should be at midnight, got %02d:%02d:%02d", today.Hour(), today.Minute(), today.Second())
	}
}

func TestNowInTodayIn(t *testing.T) {
	loc := time.FixedZone("TEST", 3*3600)
	n := NowIn(loc)
	if n.Location().String() != loc.String() {
		t.Errorf("NowIn location = %v, want %v", n.Location(), loc)
	}
	td := TodayIn(loc)
	if td.Hour() != 0 || td.Minute() != 0 || td.Second() != 0 {
		t.Errorf("TodayIn should be at midnight, got %02d:%02d:%02d", td.Hour(), td.Minute(), td.Second())
	}
}

func TestStartEndOfDay(t *testing.T) {
	d := FromBuddhistTime(2569, time.August, 26, 14, 30, 0, 0, time.UTC)
	start := d.StartOfDay()
	if start.Hour() != 0 || start.Minute() != 0 || start.Second() != 0 {
		t.Errorf("StartOfDay = %v, want midnight", start)
	}
	end := d.EndOfDay()
	if end.Hour() != 23 || end.Minute() != 59 || end.Second() != 59 {
		t.Errorf("EndOfDay = %v, want 23:59:59", end)
	}

	if got := StartOfDay(d); got.Hour() != 0 {
		t.Errorf("package StartOfDay mismatch")
	}
	if got := EndOfDay(d); got.Hour() != 23 {
		t.Errorf("package EndOfDay mismatch")
	}
}

func TestStartEndOfMonth(t *testing.T) {
	d := New(2569, time.February, 15)
	start := d.StartOfMonth()
	if start.Day() != 1 || start.Month() != time.February {
		t.Errorf("StartOfMonth = %v, want Feb 1", start)
	}
	end := d.EndOfMonth()
	if end.Day() != 28 || end.Month() != time.February {
		t.Errorf("EndOfMonth = %v, want Feb 28", end)
	}

	leap := New(ToBuddhistYear(2024), time.February, 10)
	leapEnd := leap.EndOfMonth()
	if leapEnd.Day() != 29 {
		t.Errorf("leap EndOfMonth day = %d, want 29", leapEnd.Day())
	}

	if got := StartOfMonth(d); got.Day() != 1 {
		t.Errorf("package StartOfMonth mismatch")
	}
	if got := EndOfMonth(d); got.Day() != 28 {
		t.Errorf("package EndOfMonth mismatch")
	}
}

func TestStartEndOfYear(t *testing.T) {
	d := New(2569, time.August, 26)
	start := d.StartOfYear()
	if start.Month() != time.January || start.Day() != 1 {
		t.Errorf("StartOfYear = %v, want Jan 1", start)
	}
	end := d.EndOfYear()
	if end.Month() != time.December || end.Day() != 31 {
		t.Errorf("EndOfYear = %v, want Dec 31", end)
	}

	if got := StartOfYear(d); got.Month() != time.January {
		t.Errorf("package StartOfYear mismatch")
	}
	if got := EndOfYear(d); got.Month() != time.December {
		t.Errorf("package EndOfYear mismatch")
	}
}

func TestDateString(t *testing.T) {
	d := New(2569, time.August, 26)
	if got, want := d.String(), "26/08/2569"; got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}
