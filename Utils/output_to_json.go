package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// OutputToJSON writes the data to a JSON file
func OutputToJSON(data interface{}, filename string) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	log.Printf("Output saved to %s\n", filename)
	return nil
}
