package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"
)

type Lesson struct {
	Date        time.Time
	Time        time.Time
	TeacherName string
}

func main() {
	in := bufio.NewReader(os.Stdin)
	input, _ := in.ReadString('\n')
	input = string([]rune(input)[:len([]rune(input))-1])

	lesson, err := ProcessInput(input)
	if err != nil {
		slog.Error("ProcessInput failed", "err", err)
		os.Exit(1)
	}

	fmt.Printf("Result lesson: %v\n", lesson)
}

func ProcessInput(input string) (Lesson, error) {
	lesson := Lesson{}
	splitted := strings.Split(input, "\"")
	if len(splitted) != 3 {
		return Lesson{}, fmt.Errorf("Bad number of quotation marks")
	}

	lesson.TeacherName = splitted[1]

	timePart := ""
	if splitted[0] == "" {
		timePart = splitted[2]
	} else if splitted[1] == "" {
		timePart = splitted[0]
	} else {
		timePart = strings.Join([]string{splitted[0], splitted[2]}, " ")
	}

	splittedTime := strings.Split(timePart, " ")
	for i := range splittedTime {
		if splittedTime[i] == "" {
			continue
		}

		if len([]rune(splittedTime[i])) == 5 {
			time, err := time.Parse("15:04", splittedTime[i])
			if err != nil {
				return Lesson{}, err
			}
			lesson.Time = time
		} else {
			time, err := time.Parse("2006.01.02", splittedTime[i])
			if err != nil {
				return Lesson{}, err
			}
			lesson.Date = time
		}
	}

	return lesson, nil
}
