package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {
	var err error

	db, err = sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}

	return db.Ping()
}

func CreateTableQuizScores() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS quiz_scores (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"id_player" INTEGER NOT NULL,
		"name" TEXT,
		"score" DOUBLE,
		CONSTRAINT fk_quiz_scores__id_player FOREIGN KEY (id_player) REFERENCES players(id)
	  );`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("quiz_scores table created")
}
func CreateTablePlayers() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS players (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"username" TEXT,
		"score" DOUBLE
	  );`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("players table created")
}
