package main

import (
	"encoding/json"
	"fmt"

	"github.com/robfig/cron/v3"

	"raspberrypi.local/cronTrigger/checkpointReceiver"
	"raspberrypi.local/cronTrigger/mqttHandler"
)

var checkpoints *checkpointReceiver.Response

func init() {
	checkpoints = checkpointReceiver.GetCheckpoints()
}

func main() {
	automaticallyRefreshDataWhenDayStarts()

	message := mqttHandler.Message{
		IsLightOn: true,
		Place:     "cron",
	}
	m, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	mqttHandler.PublishMessage(string(m))
}

func automaticallyRefreshDataWhenDayStarts() {
	c := cron.New()
	_, err := c.AddFunc("0 2 * * *", func() {
		fmt.Println("Cron function executes")
		checkpointReceiver.GetCheckpoints()
	})
	if err != nil {
		panic(err)
	}
	c.Start()
}
