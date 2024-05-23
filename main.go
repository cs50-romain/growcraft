package main

import (
	"flag"
	"fmt"

	"github.com/cs50-romain/growcraft/cmd"
)

func main() {
	filePtr := flag.String("file", "default.csv", "A csv file to save your progress")
	subjectPtr := flag.String("subj", "default", "What subject within your craft did your do today")
	timePtr := flag.Int("time", 000, "The amount of time spent working on your craft today in minutes")
	multiplierPtr := flag.Float64("mult", 0.0, "The multiplier effect. From 1.0 to 2.0 how much were you working on something out of your comfort zone")

	flag.Parse()
	if err := cmd.Run(*filePtr, *subjectPtr, *timePtr, *multiplierPtr); err != nil {
		fmt.Println("\u001b[1;31m" + err.Error() + "\u001b[0m")
	}
}
