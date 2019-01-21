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
var TT2 bool
var TTBOOM bool
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

	TT = false     //  t01 is True..  t02 is False.  Just togle depending on which ever fired last.
	TT2 = false    // This is a more assignment specific solution.  It assumes a constant 1 second tick.
	TTBOOM = false // will toggle on BOOM
	for {
		select {
		case <-tick:
			ElapsedTic = ElapsedTic + 1
			RemainingTicks = TotalTicks - ElapsedTic
			fmt.Println("---------------------------------------------")
			fmt.Println("-----------      Elapsed Tick ", ElapsedTic)
			fmt.Println("-----------   Remaining Ticks ", RemainingTicks)

			if !TTBOOM && !TT2 {
				//  Tick event
				fmt.Println("--  Event :  ", CC.MesgTick)
			}
			if !TTBOOM && TT2 {
				//  Tock event
				fmt.Println("--  Event :  ", CC.MesgTock)
			}
			if TTBOOM {
				//  BOOM event
				fmt.Println("--  Event :  ", CC.Boom.Message)
			}


			
			//  Traditional
			//fmt.Println("ticks remaining........", RemainingTicks)
			//fmt.Println("--------------------------------------")

			//  Only 1 response is permitted per Tick.
			//  A tick is every second
			//  So this is acting like a state machine.
			//  The Ticks and Tocks are not every other second..
			//     	- CC.Timer01.TickEverySecond
			//		- CC.Timer02.TickEverySecond
			//		Control the fire events.  So T01 could be every 2 seconds and T02 could be every 20 seconds.
			//		This Does create a race Condition, however, so there is logic in the TICK routine
			//		To select the best responce.
			//		I will create a new system that does not have the same race conditon.

		case <-tickreset:
			CC = GetShellDefaults()
			// 	Changes in the responces, can be set at any time, by editing the SHELL.JSON file.
			//	Once the timer starts, The intervals can not be changed with out a restart.
			//	The changes take effect every 10 seconds after the edited SHELL.json FIle has been saved.

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
