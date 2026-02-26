package tabler

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"window/pkg/models"
	"window/pkg/parser"
)

type Tabler struct {
	tableName string
}

var (
	ErrBadLine = errors.New("err cannot save line ")
)

func NewTabler(tableName string) (*Tabler, error) {
	if tableName == "" {
		return nil, fmt.Errorf("No table name")
	}
	tabler := Tabler{tableName: tableName}
	return &tabler, nil
}

func (t *Tabler) InitTable() []models.Lesson {
	file, err := os.Open(t.tableName + ".txt")
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

	lessons := make([]models.Lesson, 0, len(splitted))
	for i := range splitted {
		lesson, err := parser.ProcessInput(splitted[i])
		if err != nil {
			slog.Info("Init: cannot process line", "line", splitted[i], "err", err)
			continue
		}
		lessons = append(lessons, lesson)
	}
	return lessons
}

func (t *Tabler) AddElement(input string) (models.Lesson, error) {
	return parser.ProcessInput(input)
}

func (t *Tabler) SaveTable(table [][]string) []error {
	var errs []error
	slog.Info("Got table", "table", table)
	data := ""
	for i := range table {
		line := fmt.Sprintf("%v %v \"%v\"\n", table[i][0], table[i][1], table[i][2])
		if _, err := parser.ProcessInput(line); err != nil {
			slog.Error("SaveTable: process line error", "err", errors.New(ErrBadLine.Error()+line))
			errs = append(errs, errors.New(ErrBadLine.Error()+line))
			continue
		}
		data += line
	}

	file, err := os.Create(t.tableName + ".txt")
	if err != nil {
		slog.Error("SaveTable: open file error", "err", err)
		return errs
	}
	defer file.Close()

	_, err = file.Write([]byte(data))
	if err != nil {
		slog.Error("SaveTable: write file error", "err", err)
		return errs
	}

	return errs
}
