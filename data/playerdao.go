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
}

func UpdatePlayer(score int, idPlayer float64) {
	insertScoreSQL := `UPDATE players SET score  = ? where id = ? `
	statement, err := db.Prepare(insertScoreSQL)

	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(score, idPlayer)

	if err != nil {
		log.Fatalln(err)
	}
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
		playersResult = append(playersResult, Player{Id: id, Username: username, Score: score})
	}
	return playersResult
}
func GetPlayer(idPlayer float64) Player {
	row, err := db.Query("SELECT * FROM players where id = $1", idPlayer)
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()
	var playersResult Player
	for row.Next() {
		var id float64
		var username string
		var score float64
		row.Scan(&id, &username, &score)
		playersResult = Player{Id: id, Username: username, Score: score}
	}
	return playersResult
}
