
package main

import (
	"fmt"
	"log"
	//"os"
	"lib"
	"config"
	"toml"
	"io"
	"io/ioutil"
	"analyze"
	
) 

// file prefix for --all option: "-all" for all parameters, "" empty by default
var prefixAll = ""

// use for debug mode
var debug io.Writer = ioutil.Discard
// use for echo mode
var echo io.Writer = ioutil.Discard

// usefull macro
var p = fmt.Println
var f = fmt.Printf


var cfg toml.Configtoml
var filestruct analyze.Structfile

type AllData_2D lib.AllData_2D

var nc lib.Nc
var m config.Map

var fileconfig string

// main body
func main() {
	
	//init variable cfg with config file in TOML
	fileconfig,cfg = toml.InitToml()
	//init variable m with empty map 
	m = config.InitMap()
	
	var files []string
	// to change the flags on the default logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	
	files, optCfgfile := GetOptions()
	
	//analyse the file to know contructor, instrument and instrument type
	filestruct = analyze.AnalyzeFile(cfg,files)
	
	switch{
		
		case filestruct.Constructeur.Name == "Seabird" :
			ReadSeabird(&nc,&m,files,optCfgfile)
		}
	
}
