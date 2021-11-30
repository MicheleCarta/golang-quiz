package data

import "log"

func InsertScore(idPlayer float64, question string, outcome bool) {
	log.Println("InsertScore --> ", idPlayer, " ", question, " ", outcome)

	insertScoreSQL := `INSERT INTO quiz_scores(id_player, question, outcome) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertScoreSQL)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(idPlayer, question, outcome)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Inserted quiz_scores successfully")
}

func DisplayAllScores() {
	row, err := db.Query("SELECT * FROM quiz_scores ORDER BY id_player")
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	for row.Next() {
		var id_player float64
		var question string
		var outcome int
		row.Scan(&id_player, &question, &outcome)
		log.Println("[", id_player, "] ", question, "—", outcome)
	}
}

func GetScoresPlayer(idPlayer float64) []Scores {
	log.Println("playerId is -> ", idPlayer)
	row, err := db.Query("SELECT * FROM quiz_scores where id_player = $1", idPlayer)
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()
	var scoresResult []Scores
	for row.Next() {
		var id_player float64
		var question string
		var outcome bool
		row.Scan(&id_player, &question, &outcome)
		log.Println("[", id_player, "] ", question, "—", outcome)
		scoresResult = append(scoresResult, Scores{PlayerId: id_player, Question: question, Outcome: outcome})
	}
	return scoresResult
}
