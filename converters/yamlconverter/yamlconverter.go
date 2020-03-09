package yamlconverter

import (
	"fmt"
	"log"
	"os"
	hotelType "trivago-BackendSoftwareEngineer-UserFlow/hotel"
)

func ToYaml(hotels []hotelType.Hotel) {
	yamlData, err := Marshal(hotels)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	yamlFile, err := os.Create("./output/hotels.yaml")
	if err != nil {
		fmt.Println(err)
	}
	defer yamlFile.Close()
	yamlFile.Write(yamlData)
	log.Println("Converted to YAML format")
	yamlFile.Close()
}
