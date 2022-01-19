package sunday

type Reason struct {
	Id      int
	Message string
}

var (
	isNotSunday = Reason{
		Id:      1,
		Message: "is not a Sunday",
	}
	monthNotValid = Reason{
		Id:      2,
		Message: "month is not equal to January, April, June, August",
	}
	notLastSundayOfMonth = Reason{
		Id:      3,
		Message: "the day is not the last Sunday of month",
	}
	notBeforeChristmas = Reason{
		Id:      4,
		Message: "is not within two Sundays before Christmas",
	}
	notBeforeEaster = Reason{
		Id:      5,
		Message: "is not a Sunday before Easter",
	}
)
