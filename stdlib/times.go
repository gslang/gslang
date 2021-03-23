package stdlib

import (
	"time"

	"github.com/gslang/gslang"
)

var timesModule = map[string]gslang.Object{
	"format_ansic":        &gslang.String{Value: time.ANSIC},
	"format_unix_date":    &gslang.String{Value: time.UnixDate},
	"format_ruby_date":    &gslang.String{Value: time.RubyDate},
	"format_rfc822":       &gslang.String{Value: time.RFC822},
	"format_rfc822z":      &gslang.String{Value: time.RFC822Z},
	"format_rfc850":       &gslang.String{Value: time.RFC850},
	"format_rfc1123":      &gslang.String{Value: time.RFC1123},
	"format_rfc1123z":     &gslang.String{Value: time.RFC1123Z},
	"format_rfc3339":      &gslang.String{Value: time.RFC3339},
	"format_rfc3339_nano": &gslang.String{Value: time.RFC3339Nano},
	"format_kitchen":      &gslang.String{Value: time.Kitchen},
	"format_stamp":        &gslang.String{Value: time.Stamp},
	"format_stamp_milli":  &gslang.String{Value: time.StampMilli},
	"format_stamp_micro":  &gslang.String{Value: time.StampMicro},
	"format_stamp_nano":   &gslang.String{Value: time.StampNano},
	"nanosecond":          &gslang.Int{Value: int64(time.Nanosecond)},
	"microsecond":         &gslang.Int{Value: int64(time.Microsecond)},
	"millisecond":         &gslang.Int{Value: int64(time.Millisecond)},
	"second":              &gslang.Int{Value: int64(time.Second)},
	"minute":              &gslang.Int{Value: int64(time.Minute)},
	"hour":                &gslang.Int{Value: int64(time.Hour)},
	"january":             &gslang.Int{Value: int64(time.January)},
	"february":            &gslang.Int{Value: int64(time.February)},
	"march":               &gslang.Int{Value: int64(time.March)},
	"april":               &gslang.Int{Value: int64(time.April)},
	"may":                 &gslang.Int{Value: int64(time.May)},
	"june":                &gslang.Int{Value: int64(time.June)},
	"july":                &gslang.Int{Value: int64(time.July)},
	"august":              &gslang.Int{Value: int64(time.August)},
	"september":           &gslang.Int{Value: int64(time.September)},
	"october":             &gslang.Int{Value: int64(time.October)},
	"november":            &gslang.Int{Value: int64(time.November)},
	"december":            &gslang.Int{Value: int64(time.December)},
	"sleep": &gslang.UserFunction{
		Name:  "sleep",
		Value: timesSleep,
	}, // sleep(int)
	"parse_duration": &gslang.UserFunction{
		Name:  "parse_duration",
		Value: timesParseDuration,
	}, // parse_duration(str) => int
	"since": &gslang.UserFunction{
		Name:  "since",
		Value: timesSince,
	}, // since(time) => int
	"until": &gslang.UserFunction{
		Name:  "until",
		Value: timesUntil,
	}, // until(time) => int
	"duration_hours": &gslang.UserFunction{
		Name:  "duration_hours",
		Value: timesDurationHours,
	}, // duration_hours(int) => float
	"duration_minutes": &gslang.UserFunction{
		Name:  "duration_minutes",
		Value: timesDurationMinutes,
	}, // duration_minutes(int) => float
	"duration_nanoseconds": &gslang.UserFunction{
		Name:  "duration_nanoseconds",
		Value: timesDurationNanoseconds,
	}, // duration_nanoseconds(int) => int
	"duration_seconds": &gslang.UserFunction{
		Name:  "duration_seconds",
		Value: timesDurationSeconds,
	}, // duration_seconds(int) => float
	"duration_string": &gslang.UserFunction{
		Name:  "duration_string",
		Value: timesDurationString,
	}, // duration_string(int) => string
	"month_string": &gslang.UserFunction{
		Name:  "month_string",
		Value: timesMonthString,
	}, // month_string(int) => string
	"date": &gslang.UserFunction{
		Name:  "date",
		Value: timesDate,
	}, // date(year, month, day, hour, min, sec, nsec) => time
	"now": &gslang.UserFunction{
		Name:  "now",
		Value: timesNow,
	}, // now() => time
	"parse": &gslang.UserFunction{
		Name:  "parse",
		Value: timesParse,
	}, // parse(format, str) => time
	"unix": &gslang.UserFunction{
		Name:  "unix",
		Value: timesUnix,
	}, // unix(sec, nsec) => time
	"add": &gslang.UserFunction{
		Name:  "add",
		Value: timesAdd,
	}, // add(time, int) => time
	"add_date": &gslang.UserFunction{
		Name:  "add_date",
		Value: timesAddDate,
	}, // add_date(time, years, months, days) => time
	"sub": &gslang.UserFunction{
		Name:  "sub",
		Value: timesSub,
	}, // sub(t time, u time) => int
	"after": &gslang.UserFunction{
		Name:  "after",
		Value: timesAfter,
	}, // after(t time, u time) => bool
	"before": &gslang.UserFunction{
		Name:  "before",
		Value: timesBefore,
	}, // before(t time, u time) => bool
	"time_year": &gslang.UserFunction{
		Name:  "time_year",
		Value: timesTimeYear,
	}, // time_year(time) => int
	"time_month": &gslang.UserFunction{
		Name:  "time_month",
		Value: timesTimeMonth,
	}, // time_month(time) => int
	"time_day": &gslang.UserFunction{
		Name:  "time_day",
		Value: timesTimeDay,
	}, // time_day(time) => int
	"time_weekday": &gslang.UserFunction{
		Name:  "time_weekday",
		Value: timesTimeWeekday,
	}, // time_weekday(time) => int
	"time_hour": &gslang.UserFunction{
		Name:  "time_hour",
		Value: timesTimeHour,
	}, // time_hour(time) => int
	"time_minute": &gslang.UserFunction{
		Name:  "time_minute",
		Value: timesTimeMinute,
	}, // time_minute(time) => int
	"time_second": &gslang.UserFunction{
		Name:  "time_second",
		Value: timesTimeSecond,
	}, // time_second(time) => int
	"time_nanosecond": &gslang.UserFunction{
		Name:  "time_nanosecond",
		Value: timesTimeNanosecond,
	}, // time_nanosecond(time) => int
	"time_unix": &gslang.UserFunction{
		Name:  "time_unix",
		Value: timesTimeUnix,
	}, // time_unix(time) => int
	"time_unix_nano": &gslang.UserFunction{
		Name:  "time_unix_nano",
		Value: timesTimeUnixNano,
	}, // time_unix_nano(time) => int
	"time_format": &gslang.UserFunction{
		Name:  "time_format",
		Value: timesTimeFormat,
	}, // time_format(time, format) => string
	"time_location": &gslang.UserFunction{
		Name:  "time_location",
		Value: timesTimeLocation,
	}, // time_location(time) => string
	"time_string": &gslang.UserFunction{
		Name:  "time_string",
		Value: timesTimeString,
	}, // time_string(time) => string
	"is_zero": &gslang.UserFunction{
		Name:  "is_zero",
		Value: timesIsZero,
	}, // is_zero(time) => bool
	"to_local": &gslang.UserFunction{
		Name:  "to_local",
		Value: timesToLocal,
	}, // to_local(time) => time
	"to_utc": &gslang.UserFunction{
		Name:  "to_utc",
		Value: timesToUTC,
	}, // to_utc(time) => time
}

