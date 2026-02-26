package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"regexp"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

var rdate *regexp.Regexp
var rname *regexp.Regexp
var rtime *regexp.Regexp

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	rdate, _ = regexp.Compile(`(\d{4})+\.+(\d{2})+\.+(\d{2})`)
	rtime, _ = regexp.Compile(`(\d{2})+\:+(\d{2})`)
	rname, _ = regexp.Compile(`\"(([A-z]){1,}\s?)*\"`)
	a.ctx = ctx

}

type Lesson struct {
	Date        string `json:"date"`
	Time        string `json:"time"`
	TeacherName string `json:"name"`
}

func processInput(input string) (Lesson, error) {
	lesson := Lesson{}
	lesson.Date = string(rdate.Find([]byte(input)))
	lesson.Time = string(rtime.Find([]byte(input)))
	lesson.TeacherName = string(rname.Find([]byte(input)))
	lesson.TeacherName = strings.ReplaceAll(lesson.TeacherName, "\"", "")
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
