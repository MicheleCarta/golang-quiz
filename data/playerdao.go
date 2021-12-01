package data

import "log"

func InsertPlayer(name string, score float64, percentage float64) {
	insertScoreSQL := `INSERT INTO players(username, score,percentage,game_match) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insertScoreSQL)

	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(name, score, percentage, 0)

	if err != nil {
		log.Fatalln(err)
	}
}

func UpdatePlayer(score int, idPlayer float64, currentScore int, percentage float64, gameMatch int) {
	insertScoreSQL := `UPDATE players SET score  = ?, percentage = ?, game_match = ? where id = ? `
	statement, err := db.Prepare(insertScoreSQL)

	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec((score + currentScore), percentage, gameMatch, idPlayer)

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
		var percentage float64
		var matches int
		row.Scan(&id, &username, &score, &percentage, &matches)
		playersResult = append(playersResult, Player{Id: id, Username: username, Score: score, Percentage: percentage, GameMatch: matches})
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
		var percentage float64
		var matches int
		row.Scan(&id, &username, &score, &percentage, &matches)
		playersResult = Player{Id: id, Username: username, Score: score, Percentage: percentage, GameMatch: matches}
	}
	return playersResult
}
