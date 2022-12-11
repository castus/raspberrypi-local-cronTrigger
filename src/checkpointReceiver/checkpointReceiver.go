package checkpointReceiver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	cacheFile = "../triggers.json"
)

type Response struct {
	Checkpoints []string `json:"checkpoints"`
}

func GetCheckpoints() *Response {
	if HasCacheHit() {
		fmt.Println("Cache hit, serving data from cache")
		data, err := GetDataFromFile()
		if err != nil {
			panic(err)
		}
		return data
	}

	URL := fmt.Sprintf("%s", os.Getenv("TRIGGER_API_SERVER_ADDRESS"))
	resp, err := http.Get(URL)

	if err != nil {
		fmt.Printf("Request Failed: %s", err)
		panic("Request Failed")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Reading body failed: %s", err)
		panic("Reading body failed")
	}
	fmt.Println(string(body))
	var s = new(Response)
	err2 := json.Unmarshal(body, &s)
	if err2 != nil {
		fmt.Println("Error reading response from Triggers API")
		fmt.Println(err2)
	}
	SaveDataToFile(s)

	return s
}

func SaveDataToFile(position *Response) {
	file, _ := json.MarshalIndent(position, "", " ")
	_ = os.WriteFile(cacheFile, file, 0644)
}

func GetDataFromFile() (*Response, error) {
	file, readErr := os.ReadFile(cacheFile)
	if readErr != nil {
		return nil, readErr
	}

	var result *Response
	err := json.Unmarshal([]byte(file), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func HasCacheHit() bool {
	now := time.Now()
	data, err := GetDataFromFile()
	if err != nil {
		return false
	}

	firstDate, err := time.Parse(time.UnixDate, data.Checkpoints[0])
	if err != nil {
		panic(err)
	}

	return now.Day() == firstDate.Day() && now.Month() == firstDate.Month() && now.Year() == firstDate.Year()
}
