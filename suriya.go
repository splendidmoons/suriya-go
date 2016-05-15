package suriya

// Eade, p.10. South Asian traditional number of days in 800 years

const (
	EraDays          = 292207
	EraYears         = 800
	EraHorakhun      = 373 // The Horakhun at the beginning of the CS Era, Ephemeris p.15, H2 element
	EraUccabala      = 2611
	EraAvoman        = 650
	EraMasaken       = 0
	MonthLength      = 30
	CycleTithi       = 703 // For every 692 solar days that elapse there are also 703 tithi = 692 + 11 / 692
	CycleSolar       = 692
	CycleDaily       = 11
	KammacubalaDaily = 800 // Daily increase
	CSdiff           = 638 // Absolute of CE - CS Era difference
	BEdiff           = 543 // Absolute of BE - CS Era difference

	// 1963 July 5, adhikavāra year, day 103, 1 day before Asalha Full Moon.
	// Eade uses this example in "Interpolation".
	horakhunRef    = 484049
	horakhunRefStr = "1963 Jul 5"
	//horakhunRefDate = time.Parse("2002 Jan 2", "1963 Jul 5")
)

// Whether to apply the (adhikavāra) exceptions where the official calendar
// differed from the formulas. Default is false, to generate calendar data that
// is "pure" in its consistency. Set to true if you want to match official past
// calendars which differed from the regular pattern.
var UseExceptions bool = false

var AdhikavaraExceptions = map[int]bool{
	1994: false,
	1997: true,
}

var monthToInt = map[string]int{
	"null":       0,
	"magasira":   1,
	"phussa":     2,
	"magha":      3,
	"phagguna":   4,
	"citta":      5,
	"vesakha":    6,
	"jettha":     7,
	"asalha":     8,
	"savana":     9,
	"bhaddapada": 10,
	"assayuja":   11,
	"kattika":    12,
	"2nd asalha": 13,
}

func MonthToInt(month string) int {
	return monthToInt[month]
}

var seasonToInt = map[string]int{
	"null":    0,
	"hemanta": 1,
	"gimhana": 2,
	"vassana": 3,
}

func SeasonToInt(season string) int {
	return seasonToInt[season]
}

var seasonName = map[int]string{
	0: "",
	1: "Hemanta",
	2: "Gimha",
	3: "Vassāna",
}

func SeasonName(number int) string {
	return seasonName[number]
}

var calendarToInt = map[string]int{
	"mahanikaya": 0,
	"dhammayut":  1,
	"srilanka":   2,
	"myanmar":    3,
}

func CalendarToInt(calendar string) int {
	return calendarToInt[calendar]
}

var statusToInt = map[string]int{
	"draft":     0,
	"predicted": 1,
	"confirmed": 2,
}

func StatusToInt(status string) int {
	return statusToInt[status]
}
