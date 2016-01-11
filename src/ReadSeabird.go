//ReadSeabird.go
//Function for read data file for constructor Seabird

package main

// read cnv files in two pass, the first to get dimensions
// second to get data
func (nc *Nc) ReadSeabird(files []string,optCfgfile string) {
	
	switch{
		case filestruct.Instrument == cfg.Instrument.Type[0] :

			nc.GetConfigCTD(optCfgfile,filestruct.TypeInstrument)
			
			// first pass, return dimensions fron cnv files
			nc.Dimensions["TIME"], nc.Dimensions["DEPTH"] = nc.firstPassCTD(files)
		
			// initialize 2D data
			nc.Variables_2D = make(AllData_2D)
			for i, _ := range map_var {
				nc.Variables_2D.NewData_2D(i, nc.Dimensions["TIME"], nc.Dimensions["DEPTH"])
			}
		
			// second pass, read files again, extract data and fill slices
			nc.secondPassCTD(files)
			// write ASCII file
			nc.WriteAsciiCTD(map_format, hdr,filestruct.Instrument)
		
			// write netcdf file
			//if err := nc.WriteNetcdf(); err != nil {
			//log.Fatal(err)
			//}
			nc.WriteNetcdf(filestruct.Instrument)
			
		case filestruct.Instrument == cfg.Instrument.Type[1] :
		
			nc.GetConfigBTL(optCfgfile,filestruct.TypeInstrument)
			// first pass, return dimensions fron btl files
			nc.Dimensions["TIME"], nc.Dimensions["DEPTH"] = nc.firstPassBTL(files)
		
			//	// initialize 2D data
			//	nc.Variables_2D = make(AllData_2D)
			//	for i, _ := range map_var {
			//		nc.Variables_2D.NewData_2D(i, nc.Dimensions["TIME"], nc.Dimensions["DEPTH"])
			//	}
		
			// second pass, read files again, extract data and fill slices
			nc.secondPassBTL(files)
			// write ASCII file
			nc.WriteAsciiBTL2(map_format, hdr,filestruct.Instrument)
		
			// write netcdf file
			//if err := nc.WriteNetcdf(); err != nil {
			//log.Fatal(err)
			//}
			nc.WriteNetcdf(filestruct.Instrument)
			}
}