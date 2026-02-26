package tabler

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"window/pkg/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const mockTableName = "mock_table"

var mockData = [][]string{{"2000.10.10", "12:30", "Valeriy Petrovich"}, {"2000.10.11", "12:30", "Valeriy Petrovich"}}
var mockBadData = [][]string{{"2000.1000.10", "12:332a0", "Valerg12iy Pe6trovich"}, {"2000....10.1123", "12::30", "ValeriyPetrovich]"}}

var mockTabler *Tabler

func TestMain(t *testing.T) {
	var err error
	t.Log("Creating mock tabler...")
	mockTabler, err = NewTabler(mockTableName)
	t.Log("Mock tabler created successfully.")
	assert.Nil(t, err)
}

func TestSaveTable(t *testing.T) {
	t.Log("Good data save test")
	err := mockTabler.SaveTable(mockData)
	assert.Nil(t, err)
	if len(err) == 0 {
		defer killTable()
	}
	t.Log("Bad data save test")
	err = mockTabler.SaveTable(mockBadData)
	if len(err) == 0 {
		t.Error("No errors")
	} else {
		require.Error(t, err[0], errors.New(ErrBadLine.Error()+fmt.Sprintf("%v %v \"%v\"\n", mockBadData[0][0], mockBadData[0][1], mockBadData[0][2])))
		require.Error(t, err[1], errors.New(ErrBadLine.Error()+fmt.Sprintf("%v %v \"%v\"\n", mockBadData[1][0], mockBadData[1][1], mockBadData[1][2])))
	}
}

func TestInitTable(t *testing.T) {
	err := mockTabler.SaveTable(mockData)
	if len(err) == 0 {
		defer killTable()
	}

	shouldLessons := []models.Lesson{models.Lesson{Date: "2000.10.10", Time: "12:30", TeacherName: "Valeriy Petrovich"}, models.Lesson{Date: "2000.10.11", Time: "12:30", TeacherName: "Valeriy Petrovich"}}
	lessons := mockTabler.InitTable()
	require.Equal(t, shouldLessons, lessons)
}

func killTable() {
	os.Remove(mockTableName + ".txt")
}
