package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//LoadJson Read param file if exist and unmarshal them
func LoadJson(pathStudent string, pathPlanning string) (st StudentRoot, pl PlanningRoot, err error) {
	raw, err := ioutil.ReadFile(pathStudent)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	json.Unmarshal(raw, &st)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	raw2, err := ioutil.ReadFile(pathPlanning)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = json.Unmarshal(raw2, &pl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
