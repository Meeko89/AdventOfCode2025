package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func day1(runExample bool) {
	dataFileName := "day1"
	if runExample {
		dataFileName += "_example"
	}
	// Open the file
	file, err := os.Open(filepath.Join("data", dataFileName+".txt"))
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	// Create a new reader
	reader := bufio.NewReader(file)
	pos := 50
	zeros := 0
	for {
		// Read until we encounter a newline character
		line, err := reader.ReadString('\n')
		if err != nil {
			// If we encounter EOF, break out of the loop
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("error reading file: %s", err)
		}

		dir := string(line[0])
		clicks, err := strconv.Atoi(line[1 : len(line)-1])
		if err != nil {
			log.Fatalf("failed to parse clicks: %s", err)
		}
		modifier := 1
		if dir == "L" {
			modifier = -1
		}

		pos += modifier * clicks
		fmt.Print("clicks: ", clicks, "\tpos: ", pos)
		if pos < 0 {
			pos += 100
		}
		pos = pos % 100
		fmt.Println("\tpos: ", pos)
		if pos == 0 {
			zeros++
		}

	}
	fmt.Println("Zeros:", zeros)
}

func main() {
	day := 1
	runExample := false
	switch day {
	case 1:
		day1(runExample)
	}
}
