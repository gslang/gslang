package stdlib

import (
	t "time"
	"github.com/gslang/gslang"
)

var timeModule = map[string]gslang.Object{
	"format_ansic":        &gslang.String{Value: t.ANSIC},
	"format_unix_date":    &gslang.String{Value: t.UnixDate},
	"format_ruby_date":    &gslang.String{Value: t.RubyDate},
	"format_rfc822":       &gslang.String{Value: t.RFC822},
	"format_rfc822z":      &gslang.String{Value: t.RFC822Z},
	"format_rfc850":       &gslang.String{Value: t.RFC850},
	"format_rfc1123":      &gslang.String{Value: t.RFC1123},
	"format_rfc1123z":     &gslang.String{Value: t.RFC1123Z},
	"format_rfc3339":      &gslang.String{Value: t.RFC3339},
	"format_rfc3339_nano": &gslang.String{Value: t.RFC3339Nano},
	"format_kitchen":      &gslang.String{Value: t.Kitchen},
	"format_stamp":        &gslang.String{Value: t.Stamp},
	"format_stamp_milli":  &gslang.String{Value: t.StampMilli},
	"format_stamp_micro":  &gslang.String{Value: t.StampMicro},
	"format_stamp_nano":   &gslang.String{Value: t.StampNano},
	"nanosecond":          &gslang.Int{Value: int64(t.Nanosecond)},
	"microsecond":         &gslang.Int{Value: int64(t.Microsecond)},
	"millisecond":         &gslang.Int{Value: int64(t.Millisecond)},
	"second":              &gslang.Int{Value: int64(t.Second)},
	"minute":              &gslang.Int{Value: int64(t.Minute)},
	"hour":                &gslang.Int{Value: int64(t.Hour)},
	"january":             &gslang.Int{Value: int64(t.January)},
	"february":            &gslang.Int{Value: int64(t.February)},
	"march":               &gslang.Int{Value: int64(t.March)},
	"april":               &gslang.Int{Value: int64(t.April)},
	"may":                 &gslang.Int{Value: int64(t.May)},
	"june":                &gslang.Int{Value: int64(t.June)},
	"july":                &gslang.Int{Value: int64(t.July)},
	"august":              &gslang.Int{Value: int64(t.August)},
	"september":           &gslang.Int{Value: int64(t.September)},
	"october":             &gslang.Int{Value: int64(t.October)},
	"november":            &gslang.Int{Value: int64(t.November)},
	"december":            &gslang.Int{Value: int64(t.December)},
	"sleep": &gslang.UserFunction{
		Name:  "sleep",
		Value: timeSleep,
	}, // sleep(int)
	"parse_duration": &gslang.UserFunction{
		Name:  "parse_duration",
		Value: timeParseDuration,
	}, // parse_duration(str) => int
	"since": &gslang.UserFunction{
		Name:  "since",
		Value: timeSince,
	}, // since(time) => int
	"until": &gslang.UserFunction{
		Name:  "until",
		Value: timeUntil,
	}, // until(time) => int
	"duration_hours": &gslang.UserFunction{
		Name:  "duration_hours",
		Value: timeDurationHours,
	}, // duration_hours(int) => float
	"duration_minutes": &gslang.UserFunction{
		Name:  "duration_minutes",
		Value: timeDurationMinutes,
	}, // duration_minutes(int) => float
	"duration_nanoseconds": &gslang.UserFunction{
		Name:  "duration_nanoseconds",
		Value: timeDurationNanoseconds,
	}, // duration_nanoseconds(int) => int
	"duration_seconds": &gslang.UserFunction{
		Name:  "duration_seconds",
		Value: timeDurationSeconds,
	}, // duration_seconds(int) => float
	"duration_string": &gslang.UserFunction{
		Name:  "duration_string",
		Value: timeDurationString,
	}, // duration_string(int) => string
	"month_string": &gslang.UserFunction{
		Name:  "month_string",
		Value: timeMonthString,
	}, // month_string(int) => string
	"date": &gslang.UserFunction{
		Name:  "date",
		Value: timeDate,
	}, // date(year, month, day, hour, min, sec, nsec) => time
	"now": &gslang.UserFunction{
		Name:  "now",
		Value: timeNow,
	}, // now() => time
	"parse": &gslang.UserFunction{
		Name:  "parse",
		Value: timeParse,
	}, // parse(format, str) => time
	"unix": &gslang.UserFunction{
		Name:  "unix",
		Value: timeUnix,
	}, // unix(sec, nsec) => time
	"add": &gslang.UserFunction{
		Name:  "add",
		Value: timeAdd,
	}, // add(time, int) => time
	"add_date": &gslang.UserFunction{
		Name:  "add_date",
		Value: timeAddDate,
	}, // add_date(time, years, months, days) => time
	"sub": &gslang.UserFunction{
		Name:  "sub",
		Value: timeSub,
	}, // sub(t time, u time) => int
	"after": &gslang.UserFunction{
		Name:  "after",
		Value: timeAfter,
	}, // after(t time, u time) => bool
	"before": &gslang.UserFunction{
		Name:  "before",
		Value: timeBefore,
	}, // before(t time, u time) => bool
	"time_year": &gslang.UserFunction{
		Name:  "time_year",
		Value: timeTimeYear,
	}, // time_year(time) => int
	"time_month": &gslang.UserFunction{
		Name:  "time_month",
		Value: timeTimeMonth,
	}, // time_month(time) => int
	"time_day": &gslang.UserFunction{
		Name:  "time_day",
		Value: timeTimeDay,
	}, // time_day(time) => int
	"time_weekday": &gslang.UserFunction{
		Name:  "time_weekday",
		Value: timeTimeWeekday,
	}, // time_weekday(time) => int
	"time_hour": &gslang.UserFunction{
		Name:  "time_hour",
		Value: timeTimeHour,
	}, // time_hour(time) => int
	"time_minute": &gslang.UserFunction{
		Name:  "time_minute",
		Value: timeTimeMinute,
	}, // time_minute(time) => int
	"time_second": &gslang.UserFunction{
		Name:  "time_second",
		Value: timeTimeSecond,
	}, // time_second(time) => int
	"time_nanosecond": &gslang.UserFunction{
		Name:  "time_nanosecond",
		Value: timeTimeNanosecond,
	}, // time_nanosecond(time) => int
	"time_unix": &gslang.UserFunction{
		Name:  "time_unix",
		Value: timeTimeUnix,
	}, // time_unix(time) => int
	"time_unix_nano": &gslang.UserFunction{
		Name:  "time_unix_nano",
		Value: timeTimeUnixNano,
	}, // time_unix_nano(time) => int
	"time_format": &gslang.UserFunction{
		Name:  "time_format",
		Value: timeTimeFormat,
	}, // time_format(time, format) => string
	"time_location": &gslang.UserFunction{
		Name:  "time_location",
		Value: timeTimeLocation,
	}, // time_location(time) => string
	"time_string": &gslang.UserFunction{
		Name:  "time_string",
		Value: timeTimeString,
	}, // time_string(time) => string
	"is_zero": &gslang.UserFunction{
		Name:  "is_zero",
		Value: timeIsZero,
	}, // is_zero(time) => bool
	"to_local": &gslang.UserFunction{
		Name:  "to_local",
		Value: timeToLocal,
	}, // to_local(time) => time
	"to_utc": &gslang.UserFunction{
		Name:  "to_utc",
		Value: timeToUTC,
	}, // to_utc(time) => time
}

