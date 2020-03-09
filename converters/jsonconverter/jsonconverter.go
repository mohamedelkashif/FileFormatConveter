package jsonconverter

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"trivago-BackendSoftwareEngineer-UserFlow/hotel"
)

func ToJson(hotels []hotel.Hotel) {
	jsonData, err := json.MarshalIndent(hotels, "", "\t")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	jsonFile, err := os.Create("./output/hotels.json")
	defer jsonFile.Close()
	jsonFile.Write(jsonData)
	log.Println("Converted to JSON format")
	jsonFile.Close()
}
