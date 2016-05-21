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
	flag.Parse()
	l := *lp
	start := time.Now()
	this := start
	durationList := make([]time.Duration, 1, l)
	lastDuration := time.Duration(0)
	for i := 1; i < l; i++ {
		this = time.Now()
		d := this.Sub(start)
		if d != lastDuration {
			durationList = append(durationList, d)
			lastDuration = d
		}
	}

	dl := len(durationList)
	if dl > 1 {
		for i := 1; i < dl; i++ {
			fmt.Println(durationList[i])
			durationList[i] = time.Duration(Gcd(int64(durationList[i]), int64(durationList[i-1])))
		}
	}
	fmt.Println("GCD: ", durationList[dl-1])
}
