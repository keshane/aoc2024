package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func problem2() {
	file, err := os.Open("inputs/problem2.txt")
	if err != nil {
		log.Fatal(err)
	}
	fscanner := bufio.NewScanner(file)
	safeCount := 0
	safeDampenedCount := 0
	for fscanner.Scan() {
		line := fscanner.Text()
		var report []int
		for _, field := range strings.Fields(line) {
			fieldVal, _ := strconv.Atoi(field)
			report = append(report, fieldVal)
		}
		if isSafe(report) {
			safeCount += 1
		}
		if isSafeDampener(report) {
			safeDampenedCount += 1
		}
	}
	fmt.Println(safeCount)
	fmt.Println(safeDampenedCount)
}

func isSafe(levels []int) bool {
	// assuming at least 2 levels
	firstDiff := levels[1] - levels[0]
	var direction int
	if firstDiff < 0 {
		direction = -1
	} else if firstDiff > 0 {
		direction = 1
	} else {
		return false
	}

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		magnitude := diff * direction
		if magnitude < 1 || magnitude > 3 {
			return false
		}
	}
	return true
}

func isSafeDampener(levels []int) bool {
	if isSafe(levels) {
		return true
	}

	for i := 0; i < len(levels); i++ {
		if isSafe(slices.Concat(levels[:i], levels[i+1:])) {
			return true
		}
	}

	return false
}
