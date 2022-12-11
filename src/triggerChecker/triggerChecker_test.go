package triggerChecker_test

import (
	"testing"
	"time"

	"raspberrypi.local/cronTrigger/checkpointReceiver"
	"raspberrypi.local/cronTrigger/triggerChecker"
)

func TestShouldNotTriggerWhenAllCheckpointsTriggered(t *testing.T) {
	checkpoints := getCheckpoints()
	now := time.Now().In(getLocation())
	alreadySentTriggers := triggerChecker.AlreadySentTriggers{
		"Sun Dec 11 06:48:25 CET 2022": true,
		"Sun Dec 11 07:48:25 CET 2022": true,
		"Sun Dec 11 18:58:01 CET 2022": true,
		"Sun Dec 11 17:58:01 CET 2022": true,
		"Sun Dec 11 13:00:00 CET 2022": true,
		"Sun Dec 11 11:00:00 CET 2022": true,
		"Sun Dec 11 22:00:00 CET 2022": true,
		"Sun Dec 11 23:00:00 CET 2022": true,
		"Sun Dec 11 03:00:00 CET 2022": true,
	}
	if triggerChecker.ShouldTriggerLight(alreadySentTriggers, now, checkpoints) != false {
		t.Fail()
	}
}

func TestShouldTriggerWhenSomeCheckpointNotTriggered(t *testing.T) {
	checkpoints := getCheckpoints()
	now := time.Now().In(getLocation())
	alreadySentTriggers := triggerChecker.AlreadySentTriggers{
		"Sun Dec 11 06:48:25 CET 2022": true,
		"Sun Dec 11 07:48:25 CET 2022": true,
		"Sun Dec 11 18:58:01 CET 2022": true,
		"Sun Dec 11 17:58:01 CET 2022": true,
		"Sun Dec 11 13:00:00 CET 2022": true,
		"Sun Dec 11 11:00:00 CET 2022": false,
		"Sun Dec 11 22:00:00 CET 2022": true,
		"Sun Dec 11 23:00:00 CET 2022": true,
		"Sun Dec 11 03:00:00 CET 2022": true,
	}
	if triggerChecker.ShouldTriggerLight(alreadySentTriggers, now, checkpoints) != true {
		t.Fail()
	}
}

func getLocation() *time.Location {
	l, _ := time.LoadLocation("Europe/Warsaw")
	return l
}

func getCheckpoints() *checkpointReceiver.Response {
	var checkpoints = []string{
		"Sun Dec 11 06:48:25 CET 2022",
		"Sun Dec 11 07:48:25 CET 2022",
		"Sun Dec 11 18:58:01 CET 2022",
		"Sun Dec 11 17:58:01 CET 2022",
		"Sun Dec 11 13:00:00 CET 2022",
		"Sun Dec 11 11:00:00 CET 2022",
		"Sun Dec 11 22:00:00 CET 2022",
		"Sun Dec 11 23:00:00 CET 2022",
		"Sun Dec 11 03:00:00 CET 2022",
	}

	var response = &checkpointReceiver.Response{
		Checkpoints: checkpoints,
	}

	return response
}
