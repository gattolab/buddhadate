package buddha_test

import (
	"fmt"
	"time"

	"github.com/gattolab/buddhadate"
)

func ExampleNow() {
	now := buddha.Now()
	fmt.Println(now.Year() >= 2500)
	// Output: true
}

func ExampleFromTime() {
	gregorian := time.Date(2026, time.August, 26, 0, 0, 0, 0, time.UTC)
	d := buddha.FromTime(gregorian)
	fmt.Println(d)
	// Output: 26/08/2569
}

func ExampleToBuddhistYear() {
	fmt.Println(buddha.ToBuddhistYear(2026))
	// Output: 2569
}

func ExampleDate_Format() {
	d := buddha.New(2569, time.August, 26)
	fmt.Println(d.Format("02/01/2006"))
	// Output: 26/08/2569
}

func ExampleParseBuddhist() {
	d, err := buddha.ParseBuddhist("02/01/2006", "26/08/2569")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
	// Output: 26/08/2569
}

func ExampleBirthDate() {
	d := buddha.BirthDate()
	fmt.Println(d.Month(), d.Day())
	// Output: May 15
}

func ExampleToBuddhistYearIn() {
	fmt.Println(buddha.ToBuddhistYearIn(2026, buddha.ThaiEra))
	fmt.Println(buddha.ToBuddhistYearIn(2026, buddha.TheravadaEra))
	// Output:
	// 2569
	// 2570
}

func ExampleFormatLocale() {
	gregorian := time.Date(2026, time.August, 26, 0, 0, 0, 0, time.UTC)
	fmt.Println(buddha.FormatLocale(gregorian, buddha.ThaiLocale{}, buddha.ThaiEra))
	// Output: 26 สิงหาคม 2569
}
