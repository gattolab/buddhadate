# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased

### Added

- Initial Buddhist Era date utilities.
- `Date` type wrapping `time.Time`, with `Now`, `Today`, `NowIn`, `TodayIn`,
  `New`, `FromTime`, and `FromBuddhistTime` constructors.
- Gregorian ↔ Buddhist Era year conversion (`ToBuddhistYear`,
  `ToGregorianYear`).
- Buddhist Era aware formatting (`Format`, `FormatDate`, `FormatDateTime`,
  `FormatShortDate`, `FormatThai`) and parsing (`Parse`, `ParseBuddhist`,
  `ParseGregorian`).
- Calendar helpers (`IsLeapYear`, `DaysInMonth`, `DaysInYear`) and
  start/end-of period utilities (day, month, year).
- Thai month and weekday localization (`MonthName`, `MonthShortName`,
  `WeekdayName`, `WeekdayShortName`).
- Playful conventional Buddhist calendar dates (`BirthDate`,
  `EnlightenmentDate`, `DeathDate`, `VisakhaBucha`).
- `Era` type with `ThaiEra` (+543, default) and `TheravadaEra` (+544, used
  by Sri Lanka, Myanmar, Cambodia, and Laos), plus explicit `"*In"`
  variants of the era-sensitive functions: `ToBuddhistYearIn`,
  `ToGregorianYearIn`, `NewIn`, `FromBuddhistTimeIn`, `Date.YearIn`,
  `FormatIn`, `Date.FormatIn`, `ParseIn`, `IsLeapYearIn`, `DaysInMonthIn`,
  `DaysInYearIn`.
- `Locale` interface and `ThaiLocale` implementation, plus `FormatLocale`,
  so other Buddhist-Era-using countries can supply their own month/weekday
  names instead of relying on unverified translations bundled by this
  package.
