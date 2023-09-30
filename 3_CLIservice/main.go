package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Student struct {
	// id      int
	name    string
	address string
	job     string
	reason  string
}

func (student Student) prettyPrint(i int) string {
	s := reflect.ValueOf(student)
	typeOfS := reflect.TypeOf(student)
	res := fmt.Sprintf("%-12s : %v\n", "id", i)
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
		{name: "xever", address: "66209 Surrey Way", job: "Internal Auditor", reason: "Re-engineered reciprocal artificial intelligence"},
		{name: "stanfield", address: "90648 Londonderry Place", job: "Payment Adjustment Coordinator", reason: "Total solution-oriented open system"},
		{name: "jeromy", address: "506 Schlimgen Plaza", job: "Tax Accountant", reason: "Cross-platform clear-thinking software"},
		{name: "simon", address: "2 Manitowish Trail", job: "Human Resources Manager", reason: "User-centric even-keeled solution"},
		{name: "hersh", address: "348 Nancy Court", job: "Programmer III", reason: "Centralized responsive workforce"},
		{name: "edd", address: "3368 Raven Terrace", job: "Geologist I", reason: "Vision-oriented zero defect implementation"},
		{name: "berkeley", address: "1224 Village Green Terrace", job: "Nurse Practicioner", reason: "Intuitive disintermediate flexibility"},
		{name: "marcello", address: "843 Northview Circle", job: "Help Desk Operator", reason: "Grass-roots logistical conglomeration"},
		{name: "kirby", address: "1 Carpenter Crossing", job: "Senior Cost Accountant", reason: "Advanced 6th generation framework"},
		{name: "guntar", address: "6116 Corscot Terrace", job: "Health Coach III", reason: "Function-based foreground toolset"},
	}

	id, err := strconv.Atoi(q)

	switch {
	case err == nil:
		if 0 <= id && id < len(students) {
			return students[id].prettyPrint(id)
		}
		return fmt.Sprintf("Tidak ada yang memiliki nomor absen %+v\n", id)
	default:
		for i, s := range students {
			if q == s.name {
				return s.prettyPrint(i)
			}
		}
		return fmt.Sprintf("Tidak ada yang bernama %+v\n", q)
	}
}
