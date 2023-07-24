package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	apiurlone = "https://india-pincode-with-latitude-and-longitude.p.rapidapi.com/api/v1/pincode/400701"
	apiurltwo = "https://carbonfootprint1.p.rapidapi.com/TreeEquivalent?weight=200&unit=kg"
)

type PincodeResponse struct {
	GetDistricts string `json:"getDistricts"`
}

type WeatherResponse struct {
	AirQualityIndex float32 `json:"airQualityIndex"`
}

func fetchPinCodeData() (*PincodeResponse, error) {
	resp, err := http.Get(apiurlone)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data PincodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func fetchWeatherData() (*WeatherResponse, error) {
	resp, err := http.Get(apiurltwo)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func main() {
	// println("Calling two apis")

	pincodeChan := make(chan *PincodeResponse)
	weatherChan := make(chan *WeatherResponse)

	go func() {
		data, err := fetchPinCodeData()
		if err != nil {
			println(err)
			pincodeChan <- nil
			return
		}
		pincodeChan <- data
	}()

	go func() {
		data, err := fetchWeatherData()
		if err != nil {
			fmt.Println("Error fetching weather data:", err)
			weatherChan <- nil
			return
		}
		weatherChan <- data
	}()

	var pincodeData *PincodeResponse
	var weatherData *WeatherResponse

	select {
	case pincodeData = <-pincodeChan:
	case weatherData = <-weatherChan:
	}

	// Process the responses or combine them as needed
	if pincodeData != nil && weatherData != nil {
		fmt.Println(pincodeData)
		fmt.Println(weatherData)
	} else if pincodeData != nil {
		fmt.Println(pincodeData)
	} else if weatherData != nil {
		fmt.Println(weatherData)
	} else {
		fmt.Println("Both api calls failed")
	}
}
