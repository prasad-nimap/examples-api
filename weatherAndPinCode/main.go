package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	url := "https://get-details-by-pin-code-india.p.rapidapi.com/detailsbypincode"

	payload := strings.NewReader("{\r\n    \"pincode\": \"415001\"\r\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", "4d2c003e3emsh82512a559ce6b89p16ac42jsnc0e7dc1d61ac")
	req.Header.Add("X-RapidAPI-Host", "get-details-by-pin-code-india.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
