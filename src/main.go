package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/robfig/cron/v3"

	"raspberrypi.local/cronTrigger/checkpointReceiver"
	"raspberrypi.local/cronTrigger/mqttHandler"
	"raspberrypi.local/cronTrigger/triggerChecker"
)

var checkpoints *checkpointReceiver.Response

func init() {
	checkpoints = checkpointReceiver.GetCheckpoints()
}

func main() {
	automaticallyRefreshDataWhenDayStarts()
	periodicallyCheckForLightTrigger()

	fmt.Println("Cron Trigger is up and running")

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func automaticallyRefreshDataWhenDayStarts() {
	c := cron.New()
	_, err := c.AddFunc("0 2 * * *", func() { // At 2:00 AM
		fmt.Println("Executing RefreshData function")
		checkpointReceiver.GetCheckpoints()
	})
	if err != nil {
		panic(err)
	}
	c.Start()
}

func periodicallyCheckForLightTrigger() {
	c := cron.New()
	_, err := c.AddFunc("@every "+triggerChecker.CheckFrequencyDurationString, func() {
		fmt.Println("Executing CheckForLightTrigger function")
		checkpoints := checkpointReceiver.GetCheckpoints()
		now := time.Now()
		if triggerChecker.ShouldTriggerLight(now, checkpoints) {
			fmt.Println("Current time: " + now.String() + ", triggering lightController")
			go mqttHandler.PublishMessage(getMessage())
		} else {
			fmt.Println("Current time: " + now.String() + ", no need to trigger lightController")
		}
	})
	if err != nil {
		panic(err)
	}
	c.Start()
}

func getMessage() string {
	message := mqttHandler.Message{
		IsLightOn: true,
		Place:     "cron",
	}
	m, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	return string(m)
}
