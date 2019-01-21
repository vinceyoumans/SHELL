package main

import "time"

type Config struct {
	// Boom    time.Duration `json:"boom"`
	Timer01 struct {
		TickEverySecond time.Duration `json:"tick_every_second"`
		Count           int           `json:"count"`
		Message         string        `json:"message"`
	} `json:"timer01"`
	Timer02 struct {
		TickEverySecond time.Duration `json:"tick_every_second"`
		Count           int           `json:"count"`
		Message         string        `json:"message"`
	} `json:"timer02"`
	Timer03 struct {
		TickEverySecond time.Duration `json:"tick_every_second"`
		Count           int           `json:"count"`
		Message         string        `json:"message"`
	} `json:"timer03"`
	Boom struct {
		BoomTime time.Duration `json:"boom_time_second"`
		Message  string        `json:"message"`
	} `json:"boom"`
}
