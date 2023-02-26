package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/robfig/cron/v3"

	"raspberrypi.local/cronTrigger/checkpointReceiver"
	"raspberrypi.local/cronTrigger/mqttHandler"
	"raspberrypi.local/cronTrigger/triggerChecker"
)

func main() {
	periodicallyCheckForLightTrigger()

	log.Println("type=success msg=\"Cron Trigger is up and running\"")

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func periodicallyCheckForLightTrigger() {
	c := cron.New()
	_, err := c.AddFunc("@every "+triggerChecker.CheckFrequencyDurationString, func() {
		checkpoints := checkpointReceiver.GetCheckpoints()
		log.Printf("type=debug croncheck=true")
		l, _ := time.LoadLocation("Europe/Warsaw")
		now := time.Now().In(l)
		log.Printf("type=debug msg=\"Get checkpoints\" checkpoints=\"%s\"\n", checkpoints)
		log.Printf("type=debug msg=\"Current time is: %s\"\n", now.String())
		if triggerChecker.ShouldTriggerLight(now, checkpoints) {
			log.Printf("type=debug crontrigger=true msg=\"Reached a checkpoint, triggerring lightController, time: %s\"\n", now.String())
			go mqttHandler.PublishMessage(getMessage())
		} else {
			log.Printf("type=debug crontrigger=false msg=\"No need to trigger lightController, time: %s\"\n", now.String())
		}
	})
	if err != nil {
		panic(err)
	}
	c.Start()
}

func getMessage() string {
	message := mqttHandler.Message{
		Place: "cron",
	}
	m, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	return string(m)
}
