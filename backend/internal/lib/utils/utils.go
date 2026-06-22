package utils

import (
	"encoding/json"
	"fmt"
)

// Takes an interface - could me map , slice 
// Marshals it in json format
// Prints it in string
// allows us to view content of map,slice

func PrintJSON(v interface{}) {
	json, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println("JSON:", string(json))
}