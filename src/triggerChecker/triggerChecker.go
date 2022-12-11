package triggerChecker

import (
	"encoding/json"
	"os"
	"time"

	"raspberrypi.local/cronTrigger/checkpointReceiver"
)

const (
	cacheFile = "../alreadySentTriggers.json"
)

type AlreadySentTriggers map[string]bool

func ShouldTriggerLight(alreadySentTriggers AlreadySentTriggers, now time.Time, checkpoints *checkpointReceiver.Response) bool {
	for _, element := range checkpoints.Checkpoints {
		if alreadySentTriggers[element] != true {
			return true
		}
	}

	return false
}

func SaveDataToFile(position *AlreadySentTriggers) {
	file, _ := json.MarshalIndent(position, "", " ")
	_ = os.WriteFile(cacheFile, file, 0644)
}

func GetDataFromFile() (*AlreadySentTriggers, error) {
	file, readErr := os.ReadFile(cacheFile)
	if readErr != nil {
		return nil, readErr
	}

	var result *AlreadySentTriggers
	err := json.Unmarshal([]byte(file), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func RemoveDataFile(position *AlreadySentTriggers) {
	err := os.Remove(cacheFile)
	if err != nil {
		panic(err)
	}
}
