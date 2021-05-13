package api

import "time"

// NewStartScenarioSchedule represent new request for scheduling a scenario.
// NewStartScenarioSchedule is used in scheduler package.
// ScheduledTime is the time that scenario should be started.
// SystemTime is the time that this scheduling occurred.
// Actually client send a request with the user SystemTime and ScheduledTime.
// So it considers ScheduledTime - SystemTime as a delay.
type NewStartScenarioSchedule struct {
	ScheduledTime time.Time `json:"scheduledTime"`
	SystemTime    time.Time `json:"systemTime"`
}
