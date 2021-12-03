package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// func userInput() (string, string) {
// 	var user string
// 	var password string
// 	fmt.Print("Enter new user: ")
// 	fmt.Scanln(&user)
// 	fmt.Print("Enter password: ")
// 	fmt.Scanln(&password)
// 	return user, password
// }

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/myDB?parseTime=true")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to DB!")
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db, err
}

// func createTable() {
// 	db, _ := connectDB()
// 	query := `
// 	CREATE TABLE users (
// 	    id INT AUTO_INCREMENT,
// 	    username TEXT NOT NULL,
// 	    password TEXT NOT NULL,
// 	    created_at DATETIME,
// 	    PRIMARY KEY (id)
// 	);`

// 	_, err := db.Exec(query)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		fmt.Println("Table created")
// 	}
// }

func insertRow(username string, password string) {
	db, _ := connectDB()

	// username := "admasalim"
	// password := "hellopwd"
	//username, password := userInput()
	createdAt := time.Now()

	result, _ := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	userID, _ := result.LastInsertId()
	fmt.Printf("User %v has been added to the DB with user ID %v \n", username, userID)
}

func deleteRow(id int64) {
	db, _ := connectDB()
	db.Exec(`DELETE FROM users WHERE id=?`, id)
}

func queryDB() {
	db, _ := connectDB()
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)

	var queryStr string = "SELECT * FROM users"
	rows, err := db.Query(queryStr)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &username, &password, &createdAt)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(id, username, password, createdAt)
	}

	// query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
	// db_query := db.QueryRow(query, 6).Scan(&id, &username, &password, &createdAt)
	// if db_query != nil {
	// 	log.Fatal(db_query)
	// } else {
	// 	fmt.Println(id, username, password, createdAt)
	// }

}

// func main() {
// 	createTable()
// 	insertRow()
// 	deleteRow()
// 	queryDB()
// }
