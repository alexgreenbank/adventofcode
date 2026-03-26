package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

// Find total of numbers in [n1,n2] are divisible by div
func sumDivs(n1, n2, div int) int {
	min := -1
	max := -1
	if n1%div == 0 {
		// 90 % 10 == 0 so 90 is the minimum value divisible by div
		min = n1
	} else {
		// 99 % 10 == 9 so 99-9+10 = 100 is the minimum value divisble by div
		min = n1 - (n1 % div) + div
	}
	if min > n2 {
		// If the first value divisible by div is >n2 then we are done
		// e.g. 85,95 div 20 would give us a min of 100 which is >n2
		return 0
	}
	if n2%div == 0 {
		// 100 % 10 == 0 so 100 is the max value divisible by div
		max = n2
	} else {
		// 105 % 10 == 5 so 100-5 = 100 is the max value divisible by div
		max = n2 - (n2 % div)
	}
	if max < n1 {
		// If the last value divisible by div is <n1 then we are done
		// Don't think this can happen otherwise we would have bailed on min by now anyway
		return 0
	}
	ran := max - min
	// ran must be a multiple of div so save to divide here
	if ran%div != 0 {
		log.Fatalf("Somehow got ran=%d not divisible by %d", ran, div)
	}
	nosRan := ran / div
	//fmt.Printf("DEBUG: %d,%d -> %d,%d ran=%d nosRan=%d\n", n1, n2, min, max, ran, nosRan)
	// nosRan is the number of values which match
	return (min * (nosRan + 1)) + ((nosRan * (nosRan + 1) / 2) * div)
}

var divs = []int{0, 0, 11, 0, 101, 0, 1001, 0, 10001, 0, 100001, 0, 1000001}

// Find total of numbers in [n1,n2] that are formed of two sets of digits repeated
// We do this by checking for divisibility using numbers of the form 11, 101, 1001, etc
// e.g. 88 and 99 are divisible by 11
// e.g. 1010 and 1111 are divisible by 101
// etc
// Assumes: n1 <= n2
func doRange(n1, n2 int) int {
	n1len := len(fmt.Sprintf("%d", n1))
	n2len := len(fmt.Sprintf("%d", n2))

	if n1 > n2 {
		log.Fatalf("FATAL: %d > %d", n1, n2)
	}
	if n2len > 10 {
		log.Fatalf("FATAL: Only handles numbers up to 10 digits long, got %d", n2)
	}
	p1 := 0
	// Starting at length 1 we check each digit length individually
	baselen := 1
	basemax := 10
	bottom := 0
	for baselen <= n2len {
		if baselen == n1len {
			// We start at n1
			bottom = n1
		} else if baselen > n1len {
			// The bottom value for this range is the top of the previous range
			// that was updated at the end of the previous loop
		}
		// Default the top to one below the next order of magnitude
		top := basemax - 1
		// Check if we have reached the top of the range
		if n2 < basemax {
			top = n2
		}
		if baselen >= n1len {
			// We have reached the start of the range
			if baselen%2 == 0 {
				// We have an even number of digits
				r := sumDivs(bottom, top, divs[baselen])
				// fmt.Printf("DEBUG: %d-%d len %d gives %d\n", bottom, top, baselen, r)
				p1 += r
			}
		} else {
			// fmt.Printf("\tUNHANDLED: %d-%d len %d\n", bottom, top-1, baselen)
		}
		// That was our loop, we iterate by bumping up the values
		baselen++
		bottom = top
		basemax *= 10
	}
	return p1
}

func main() {
	start := time.Now()
	if len(os.Args) != 2 {
		log.Fatal("Need to supply filename as first arg")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	p1 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		// We have one big line, split it by , first
		ents := strings.Split(str, ",")
		for _, ent := range ents {
			var n1 int
			var n2 int
			// Parse out both numbers
			n, err := fmt.Sscanf(ent, "%d-%d", &n1, &n2)
			if err != nil {
				log.Fatal(err)
			}
			if n != 2 {
				log.Fatalf("Expected Sscanf to return 2 but got %d", n)
			}
			p1 += doRange(n1, n2)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("part1: %d\n", p1)
	fmt.Printf("dur=%s\n", time.Since(start))
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