func timeSleep(args ...gslang.Object) (ret gslang.Object, err error) {
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

	t.Sleep(t.Duration(i1))
	ret = gslang.NilValue

	return
}

func timeParseDuration(args ...gslang.Object) (
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

	dur, err := t.ParseDuration(s1)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &gslang.Int{Value: int64(dur)}

	return
}

func timeSince(args ...gslang.Object) (
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

	ret = &gslang.Int{Value: int64(t.Since(t1))}

	return
}

func timeUntil(args ...gslang.Object) (
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

	ret = &gslang.Int{Value: int64(t.Until(t1))}

	return
}

func timeDurationHours(args ...gslang.Object) (
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

	ret = &gslang.Float{Value: t.Duration(i1).Hours()}

	return
}

func timeDurationMinutes(args ...gslang.Object) (
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

	ret = &gslang.Float{Value: t.Duration(i1).Minutes()}

	return
}

func timeDurationNanoseconds(args ...gslang.Object) (
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

	ret = &gslang.Int{Value: t.Duration(i1).Nanoseconds()}

	return
}

func timeDurationSeconds(args ...gslang.Object) (
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

	ret = &gslang.Float{Value: t.Duration(i1).Seconds()}

	return
}

func timeDurationString(args ...gslang.Object) (
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

	ret = &gslang.String{Value: t.Duration(i1).String()}

	return
}

