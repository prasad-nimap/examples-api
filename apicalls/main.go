package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Person struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Age   int    `json:"age"`
}

func checkForFile(fileName string) ([]byte, error) {
	_, err := os.Stat(fileName)
	checkerror(err)
	//	println(status)

	//	check if the file exist and !exist create one
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

func appendToJSONFile(data []Person, filename string) error {
	var jsondata []Person

	// Check for file and if !exist create
	file, err := checkForFile(filename)

	// Unmarshal the existing JSON data
	err = json.Unmarshal(file, &jsondata)
	checkerror(err)

	//	data := []Person{}

	newStruct := &Person{
		Fname: "Prasad",
		Lname: "Junghare",
		Age:   21,
	}

	//	append the new data
	//	jsondata = append(jsondata, data...)
	data = append(data, *newStruct)

	//	Marshal the updated JSON data
	//	data, err := append(jsondata)
	//	checkerror(err)

	/*
		//	Preparing the data to be marshalled and written
		dataBytes, err := json.Marshal(data)
		checkerror(err)

		err = ioutil.WriteFile("result.json", dataBytes, 0644)
		checkerror(err)
	*/

	databyte, err := json.Marshal(data)

	//	write the updated data to the file
	err = ioutil.WriteFile(filename, databyte, 0644)
	checkerror(err)

	return nil
}

func main() {
	//	println("api call")

	//	append using the apppend function
	appendToJSONFile(data, "result.json")
	
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

	/*
	   //	rapid api link
	   url := "https://carbonfootprint1.p.rapidapi.com/AirQualityHealthIndex?O3=10&NO2=10&PM=10"

	   //	Make a new request
	   req, err := http.NewRequest("GET", url, nil)
	   checkerror(err)

	   //Adding key
	   req.Header.Add("X-RapidAPI-Key", "4d2c003e3emsh82512a559ce6b89p16ac42jsnc0e7dc1d61ac")

	   // sending the request to receive the respone
	   res, err := http.DefaultClient.Do(req)
	   checkerror(err)

	   // closing the repsone body
	   defer res.Body.Close()

	   //Accessing the response body
	   body, err := ioutil.ReadAll(res.Body)
	   checkerror(err)

	   //	print the response
	   println(string(body))
	*/
}
