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

func DropTableQuizScore() {
	sql := `DROP TABLE quiz_scores;`

	statement, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("quiz_scores table dropped")
}
func DropTablePlayer() {
	sql := `DROP TABLE players;`

	statement, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("players table dropped")
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
func CreateTableQuizScores() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS quiz_scores (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"id_player" INTEGER NOT NULL,
		"question" TEXT,
		"outcome" TINYINT UNSIGNED,
		CONSTRAINT fk_quiz_scores__id_player FOREIGN KEY (id_player) REFERENCES players(id)
	  );`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("quiz_scores table created")
}
