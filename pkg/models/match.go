package models

type Match struct {
	ID         int
	TotalKills int
	Players    map[int]*Player
	Kills      map[int]int
	KillsByMod map[string]int
}
