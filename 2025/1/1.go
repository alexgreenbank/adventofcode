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
// Returns new dial position and p1 and p2 score increase
func doMove(curr int, move string) (int, int, int) {
	p1 := 0
	p2 := 0

	// Ensure inputs are valid
	if curr < 0 || curr >= 100 {
		log.Fatal("Invalid curr value in doMove, must be [0,99]")
	}
	if len(move) < 2 {
		log.Fatal("Invalid move string in doMove")
	}
	if move[0] != 'L' && move[0] != 'R' {
		log.Fatal("Invalid move letter in doMove")
	}

	var n int
	n, err := strconv.Atoi(move[1:])
	if err != nil {
		log.Fatal(err)
	}
	if n < 0 {
		log.Fatalf("Does not handle negative n values in input [%s]", move)
	}
	// We add on any multiples of 100 to the p2inc score
	p2 += int(n / 100)
	// We can now strip n down to the range [0,99]
	n %= 100
	if move[0] == 'L' {
		n = 0 - n
	}
	if (curr > 0 && curr+n < 0) || (curr+n > 100) {
		// This move crosses over the 0 ticker
		// We handle the case of it ending on 0 later
		// We ignore the case of it starting on 0 and ending -ve as it hasn't passed 0 on this move
		p2++
	}
	curr += n
	// % on negative numbers will produce a negative number
	// -47 % 100 = -47
	// -147 % 100 = -147
	// So we perform the modulus operation, add on 100 and perform it again
	// This covers both positive and negative values of any magnitude
	curr = ((curr % 100) + 100) % 100
	// fmt.Printf("AFTER [%s] dial = %d\n", str, curr)
	if curr == 0 {
		p1++
		if n != 0 {
			// If we did an R100 starting on 0 then we've counted the 0 for part2 already above
			// so don't double count it here
			p2++
		}
	}
	return curr, p1, p2
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
	p2 := 0
	curr := 50

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		newDial, p1inc, p2inc := doMove(curr, str)

		p1 += p1inc
		p2 += p2inc
		curr = newDial
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("part1: %d\n", p1)
	fmt.Printf("part2: %d\n", p2)
	fmt.Printf("dur=%s\n", time.Since(start))
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
