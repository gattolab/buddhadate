package buddha_test

import (
	"testing"
	"time"

	"github.com/OWNER/buddha"
)

// This file contains boundary-condition and cross-cutting tests that
// exercise the package's public API as an external consumer would.

func TestYearBoundary(t *testing.T) {
	d := buddha.New(2569, time.December, 31)
	next := d.AddDate(0, 0, 1)
	if next.Year() != 2570 || next.Month() != time.January || next.Day() != 1 {
		t.Errorf("crossing year boundary: got %v", next)
	}
}

func TestMonthBoundary(t *testing.T) {
	d := buddha.New(2569, time.August, 31)
	next := d.AddDate(0, 0, 1)
	if next.Month() != time.September || next.Day() != 1 {
		t.Errorf("crossing month boundary: got %v", next)
	}
}

func TestDayBoundaryMidnight(t *testing.T) {
	d := buddha.FromBuddhistTime(2569, time.August, 26, 23, 59, 59, 0, time.UTC)
	next := d.Add(time.Second)
	if next.Day() != 27 || next.Hour() != 0 || next.Minute() != 0 || next.Second() != 0 {
		t.Errorf("crossing midnight boundary: got %v %02d:%02d:%02d", next, next.Hour(), next.Minute(), next.Second())
	}
}

func TestUTCAndNonUTCLocations(t *testing.T) {
	utc := buddha.FromBuddhistTime(2569, time.August, 26, 12, 0, 0, 0, time.UTC)
	if utc.Location() != time.UTC {
		t.Errorf("expected UTC location, got %v", utc.Location())
	}

	bangkok := time.FixedZone("Asia/Bangkok", 7*3600)
	local := buddha.FromBuddhistTime(2569, time.August, 26, 12, 0, 0, 0, bangkok)
	if local.Location().String() != bangkok.String() {
		t.Errorf("expected %v location, got %v", bangkok, local.Location())
	}
}
