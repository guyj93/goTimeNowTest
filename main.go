// hello project main.go
package main

import (
	"flag"
	"fmt"
	"time"
)

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func Gcd(a, b int64) int64 {
	if a > b {
		return gcd(a, b)
	} else {
		return gcd(b, a)
	}
}

func main() {
	lp := flag.Int("l", 100000, "Length of Loop")
	verboseP := flag.Bool("v", false, "verbose mode: output each duration")
	flag.Parse()
	l := *lp
	verbose := *verboseP

	durationList := make([]time.Duration, l)
	//lastDuration := time.Duration(0)
	start := time.Now()
	for i := 0; i < l; i++ {
		durationList[i] = time.Since(start)
	}
	if verbose {
		for k, d := range durationList {
			fmt.Println(k, "\t", d)
		}
	}

	dl := len(durationList)
	if dl > 1 {
		minDurationDiff := durationList[1] - durationList[0]
		minDurationDiffNotZeroInit := time.Hour
		minDurationDiffNotZero := minDurationDiffNotZeroInit
		noRepeatdurationList := []time.Duration{durationList[0]}
		avgDurationPerLoop := time.Duration(0)
		for i := 1; i < dl; i++ {
			d := durationList[i] - durationList[i-1]
			if d < minDurationDiff {
				minDurationDiff = d
			}
			if d != 0 {
				noRepeatdurationList = append(noRepeatdurationList, durationList[i])
				avgDurationPerLoop = durationList[i] / time.Duration(i)
				if d < minDurationDiffNotZero {
					minDurationDiffNotZero = d
				}
			}

		}
		fmt.Println("min duration diff: ", minDurationDiff)
		if minDurationDiffNotZero == minDurationDiffNotZeroInit {
			fmt.Println("min duration diff (not zero): N/A")
		} else {
			fmt.Println("min duration diff (not zero): ", minDurationDiffNotZero)
		}
		nrdl := len(noRepeatdurationList)
		for i := 1; i < nrdl; i++ {
			noRepeatdurationList[i] = time.Duration(Gcd(int64(noRepeatdurationList[i]), int64(noRepeatdurationList[i-1])))
		}
		fmt.Println("GCD: ", noRepeatdurationList[nrdl-1])
		fmt.Println("avg duration per loop:", avgDurationPerLoop)
	}

}
