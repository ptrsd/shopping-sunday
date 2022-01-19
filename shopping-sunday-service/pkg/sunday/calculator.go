package sunday

import (
	"time"
)

const ShoppingSundayFormat = "2006-01-02"

func IsShopping(date time.Time) (isShopping bool, reasons []Reason) {
	precheck := isSunday{}
	if result, reason := precheck.Calculate(date); !result {
		return result, []Reason{reason}
	}

	checks := []checkDate{isLastSundayOfMonth{}, beforeChristmas{}, beforeEaster{}}
	for _, check := range checks {
		isShopping, reason := check.Calculate(date)
		if isShopping {
			return isShopping, nil
		}
		reasons = append(reasons, reason)
	}
	return false, reasons
}
