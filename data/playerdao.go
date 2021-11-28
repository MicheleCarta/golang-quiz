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

func DisplayAllPlayers() {
	row, err := db.Query("SELECT * FROM players ORDER BY id_player")
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	for row.Next() {
		var id int
		var name string
		var score float64
		row.Scan(&id, &name, &score)
		log.Println("[", id, "] ", name, "—", score)
	}
}
