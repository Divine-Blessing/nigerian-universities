package main

import (
	"fmt"
)

type University struct {
	Id int
	Abbrv string
	UniversityName string
}

func UniList() []*University{
	return []*University {
		&University{Id: 1, Abbrv: "ABSU", UniversityName: "Abia State University"},
		&University{Id: 2, Abbrv: "AUO", UniversityName: "Achievers University"},
		&University{Id: 3, Abbrv: "ADSU", UniversityName: "Adamawa State University"},
	}
}

func main() {
	listOfUni := UniList()
	
	for _, list := range listOfUni {
		fmt.Println(list.Id, list.Abbrv, list.UniversityName)
	}
}