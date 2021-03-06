package data

type Player struct {
	Id         float64 `json:"id"`
	Username   string  `json:"username"`
	Score      float64 `json:"score"`
	Percentage float64 `json:"percentage"`
	GameMatch  int     `json:"matches"`
}

type Scores struct {
	PlayerId float64 `json:"playerId"`
	Question string  `json:"question"`
	Outcome  bool    `json:"outcome"`
}
