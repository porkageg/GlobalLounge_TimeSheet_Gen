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
