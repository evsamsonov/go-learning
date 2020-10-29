package main

import (
	"bufio"
	"fmt"
	"github.com/apex/log"
	"io"
	"os"
	"regexp"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Provide text file")
		os.Exit(1)
	}

	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Failed to open file")
		os.Exit(1)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("Failed to close file")
		}
	}()

	i := 0
	reader := bufio.NewReader(f)
	dateRegexp := regexp.MustCompile(`.*\[(\d{2}/\w+/\d{4}:\d{2}:\d{2}:\d{2}.*)] .*`)
	dateRegexp2 := regexp.MustCompile(`.*\[(\w+-\d{2}-\d{2}:\d{2}:\d{2}:\d{2}.*)] .*`)
	for {
		i++
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("%d: failed to read string: %s", i, err)
			return
		}

		var date time.Time
		switch {
		case dateRegexp.MatchString(line):
			match := dateRegexp.FindStringSubmatch(line)
			date, err = time.Parse("02/Jan/2006:15:04:05 -0700", match[1])
			if err != nil {
				fmt.Printf("%d: failed to parse date: %s\n", i, match[1])
				continue
			}
		case dateRegexp2.MatchString(line):
			match := dateRegexp2.FindStringSubmatch(line)
			date, err = time.Parse("Jan-02-06:15:04:05 -0700", match[1])
			if err != nil {
				fmt.Printf("%d: failed to parse date: %s\n", i, match[1])
				continue
			}
		default:
			log.Errorf("%d: date not found\n", i)
		}
		fmt.Printf("%d: %s\n", i, date)
	}

}
