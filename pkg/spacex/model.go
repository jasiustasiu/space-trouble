package spacex

import "time"

type Launch struct {
	LaunchDateLocal time.Time `json:"launch_date_local"`
}

type LaunchPad struct {
	Status string `json:"status"`
}
