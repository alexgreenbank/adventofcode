package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	//"strconv"
	"time"
)

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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//str := scanner.Text()
		//var n int
		//n, err = strconv.Atoi(str)
		//if err != nil {
		//log.Fatal(err)
		//}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("part1: \n")
	fmt.Printf("dur=%s\n", time.Since(start))
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
