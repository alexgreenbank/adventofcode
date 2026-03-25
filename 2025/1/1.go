package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

// Perform a move from current dial position
// Returns new dial position and p1 score increase
func doMove(curr int, move string) (int, int) {
	p1 := 0

	var n int
	n, err := strconv.Atoi(move[1:])
	if err != nil {
		log.Fatal(err)
	}
	if move[0] == 'L' {
		n = 0-n
	} else if move[0] != 'R' {
		log.Fatalf("FATAL: instruction does not begin with either 'L' or 'R' [%s]", move)
	}
	curr += n
	// % on negative numbers will produce a negative number
	// -47 % 100 = -47
	// -147 % 100 = -147
	// So we perform the modulus operation, add on 100 and perform it again
	// This covers both positive and negative values of any magnitude
	curr = ( ( curr % 100 ) + 100 ) % 100
	// fmt.Printf("AFTER [%s] dial = %d\n", str, curr)
	if curr == 0 {
		p1++
	}
	return curr, p1
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
	curr := 50

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		newDial, p1inc := doMove(curr, str)

		p1 += p1inc
		curr = newDial
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
