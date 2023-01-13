package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//this is a modification of https://golangbyexample.com/http-basic-auth-golang/
func main() {
	call("https://jsantore.wufoo.com/api/v3/forms/cubes-project-proposal-submission/entries/json", "GET")
}

func call(url, method string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	req.SetBasicAuth(wufooKey, "password")
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()
	body, error := ioutil.ReadAll(response.Body)
	var prettyJSON bytes.Buffer
	error = json.Indent(&prettyJSON, body, "", "\t")
	if error != nil {
		log.Println("JSON parse error: ", error)
		return error
	}
	var jsonResponse WufooResponse
	json.Unmarshal(prettyJSON.Bytes(), &jsonResponse)
	var usefulData = jsonResponse.Entries
	processData(usefulData)
	return nil
}

func processData(webData []WuFooData) {
	for _, entry := range webData {
		entry.prettyPrint()
	}
}
