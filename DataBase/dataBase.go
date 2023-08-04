package DataBase

import (
	"IntrovertBot/object"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	sql *sql.DB
}

func New(dbPath string) (*DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &DB{sql: db}, nil
}

func (db *DB) SaveFB(tableName string, obj *object.Object) error {
	query := `INSERT INTO films (name, author, genre, year, adviser, done, userID) VALUES (?, ?, ?, ?, ?, ?, ?)`
	if _, err := db.sql.Exec(
		query,
		obj.Name,
		obj.Author,
		obj.Genre,
		obj.Year,
		obj.Adviser,
		obj.Done,
		obj.UserId); err != nil {
		return err
	}
	return nil
}

//SaveLink(tableName string, obj *Object) error

func (db *DB) ShowAllFB(tableName string, userID int) (*[]object.Object, error) {
	query := `SELECT * FROM ? WHERE id = ?`

	rows, err := db.sql.Query(query, tableName, userID)
	if err != nil {
		return nil, err
	}
	//defer rows.Close()

	objects := []object.Object{}

	for rows.Next() {
		obj := object.Object{}
		if err := rows.Scan(&obj.Id, &obj.Name, &obj.Author, &obj.Genre, &obj.Year, &obj.Adviser, &obj.Done, &obj.UserId); err != nil {
			return nil, err
		}
		objects = append(objects, obj)
	}
	return &objects, nil
}

//	ShowOne(tableName string, userName string) (**Object, error)
//	NoteDone(tableName string, obj *Object) error
//	Remove(tableName string, obj *Object) error
//	PickRandom(tableName string, obj *Object) error
//	IsExists(tableName string, obj *Object) (bool, error)
