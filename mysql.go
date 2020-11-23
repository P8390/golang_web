package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	query := `
	CREATE TABLE golang (
	    id INT AUTO_INCREMENT,
	    username TEXT NOT NULL,
	    password TEXT NOT NULL,
	    created_at DATETIME,
	    PRIMARY KEY (id)
	);`
	query := `INSERT INTO golang (username, password, created_at) VALUES (?, ?, ?)`
	res, _ := db.Exec(query, "Pankaj", "test", time.Now())
	userId, _ := res.LastInsertId()
	fmt.Println("Last Id - ", userId)

	// Query to Single Row
	var (
		id         int
		username   string
		password   string
		created_at time.Time
	)
	query := `SELECT id, username, password, created_at FROM golang WHERE id = ?`
	err = db.QueryRow(query, 1).Scan(&id, &username, &password, &created_at)

	// Query all rows
	type user struct {
		id         int
		username   string
		password   string
		created_at time.Time
	}
	rows, err := db.Query(`SELECT id, username, password, created_at FROM golang`)
	defer rows.Close()

	var users []user

	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password, &u.created_at)
		fmt.Println(err)
		users = append(users, u)
	}
	err1 := rows.Err() // check err
	fmt.Println(err1)
	fmt.Println(users)
	_, err2 := db.Exec(`DELETE FROM users WHERE id = ?`, 1) // check err
	fmt.Println(err2)
}
