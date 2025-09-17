package main

import "flag"

func main() {
	fileName := flag.String("file", "problems.csv", "a csv file in the format of 'question,answer'")
	Quiz1(*fileName)
}
