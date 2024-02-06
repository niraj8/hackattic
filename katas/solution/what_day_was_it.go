package solution

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func WhatDayWasIt() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		numDays, err := strconv.Atoi(input)
		if err != nil {
			log.Fatalf("error parsing input to int: %v", err)
		}
		var dayZero time.Time = time.Unix(0, 0)
		result := dayZero.Add(time.Hour * 24 * time.Duration(numDays))
		fmt.Println(result.Weekday())
	}
}
