package main

import (
	"fmt"
	"time"

	"github.com/araddon/dateparse"
	"github.com/tealeg/xlsx"
	"github.com/tsGen/model"
)

var (
	weekMap = map[time.Weekday]int32{
		time.Monday:    1,
		time.Tuesday:   2,
		time.Wednesday: 3,
		time.Thursday:  4,
		time.Friday:    5,
		time.Saturday:  6,
		time.Sunday:    7,
	}
)

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%d:%02d", h, m)
}

func generateTimeSheet(st *model.StudentRoot, pl *model.PlanningRoot) (lines []*model.Line, err error) {
	today, err := dateparse.ParseLocal(pl.Start)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	end, err := dateparse.ParseLocal(pl.Stop)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var i int32
	for end.Sub(today) >= 0 {
		if today.Weekday() != time.Saturday && today.Weekday() != time.Sunday {
			planning := pl.Days[today.Weekday()-1].Planning
			for _, p := range planning {
				i++
				student := st.Students[p.StudentID-1]                                           //Get Student Info
				date := fmt.Sprintf("%04d/%02d/%02d", today.Year(), today.Month(), today.Day()) //Format Date
				begin, err := dateparse.ParseLocal(date + " " + p.Begin)                        // Parse begin hour
				if err != nil {
					fmt.Println(err.Error())
					return nil, err
				}
				end, err := dateparse.ParseLocal(date + " " + p.End) // Parse end hour
				if err != nil {
					fmt.Println(err.Error())
					return nil, err
				}
				timeWork := end.Sub(begin).Round(time.Minute) // Compute time worked
				workHour := fmtDuration(timeWork)             // Format time worked
				line := model.Line{                           // Create Line Data
					ID:        i,
					StartTime: p.Begin,
					EndTime:   p.End,
					Name:      student.Name,
					StudentID: student.IDCard,
					Major:     student.Major,
					Date:      date,
					WorkHour:  workHour,
				}
				lines = append(lines, &line)
			}
		}
		today = today.AddDate(0, 0, 1)
	}

	return
}

func main() {
	st, pl, err := model.LoadJson("./input/student.json", "./input/planning.json")
	if err != nil {
		return // error already printed
	}

	lines, err := generateTimeSheet(&st, &pl)
	if err != nil {
		return // error already printed
	}

	writeXlsx(lines)

	return
}

func addCell(row *xlsx.Row, value string) {
	cell := row.AddCell()
	cell.Value = value
}

func createHeader(sheet *xlsx.Sheet) {
	row := sheet.AddRow()
	addCell(row, "Global Lounge Operator Attendance Book")

	sheet.AddRow()

	rowTable := sheet.AddRow()
	addCell(rowTable, "NO")
	addCell(rowTable, "Date")
	addCell(rowTable, "Name")
	addCell(rowTable, "Student ID")
	addCell(rowTable, "Major")
	addCell(rowTable, "Start Time")
	addCell(rowTable, "End Time")
	addCell(rowTable, "Work Hour")
	addCell(rowTable, "Reason")
}

func writeBody(line *model.Line, sheet *xlsx.Sheet) {
	row := sheet.AddRow()

	addCell(row, fmt.Sprintf("%02d", line.ID))
	addCell(row, line.Date)
	addCell(row, line.Name)
	addCell(row, line.StudentID)
	addCell(row, line.Major)
	addCell(row, line.StartTime)
	addCell(row, line.EndTime)
	addCell(row, line.WorkHour)
	addCell(row, " ")
}

func writeXlsx(lines []*model.Line) (err error) {
	var file *xlsx.File
	var sheet *xlsx.Sheet

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	//Header
	createHeader(sheet)

	//Body
	for _, line := range lines {
		writeBody(line, sheet)
	}

	err = file.Save("GlobalLoungeOperatorAttendanceBook.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
	return
}
