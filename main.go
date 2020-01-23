package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	filename := "/Users/sgodasi/Downloads/6403_100.json"
	plan, err := ioutil.ReadFile(filename) // filename is the JSON file to read
	var data map[string]interface{}
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(plan, &data)
	if err != nil {
		log.Println("Cannot unmarshal the json ", err)
	}
	list := make([]map[string]interface{}, 0)
	for _, responseList := range data["response"].([]interface{}) {
		var result map[string]interface{}
		var raw map[string]interface{}
		var body map[string]interface{}
		result = (responseList.(map[string]interface{}))["result"].(map[string]interface{})
		err = json.Unmarshal([]byte(result["_raw"].(string)), &raw)
		err = json.Unmarshal([]byte(raw["body"].(string)), &body)

		list = append(list, body)
	}
	writeJosnOutPutFile(map[string]interface{}{"payloads": list})
}

func writeJosnOutPutFile(data map[string]interface{}) {
	outputFileName := "/Users/sgodasi/Downloads/FormattedJson1.json"
	jsonData, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}
	jsonFile, err := os.Create(outputFileName)

	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
	fmt.Println("JSON data written to ", jsonFile.Name())
}
