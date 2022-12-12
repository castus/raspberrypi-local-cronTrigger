package triggerChecker

import (
	"sort"
	"time"

	"raspberrypi.local/cronTrigger/checkpointReceiver"
)

const (
	CheckFrequencyDurationString = "10m"
)

func ShouldTriggerLight(now time.Time, checkpoints *checkpointReceiver.Response) bool {
	sort.Strings(checkpoints.Checkpoints)

	for _, element := range checkpoints.Checkpoints {
		t, _ := time.Parse(time.UnixDate, element)
		diff := now.Sub(t).Minutes()
		if diff < 0 {
			continue
		}

		if diff == 0 || diff <= getFrequencyDuration().Minutes() {
			return true
		}
	}

	return false
}

func getFrequencyDuration() time.Duration {
	t, err := time.ParseDuration(CheckFrequencyDurationString)
	if err != nil {
		panic("Error parsing Check frequency duration string")
	}

	return t
}
