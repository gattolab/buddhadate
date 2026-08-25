package buddha

import (
	"testing"
	"time"
)

func TestBirthDate(t *testing.T) {
	d := BirthDate()
	if d.Month() != time.May || d.Day() != 15 {
		t.Errorf("BirthDate() = %v, want 15 May", d)
	}
}

func TestEnlightenmentDate(t *testing.T) {
	d := EnlightenmentDate()
	if d.Month() != time.May || d.Day() != 15 {
		t.Errorf("EnlightenmentDate() = %v, want 15 May", d)
	}
}

func TestDeathDate(t *testing.T) {
	d := DeathDate()
	if d.Month() != time.May || d.Day() != 15 {
		t.Errorf("DeathDate() = %v, want 15 May", d)
	}
	if d.Time().Year() != -542 {
		t.Errorf("DeathDate() Gregorian year = %d, want -542", d.Time().Year())
	}
	if d.Year() != 1 {
		t.Errorf("DeathDate().Year() = %d, want 1 (start of the Buddhist Era)", d.Year())
	}
}

func TestSpecialDatesConsistentAges(t *testing.T) {
	birth := BirthDate().Time().Year()
	enlightenment := EnlightenmentDate().Time().Year()
	death := DeathDate().Time().Year()

	if enlightenment-birth != 35 {
		t.Errorf("age at enlightenment = %d, want 35", enlightenment-birth)
	}
	if death-birth != 80 {
		t.Errorf("age at death = %d, want 80", death-birth)
	}
}

func TestVisakhaBucha(t *testing.T) {
	d := VisakhaBucha()
	if d.Month() != time.May || d.Day() != 15 {
		t.Errorf("VisakhaBucha() = %v, want 15 May", d)
	}
	if d.Year() != ToBuddhistYear(time.Now().Year()) {
		t.Errorf("VisakhaBucha() year = %d, want current Buddhist Era year", d.Year())
	}
}
