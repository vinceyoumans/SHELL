package main

import (
	"encoding/json"
	"fmt"
	"time"
)

const Minutez time.Duration = 1000 * time.Millisecond

var CC Config

var TotalTicks int = 0
var ElapsedTic int = 0
var RemainingTicks = 0
var TT bool
var TTLock bool // True when there is a Boom
func main() {


	fmt.Println("========   starting project ==========")
	CC = GetShellDefaults()
	TotalTicks = int(CC.Boom.BoomTime) / 1000000000
	fmt.Println(CC)
	tick := time.Tick(1000 * time.Millisecond)
	tickreset := time.Tick(10000 * time.Millisecond) // will reset messages every 10 seconds
	timer01 := time.Tick(CC.Timer01.TickEverySecond * Minutez)
	timer02 := time.Tick(CC.Timer02.TickEverySecond * Minutez)
	timer03 := time.Tick(CC.Timer03.TickEverySecond * Minutez)
	boom := time.After(CC.Boom.BoomTime)

	TT = false

	for {
		select {
		case <-tick:
			ElapsedTic = ElapsedTic + 1
			RemainingTicks = TotalTicks - ElapsedTic
			//fmt.Println("ticks remaining........", RemainingTicks)
			//fmt.Println("--------------------------------------")

		case <-tickreset:
			CC = GetShellDefaults()
			//  Every 10 seconds, the shell.json file will be 
			//  reviewed.  Messages will be changed and the count will be set to 0

		case <-timer01:
			//fmt.Println("tick01.")
			CC.Timer01.Count = CC.Timer01.Count + 1
			if TT {
				//fmt.Println("------  Tick01 ", CC.Timer01.Count)
				fmt.Println("--------------------------------------", CC.Timer01.Message)
				TT = false
			}
		case <-timer02:
			CC.Timer02.Count = CC.Timer02.Count + 1
			if !TT {
				//fmt.Println("------  Tick02 ", CC.Timer02.Count)
				fmt.Println("--------------------------------------", CC.Timer02.Message)
				TT = true
			}

		case <-timer03:
			CC.Timer03.Count = CC.Timer03.Count + 1
			//fmt.Println("------  Tick03 ", CC.Timer03.Count)

			fmt.Println("--------------------------------------", CC.Timer03.Message)

		case <-boom:
			fmt.Println(CC.Boom.Message)
			fmt.Println("--------------------------------------")
			b, _ := json.MarshalIndent(CC, "", "  ")
			b, _ = prettyprint(b)
			fmt.Printf("%s", b)
			fmt.Println("--------------------------------------")

			return
		default:
			time.Sleep(1000 * time.Millisecond)
		}
	}

}
