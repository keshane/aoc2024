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

func problem1() {
	file, err := os.Open("inputs/problem1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fscanner := bufio.NewScanner(file)
	var left []int
	var right []int
	for fscanner.Scan() {
		line := fscanner.Text()
		row := strings.Fields(line)
		leftVal, _ := strconv.Atoi(row[0])
		rightVal, _ := strconv.Atoi(row[1])
		left = append(left, leftVal)
		right = append(right, rightVal)
	}
	slices.Sort(left)
	slices.Sort(right)

	dist := 0
	rightCount := map[int]int{}
	for i := 0; i < len(left); i++ {
		dist += max(left[i], right[i]) - min(left[i], right[i])
		rightCount[right[i]] += 1
	}

	fmt.Println(dist)

	similarity := 0
	for i := 0; i < len(left); i++ {
		similarity += left[i] * rightCount[left[i]]
	}
	fmt.Println(similarity)
}
