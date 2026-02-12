package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type Lesson struct {
	Date        string `json:"date"`
	Time        string `json:"time"`
	TeacherName string `json:"name"`
}

func processInput(input string) (Lesson, error) {
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
			lesson.Time = time.Format("15:04")
		} else {
			time, err := time.Parse("2006.01.02", splittedTime[i])
			if err != nil {
				return Lesson{}, err
			}
			lesson.Date = time.Format("2006.01.02")
		}
	}

	return lesson, nil
}

func (a *App) Init() []Lesson {
	file, err := os.Open("table.txt")
	if err != nil {
		slog.Info("Init: open file error", "err", err)
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		slog.Info("Init: read file error", "err", err)
	}

	splitted := strings.Split(string(b), "\n")
	if splitted[len(splitted)-1] == "" {
		splitted = splitted[:len(splitted)-1]
	}

	lessons := make([]Lesson, 0, len(splitted))
	for i := range splitted {
		lesson, err := processInput(splitted[i])
		if err != nil {
			slog.Info("Init: process file line error", "err", err)
		}
		lessons = append(lessons, lesson)
	}
	return lessons
}

func (a *App) AddElement(input string) (Lesson, error) {
	return processInput(input)
}

func (a *App) SaveTable(table [][]string) error {
	slog.Info("Got table", "table", table)
	data := ""
	for i := range table {
		if i == 0 {
			continue
		}
		data += fmt.Sprintf("%v %v \"%v\"\n", table[i][0], table[i][1], table[i][2])
	}

	file, err := os.Create("table.txt")
	if err != nil {
		slog.Info("Init: open file error", "err", err)
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(data))
	if err != nil {
		slog.Info("Init: open file error", "err", err)
		return err
	}

	return nil
}
