<p align="center">
  <img
    src="https://cdn.tubb.me/static/buddhadate.jpeg"
    alt="Buddha"
    width="320"
  />
</p>

<h1 align="center">buddha</h1>

<p align="center">
  A small, dependency-free Go library for working with Buddhist Era
  (B.E./พ.ศ.) dates.
</p>

<p align="center">
  <a href="https://pkg.go.dev/github.com/gattolab/buddhadate"><img src="https://pkg.go.dev/badge/github.com/gattolab/buddhadate.svg" alt="Go Reference"></a>
  <a href="https://goreportcard.com/report/github.com/gattolab/buddhadate"><img src="https://goreportcard.com/badge/github.com/gattolab/buddhadate" alt="Go Report Card"></a>
  <a href="https://github.com/gattolab/buddhadate/blob/main/LICENSE"><img src="https://img.shields.io/github/license/gattolab/buddhadate" alt="License"></a>
</p>

```go
now := buddha.Now()

fmt.Println(now)
// 26/08/2569
```

`buddha` is serious underneath — it's just Go's `time` package plus 543
years — and a little playful on the surface.

---

## Features

- Zero external dependencies; built entirely on the standard `time` package.
- No custom calendar implementation — Buddhist Era is always the Gregorian
  year plus a fixed offset.
- Explicit, unambiguous APIs: no function silently guesses whether a year is
  Buddhist or Gregorian, and no function silently guesses which country's
  Buddhist Era convention applies.
- Defaults to Thailand's convention (+543, Thai names), with explicit
  `Era`/`"*In"` functions and a pluggable `Locale` interface for other
  Buddhist-Era-using countries (Sri Lanka, Myanmar, Cambodia, Laos, ...).
- Buddhist-aware formatting and parsing.
- Thai month/weekday localization.
- Date boundary helpers (start/end of day, month, year).
- A few deliberately fun, clearly-documented "traditional date" APIs.

## Installation

```bash
go get github.com/gattolab/buddhadate
```

```go
import "github.com/gattolab/buddhadate"
```

```bash
go mod tidy
```

## Quick Start

```go
package main

import (
	"fmt"

	"github.com/gattolab/buddhadate"
)

func main() {
	fmt.Println(buddha.Now())
}
```

```go
now := buddha.Now()

fmt.Println(now.Year())
// 2569

fmt.Println(now.Format("02/01/2006"))
// 26/08/2569
```

```go
gregorian := time.Date(
	2026,
	time.August,
	26,
	0, 0, 0, 0,
	time.Local,
)

d := buddha.FromTime(gregorian)

fmt.Println(d)
// 26/08/2569
```

## Usage

### Current date

```go
buddha.Now()
buddha.Today()

buddha.NowIn(time.UTC)
buddha.TodayIn(time.FixedZone("ICT", 7*3600))
```

### Buddhist Era Conversion

```go
buddha.ToBuddhistYear(2026)
// 2569

buddha.ToGregorianYear(2569)
// 2026
```

### Formatting

```go
buddha.FormatDate(time.Now())
// 26/08/2569

buddha.FormatDateTime(time.Now())
// 26/08/2569 14:30:00

buddha.FormatShortDate(time.Now())
// 26/08/69

buddha.FormatThai(time.Now())
// 26 สิงหาคม 2569
```

`buddha.Format` follows the same reference layout rules as
`time.Time.Format`, except that the `2006` and `06` year tokens are
replaced with the Buddhist Era year:

```go
buddha.Format(
	time.Date(2026, 8, 26, 0, 0, 0, 0, time.Local),
	"02/01/2006",
)
// 26/08/2569
```

### Parsing

`Parse`/`ParseBuddhist` and `ParseGregorian` are always explicit about which
calendar era the year in the input string represents:

```go
buddha.ParseBuddhist(
	"02/01/2006",
	"26/08/2569",
)

buddha.ParseGregorian(
	"2006-01-02",
	"2026-08-26",
)
```

### Thai Localization

```go
buddha.MonthName(time.August)
// สิงหาคม

buddha.MonthShortName(time.August)
// ส.ค.

buddha.WeekdayName(time.Monday)
// วันจันทร์

buddha.WeekdayShortName(time.Monday)
// จ.
```

### Other Countries / Other Buddhist Era Conventions

"The" Buddhist Era isn't a single universal system. Every function above
defaults to **Thailand's** convention: a +543 year offset (`ThaiEra`) and
Thai month/weekday names (`ThaiLocale`). Sri Lanka, Myanmar, Cambodia, and
Laos also use a Buddhist Era, but commonly count one year earlier (+544),
and each has its own language for month/weekday names.

Rather than guessing, `buddha` makes the convention explicit:

