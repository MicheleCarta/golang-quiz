package data

import "log"

func InsertPlayer(name string, score float64) {
	insertScoreSQL := `INSERT INTO players(username, score) VALUES (?, ?)`
	statement, err := db.Prepare(insertScoreSQL)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(name, score)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Inserted player successfully")
}

func DisplayAllPlayers() []Player {
	row, err := db.Query("SELECT * FROM players ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()
	var playersResult []Player
	for row.Next() {
		var id float64
		var username string
		var score float64
		row.Scan(&id, &username, &score)
		log.Println("[", id, "] ", username, "—", score)
		playersResult = append(playersResult, Player{Id: id, Username: username, Score: score})
	}
	return playersResult
}