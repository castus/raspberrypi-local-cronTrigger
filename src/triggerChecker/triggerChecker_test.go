package triggerChecker_test

import (
	"testing"
	"time"

	"raspberrypi.local/cronTrigger/checkpointReceiver"
	"raspberrypi.local/cronTrigger/triggerChecker"
)

func TestShouldTriggerAtTheCheckpointTime(t *testing.T) {
	checkpoints := getCheckpoints()
	theSameTimeAsCheckpoints := checkpoints.Checkpoints
	for _, element := range theSameTimeAsCheckpoints {
		now, _ := time.Parse(time.UnixDate, element)
		if triggerChecker.ShouldTriggerLight(now, checkpoints) != true {
			t.Fail()
		}
	}
}

func TestShouldTriggerAfterCheckpointAndWhenDifferenceIsAtMaxAsTheFrequency(t *testing.T) {
	checkpoints := getCheckpoints()

	if triggerChecker.ShouldTriggerLight(getTime("Sun Dec 11 03:01:00 CET 2022"), checkpoints) != true {
		t.Fail()
	}
	if triggerChecker.ShouldTriggerLight(getTime("Sun Dec 11 03:09:00 CET 2022"), checkpoints) != true {
		t.Fail()
	}
	if triggerChecker.ShouldTriggerLight(getTime("Sun Dec 11 03:10:00 CET 2022"), checkpoints) != true {
		t.Fail()
	}
}

func TestShouldNotTriggerBeforeCheckpoint(t *testing.T) {
	checkpoints := getCheckpoints()

	if triggerChecker.ShouldTriggerLight(getTime("Sun Dec 11 01:00:00 CET 2022"), checkpoints) == true {
		t.Fail()
	}
	if triggerChecker.ShouldTriggerLight(getTime("Sun Dec 11 02:59:59 CET 2022"), checkpoints) == true {
		t.Fail()
	}
	if triggerChecker.ShouldTriggerLight(getTime("Sun Dec 11 06:48:24 CET 2022"), checkpoints) == true {
		t.Fail()
	}
	if triggerChecker.ShouldTriggerLight(getTime("Sun Dec 11 07:48:24 CET 2022"), checkpoints) == true {
		t.Fail()
	}
}

func TestShouldNotTriggerWhenDifferenceIsGreaterThenTheFrequency(t *testing.T) {
	checkpoints := getCheckpoints()

	if triggerChecker.ShouldTriggerLight(getTime("Sun Dec 11 03:11:00 CET 2022"), checkpoints) == true {
		t.Fail()
	}
	if triggerChecker.ShouldTriggerLight(getTime("Sun Dec 11 04:00:00 CET 2022"), checkpoints) == true {
		t.Fail()
	}
	if triggerChecker.ShouldTriggerLight(getTime("Sun Dec 11 05:48:00 CET 2022"), checkpoints) == true {
		t.Fail()
	}
}

func getTime(dateString string) time.Time {
	t, _ := time.Parse(time.UnixDate, dateString)
	return t
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
