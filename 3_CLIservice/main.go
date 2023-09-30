package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Student struct {
	id      int
	name    string
	address string
	job     string
	reason  string
}

func (student Student) prettyPrint() string {
	s := reflect.ValueOf(student)
	typeOfS := reflect.TypeOf(student)
	res := ""
	for i := 0; i < s.NumField(); i++ {
		res += fmt.Sprintf("%-12s : %v\n", typeOfS.Field(i).Name, s.Field(i))
	}
	return res
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Tolong masukan nama atau nomor absen")
		fmt.Println("Contoh: 'go run main.go jamal' atau 'go run main.go 3'")
		return
	}
	student := getStudent(os.Args[1])
	fmt.Println(student)
}

func getStudent(q string) string {

	var students = []Student{
		{id: 0, name: "xever", address: "66209 Surrey Way", job: "Internal Auditor", reason: "Re-engineered reciprocal artificial intelligence"},
		{id: 1, name: "stanfield", address: "90648 Londonderry Place", job: "Payment Adjustment Coordinator", reason: "Total solution-oriented open system"},
		{id: 2, name: "jeromy", address: "506 Schlimgen Plaza", job: "Tax Accountant", reason: "Cross-platform clear-thinking software"},
		{id: 3, name: "simon", address: "2 Manitowish Trail", job: "Human Resources Manager", reason: "User-centric even-keeled solution"},
		{id: 4, name: "hersh", address: "348 Nancy Court", job: "Programmer III", reason: "Centralized responsive workforce"},
		{id: 5, name: "edd", address: "3368 Raven Terrace", job: "Geologist I", reason: "Vision-oriented zero defect implementation"},
		{id: 6, name: "berkeley", address: "1224 Village Green Terrace", job: "Nurse Practicioner", reason: "Intuitive disintermediate flexibility"},
		{id: 7, name: "marcello", address: "843 Northview Circle", job: "Help Desk Operator", reason: "Grass-roots logistical conglomeration"},
		{id: 8, name: "kirby", address: "1 Carpenter Crossing", job: "Senior Cost Accountant", reason: "Advanced 6th generation framework"},
		{id: 9, name: "guntar", address: "6116 Corscot Terrace", job: "Health Coach III", reason: "Function-based foreground toolset"},
	}

	id, err := strconv.Atoi(q)

	switch {
	case err == nil:
		for _, s := range students {
			if id == s.id {
				return s.prettyPrint()
			}
		}
		return fmt.Sprintf("Tidak ada yang memiliki nomor absen %+v\n", id)
	default:
		for _, s := range students {
			if q == s.name {
				return s.prettyPrint()
			}
		}
		return fmt.Sprintf("Tidak ada yang bernama %+v\n", q)
	}
}
