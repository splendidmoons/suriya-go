package suriya

import "testing"

func TestAdhikamasa(t *testing.T) {
	adhikamasaYears := map[int]bool{
		1998: false, //
		1999: true,  // 3
		2000: false, //
		2001: true,  // 2
		2002: false, //
		2003: false, //
		2004: true,  // 3
		2005: false, //
		2006: false, //
		2007: true,  // 3
		2008: false, //
		2009: true,  // 2
		2010: false, //
		2011: false, //
		2012: true,  // 3
		2013: false, //
		2014: false, //
		2015: true,  // 3
		2016: false, //
	}

	for year, expect := range adhikamasaYears {
		su := SuriyaYear{}
		su.Init(year)
		res := su.Is_Adhikamasa()
		if res != expect {
			t.Errorf("%d.Is_Adhikamasa() should be %v, but got %v", su.Year, expect, res)
		}
	}
}

func TestAdhikavara(t *testing.T) {
	// Adhikavāra in FS-Cal: 2005, 2010, 2016.
	adhikavaraYears := map[int]bool{
		1998: false, //
		1999: false, //
		2000: true,  // 6
		2001: false, //
		2002: false, //
		2003: false, //
		2004: false, //
		2005: true,  // 5
		2006: false, //
		2007: false, //
		2008: false, //
		2009: false, //
		2010: true,  // 5
		2011: false, //
		2012: false, //
		2013: false, //
		2014: false, //
		2015: false, //
		2016: true,  // 6
	}

	for year, expect := range adhikavaraYears {
		su := SuriyaYear{}
		su.Init(year)
		res := su.Is_Adhikavara()
		if res != expect {
			t.Errorf("%d.Is_Adhikavara() should be %v, but got %v", su.Year, expect, res)
		}
	}
}
