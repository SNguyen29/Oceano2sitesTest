// GetProfileNumber.go
//Function for get the profil number of a data file for Seabird Constructor
package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var CruisePrefix string = "csp"

func (nc *Nc) GetProfileNumber(str string) float64 {
	var value float64
	var err error
	if strings.Contains(str,CruisePrefix) {
		res := strings.Split(str,CruisePrefix)
		res = strings.Split(res[1],".")
		if value, err = strconv.ParseFloat(res[0], 64); err == nil {
			// get profile name, eg: csp00101
			nc.Extras_s[fmt.Sprintf("PRFL_NAME:%d", int(value))] = res[0]
		} else {
			log.Fatal(err)
		}

	} else {
		log.Fatal("func GetProfileNumber", err)
	}
	return value

}