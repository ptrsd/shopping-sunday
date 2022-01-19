package sunday

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var tests = []struct {
	name            string
	input           time.Time
	expectedResult  bool
	expectedReasons []Reason
}{
	{
		name:            "Success: Sunday before Christmas 2021",
		input:           time.Date(2021, time.December, 19, 0, 0, 0, 0, time.UTC),
		expectedResult:  true,
		expectedReasons: []Reason{},
	}, {
		name:            "Fail: Christmas is not a shopping day",
		input:           time.Date(2021, time.December, 25, 0, 0, 0, 0, time.UTC),
		expectedResult:  false,
		expectedReasons: []Reason{isNotSunday},
	}, {
		name:            "Success: The last day of January",
		input:           time.Date(2022, time.January, 30, 0, 0, 0, 0, time.UTC),
		expectedResult:  true,
		expectedReasons: []Reason{},
	}, {
		name:            "Fail: Is not the last day of January",
		input:           time.Date(2022, time.January, 23, 0, 0, 0, 0, time.UTC),
		expectedResult:  false,
		expectedReasons: []Reason{notLastSundayOfMonth, notBeforeChristmas, notBeforeEaster},
	}, {
		name:            "Fail: No shopping Sundays in February",
		input:           time.Date(2022, time.February, 6, 0, 0, 0, 0, time.UTC),
		expectedResult:  false,
		expectedReasons: []Reason{monthNotValid, notBeforeChristmas, notBeforeEaster},
	}, {
		name:            "Success: Sunday before Easter",
		input:           time.Date(2022, time.April, 10, 0, 0, 0, 0, time.UTC),
		expectedResult:  true,
		expectedReasons: []Reason{},
	}, {
		name:            "Fail: is not the Sunday before Easter",
		input:           time.Date(2022, time.April, 17, 0, 0, 0, 0, time.UTC),
		expectedResult:  false,
		expectedReasons: []Reason{notLastSundayOfMonth, notBeforeChristmas, notBeforeEaster},
	}, {
		name:            "Success: The last Sunday in April",
		input:           time.Date(2022, time.April, 24, 0, 0, 0, 0, time.UTC),
		expectedResult:  true,
		expectedReasons: []Reason{},
	}, {
		name:            "Success: The last Sunday in June",
		input:           time.Date(2022, time.June, 26, 0, 0, 0, 0, time.UTC),
		expectedResult:  true,
		expectedReasons: []Reason{},
	}, {
		name:            "Success: The last Sunday in August",
		input:           time.Date(2022, time.August, 28, 0, 0, 0, 0, time.UTC),
		expectedResult:  true,
		expectedReasons: []Reason{},
	}, {
		name:            "Success: Two Sundays before Christmas 2022",
		input:           time.Date(2022, time.December, 11, 0, 0, 0, 0, time.UTC),
		expectedResult:  true,
		expectedReasons: []Reason{},
	}, {
		name:            "Success: Sunday before Christmas 2022",
		input:           time.Date(2022, time.December, 18, 0, 0, 0, 0, time.UTC),
		expectedResult:  true,
		expectedReasons: []Reason{},
	},
}

func TestName(t *testing.T) {
	for _, tc := range tests {
		actualResult, actualReasons := IsShopping(tc.input)
		assert.Equal(t, tc.expectedResult, actualResult, tc.name)
		assert.EqualValues(t, tc.expectedReasons, actualReasons, tc.name)
	}
}
