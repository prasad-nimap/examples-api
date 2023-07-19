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

func checkerror(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func checkForFile(fileName string) ([]byte, error) {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			return nil, err
		}
		file.Close()
	} else if err != nil {
		return nil, err
	}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func appendToJSONFile(filename string) error {
	var jsondata []Person

	file, err := checkForFile(filename)
	if err != nil {
		return err
	}

	// Handle empty file
	if len(file) == 0 {
		data := []Person{
			{
				Fname: "Prasad",
				Lname: "Junghare",
				Age:   21,
			},
		}
		databyte, err := json.Marshal(data)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(filename, databyte, 0644)
		if err != nil {
			return err
		}

		return nil
	}

	err = json.Unmarshal(file, &jsondata)
	if err != nil {
		return err
	}

	data := append(jsondata, Person{
		Fname: "Prasad",
		Lname: "Junghare",
		Age:   21,
	})

	databyte, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, databyte, 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]

	err := appendToJSONFile(filename)
	if err != nil {
		println("Error:", err.Error())
		os.Exit(1)
	}
}
