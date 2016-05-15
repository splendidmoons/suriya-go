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

// TODO: use env var verbose
var verbose bool = false
