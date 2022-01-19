package sunday

import (
	"time"
)

const weekInDays = 7

type checkDate interface {
	Calculate(time time.Time) (bool, Reason)
}

type isSunday struct {
}

func (i isSunday) Calculate(date time.Time) (bool, Reason) {
	result := time.Sunday == date.Weekday()
	if !result {
		return result, isNotSunday
	}
	return result, Reason{}
}

type isLastSundayOfMonth struct{}

func (i isLastSundayOfMonth) Calculate(date time.Time) (bool, Reason) {
	if !i.monthWithin(date, []time.Month{time.January, time.April, time.June, time.August}) {
		return false, monthNotValid
	}

	addDate := date.AddDate(0, 0, 7)
	if !(time.Sunday == date.Weekday() && addDate.Month() != date.Month()) {
		return false, notLastSundayOfMonth
	}
	return true, Reason{}
}

func (i isLastSundayOfMonth) monthWithin(date time.Time, months []time.Month) bool {
	for _, month := range months {
		if date.Month() == month {
			return true
		}
	}
	return false
}

type beforeChristmas struct{}

func (b beforeChristmas) Calculate(date time.Time) (bool, Reason) {
	christmas := time.Date(date.Year(), time.December, 24, 0, 0, 0, 0, date.Location())
	daysBeforeChristmas := -1 * int(christmas.Weekday())
	sundayBefore := christmas.AddDate(0, 0, daysBeforeChristmas)
	twoSundaysBefore := christmas.AddDate(0, 0, daysBeforeChristmas-weekInDays)
	if !(date.Equal(sundayBefore) || date.Equal(twoSundaysBefore)) {
		return false, notBeforeChristmas
	}
	return true, Reason{}
}

type beforeEaster struct{}

func (b beforeEaster) Calculate(date time.Time) (bool, Reason) {
	easter := b.easterForYear(date)
	sundayBefore := easter.AddDate(0, 0, -int(easter.Weekday())-weekInDays)
	if !date.Equal(sundayBefore) {
		return false, notBeforeEaster
	}
	return true, Reason{}
}

func (b beforeEaster) easterForYear(date time.Time) time.Time {
	year := date.Year()

	month := 3
	golden := (year % 19) + 1
	century := year/100 + 1
	xx := (3*century)/4 - 12
	yy := (8*century+5)/25 - 5
	zz := (5*year)/4 - xx - 10
	ee := (11*golden + 20 + yy - xx) % 30
	if ee == 24 {
		ee += 1
	}
	if (ee == 25) && (golden > 11) {
		ee += 1
	}
	moon := 44 - ee
	if moon < 21 {
		moon += 30
	}
	day := (moon + 7) - ((zz + moon) % 7)
	if day > 31 {
		day -= 31
		month = 4
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, date.Location())
}
