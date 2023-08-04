package main

import (
	"IntrovertBot/client"
	"IntrovertBot/object"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	user := client.Client{Id: 1, Name: "Elina"}
	obj := object.Object{Id: 1, Name: "fffffff", Author: "dddddddd", Genre: "ssssss", Year: 2021, Adviser: "aaaaaaaaa", UserId: user.Id}

	db, err := sql.Open("sqlite3", "DataBase/test.db")
	if err != nil {
		fmt.Errorf("can't connect with data")
	}
	table := "films"
	query2 := `INSERT INTO table (name, author, genre, year, adviser, done, user_id) VALUES (?, ?, ?, ?, ?, ?, (SELECT id FROM users WHERE name = ?))`
	if _, err := db.Exec(query2, table, obj.Name, obj.Author, obj.Genre, obj.Year, obj.Adviser, obj.Done, user.Name); err != nil {
		fmt.Errorf("can't connect with data")
	}

	//if _, err := db.Exec(
	//	`INSERT INTO films (name, author, genre, year, adviser, done, user_id) VALUES ($1, $2, $3, $4, $5, $6, (SELECT id FROM users WHERE name = "Elina"))`,
	//	obj.Name, obj.Author, obj.Genre, obj.Year, obj.Adviser, obj.Done); err != nil {
	//	fmt.Errorf("can't connect with data")
	//}

	//db, err := DataBase.New("DataBase/test.db")
	//if err != nil {
	//	fmt.Errorf("can't connect with data")
	//}
	//
	//err = db.SaveFB("films", &obj)
	//if err != nil {
	//	fmt.Errorf("can't save data")
	//}
	//
	//slice, err1 := db.ShowAllFB("films", 1)
	//if err1 != nil {
	//	fmt.Errorf("can't show data")
	//}
	//
	//for _, e := range *slice {
	//	fmt.Println(e.Name, e.Author)
	//}
}
