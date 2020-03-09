package xmlconverter

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"trivago-BackendSoftwareEngineer-UserFlow/hotel"
)

func ToXml(hotels []hotel.Hotel) {
	xmlDate, err := xml.MarshalIndent(hotels, "", "\t")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	xmlFile, err := os.Create("./output/hotels.xml")
	defer xmlFile.Close()
	xmlFile.Write(xmlDate)
	log.Println("Converted to XML format")
	xmlFile.Close()
}
