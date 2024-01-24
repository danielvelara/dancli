package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	id   int
	name string
	age  int
}

func RunQuery() {
	db := initDB()
	defer db.Close()

	err := CreateUser(db, User{
		name: "John",
		age:  42,
	})
	if err != nil {
		log.Fatal(err)
	}

	users, err := GetUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range users {
		fmt.Println(u)
	}
	fmt.Println("Done")
}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		log.Fatal(err)
	}

	createTableQuery := `
  CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    age INTEGER NOT NULL
  );
  `
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetUsers(db *sql.DB) ([]User, error) {
	// q := "SELECT id, name, age FROM users"
	q := "SELECT id, name, age FROM users"
	// q := "SELECT id, name, age FROM users; DROP TABLE users;"
	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.id, &user.name, &user.age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func CreateUser(db *sql.DB, user User) error {
	// _injection := "text'); DROP TABLE logs; --"
	// insertQuery := "INSERT INTO users(name, age) VALUES (?, ?);"
	insertQuery := "INSERT INTO users(name, age) VALUES ($1, $2);"
	// stmt, err := db.Prepare(insertQuery)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer stmt.Close()

	// _, err = stmt.Exec(user.Name, user.Age)
	_, err := db.Exec(insertQuery, user.name, user.age)
	return err
}

func UpdateUser(db *sql.DB, user User) error {
	updateQuery := `
  UPDATE users
  SET name=?, age=? WHERE id=?;`
	stmt, err := db.Prepare(updateQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.name, user.age, user.id)
	return err
}

func DeleteUser(db *sql.DB, id int) error {
	deleteQuery := `DELETE FROM users WHERE id=?`
	stmt, err := db.Prepare(deleteQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}
