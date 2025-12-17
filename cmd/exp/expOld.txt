package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode,
	)
}

func main() {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5555",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}

	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	// Create tables FIRST
	_, err = db.Exec(`
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  age INT NOT NULL,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  email TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS orders (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  amount INT,
  description TEXT
);`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tables created.")

	// age := 25
	// firstName := "Jon"
	// lastName := "Calhoun"
	// email := "jfon@calhoun.iof"

	// 	// Insert + get id back
	// 	var newID int
	// 	err = db.QueryRow(`
	//   INSERT INTO users(age, first_name, last_name, email)
	//   VALUES ($1, $2, $3, $4)
	//   RETURNING id;
	// `, age, firstName, lastName, email).Scan(&newID)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("User created. id =", newID)

	// 	// Query just the new user
	// 	var (
	// 		id       int
	// 		gotEmail string
	// 	)
	// 	err = db.QueryRow(`SELECT id, email FROM users WHERE id = $1`, newID).Scan(&id, &gotEmail)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("id=%d email=%s\n", id, gotEmail)

	userID := 5
	for i := 1; i <= 5; i++ {
		amount := i * 100
		desc := fmt.Sprintf("Fake order #%d", i)
		_, err := db.Exec(`
    INSERT INTO orders(user_id, amount, description)
    VALUES($1, $2, $3)`, userID, amount, desc)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Created fake orders.")
}
