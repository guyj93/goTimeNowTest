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
		for i := 1; i < dl; i++ {
			d := durationList[i] - durationList[i-1]
			if d < minDurationDiff {
				minDurationDiff = d
			}
		}
		fmt.Println("min duration diff: ", minDurationDiff)

		minDurationDiffNotZeroInit := time.Hour
		minDurationDiffNotZero := minDurationDiffNotZeroInit
		for i := 1; i < dl; i++ {
			d := durationList[i] - durationList[i-1]
			if d < minDurationDiffNotZero && d != 0 {
				minDurationDiffNotZero = d
			}
		}
		if minDurationDiffNotZero == minDurationDiffNotZeroInit {
			fmt.Println("min duration diff (not zero): N/A")
		} else {
			fmt.Println("min duration diff (not zero): ", minDurationDiffNotZero)
		}

		for i := 1; i < dl; i++ {
			durationList[i] = time.Duration(Gcd(int64(durationList[i]), int64(durationList[i-1])))
		}
		fmt.Println("GCD: ", durationList[dl-1])
	}

}