func timesSleep(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	i1, ok := gslang.ToInt64(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	time.Sleep(time.Duration(i1))
	ret = gslang.UndefinedValue

	return
}

func timesParseDuration(args ...gslang.Object) (
	ret gslang.Object,
	err error,
) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	s1, ok := gslang.ToString(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	dur, err := time.ParseDuration(s1)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &gslang.Int{Value: int64(dur)}

	return
}

func timesSince(args ...gslang.Object) (
	ret gslang.Object,
	err error,
) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.Int{Value: int64(time.Since(t1))}

	return
}

func timesUntil(args ...gslang.Object) (
	ret gslang.Object,
	err error,
) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.Int{Value: int64(time.Until(t1))}

	return
}

func timesDurationHours(args ...gslang.Object) (
	ret gslang.Object,
	err error,
) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	i1, ok := gslang.ToInt64(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.Float{Value: time.Duration(i1).Hours()}

	return
}

func timesDurationMinutes(args ...gslang.Object) (
	ret gslang.Object,
	err error,
) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	i1, ok := gslang.ToInt64(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.Float{Value: time.Duration(i1).Minutes()}

	return
}

func timesDurationNanoseconds(args ...gslang.Object) (
	ret gslang.Object,
	err error,
) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	i1, ok := gslang.ToInt64(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.Int{Value: time.Duration(i1).Nanoseconds()}

	return
}

func timesDurationSeconds(args ...gslang.Object) (
	ret gslang.Object,
	err error,
) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	i1, ok := gslang.ToInt64(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.Float{Value: time.Duration(i1).Seconds()}

	return
}

func timesDurationString(args ...gslang.Object) (
	ret gslang.Object,
	err error,
) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	i1, ok := gslang.ToInt64(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.String{Value: time.Duration(i1).String()}

	return
}

func timesMonthString(args ...gslang.Object) (
	ret gslang.Object,
	err error,
) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	i1, ok := gslang.ToInt64(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.String{Value: time.Month(i1).String()}

	return
}

