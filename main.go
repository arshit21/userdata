package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

// type User struct {
// 	ID      uuid.UUID
// 	S_ID    string
// 	IDNO    string
// 	SNAME   string
// 	HOSCODE string
// 	ROOM    int
// 	// BMAIL   string    `json:"bmail" bson:"bmail"`
// }

func createUser() {

	f, err := excelize.OpenFile("student.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Fatalf("%v", err)
	}
	// length := len(rows)

	for _, row := range rows[1:] { // Skip header row
		user := User{
			S_ID:    row[0],
			IDNO:    row[1],
			SNAME:   row[2],
			HOSCODE: row[4],
			ROOM:    row[5],
		}

		err := InsertUser(user)
		if err != nil {
			log.Println(err)
			// Handle error if necessary
		}
	}
}

func main() {
	createUser()
}
