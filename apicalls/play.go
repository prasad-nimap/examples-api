package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

/*type Person struct {
	Fname string
	Lname string
	Age   int
}*/

type Personplay struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Age   int    `json:"age"`
}

type APIResponseplay struct {
	AirQualityIndex float32 `json:"airQualityIndex"`
}

func checkForFileplay(filename string) (string, error) {
	_, err := os.Stat(filename)
	checkerrorplay(err)
	//	println(status)

	// check if the file exist and !exist create one
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		checkerrorplay(err)
	}
	file, err := ioutil.ReadFile("output.json")
	checkerrorplay(err)
	println(file)

	return filename, nil
}

func checkerrorplay(err error) {
	if err != nil {
		println(err.Error())
	}
}

/*
func appendToJSONFile(filename string) error {
	var jsondata []Person

	// Check for file and if !exist create
	file, err := checkForFile(filename)

	// Unmarshal the existing JSON data
	err = json.Unmarshal(file, &jsondata)
	checkerror(err)

	data := []Person{}

	newStruct := Person{
		Fname: "Prasad",
		Lname: "Junghare",
		Age:   21,
	}
	//	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 key
	// append the new data
	// jsondata = append(jsondata, data...)
	data = append(data, newStruct)

	//	Marshal the updated JSON data
	//	data, err := append(jsondata)
	//	checkerror(err)

	/*
	   //	Preparing the data to be marshalled and written
	   dataBytes, err := json.Marshal(data)
	   checkerror(err)

	   err = ioutil.WriteFile("result.json", dataBytes, 0644)
	   checkerror(err)
//
	databyte, err := json.Marshal(data)

	// write the updated data to the file
	err = ioutil.WriteFile(filename, databyte, 0644)
	checkerror(err)

	return nil
}
*/

func appendToJSONFileplay(filepath string, apiResponse APIResponseplay) {
	// filepath := checkForFileplay()
	// Instead of using map use struct
	existingData := APIResponseplay{}

	// Read and unmarshal the existing data into the apiresponse struct
	file, err := ioutil.ReadFile(filepath)
	checkerrorplay(err)

	err = json.Unmarshal(file, &existingData)

	// Updated the AQI filed
	existingData.AirQualityIndex = apiResponse.AirQualityIndex

	// Marshal back to json format
	updatedJSONData, err := json.Marshal(existingData)
	checkerrorplay(err)

	// Write the updated to file
	err = ioutil.WriteFile(filepath, updatedJSONData, 0644)

	/* // var apiResponse APIResponse
	err := json.NewDecoder(res.Body).Decode(&apiResponse)
	checkerrorplay(err)

	existingData["apiresponse"] = apiResponse

	//	Marshal the updated data back to json format
	updatedJSONData, err := json.Marshal(existingData)
	checkerrorplay(err)

	// updated the JSON file
	err = ioutil.WriteFile(filepath, updatedJSONData, 0644) */
}

func main() {
	//	println("api call")
	//	filename := os.Args[1]

	//	append using the apppend function
	//	appendToJSONFile(filename)

	// some file stuff going on
	/*
		//	 Check if the file exist
		// if !exist create new file
		err := checkForFile("result.json")
		checkerror(err)

		file, err := ioutil.ReadFile("result.json")
		checkerror(err)

		data := []Person{}

		json.Unmarshal(file, &data)

		newStruct := &Person{
		Fname: "Prasad",
		Lname: "Junghare",
		Age:   21,
		}

		data = append(data, *newStruct)

		//	Preparing the data to be marshalled and written
		dataBytes, err := json.Marshal(data)
		checkerror(err)

		err = ioutil.WriteFile("result.json", dataBytes, 0644)
		checkerror(err)

	*/

	//	search endpoints
	searchendpoint := os.Args[1]
	// weight := os.Args[2]
	// unit := os.Args[3]
	file := os.Args[2]

	//	rapid api link
	url := "https://carbonfootprint1.p.rapidapi.com/"

	//	queryString := "TreeEquivalent" + "?weight=200&unit=kg"
	queryString := searchendpoint + "?weight=200&unit=kg"
	// queryString := searchendpoint + "?weight=" + weight + "&unit=" + unit

	//	Make a new request
	req, err := http.NewRequest("GET", url+queryString, nil)
	checkerrorplay(err)

	//Adding key
	req.Header.Add("X-RapidAPI-Key", "4d2c003e3emsh82512a559ce6b89p16ac42jsnc0e7dc1d61ac")

	// sending the request to receive the respone
	res, err := http.DefaultClient.Do(req)
	checkerrorplay(err)

	var apiResponse APIResponseplay
	err = json.NewDecoder(res.Body).Decode(&apiResponse)
	checkerrorplay(err)

	filepath := file

	// closing the repsone body
	defer res.Body.Close()

	appendToJSONFileplay(filepath, apiResponse)

	//Accessing the response body
	body, err := ioutil.ReadAll(res.Body)
	checkerrorplay(err)

	//	print the response
	println(string(body))

}
