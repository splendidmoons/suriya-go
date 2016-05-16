package suriya

import (
	"encoding/json"
	"fmt"
	"github.com/soh335/ical"
	"log"
	"path/filepath"
	s "strings"
	"time"
)

const (
	AstroMoonDir = "./data/astro"
)

var useLocal = false

type AstroMoon struct {
	Phase string
	Date  time.Time
}

func (m AstroMoon) AddToDay(day_p *CalDay) {
	day_p.Date = m.Date
	day_p.SetAstroMoon(m)
	return
}

func (m AstroMoon) GetDate() time.Time {
	return m.Date
}

func (m AstroMoon) String() string {
	if len(m.Phase) != 0 {
		return fmt.Sprintf("%s Moon", s.Title(m.Phase))
	}
	return ""
}

func (m AstroMoon) IcalEvent() ical.VEvent {
	return icalEvent(m)
}

func GetAstroMoons(fromDate time.Time, toDate time.Time) (moons []AstroMoon) {
	var err error

	for year := fromDate.Year(); year <= toDate.Year(); year++ {

		// filenames are astro-YYYY.json

		filename := fmt.Sprintf("astro-%d.json", year)
		filepath := "/" + filepath.Join(AstroMoonDir, filename)

		var data []byte

		if data, err = FSByte(useLocal, filepath); err != nil {
			if verbose {
				log.Printf("%v, %s\n", err, filepath)
			}
			return moons
		}

		var resp AerisResp
		if err := json.Unmarshal(data, &resp); err != nil {
			panic(err)
			return moons
		}

		if resp.Success != true {
			if verbose {
				log.Printf("%s was not successful", filepath)
			}
			continue
		}

		if len(resp.Error.Code) != 0 {
			if verbose {
				log.Println("%s has error: %s", filepath, resp.Error.Description)
			}
			continue
		}

		for _, aemoon := range resp.Response {
			var m AstroMoon
			m.Date = aemoon.DateTimeISO.UTC()
			if m.Date.Before(fromDate) || m.Date.After(toDate) {
				continue
			}
			m.Phase = phaseCodes[aemoon.Code]
			if m.Phase == "full" || m.Phase == "new" {
				moons = append(moons, m)
			}
		}
	}

	return moons
}
