package parser

import (
	"quake_log_parser/pkg/models"
	"strconv"
	"strings"
)
func parsePlayer(line string) *models.Player {
	parts := strings.Split(line, " ")
	playerID, _ := strconv.Atoi(parts[2])
	infoParts := strings.Split(parts[3], "\\")
	playerName := infoParts[1]

	return &models.Player{ID: playerID, Name: playerName}
}

func parseKill() {
	return
}

func processKill() {
	return
}

func ParseLogFile() {
	return
}
