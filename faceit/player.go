package faceit

type Game struct {
	Level int `json:"skill_level"`
	Elo   int `json:"faceit_elo"`
}

type Player struct {
	Id      string          `json:"player_id"`
	Name    string          `json:"nickname"`
	SteamId string          `json:"steam_id_64"`
	Games   map[string]Game `json:"games"`
}

type MatchTab struct {
	Matches []Match `json:"items"`
}

type Match struct {
	Id   string   `json:"match_id"`
	Demo []string `json:"demo_url"`
}
