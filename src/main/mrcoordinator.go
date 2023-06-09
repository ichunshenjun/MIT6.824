package main

//
// start the coordinator process, which is implemented
// in ../mr/coordinator.go
//
// go run mrcoordinator.go pg*.txt
//
// Please do not change this file.
//

import "6.5840/mr"
import "time"
import "os"
import "fmt"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: mrcoordinator inputfiles...\n")
		os.Exit(1)
	}

	m := mr.MakeCoordinator(os.Args[1:], 10)
	for m.Done() == false {
		if m.State == mr.MAPING {
			for i := 0; i < len(m.Maps); i++ {
				if time.Since(m.Maps[i].Time).Seconds() > 20 && m.Maps[i].MapState == mr.DOING {
					m.Maps[i].MapState = mr.UNSTART
				}
			}
		} else if m.State == mr.REDUCING {
			for i := 0; i < len(m.Reduces); i++ {
				if time.Since(m.Reduces[i].Time).Seconds() > 20 && m.Reduces[i].ReduceState == mr.DOING {
					m.Reduces[i].ReduceState = mr.UNSTART
				}
			}
		}
		time.Sleep(time.Second)
	}

	time.Sleep(time.Second)
}
