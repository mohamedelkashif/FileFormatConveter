package main

import (
	"log"
	"os"
	"strconv"
	jsonConverter "trivago-BackendSoftwareEngineer-UserFlow/converters/jsonconverter"
	xmlConverter "trivago-BackendSoftwareEngineer-UserFlow/converters/xmlconverter"
	yamlConverter "trivago-BackendSoftwareEngineer-UserFlow/converters/yamlconverter"
	"trivago-BackendSoftwareEngineer-UserFlow/helpers"
	hotelType "trivago-BackendSoftwareEngineer-UserFlow/hotel"
)

func main() {
	file, err := os.Open("hotels.csv")
	arg := os.Args[1:]
	if len(arg) > 3 {
		log.Fatalf("Only 3 args is allowed: json, xml and yaml")
	} else if len(arg) < 1 {
		log.Fatalf("You need to specify at least one format (JSON, XML or YAML) to convert to!")

	}

	//fmt.Println(arg)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	rows, err := helpers.ReadCsvFile(file)

	_ = file.Close()

	var hotel hotelType.Hotel
	var hotels []hotelType.Hotel
	for _, each := range rows {
		if helpers.ValidateStars(strconv.Atoi(each[2])) && helpers.ValidateNameUtf8(each[0]) && helpers.ValidateURI(each[5]) {
			hotel.Name = each[0]
			hotel.Address = each[1]
			hotel.Stars, _ = strconv.Atoi(each[2])
			hotel.Contact = each[3]
			hotel.Phone = each[4]
			hotel.Uri = each[5]
			hotels = append(hotels, hotel)
		}
	}
	if helpers.Contains("xml", arg) || helpers.Contains("XML", arg) {
		xmlConverter.ToXml(hotels)
	}
	if helpers.Contains("json", arg) || helpers.Contains("JSON", arg) {
		jsonConverter.ToJson(hotels)
	}
	if helpers.Contains("yaml", arg) || helpers.Contains("YAML", arg) {
		yamlConverter.ToYaml(hotels)
	}
}
