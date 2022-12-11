package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/robfig/cron/v3"

	"raspberrypi.local/cronTrigger/checkpointReceiver"
)

var checkpoints *checkpointReceiver.Response

func init() {
	checkpoints = checkpointReceiver.GetCheckpoints()
}

func main() {
	automaticallyRefreshDataWhenDayStarts()
	periodicallyCheckForLightTrigger()
	//
	// message := mqttHandler.Message{
	// 	IsLightOn: true,
	// 	Place:     "cron",
	// }
	// m, err := json.Marshal(message)
	// if err != nil {
	// 	panic(err)
	// }
	// go mqttHandler.PublishMessage(string(m))

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func automaticallyRefreshDataWhenDayStarts() {
	c := cron.New()
	_, err := c.AddFunc("0 2 * * *", func() { // At 2:00 AM
		fmt.Println("Cron function executes")
		checkpointReceiver.GetCheckpoints()
	})
	if err != nil {
		panic(err)
	}
	c.Start()
}

func periodicallyCheckForLightTrigger() {
	c := cron.New()
	_, err := c.AddFunc("@every 1m", func() {
		fmt.Println("Cron function executes")
		// checkpointReceiver.GetCheckpoints()
	})
	if err != nil {
		panic(err)
	}
	c.Start()
}
