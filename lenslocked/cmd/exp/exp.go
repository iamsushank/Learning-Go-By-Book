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

// String returns a string representation of the PostgresConfig
func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func main() {

	db, err := sql.Open("pgx", "host=localhost port=5432 user=baloo password=junglebook dbname=lenslocked")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
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

	name := "New User"
	email := "mandal678@gmail.com"
	row := db.QueryRow(`
		INSERT INTO users (name, email)
		VALUES ($1, $2) RETURNING id`, name, email)
	var id int
	err = row.Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("User created. id=", id)

}

type Order struct {
	ID          int
	UserID      int
	Amount      int
	Description string
}

