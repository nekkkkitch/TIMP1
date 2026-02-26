package parser

import (
	"errors"
	"log/slog"
	"regexp"
	"strings"
	"time"
	"window/pkg/models"
)

var (
	ErrNoDate  = errors.New("err cannot find date")
	ErrNoTime  = errors.New("err cannot find time")
	ErrNoName  = errors.New("err cannot find name or name is invalid")
	ErrBadDate = errors.New("err bad date format")
	ErrBadTime = errors.New("err bad time format")
)

var (
	rdate, _ = regexp.Compile(`(\d{4})+\.+(\d{2})+\.+(\d{2})`)
	rtime, _ = regexp.Compile(`(\d{2})+\:+(\d{2})`)
	rname, _ = regexp.Compile(`\"([A-Z]([a-z]){1,}\s?){1,3}\"`)
)

const (
	fdate = "2006.01.02"
	ftime = "15:04"
)

func ProcessInput(input string) (models.Lesson, error) {
	lesson := models.Lesson{}
	lesson.Date = string(rdate.Find([]byte(input)))
	if lesson.Date == "" {
		slog.Error("ProcessInput: cannot find date", "input", input)
		return models.Lesson{}, ErrNoDate
	}
	if _, err := time.Parse(fdate, lesson.Date); err != nil {
		slog.Error("ProcessInput: cannot parse date", "err", err, "date", lesson.Date)
		return models.Lesson{}, ErrBadDate
	}

	lesson.Time = string(rtime.Find([]byte(input)))
	if lesson.Time == "" {
		slog.Error("ProcessInput: cannot find time", "input", input)
		return models.Lesson{}, ErrNoTime
	}
	if _, err := time.Parse(ftime, lesson.Time); err != nil {
		slog.Error("ProcessInput: cannot parse time", "err", err, "time", lesson.Time)
		return models.Lesson{}, ErrBadTime
	}

	lesson.TeacherName = string(rname.Find([]byte(input)))
	if lesson.TeacherName == "" {
		slog.Error("ProcessInput: cannot find name", "input", input)
		return models.Lesson{}, ErrNoName
	}

	lesson.TeacherName = strings.ReplaceAll(lesson.TeacherName, "\"", "")
	return lesson, nil
}