func timeMonthString(args ...gslang.Object) (
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

	ret = &gslang.String{Value: t.Month(i1).String()}

	return
}

func timeDate(args ...gslang.Object) (
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
		Value: t.Date(i1,
			t.Month(i2), i3, i4, i5, i6, i7, t.Now().Location()),
	}

	return
}

func timeNow(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 0 {
		err = gslang.ErrWrongNumArguments
		return
	}

	ret = &gslang.Time{Value: t.Now()}

	return
}

func timeParse(args ...gslang.Object) (ret gslang.Object, err error) {
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

	parsed, err := t.Parse(s1, s2)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &gslang.Time{Value: parsed}

	return
}

func timeUnix(args ...gslang.Object) (ret gslang.Object, err error) {
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

	ret = &gslang.Time{Value: t.Unix(i1, i2)}

	return
}

func timeAdd(args ...gslang.Object) (ret gslang.Object, err error) {
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

	ret = &gslang.Time{Value: t1.Add(t.Duration(i2))}

	return
}

func timeSub(args ...gslang.Object) (ret gslang.Object, err error) {
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

func timeAddDate(args ...gslang.Object) (ret gslang.Object, err error) {
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

func timeAfter(args ...gslang.Object) (ret gslang.Object, err error) {
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

func timeBefore(args ...gslang.Object) (ret gslang.Object, err error) {
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

func timeTimeYear(args ...gslang.Object) (ret gslang.Object, err error) {
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

func timeTimeMonth(args ...gslang.Object) (ret gslang.Object, err error) {
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

func timeTimeDay(args ...gslang.Object) (ret gslang.Object, err error) {
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

func timeTimeWeekday(args ...gslang.Object) (ret gslang.Object, err error) {
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

func timeTimeHour(args ...gslang.Object) (ret gslang.Object, err error) {
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

func timeTimeMinute(args ...gslang.Object) (ret gslang.Object, err error) {
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

func timeTimeSecond(args ...gslang.Object) (ret gslang.Object, err error) {
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

func timeTimeNanosecond(args ...gslang.Object) (
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

func timeTimeUnix(args ...gslang.Object) (ret gslang.Object, err error) {
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

func timeTimeUnixNano(args ...gslang.Object) (
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

func timeTimeFormat(args ...gslang.Object) (ret gslang.Object, err error) {
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

func timeIsZero(args ...gslang.Object) (ret gslang.Object, err error) {
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

func timeToLocal(args ...gslang.Object) (ret gslang.Object, err error) {
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

func timeToUTC(args ...gslang.Object) (ret gslang.Object, err error) {
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

func timeTimeLocation(args ...gslang.Object) (
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

func timeTimeString(args ...gslang.Object) (ret gslang.Object, err error) {
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
