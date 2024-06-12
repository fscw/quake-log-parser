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

func parseKill(line string) models.Kill {
	parts := strings.Split(line, " ")
	killerID, _ := strconv.Atoi(parts[3])
	victimID, _ := strconv.Atoi(parts[4])
	modID, _ := strconv.Atoi(parts[5])

	return models.Kill{KillerID: killerID, VictimID: victimID, ModID: modID}
}

func processKill() {
	return
}

func ParseLogFile() {
	return
}
