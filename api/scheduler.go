package api

import "time"

// NewStartScenarioSchedule represent new request for scheduling a scenario.
// NewStartScenarioSchedule is used in scheduler package.
// ScheduledTime is the time that scenario should be started.
// SystemTime is the time that this scheduling occurred.
// Actually client send a request with the user SystemTime and ScheduledTime.
// So it considers ScheduledTime - SystemTime as a delay.
// Interval is the period time to repeat the task from the point that
// the scenario is scheduled. Zero interval meant to start the scenario
// just once.
// Token is used when communication established between legato scheduler.
type NewStartScenarioSchedule struct {
	ScheduledTime time.Time `json:"scheduledTime"`
	SystemTime    time.Time `json:"systemTime"`
	Interval      int32     `json:"interval"`
	Token         []byte     `json:"token"`
}
