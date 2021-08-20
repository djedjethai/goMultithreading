package mainaa

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

var (
	// windRegex     = regexp.MustCompile(`\d* METAR.*EGLL \d*Z [A-Z ]*(\d{5}KT|VRB\d{2}KT).*=`)
	tafValidation = regexp.MustCompile(`.*TAF.*`)
	comment       = regexp.MustCompile(`\w*#.*`)
	metarClose    = regexp.MustCompile(`.*=`)
	variableWind  = regexp.MustCompile(`.*VRB\d{2}KT`)
	validWind     = regexp.MustCompile(`\d{5}KT`)
	windDirOnly   = regexp.MustCompile(`(\d{3})\d{2}KT`)
	windDist      [8]int

	metar = "201208010620 METAR COR EGLL 010620Z 12005KT 8000 NSC 15/14 Q1010 NOSIG="
)

func main() {

	var windRegex = regexp.MustCompile(`\d* METAR.*EGLL \d*Z [A-Z ]*(\d{5}KT|VRB\d{2}KT).*(N.*G)=`)
	var wind = windRegex.FindAllStringSubmatch(metar, -1)[0][1]
	var windtest = windRegex.FindAllStringSubmatch(metar, -1)[0][2]
	fmt.Println("tesst: " + windtest)
	fmt.Println(wind)
	if variableWind.MatchString(wind) {
		windDist[0]++
		fmt.Printf("the prrint: %v\n", windDist)
	} else if validWind.MatchString(wind) {
		windStr := windDirOnly.FindAllStringSubmatch(wind, -1)[0][1]
		fmt.Printf("the prrint1: %v\n", windStr)
		if d, err := strconv.ParseFloat(windStr, 64); err == nil {
			dirIndex := int(math.Round(d/45.0)) % 8
			fmt.Printf("the prrint2: %v\n", dirIndex)
			windDist[dirIndex]++
		}
	}

}
