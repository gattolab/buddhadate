package buddha

import "time"

// thaiMonthNames and thaiMonthShortNames are indexed by time.Month (1-12);
// index 0 is unused.
var thaiMonthNames = [...]string{
	"",
	"มกราคม",
	"กุมภาพันธ์",
	"มีนาคม",
	"เมษายน",
	"พฤษภาคม",
	"มิถุนายน",
	"กรกฎาคม",
	"สิงหาคม",
	"กันยายน",
	"ตุลาคม",
	"พฤศจิกายน",
	"ธันวาคม",
}

var thaiMonthShortNames = [...]string{
	"",
	"ม.ค.",
	"ก.พ.",
	"มี.ค.",
	"เม.ย.",
	"พ.ค.",
	"มิ.ย.",
	"ก.ค.",
	"ส.ค.",
	"ก.ย.",
	"ต.ค.",
	"พ.ย.",
	"ธ.ค.",
}

// thaiWeekdayNames and thaiWeekdayShortNames are indexed by time.Weekday
// (0 = Sunday .. 6 = Saturday).
var thaiWeekdayNames = [...]string{
	"วันอาทิตย์",
	"วันจันทร์",
	"วันอังคาร",
	"วันพุธ",
	"วันพฤหัสบดี",
	"วันศุกร์",
	"วันเสาร์",
}

var thaiWeekdayShortNames = [...]string{
	"อา.",
	"จ.",
	"อ.",
	"พ.",
	"พฤ.",
	"ศ.",
	"ส.",
}

// MonthName returns the full Thai name of month, for example "สิงหาคม" for
// time.August.
func MonthName(month time.Month) string {
	if month < time.January || month > time.December {
		return ""
	}
	return thaiMonthNames[month]
}

// MonthShortName returns the abbreviated Thai name of month, for example
// "ส.ค." for time.August.
func MonthShortName(month time.Month) string {
	if month < time.January || month > time.December {
		return ""
	}
	return thaiMonthShortNames[month]
}

// WeekdayName returns the full Thai name of day, for example "วันจันทร์"
// for time.Monday.
func WeekdayName(day time.Weekday) string {
	if day < time.Sunday || day > time.Saturday {
		return ""
	}
	return thaiWeekdayNames[day]
}

// WeekdayShortName returns the abbreviated Thai name of day, for example
// "จ." for time.Monday.
func WeekdayShortName(day time.Weekday) string {
	if day < time.Sunday || day > time.Saturday {
		return ""
	}
	return thaiWeekdayShortNames[day]
}

// ThaiLocale implements Locale using Thai month and weekday names. It is
// the locale used internally by FormatThai.
type ThaiLocale struct{}

// MonthName returns the full Thai name of month. See the package-level
// MonthName.
func (ThaiLocale) MonthName(month time.Month) string { return MonthName(month) }

// MonthShortName returns the abbreviated Thai name of month. See the
// package-level MonthShortName.
func (ThaiLocale) MonthShortName(month time.Month) string { return MonthShortName(month) }

// WeekdayName returns the full Thai name of day. See the package-level
// WeekdayName.
func (ThaiLocale) WeekdayName(day time.Weekday) string { return WeekdayName(day) }

// WeekdayShortName returns the abbreviated Thai name of day. See the
// package-level WeekdayShortName.
func (ThaiLocale) WeekdayShortName(day time.Weekday) string { return WeekdayShortName(day) }