func timesDate(args ...gslang.Object) (
	ret gslang.Object,
	err error,
) {
	if len(args) != 7 {
		err = gslang.ErrWrongNumArguments
		return
	}

	i1, ok := gslang.ToInt(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}
	i2, ok := gslang.ToInt(args[1])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}
	i3, ok := gslang.ToInt(args[2])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}
	i4, ok := gslang.ToInt(args[3])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}
	i5, ok := gslang.ToInt(args[4])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "fifth",
			Expected: "int(compatible)",
			Found:    args[4].TypeName(),
		}
		return
	}
	i6, ok := gslang.ToInt(args[5])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "sixth",
			Expected: "int(compatible)",
			Found:    args[5].TypeName(),
		}
		return
	}
	i7, ok := gslang.ToInt(args[6])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "seventh",
			Expected: "int(compatible)",
			Found:    args[6].TypeName(),
		}
		return
	}

	ret = &gslang.Time{
		Value: time.Date(i1,
			time.Month(i2), i3, i4, i5, i6, i7, time.Now().Location()),
	}

	return
}

func timesNow(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 0 {
		err = gslang.ErrWrongNumArguments
		return
	}

	ret = &gslang.Time{Value: time.Now()}

	return
}

func timesParse(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 2 {
		err = gslang.ErrWrongNumArguments
		return
	}

	s1, ok := gslang.ToString(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := gslang.ToString(args[1])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	parsed, err := time.Parse(s1, s2)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &gslang.Time{Value: parsed}

	return
}

func timesUnix(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 2 {
		err = gslang.ErrWrongNumArguments
		return
	}

	i1, ok := gslang.ToInt64(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := gslang.ToInt64(args[1])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &gslang.Time{Value: time.Unix(i1, i2)}

	return
}

func timesAdd(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 2 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := gslang.ToInt64(args[1])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &gslang.Time{Value: t1.Add(time.Duration(i2))}

	return
}

func timesSub(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 2 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := gslang.ToTime(args[1])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &gslang.Int{Value: int64(t1.Sub(t2))}

	return
}

func timesAddDate(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 4 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := gslang.ToInt(args[1])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	i3, ok := gslang.ToInt(args[2])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	i4, ok := gslang.ToInt(args[3])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}

	ret = &gslang.Time{Value: t1.AddDate(i2, i3, i4)}

	return
}

func timesAfter(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 2 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := gslang.ToTime(args[1])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	if t1.After(t2) {
		ret = gslang.TrueValue
	} else {
		ret = gslang.FalseValue
	}

	return
}

func timesBefore(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 2 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := gslang.ToTime(args[1])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	if t1.Before(t2) {
		ret = gslang.TrueValue
	} else {
		ret = gslang.FalseValue
	}

	return
}

func timesTimeYear(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.Int{Value: int64(t1.Year())}

	return
}

func timesTimeMonth(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.Int{Value: int64(t1.Month())}

	return
}

func timesTimeDay(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.Int{Value: int64(t1.Day())}

	return
}

func timesTimeWeekday(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.Int{Value: int64(t1.Weekday())}

	return
}

func timesTimeHour(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.Int{Value: int64(t1.Hour())}

	return
}

func timesTimeMinute(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.Int{Value: int64(t1.Minute())}

	return
}

func timesTimeSecond(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.Int{Value: int64(t1.Second())}

	return
}

func timesTimeNanosecond(args ...gslang.Object) (
	ret gslang.Object,
	err error,
) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.Int{Value: int64(t1.Nanosecond())}

	return
}

func timesTimeUnix(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.Int{Value: t1.Unix()}

	return
}

func timesTimeUnixNano(args ...gslang.Object) (
	ret gslang.Object,
	err error,
) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.Int{Value: t1.UnixNano()}

	return
}

func timesTimeFormat(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 2 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := gslang.ToString(args[1])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	s := t1.Format(s2)
	if len(s) > gslang.MaxStringLen {

		return nil, gslang.ErrStringLimit
	}

	ret = &gslang.String{Value: s}

	return
}

func timesIsZero(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	if t1.IsZero() {
		ret = gslang.TrueValue
	} else {
		ret = gslang.FalseValue
	}

	return
}

func timesToLocal(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.Time{Value: t1.Local()}

	return
}

func timesToUTC(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.Time{Value: t1.UTC()}

	return
}

func timesTimeLocation(args ...gslang.Object) (
	ret gslang.Object,
	err error,
) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.String{Value: t1.Location().String()}

	return
}

func timesTimeString(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 1 {
		err = gslang.ErrWrongNumArguments
		return
	}

	t1, ok := gslang.ToTime(args[0])
	if !ok {
		err = gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gslang.String{Value: t1.String()}

	return
}
