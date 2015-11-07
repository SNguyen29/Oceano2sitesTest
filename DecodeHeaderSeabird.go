// DecodeHeaderSeabird.go
package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// define regexp

var regCruise = regexp.MustCompile(`Cruise\s*:\s*(.*)`)
var regShip = regexp.MustCompile(`Ship\s*:\s*(.*)`)
var regStation = regexp.MustCompile(`Station\s*:\s*\D*(\d*)`)
var regType = regexp.MustCompile(`Type\s*:\s*(.*)`)
var regOperator = regexp.MustCompile(`Operator\s*:\s*(.*)`)
var regBottomDepth = regexp.MustCompile(`Bottom Depth\s*:\s*(\d*\.?\d+?)\s*\S*`)
var regDummyBottomDepth = regexp.MustCompile(`Bottom Depth\s*:\s*$`)
var regDate = regexp.MustCompile(`Date\s*:\s*(\d+)/(\d+)/(\d+)`)
var regHour = regexp.MustCompile(`[Heure|Hour]\s*:\s*(\d+)[:hH](\d+):(\d+)`)
var regLatitude = regexp.MustCompile(`Latitude\s*:\s*(\d+)\s+(\d+.\d+)\s+(\w)`)
var regLongitude = regexp.MustCompile(`Longitude\s*:\s*(\d+)\s+(\d+.\d+)\s+(\w)`)
var regSystemTime = regexp.MustCompile(`System UpLoad Time =\s+(.*)`)
var regNmeaLatitude = regexp.MustCompile(`NMEA Latitude\s*=\s*(\d+\s+\d+.\d+\s+\w)`)
var regNmeaLongitude = regexp.MustCompile(`NMEA Longitude\s*=\s*(\d+\s+\d+.\d+\s+\w)`)

// parse header line from .cnv and extract correct information
// use regular expression
// to parse time with non standard format, see:
// http://golang.org/src/time/format.go

func (nc *Nc) DecodeHeaderSeabird(str string, profileTest float64) {
	switch {
	// decode Systeme Upload Time
		case regSystemTime.MatchString(str) : 
			res := regSystemTime.FindStringSubmatch(str)
			value := res[1]
			fmt.Fprintf(debug, "%s -> ", value)
			// create new Time object, see tools.go
			var t = NewTimeFromString("Jan 02 2006 15:04:05", value)
			v := t.Time2JulianDec()
			nc.Variables_1D["TIME"] = append(nc.Variables_1D["TIME"].([]float64),v)
	
		case regNmeaLatitude.MatchString(str) :
			if v, err := Position2Decimal(str); err == nil {
				nc.Variables_1D["LATITUDE"] = append(nc.Variables_1D["LATITUDE"].([]float64), v)
			} else {
				nc.Variables_1D["LATITUDE"] = append(nc.Variables_1D["LATITUDE"].([]float64), 1e36)
			}
			
			
		case regNmeaLongitude.MatchString(str) :
			if v, err := Position2Decimal(str); err == nil {
				nc.Variables_1D["LONGITUDE"] = append(nc.Variables_1D["LONGITUDE"].([]float64), v)
				fmt.Println(v)
			} else {
				nc.Variables_1D["LONGITUDE"] = append(nc.Variables_1D["LONGITUDE"].([]float64), 1e36)
			}
			
			
		case regCruise.MatchString(str) :
			res := regCruise.FindStringSubmatch(str)
			value := res[1]
			fmt.Println(value)
			fmt.Fprintln(debug, value)
			nc.Attributes["cycle_mesure"] = value

		case regShip.MatchString(str) :
			res := regShip.FindStringSubmatch(str)
			value := res[1]
			fmt.Fprintln(debug, value)
			nc.Attributes["plateforme"] = value
			fmt.Println(value)
			
		case regStation.MatchString(str) :
			res := regStation.FindStringSubmatch(str)
			value := res[1]
			if v, err := strconv.ParseFloat(value, 64); err == nil {
				fmt.Fprintln(debug, v)
				// ch
				//			format := "%0" + cfg.Ctd.StationPrefixLength + ".0f"
				//			p := fmt.Sprintf(format, profile)
				//			//s := fmt.Sprintf(format, v)
				//			fmt.Println(p, v)
				//			if p != v {
				//				fmt.Printf("Warning: profile for header differ from file name: %s <=> %s\n", p, v)
				//			}
				nc.Variables_1D["PROFILE"] = append(nc.Variables_1D["PROFILE"].([]float64), profileTest)
			} else {
				nc.Variables_1D["PROFILE"] = append(nc.Variables_1D["PROFILE"].([]float64), 1e36)
			}
			fmt.Println(value)

		case regBottomDepth.MatchString(str) :
			res := regBottomDepth.FindStringSubmatch(str)
			value := res[1]
			if v, err := strconv.ParseFloat(value, 64); err == nil {
				fmt.Fprintf(debug, "Bath: %f\n", v)
				nc.Variables_1D["BATH"] = append(nc.Variables_1D["BATH"].([]float64), v)
			} else {
				fmt.Fprintf(debug, "Bath: %f\n", v)
				nc.Variables_1D["BATH"] = append(nc.Variables_1D["BATH"].([]float64), 1e36)
			}
			fmt.Println(value)
			
		case regDummyBottomDepth.MatchString(str) ://?
			nc.Variables_1D["BATH"] = append(nc.Variables_1D["BATH"].([]float64), 1e36)
			fmt.Fprintf(debug, "Bath: %g\n", 1e36)
			fmt.Println("1e36")

		case regOperator.MatchString(str) :
			res := regOperator.FindStringSubmatch(str)
			value := res[1]
			if *optDebug {
				fmt.Println(value)
			}
			fmt.Println(value)
			
	// TODOS: uncomment, add optionnal value from seabird header
		case regType.MatchString(str) :
			res := regType.FindStringSubmatch(str)
			value := strings.ToUpper(res[1]) // convert to upper case
			var v float64
			switch value {
			case "PHY":
				v = float64(PHY)
			case "GEO":
				v = float64(GEO)
			case "BIO":
				v = float64(BIO)
			default:
				v = float64(UNKNOW)
			}
			//f("Type: %f\n", v)
			nc.Variables_1D["TYPECAST"] = append(nc.Variables_1D["TYPECAST"].([]float64), v)
			if *optDebug {
				fmt.Println(value)
			}
			nc.Extras_s[fmt.Sprintf("TYPE:%d", int(profileTest))] = value
		
	}
}
