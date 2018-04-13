package model

//StudentRoot is the root struct use to load student.json
type StudentRoot struct {
	Students []student `json:"students"`
}

type student struct {
	ID     int32  `json:"id"`
	Name   string `json:"name"`
	IDCard string `json:"idCard"`
	Major  string `json:"major"`
}

//PlanningRoot is the root struct use to load planning.json
type PlanningRoot struct {
	Start string `json:"start"`
	Stop  string `json:"stop"`
	Days  []day  `json:"days"`
}

type day struct {
	Name     string     `json:"name"`
	ID       int32      `json:"id"`
	Planning []planning `json:"planning"`
}

type planning struct {
	Begin     string `json:"begin"`
	End       string `json:"end"`
	StudentID int32  `json:"studentId"`
}

//Line is the data that is going to be write in a line of the xlsx files
type Line struct {
	ID        int32
	Date      string
	Name      string
	StudentID string
	Major     string
	StartTime string
	EndTime   string
	WorkHour  string
}
