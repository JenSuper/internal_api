package sqllite_base

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// 定义一个 User 结构体，用于表示数据库中的用户
type User struct {
	ID    int
	Name  string
	Email string
}

func Init() {
	// 打开或创建 SQLite 数据库文件
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建一个表格
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	// 插入数据
	insertSQL := `INSERT INTO users (name, email) VALUES (?, ?)`
	_, err = db.Exec(insertSQL, "Alice", "alice@example.com")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(insertSQL, "Bob", "bob@example.com")
	if err != nil {
		log.Fatal(err)
	}

	// 查询数据
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}

	// 输出查询结果
	fmt.Println("Users:")
	for _, u := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", u.ID, u.Name, u.Email)
	}

	// 更新数据
	updateSQL := `UPDATE users SET email = ? WHERE name = ?`
	_, err = db.Exec(updateSQL, "alice_new@example.com", "Alice")
	if err != nil {
		log.Fatal(err)
	}

	// 删除数据
	deleteSQL := `DELETE FROM users WHERE name = ?`
	_, err = db.Exec(deleteSQL, "Bob")
	if err != nil {
		log.Fatal(err)
	}

	// 查询更新后的数据
	rows, err = db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Updated Users:")
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", u.ID, u.Name, u.Email)
	}
}
