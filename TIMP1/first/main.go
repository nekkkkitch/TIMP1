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
	splitInput := strings.Split(input, " ")
	var lesson Lesson
	writingString := false
	for i := range splitInput {
		if splitInput[i] == "" {
			continue
		}

		if !writingString {
			t, err := time.Parse("2006.01.02", splitInput[i])
			if err == nil {
				lesson.Date = t
				continue
			} else {
				slog.Info("Property cannot be processed as year of the day.", "property", splitInput[i])
			}

			t, err = time.Parse("15:04", splitInput[i])
			if err == nil {
				lesson.Time = t
				continue
			} else {
				slog.Info("Property cannot be processed as time of the day.", "property", splitInput[i])
			}
		}

		runedFirst := []rune(splitInput[i])
		if runedFirst[0] == '"' {
			if len(runedFirst) == 1 {
				slog.Error("Property contain single quotation mark", "property", splitInput[i], "full input", input)
				return Lesson{}, fmt.Errorf("Property contain single quotation mark")
			}
			if lesson.TeacherName != "" {
				slog.Error("More than 1 string in input", "property", splitInput[i], "full input", input)
				return Lesson{}, fmt.Errorf("More than 1 string in input")
			}
			if runedFirst[len(runedFirst)-1] == '"' {
				lesson.TeacherName = string(runedFirst[1 : len(runedFirst)-1])
				continue
			}
			lesson.TeacherName = string(runedFirst[1:]) + " "
			writingString = true
			continue
		}

		if writingString {
			runedString := []rune(splitInput[i])
			if runedString[len(runedString)-1] == '"' {
				writingString = false
				lesson.TeacherName += string(runedString[:len(runedString)-1])
				continue
			}
			lesson.TeacherName += string(runedString) + " "
			continue
		}
		slog.Error("Property cannot be processed.", "property", splitInput[i], "full input", input)
		return Lesson{}, fmt.Errorf("Unpocessable property: %v", splitInput[i])
	}
	return lesson, nil
}
