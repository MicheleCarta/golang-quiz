package persistence

import "log"

func InsertNote(name string, score float64) {
	insertScoreSQL := `INSERT INTO quiz_scores(name, score) VALUES (?, ?)`
	statement, err := db.Prepare(insertScoreSQL)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(name, score)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Inserted quiz_scores successfully")
}

func DisplayAllScores() {
	row, err := db.Query("SELECT * FROM quiz_scores ORDER BY word")
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	for row.Next() {
		var playerId int
		var name string
		var score float64
		row.Scan(&playerId, &name, &score)
		log.Println("[", playerId, "] ", name, "—", score)
	}
}