```go
// Year offset: use the "*In" functions with an explicit Era.
buddha.ToBuddhistYearIn(2026, buddha.TheravadaEra)
// 2570

d := buddha.NewIn(buddha.TheravadaEra, 2570, time.August, 26)
d.YearIn(buddha.TheravadaEra)
// 2570

buddha.FormatIn(time.Now(), "02/01/2006", buddha.TheravadaEra)
buddha.ParseIn("02/01/2006", "26/08/2570", buddha.TheravadaEra)
```

For a language other than Thai, implement the small `Locale` interface
with the correct month/weekday names for that language and pass it to
`FormatLocale` — `buddha` intentionally does not ship unverified
translations for languages it can't guarantee the accuracy of:

```go
type khmerLocale struct{ /* ... */ }

func (khmerLocale) MonthName(m time.Month) string { /* ... */ }
// ... MonthShortName, WeekdayName, WeekdayShortName

buddha.FormatLocale(time.Now(), khmerLocale{}, buddha.TheravadaEra)
```

### Date Utilities

```go
buddha.IsLeapYear(2569)
buddha.DaysInMonth(2569, time.February)
buddha.DaysInYear(2569)

buddha.StartOfMonth(d)
buddha.EndOfMonth(d)

d.StartOfDay()
d.EndOfDay()
d.StartOfYear()
d.EndOfYear()
```

### Buddhist Special Dates

```go
fmt.Println(buddha.BirthDate())
fmt.Println(buddha.EnlightenmentDate())
fmt.Println(buddha.DeathDate())
fmt.Println(buddha.VisakhaBucha())
```

These follow one commonly cited traditional Thai narrative and are **not**
historically verified Gregorian dates — see the doc comments on each
function for the exact convention used.

## API Overview

```go
// Core
Now()
Today()
NowIn()
TodayIn()

// Date
New()
NewIn()
FromTime()
FromBuddhistTime()
FromBuddhistTimeIn()
ToTime()

Date.Year()
Date.YearIn()
Date.Month()
Date.Day()
Date.Hour()
Date.Minute()
Date.Second()
Date.Location()
Date.Time()
Date.Add()
Date.AddDate()
Date.Before()
Date.After()
Date.Equal()
Date.Format()
Date.FormatIn()

// Conversion
ToBuddhistYear()   / ToBuddhistYearIn()
ToGregorianYear()  / ToGregorianYearIn()

// Formatting
Format()   / FormatIn()
FormatDate()
FormatDateTime()
FormatShortDate()
FormatThai()
FormatLocale()

// Parsing
Parse()
ParseBuddhist()
ParseGregorian()
ParseIn()

// Calendar
IsLeapYear()   / IsLeapYearIn()
DaysInMonth()  / DaysInMonthIn()
DaysInYear()   / DaysInYearIn()

// Thai
MonthName()
MonthShortName()
WeekdayName()
WeekdayShortName()
ThaiLocale{}

// Era & Locale (other Buddhist-Era-using countries)
Era, ThaiEra, TheravadaEra
Locale

// Date boundaries
StartOfDay() / Date.StartOfDay()
EndOfDay()   / Date.EndOfDay()
StartOfMonth() / Date.StartOfMonth()
EndOfMonth()   / Date.EndOfMonth()
StartOfYear()  / Date.StartOfYear()
EndOfYear()    / Date.EndOfYear()

// Buddha
BirthDate()
EnlightenmentDate()
DeathDate()
VisakhaBucha()
```

Full documentation with examples is available on
[pkg.go.dev](https://pkg.go.dev/github.com/gattolab/buddhadate).

## Design Philosophy

`buddha` does not replace Go's `time` package — it sits on top of it:

```text
time.Time
    ↓
buddha.FromTime()
    ↓
buddha.Date
    ↓
Buddhist Era presentation
```

A `Date` always keeps its underlying Gregorian `time.Time` intact. The
Buddhist Era year only comes into play when you call `Year()`, `Format()`,
`String()`, or one of the `ParseBuddhist`/`Format*` helpers. Every
computation — leap years, days in month, duration arithmetic, timezones —
is delegated to the standard library; `buddha` never implements its own
calendar rules.

## Compatibility

```text
Go version:  1.26.5+ (tested with the Go toolchain installed in this repo)
Timezone:    uses Go's time.Location; no implicit timezone behavior
Calendar:    Gregorian calendar with Buddhist Era year representation
Dependencies: none
```

## Examples

Runnable, `go test`-verified examples live in
[`examples_test.go`](./examples_test.go) and appear in the generated
documentation on pkg.go.dev.

## Contributing

Issues and pull requests are welcome. Before submitting a change, please run:

```bash
go fmt ./...
go vet ./...
go test ./...
go test -race ./...
```

## License

[MIT](./LICENSE)
