package models

type Match struct {
	ID         int
	TotalKills int
	Players    map[int]*Player
	Kills      map[int]int
	KillsByMOD map[string]int
}
