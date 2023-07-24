package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)



type APIResponse struct {
	AirQualityIndex float32 `json:"airQualityIndex"`
}

func checkForFile(fileName string) ([]byte, error) {
	_, err := os.Stat(fileName)
	checkerror(err)

	// check if the file exist and !exist create one
	if os.IsNotExist(err) {
		_, err := os.Create(fileName)
		checkerror(err)
	}
	file, err := ioutil.ReadFile("result.json")
	checkerror(err)

	return file, nil
}

func checkerror(err error) {
	if err != nil {
		println(err.Error())
	}
}

func appendToJSONFile() {
	filepath := checkForFile()
	existingData := make(map[string]interface{})

	var apiResponse APIResponse
	err := json.NewDecoder(res.Body).Decode(&apiResponse)
	checkerror(err)

	existingData["apiresponse"] = apiResponse

	//	Marshal the updated data back to json format
	updatedJSONData, err := json.Marshal(existingData)
	checkerror(err)

	// updated the JSON file
	err = ioutil.WriteFile(filepath, updatedJSONData, 0644)
}

func play() {
	//	search endpoints
	searchendpoint := os.Args[1]
	weight := os.Args[2]
	unit := os.Args[3]
	file := os.Args[4]

	//	rapid api link
	url := "https://carbonfootprint1.p.rapidapi.com/"

	//	queryString := "TreeEquivalent" + "?weight=200&unit=kg"
	queryString := searchendpoint + "?weight=" + weight + "&unit=" + unit

	//	Make a new request
	req, err := http.NewRequest("GET", url+queryString, nil)
	checkerror(err)

	//Adding key
	req.Header.Add("X-RapidAPI-Key", "4d2c003e3emsh82512a559ce6b89p16ac42jsnc0e7dc1d61ac")

	// sending the request to receive the respone
	res, err := http.DefaultClient.Do(req)
	checkerror(err)

	// closing the repsone body
	defer res.Body.Close()

	appendToJSONFile()

	//Accessing the response body
	body, err := ioutil.ReadAll(res.Body)
	checkerror(err)

	//	print the response
	println(string(body))

}
